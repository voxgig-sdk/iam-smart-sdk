package sdktest

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"

	sdk "github.com/voxgig-sdk/iam-smart-sdk"
	"github.com/voxgig-sdk/iam-smart-sdk/core"

	vs "github.com/voxgig/struct"
)

func TestRegistrationServiceCounterEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.RegistrationServiceCounter(nil)
		if ent == nil {
			t.Fatal("expected non-nil RegistrationServiceCounterEntity")
		}
	})

	t.Run("basic", func(t *testing.T) {
		setup := registration_service_counterBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"list"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "registration_service_counter." + _op, _mode); _shouldSkip {
				if _reason == "" {
					_reason = "skipped via sdk-test-control.json"
				}
				t.Skip(_reason)
				return
			}
		}
		// The basic flow consumes synthetic IDs from the fixture. In live mode
		// without an *_ENTID env override, those IDs hit the live API and 4xx.
		if setup.syntheticOnly {
			t.Skip("live entity test uses synthetic IDs from fixture — set IAMSMART_TEST_REGISTRATION_SERVICE_COUNTER_ENTID JSON to run live")
			return
		}
		client := setup.client

		// Bootstrap entity data from existing test data (no create step in flow).
		registrationServiceCounterRef01DataRaw := vs.Items(core.ToMapAny(vs.GetPath("existing.registration_service_counter", setup.data)))
		var registrationServiceCounterRef01Data map[string]any
		if len(registrationServiceCounterRef01DataRaw) > 0 {
			registrationServiceCounterRef01Data = core.ToMapAny(registrationServiceCounterRef01DataRaw[0][1])
		}
		// Discard guards against Go's unused-var check when the flow's steps
		// happen not to consume the bootstrap data (e.g. list-only flows).
		_ = registrationServiceCounterRef01Data

		// LIST
		registrationServiceCounterRef01Ent := client.RegistrationServiceCounter(nil)
		registrationServiceCounterRef01Match := map[string]any{}

		registrationServiceCounterRef01ListResult, err := registrationServiceCounterRef01Ent.List(registrationServiceCounterRef01Match, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		_, registrationServiceCounterRef01ListOk := registrationServiceCounterRef01ListResult.([]any)
		if !registrationServiceCounterRef01ListOk {
			t.Fatalf("expected list result to be an array, got %T", registrationServiceCounterRef01ListResult)
		}

	})
}

func registration_service_counterBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "registration_service_counter", "RegistrationServiceCounterTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read registration_service_counter test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse registration_service_counter test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap := vs.Transform(
		[]any{"registration_service_counter01", "registration_service_counter02", "registration_service_counter03"},
		map[string]any{
			"`$PACK`": []any{"", map[string]any{
				"`$KEY`": "`$COPY`",
				"`$VAL`": []any{"`$FORMAT`", "upper", "`$COPY`"},
			}},
		},
	)

	// Detect ENTID env override before envOverride consumes it. When live
	// mode is on without a real override, the basic test runs against synthetic
	// IDs from the fixture and 4xx's. Surface this so the test can skip.
	entidEnvRaw := os.Getenv("IAMSMART_TEST_REGISTRATION_SERVICE_COUNTER_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"IAMSMART_TEST_REGISTRATION_SERVICE_COUNTER_ENTID": idmap,
		"IAMSMART_TEST_LIVE":      "FALSE",
		"IAMSMART_TEST_EXPLAIN":   "FALSE",
		"IAMSMART_APIKEY":         "NONE",
	})

	idmapResolved := core.ToMapAny(env["IAMSMART_TEST_REGISTRATION_SERVICE_COUNTER_ENTID"])
	if idmapResolved == nil {
		idmapResolved = core.ToMapAny(idmap)
	}

	if env["IAMSMART_TEST_LIVE"] == "TRUE" {
		mergedOpts := vs.Merge([]any{
			map[string]any{
				"apikey": env["IAMSMART_APIKEY"],
			},
			extra,
		})
		client = sdk.NewIamSmartSDK(core.ToMapAny(mergedOpts))
	}

	live := env["IAMSMART_TEST_LIVE"] == "TRUE"
	return &entityTestSetup{
		client:        client,
		data:          entityData,
		idmap:         idmapResolved,
		env:           env,
		explain:       env["IAMSMART_TEST_EXPLAIN"] == "TRUE",
		live:          live,
		syntheticOnly: live && !idmapOverridden,
		now:           time.Now().UnixMilli(),
	}
}
