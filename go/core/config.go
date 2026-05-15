package core

func MakeConfig() map[string]any {
	return map[string]any{
		"main": map[string]any{
			"name": "IamSmart",
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
			"auth": map[string]any{
				"prefix": "Bearer",
			},
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
						"req": false,
						"type": "`$STRING`",
						"active": true,
						"index$": 0,
					},
					map[string]any{
						"name": "id",
						"req": false,
						"type": "`$STRING`",
						"active": true,
						"index$": 1,
					},
					map[string]any{
						"name": "latitude",
						"req": false,
						"type": "`$NUMBER`",
						"active": true,
						"index$": 2,
					},
					map[string]any{
						"name": "location",
						"req": false,
						"type": "`$STRING`",
						"active": true,
						"index$": 3,
					},
					map[string]any{
						"name": "location_en",
						"req": false,
						"type": "`$STRING`",
						"active": true,
						"index$": 4,
					},
					map[string]any{
						"name": "location_zh",
						"req": false,
						"type": "`$STRING`",
						"active": true,
						"index$": 5,
					},
					map[string]any{
						"name": "longitude",
						"req": false,
						"type": "`$NUMBER`",
						"active": true,
						"index$": 6,
					},
					map[string]any{
						"name": "name",
						"req": false,
						"type": "`$STRING`",
						"active": true,
						"index$": 7,
					},
					map[string]any{
						"name": "name_en",
						"req": false,
						"type": "`$STRING`",
						"active": true,
						"index$": 8,
					},
					map[string]any{
						"name": "name_zh",
						"req": false,
						"type": "`$STRING`",
						"active": true,
						"index$": 9,
					},
					map[string]any{
						"name": "region",
						"req": false,
						"type": "`$STRING`",
						"active": true,
						"index$": 10,
					},
					map[string]any{
						"name": "remark",
						"req": false,
						"type": "`$STRING`",
						"active": true,
						"index$": 11,
					},
					map[string]any{
						"name": "schedule",
						"req": false,
						"type": "`$ARRAY`",
						"active": true,
						"index$": 12,
					},
				},
				"name": "mobile_registration_point",
				"op": map[string]any{
					"list": map[string]any{
						"name": "list",
						"points": []any{
							map[string]any{
								"method": "GET",
								"orig": "/open_data/iam_smart/mobile-registration-points",
								"parts": []any{
									"open_data",
									"iam_smart",
									"mobile-registration-points",
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"active": true,
								"args": map[string]any{},
								"select": map[string]any{},
								"index$": 0,
							},
						},
						"input": "data",
						"key$": "list",
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
						"req": false,
						"type": "`$STRING`",
						"active": true,
						"index$": 0,
					},
					map[string]any{
						"name": "address_en",
						"req": false,
						"type": "`$STRING`",
						"active": true,
						"index$": 1,
					},
					map[string]any{
						"name": "address_zh",
						"req": false,
						"type": "`$STRING`",
						"active": true,
						"index$": 2,
					},
					map[string]any{
						"name": "district",
						"req": false,
						"type": "`$STRING`",
						"active": true,
						"index$": 3,
					},
					map[string]any{
						"name": "id",
						"req": false,
						"type": "`$STRING`",
						"active": true,
						"index$": 4,
					},
					map[string]any{
						"name": "latitude",
						"req": false,
						"type": "`$NUMBER`",
						"active": true,
						"index$": 5,
					},
					map[string]any{
						"name": "longitude",
						"req": false,
						"type": "`$NUMBER`",
						"active": true,
						"index$": 6,
					},
					map[string]any{
						"name": "name",
						"req": false,
						"type": "`$STRING`",
						"active": true,
						"index$": 7,
					},
					map[string]any{
						"name": "name_en",
						"req": false,
						"type": "`$STRING`",
						"active": true,
						"index$": 8,
					},
					map[string]any{
						"name": "name_zh",
						"req": false,
						"type": "`$STRING`",
						"active": true,
						"index$": 9,
					},
					map[string]any{
						"name": "operating_hour",
						"req": false,
						"type": "`$STRING`",
						"active": true,
						"index$": 10,
					},
					map[string]any{
						"name": "region",
						"req": false,
						"type": "`$STRING`",
						"active": true,
						"index$": 11,
					},
					map[string]any{
						"name": "remark",
						"req": false,
						"type": "`$STRING`",
						"active": true,
						"index$": 12,
					},
					map[string]any{
						"name": "service",
						"req": false,
						"type": "`$ARRAY`",
						"active": true,
						"index$": 13,
					},
					map[string]any{
						"name": "telephone",
						"req": false,
						"type": "`$STRING`",
						"active": true,
						"index$": 14,
					},
				},
				"name": "registration_service_counter",
				"op": map[string]any{
					"list": map[string]any{
						"name": "list",
						"points": []any{
							map[string]any{
								"method": "GET",
								"orig": "/open_data/iam_smart/registration-service-counters",
								"parts": []any{
									"open_data",
									"iam_smart",
									"registration-service-counters",
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"active": true,
								"args": map[string]any{},
								"select": map[string]any{},
								"index$": 0,
							},
						},
						"input": "data",
						"key$": "list",
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
						"req": false,
						"type": "`$STRING`",
						"active": true,
						"index$": 0,
					},
					map[string]any{
						"name": "address_en",
						"req": false,
						"type": "`$STRING`",
						"active": true,
						"index$": 1,
					},
					map[string]any{
						"name": "address_zh",
						"req": false,
						"type": "`$STRING`",
						"active": true,
						"index$": 2,
					},
					map[string]any{
						"name": "availability",
						"req": false,
						"type": "`$STRING`",
						"active": true,
						"index$": 3,
					},
					map[string]any{
						"name": "district",
						"req": false,
						"type": "`$STRING`",
						"active": true,
						"index$": 4,
					},
					map[string]any{
						"name": "floor",
						"req": false,
						"type": "`$STRING`",
						"active": true,
						"index$": 5,
					},
					map[string]any{
						"name": "id",
						"req": false,
						"type": "`$STRING`",
						"active": true,
						"index$": 6,
					},
					map[string]any{
						"name": "latitude",
						"req": false,
						"type": "`$NUMBER`",
						"active": true,
						"index$": 7,
					},
					map[string]any{
						"name": "longitude",
						"req": false,
						"type": "`$NUMBER`",
						"active": true,
						"index$": 8,
					},
					map[string]any{
						"name": "name",
						"req": false,
						"type": "`$STRING`",
						"active": true,
						"index$": 9,
					},
					map[string]any{
						"name": "name_en",
						"req": false,
						"type": "`$STRING`",
						"active": true,
						"index$": 10,
					},
					map[string]any{
						"name": "name_zh",
						"req": false,
						"type": "`$STRING`",
						"active": true,
						"index$": 11,
					},
					map[string]any{
						"name": "operating_hour",
						"req": false,
						"type": "`$STRING`",
						"active": true,
						"index$": 12,
					},
					map[string]any{
						"name": "region",
						"req": false,
						"type": "`$STRING`",
						"active": true,
						"index$": 13,
					},
					map[string]any{
						"name": "remark",
						"req": false,
						"type": "`$STRING`",
						"active": true,
						"index$": 14,
					},
				},
				"name": "self_registration_kiosk",
				"op": map[string]any{
					"list": map[string]any{
						"name": "list",
						"points": []any{
							map[string]any{
								"method": "GET",
								"orig": "/open_data/iam_smart/self-registration-kiosks",
								"parts": []any{
									"open_data",
									"iam_smart",
									"self-registration-kiosks",
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"active": true,
								"args": map[string]any{},
								"select": map[string]any{},
								"index$": 0,
							},
						},
						"input": "data",
						"key$": "list",
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
		},
	}
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
