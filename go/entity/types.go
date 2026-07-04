// Typed models for the IamSmart SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.
package entity

import "encoding/json"

// MobileRegistrationPoint is the typed data model for the mobile_registration_point entity.
type MobileRegistrationPoint struct {
	District *string `json:"district,omitempty"`
	Id *string `json:"id,omitempty"`
	Latitude *float64 `json:"latitude,omitempty"`
	Location *string `json:"location,omitempty"`
	LocationEn *string `json:"location_en,omitempty"`
	LocationZh *string `json:"location_zh,omitempty"`
	Longitude *float64 `json:"longitude,omitempty"`
	Name *string `json:"name,omitempty"`
	NameEn *string `json:"name_en,omitempty"`
	NameZh *string `json:"name_zh,omitempty"`
	Region *string `json:"region,omitempty"`
	Remark *string `json:"remark,omitempty"`
	Schedule *[]any `json:"schedule,omitempty"`
}

// MobileRegistrationPointListMatch mirrors the mobile_registration_point fields as an all-optional match
// filter (Go analog of Partial<MobileRegistrationPoint>).
type MobileRegistrationPointListMatch struct {
	District *string `json:"district,omitempty"`
	Id *string `json:"id,omitempty"`
	Latitude *float64 `json:"latitude,omitempty"`
	Location *string `json:"location,omitempty"`
	LocationEn *string `json:"location_en,omitempty"`
	LocationZh *string `json:"location_zh,omitempty"`
	Longitude *float64 `json:"longitude,omitempty"`
	Name *string `json:"name,omitempty"`
	NameEn *string `json:"name_en,omitempty"`
	NameZh *string `json:"name_zh,omitempty"`
	Region *string `json:"region,omitempty"`
	Remark *string `json:"remark,omitempty"`
	Schedule *[]any `json:"schedule,omitempty"`
}

// RegistrationServiceCounter is the typed data model for the registration_service_counter entity.
type RegistrationServiceCounter struct {
	Address *string `json:"address,omitempty"`
	AddressEn *string `json:"address_en,omitempty"`
	AddressZh *string `json:"address_zh,omitempty"`
	District *string `json:"district,omitempty"`
	Id *string `json:"id,omitempty"`
	Latitude *float64 `json:"latitude,omitempty"`
	Longitude *float64 `json:"longitude,omitempty"`
	Name *string `json:"name,omitempty"`
	NameEn *string `json:"name_en,omitempty"`
	NameZh *string `json:"name_zh,omitempty"`
	OperatingHour *string `json:"operating_hour,omitempty"`
	Region *string `json:"region,omitempty"`
	Remark *string `json:"remark,omitempty"`
	Service *[]any `json:"service,omitempty"`
	Telephone *string `json:"telephone,omitempty"`
}

// RegistrationServiceCounterListMatch mirrors the registration_service_counter fields as an all-optional match
// filter (Go analog of Partial<RegistrationServiceCounter>).
type RegistrationServiceCounterListMatch struct {
	Address *string `json:"address,omitempty"`
	AddressEn *string `json:"address_en,omitempty"`
	AddressZh *string `json:"address_zh,omitempty"`
	District *string `json:"district,omitempty"`
	Id *string `json:"id,omitempty"`
	Latitude *float64 `json:"latitude,omitempty"`
	Longitude *float64 `json:"longitude,omitempty"`
	Name *string `json:"name,omitempty"`
	NameEn *string `json:"name_en,omitempty"`
	NameZh *string `json:"name_zh,omitempty"`
	OperatingHour *string `json:"operating_hour,omitempty"`
	Region *string `json:"region,omitempty"`
	Remark *string `json:"remark,omitempty"`
	Service *[]any `json:"service,omitempty"`
	Telephone *string `json:"telephone,omitempty"`
}

// SelfRegistrationKiosk is the typed data model for the self_registration_kiosk entity.
type SelfRegistrationKiosk struct {
	Address *string `json:"address,omitempty"`
	AddressEn *string `json:"address_en,omitempty"`
	AddressZh *string `json:"address_zh,omitempty"`
	Availability *string `json:"availability,omitempty"`
	District *string `json:"district,omitempty"`
	Floor *string `json:"floor,omitempty"`
	Id *string `json:"id,omitempty"`
	Latitude *float64 `json:"latitude,omitempty"`
	Longitude *float64 `json:"longitude,omitempty"`
	Name *string `json:"name,omitempty"`
	NameEn *string `json:"name_en,omitempty"`
	NameZh *string `json:"name_zh,omitempty"`
	OperatingHour *string `json:"operating_hour,omitempty"`
	Region *string `json:"region,omitempty"`
	Remark *string `json:"remark,omitempty"`
}

// SelfRegistrationKioskListMatch mirrors the self_registration_kiosk fields as an all-optional match
// filter (Go analog of Partial<SelfRegistrationKiosk>).
type SelfRegistrationKioskListMatch struct {
	Address *string `json:"address,omitempty"`
	AddressEn *string `json:"address_en,omitempty"`
	AddressZh *string `json:"address_zh,omitempty"`
	Availability *string `json:"availability,omitempty"`
	District *string `json:"district,omitempty"`
	Floor *string `json:"floor,omitempty"`
	Id *string `json:"id,omitempty"`
	Latitude *float64 `json:"latitude,omitempty"`
	Longitude *float64 `json:"longitude,omitempty"`
	Name *string `json:"name,omitempty"`
	NameEn *string `json:"name_en,omitempty"`
	NameZh *string `json:"name_zh,omitempty"`
	OperatingHour *string `json:"operating_hour,omitempty"`
	Region *string `json:"region,omitempty"`
	Remark *string `json:"remark,omitempty"`
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

// typedFrom decodes a runtime value (a map[string]any produced by the op
// pipeline) into a typed model T via a JSON round-trip. On any error it
// returns the zero value of T; the op's own (value, error) tuple carries the
// real error.
func typedFrom[T any](v any) T {
	var out T
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

// typedSliceFrom decodes a runtime list value ([]any of maps) into a typed
// slice []T via a JSON round-trip, for list ops.
func typedSliceFrom[T any](v any) []T {
	var out []T
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
