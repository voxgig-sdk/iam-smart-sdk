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
			"ratelimit": map[string]any{
				"options": map[string]any{
					"active": false,
					"burst": 5,
					"rate": 5,
				},
				"optspec": map[string]any{
					"now": "`$FUNCTION`",
					"sleep": "`$FUNCTION`",
				},
				"strict": false,
				"transport": "wrap",
			},
			"retry": map[string]any{
				"options": map[string]any{
					"active": false,
					"factor": 2,
					"maxDelay": 2000,
					"minDelay": 50,
					"retries": 2,
					"statuses": []any{
						408,
						425,
						429,
						500,
						502,
						503,
						504,
					},
				},
				"optspec": map[string]any{
					"jitter": "`$BOOLEAN`",
					"sleep": "`$FUNCTION`",
				},
				"strict": false,
				"transport": "wrap",
			},
			"test": map[string]any{
				"options": map[string]any{
					"active": false,
				},
				"optspec": map[string]any{
					"entity": "`$MAP`",
					"net": "`$MAP`",
				},
				"strict": false,
				"transport": "base",
			},
			"timeout": map[string]any{
				"options": map[string]any{
					"active": false,
					"ms": 30000,
				},
				"optspec": map[string]any{
					"clearTimer": "`$FUNCTION`",
					"setTimer": "`$FUNCTION`",
				},
				"strict": false,
				"transport": "wrap",
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
						"format": "double",
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
						"format": "double",
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
				"id": map[string]any{
					"field": "id",
					"name": "id",
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
								"segments": []any{
									map[string]any{
										"lit": "open_data",
									},
									map[string]any{
										"lit": "iam_smart",
									},
									map[string]any{
										"lit": "mobile-registration-points",
									},
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"open_data",
									"iam_smart",
									"mobile-registration-points",
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
						"format": "double",
						"name": "latitude",
						"short": "Latitude coordinate",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"format": "double",
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
				"id": map[string]any{
					"field": "id",
					"name": "id",
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
								"segments": []any{
									map[string]any{
										"lit": "open_data",
									},
									map[string]any{
										"lit": "iam_smart",
									},
									map[string]any{
										"lit": "registration-service-counters",
									},
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"open_data",
									"iam_smart",
									"registration-service-counters",
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
						"format": "double",
						"name": "latitude",
						"short": "Latitude coordinate",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"format": "double",
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
				"id": map[string]any{
					"field": "id",
					"name": "id",
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
								"segments": []any{
									map[string]any{
										"lit": "open_data",
									},
									map[string]any{
										"lit": "iam_smart",
									},
									map[string]any{
										"lit": "self-registration-kiosks",
									},
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"open_data",
									"iam_smart",
									"self-registration-kiosks",
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

// The plugin definitions the model selected per feature, as []any so a
// feature package can consume them without core naming its types. Empty
// when no active feature declares active plugin groups for this target.
var featurePlugins = map[string][]any{
}

// FeaturePlugins is the definitions list for one feature's chain.
func FeaturePlugins(name string) []any {
	return featurePlugins[name]
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
	case "ratelimit":
		if NewRatelimitFeatureFunc != nil {
			return NewRatelimitFeatureFunc()
		}
	case "retry":
		if NewRetryFeatureFunc != nil {
			return NewRetryFeatureFunc()
		}
	case "test":
		if NewTestFeatureFunc != nil {
			return NewTestFeatureFunc()
		}
	case "timeout":
		if NewTimeoutFeatureFunc != nil {
			return NewTimeoutFeatureFunc()
		}
	default:
		if NewBaseFeatureFunc != nil {
			return NewBaseFeatureFunc()
		}
	}
	return nil
}
