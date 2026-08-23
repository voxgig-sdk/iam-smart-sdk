package core

import (
	"sync"
)

// MakeConfig builds a fresh, fully materialised config map. Every call
// rebuilds the whole structure, so prefer SharedConfig unless you need a
// private copy you intend to mutate.
func MakeConfig() map[string]any {
	return map[string]any{
		"main": map[string]any{
			"name": "IamSmart",
			"slug": "iam-smart",
			"version": "0.0.1",
			"target": "go",
		},
		"feature": map[string]any{
			"test": map[string]any{
				"options": map[string]any{
					"active": false,
				},
			},
		},
		"options": map[string]any{
			"base": "https://www.digitalpolicy.gov.hk",
			"headers": map[string]any{
				"content-type": "application/json",
			},
			"entity": map[string]any{
				"mobile_registration_point": map[string]any{},
				"registration_service_counter": map[string]any{},
				"self_registration_kiosk": map[string]any{},
			},
		},
		"entity": map[string]any{
			"mobile_registration_point": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "district",
						"short": "District where the mobile point operates",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"short": "Unique identifier for the mobile registration point",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "latitude",
						"short": "Latitude coordinate",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "location",
						"short": "Location description of the mobile point",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "locationEn",
						"short": "English location description",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "locationZh",
						"short": "Chinese location description",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "longitude",
						"short": "Longitude coordinate",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "name",
						"short": "Name of the mobile registration point location",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "nameEn",
						"short": "English name of the mobile registration point",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "nameZh",
						"short": "Chinese name of the mobile registration point",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "region",
						"short": "Region (Hong Kong Island, Kowloon, New Territories)",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "remarks",
						"short": "Additional remarks or notes",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "schedule",
						"short": "Schedule of mobile registration point visits",
						"type": "`$ARRAY`",
					},
				},
				"name": "mobile_registration_point",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/open_data/iam_smart/mobile-registration-points",
								"parts": []any{
									"open_data",
									"iam_smart",
									"mobile-registration-points",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"registration_service_counter": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "address",
						"short": "Full address of the service counter",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "addressEn",
						"short": "English address of the service counter",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "addressZh",
						"short": "Chinese address of the service counter",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "district",
						"short": "District where the service counter is located",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"short": "Unique identifier for the service counter",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "latitude",
						"short": "Latitude coordinate",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "longitude",
						"short": "Longitude coordinate",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "name",
						"short": "Name of the service counter location",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "nameEn",
						"short": "English name of the service counter location",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "nameZh",
						"short": "Chinese name of the service counter location",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "operatingHours",
						"short": "Operating hours of the service counter",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "region",
						"short": "Region (Hong Kong Island, Kowloon, New Territories)",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "remarks",
						"short": "Additional remarks or notes",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "services",
						"short": "List of services available at this counter",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "telephone",
						"short": "Contact telephone number",
						"type": "`$STRING`",
					},
				},
				"name": "registration_service_counter",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/open_data/iam_smart/registration-service-counters",
								"parts": []any{
									"open_data",
									"iam_smart",
									"registration-service-counters",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"self_registration_kiosk": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "address",
						"short": "Full address of the kiosk",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "addressEn",
						"short": "English address of the kiosk",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "addressZh",
						"short": "Chinese address of the kiosk",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "availability",
						"short": "Availability status of the kiosk",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "district",
						"short": "District where the kiosk is located",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "floor",
						"short": "Floor level where kiosk is located",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"short": "Unique identifier for the kiosk",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "latitude",
						"short": "Latitude coordinate",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "longitude",
						"short": "Longitude coordinate",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "name",
						"short": "Name of the kiosk location",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "nameEn",
						"short": "English name of the kiosk location",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "nameZh",
						"short": "Chinese name of the kiosk location",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "operatingHours",
						"short": "Operating hours of the kiosk",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "region",
						"short": "Region (Hong Kong Island, Kowloon, New Territories)",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "remarks",
						"short": "Additional remarks or notes",
						"type": "`$STRING`",
					},
				},
				"name": "self_registration_kiosk",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/open_data/iam_smart/self-registration-kiosks",
								"parts": []any{
									"open_data",
									"iam_smart",
									"self-registration-kiosks",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
		},
	}
}

var (
	sharedConfigOnce sync.Once
	sharedConfigVal  map[string]any
)

// SharedConfig returns the process-wide config, built once on first use.
// The SDK reads the config on every request and never writes to it, so one
// instance is shared by every client rather than rebuilt per client.
//
// The returned map is shared: treat it as read-only. Callers that need to
// mutate should use MakeConfig, which always returns a fresh copy.
func SharedConfig() map[string]any {
	sharedConfigOnce.Do(func() {
		sharedConfigVal = MakeConfig()
	})
	return sharedConfigVal
}

func makeFeature(name string) Feature {
	switch name {
	case "test":
		if NewTestFeatureFunc != nil {
			return NewTestFeatureFunc()
		}
	default:
		if NewBaseFeatureFunc != nil {
			return NewBaseFeatureFunc()
		}
	}
	return nil
}
