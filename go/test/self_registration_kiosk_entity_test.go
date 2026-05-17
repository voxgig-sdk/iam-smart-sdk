package sdktest

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"

	sdk "github.com/voxgig-sdk/iam-smart-sdk/go"
	"github.com/voxgig-sdk/iam-smart-sdk/go/core"

	vs "github.com/voxgig-sdk/iam-smart-sdk/go/utility/struct"
)

func TestSelfRegistrationKioskEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.SelfRegistrationKiosk(nil)
		if ent == nil {
			t.Fatal("expected non-nil SelfRegistrationKioskEntity")
		}
	})

	t.Run("basic", func(t *testing.T) {
		setup := self_registration_kioskBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"list"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "self_registration_kiosk." + _op, _mode); _shouldSkip {
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
			t.Skip("live entity test uses synthetic IDs from fixture — set IAMSMART_TEST_SELF_REGISTRATION_KIOSK_ENTID JSON to run live")
			return
		}
		client := setup.client

		// Bootstrap entity data from existing test data (no create step in flow).
		selfRegistrationKioskRef01DataRaw := vs.Items(core.ToMapAny(vs.GetPath("existing.self_registration_kiosk", setup.data)))
		var selfRegistrationKioskRef01Data map[string]any
		if len(selfRegistrationKioskRef01DataRaw) > 0 {
			selfRegistrationKioskRef01Data = core.ToMapAny(selfRegistrationKioskRef01DataRaw[0][1])
		}
		// Discard guards against Go's unused-var check when the flow's steps
		// happen not to consume the bootstrap data (e.g. list-only flows).
		_ = selfRegistrationKioskRef01Data

		// LIST
		selfRegistrationKioskRef01Ent := client.SelfRegistrationKiosk(nil)
		selfRegistrationKioskRef01Match := map[string]any{}

		selfRegistrationKioskRef01ListResult, err := selfRegistrationKioskRef01Ent.List(selfRegistrationKioskRef01Match, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		_, selfRegistrationKioskRef01ListOk := selfRegistrationKioskRef01ListResult.([]any)
		if !selfRegistrationKioskRef01ListOk {
			t.Fatalf("expected list result to be an array, got %T", selfRegistrationKioskRef01ListResult)
		}

	})
}

func self_registration_kioskBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "self_registration_kiosk", "SelfRegistrationKioskTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read self_registration_kiosk test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse self_registration_kiosk test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap := vs.Transform(
		[]any{"self_registration_kiosk01", "self_registration_kiosk02", "self_registration_kiosk03"},
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
	entidEnvRaw := os.Getenv("IAMSMART_TEST_SELF_REGISTRATION_KIOSK_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"IAMSMART_TEST_SELF_REGISTRATION_KIOSK_ENTID": idmap,
		"IAMSMART_TEST_LIVE":      "FALSE",
		"IAMSMART_TEST_EXPLAIN":   "FALSE",
		"IAMSMART_APIKEY":         "NONE",
	})

	idmapResolved := core.ToMapAny(env["IAMSMART_TEST_SELF_REGISTRATION_KIOSK_ENTID"])
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
