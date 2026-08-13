// Typed models for the IamSmart SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.
package entity

import (
	"encoding/json"

	"github.com/voxgig-sdk/iam-smart-sdk/go/core"
)

// MobileRegistrationPoint is the typed data model for the mobile_registration_point entity.
type MobileRegistrationPoint struct {
	District *string `json:"district,omitempty"`
	Id *string `json:"id,omitempty"`
	Latitude *float64 `json:"latitude,omitempty"`
	Location *string `json:"location,omitempty"`
	LocationEn *string `json:"locationEn,omitempty"`
	LocationZh *string `json:"locationZh,omitempty"`
	Longitude *float64 `json:"longitude,omitempty"`
	Name *string `json:"name,omitempty"`
	NameEn *string `json:"nameEn,omitempty"`
	NameZh *string `json:"nameZh,omitempty"`
	Region *string `json:"region,omitempty"`
	Remarks *string `json:"remarks,omitempty"`
	Schedule *[]any `json:"schedule,omitempty"`
}

// MobileRegistrationPointListMatch is the typed request payload for MobileRegistrationPoint.ListTyped.
type MobileRegistrationPointListMatch struct {
	District *string `json:"district,omitempty"`
	Id *string `json:"id,omitempty"`
	Latitude *float64 `json:"latitude,omitempty"`
	Location *string `json:"location,omitempty"`
	LocationEn *string `json:"locationEn,omitempty"`
	LocationZh *string `json:"locationZh,omitempty"`
	Longitude *float64 `json:"longitude,omitempty"`
	Name *string `json:"name,omitempty"`
	NameEn *string `json:"nameEn,omitempty"`
	NameZh *string `json:"nameZh,omitempty"`
	Region *string `json:"region,omitempty"`
	Remarks *string `json:"remarks,omitempty"`
	Schedule *[]any `json:"schedule,omitempty"`
}

// RegistrationServiceCounter is the typed data model for the registration_service_counter entity.
type RegistrationServiceCounter struct {
	Address *string `json:"address,omitempty"`
	AddressEn *string `json:"addressEn,omitempty"`
	AddressZh *string `json:"addressZh,omitempty"`
	District *string `json:"district,omitempty"`
	Id *string `json:"id,omitempty"`
	Latitude *float64 `json:"latitude,omitempty"`
	Longitude *float64 `json:"longitude,omitempty"`
	Name *string `json:"name,omitempty"`
	NameEn *string `json:"nameEn,omitempty"`
	NameZh *string `json:"nameZh,omitempty"`
	OperatingHours *string `json:"operatingHours,omitempty"`
	Region *string `json:"region,omitempty"`
	Remarks *string `json:"remarks,omitempty"`
	Services *[]any `json:"services,omitempty"`
	Telephone *string `json:"telephone,omitempty"`
}

// RegistrationServiceCounterListMatch is the typed request payload for RegistrationServiceCounter.ListTyped.
type RegistrationServiceCounterListMatch struct {
	Address *string `json:"address,omitempty"`
	AddressEn *string `json:"addressEn,omitempty"`
	AddressZh *string `json:"addressZh,omitempty"`
	District *string `json:"district,omitempty"`
	Id *string `json:"id,omitempty"`
	Latitude *float64 `json:"latitude,omitempty"`
	Longitude *float64 `json:"longitude,omitempty"`
	Name *string `json:"name,omitempty"`
	NameEn *string `json:"nameEn,omitempty"`
	NameZh *string `json:"nameZh,omitempty"`
	OperatingHours *string `json:"operatingHours,omitempty"`
	Region *string `json:"region,omitempty"`
	Remarks *string `json:"remarks,omitempty"`
	Services *[]any `json:"services,omitempty"`
	Telephone *string `json:"telephone,omitempty"`
}

// SelfRegistrationKiosk is the typed data model for the self_registration_kiosk entity.
type SelfRegistrationKiosk struct {
	Address *string `json:"address,omitempty"`
	AddressEn *string `json:"addressEn,omitempty"`
	AddressZh *string `json:"addressZh,omitempty"`
	Availability *string `json:"availability,omitempty"`
	District *string `json:"district,omitempty"`
	Floor *string `json:"floor,omitempty"`
	Id *string `json:"id,omitempty"`
	Latitude *float64 `json:"latitude,omitempty"`
	Longitude *float64 `json:"longitude,omitempty"`
	Name *string `json:"name,omitempty"`
	NameEn *string `json:"nameEn,omitempty"`
	NameZh *string `json:"nameZh,omitempty"`
	OperatingHours *string `json:"operatingHours,omitempty"`
	Region *string `json:"region,omitempty"`
	Remarks *string `json:"remarks,omitempty"`
}

// SelfRegistrationKioskListMatch is the typed request payload for SelfRegistrationKiosk.ListTyped.
type SelfRegistrationKioskListMatch struct {
	Address *string `json:"address,omitempty"`
	AddressEn *string `json:"addressEn,omitempty"`
	AddressZh *string `json:"addressZh,omitempty"`
	Availability *string `json:"availability,omitempty"`
	District *string `json:"district,omitempty"`
	Floor *string `json:"floor,omitempty"`
	Id *string `json:"id,omitempty"`
	Latitude *float64 `json:"latitude,omitempty"`
	Longitude *float64 `json:"longitude,omitempty"`
	Name *string `json:"name,omitempty"`
	NameEn *string `json:"nameEn,omitempty"`
	NameZh *string `json:"nameZh,omitempty"`
	OperatingHours *string `json:"operatingHours,omitempty"`
	Region *string `json:"region,omitempty"`
	Remarks *string `json:"remarks,omitempty"`
}

// asMap turns a typed request/data struct into the map[string]any the
// runtime op pipeline consumes, honouring the json tags above.
func asMap(v any) map[string]any {
	out := map[string]any{}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}

// entityData unwraps an entity to its data map.
//
// Operations resolve to the ENTITY, not the raw data (see AGENTS.md), and an
// entity's fields are UNEXPORTED — marshalling one directly yields `{}`, so
// every typed accessor would silently hand back a zero-valued struct. The
// typed boundary therefore takes the data hop first.
func entityData(v any) any {
	if ent, ok := v.(core.Entity); ok {
		return ent.Data()
	}
	return v
}

// typedFrom decodes a runtime value (an entity, or the map[string]any the op
// pipeline produced) into a typed model T via a JSON round-trip. On any error
// it returns the zero value of T; the op's own (value, error) tuple carries
// the real error.
func typedFrom[T any](v any) T {
	var out T
	v = entityData(v)
	if v == nil {
		return out
	}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}

// typedSliceFrom decodes a runtime list value into a typed slice []T via a
// JSON round-trip, for list ops. `list` resolves to a slice of ENTITY
// instances, so each element takes the data hop.
func typedSliceFrom[T any](v any) []T {
	var out []T
	if v == nil {
		return out
	}
	if list, ok := v.([]any); ok {
		unwrapped := make([]any, 0, len(list))
		for _, item := range list {
			unwrapped = append(unwrapped, entityData(item))
		}
		v = unwrapped
	}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}
