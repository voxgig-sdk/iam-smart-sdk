
import { BaseFeature } from './feature/base/BaseFeature'
import { RatelimitFeature } from './feature/ratelimit/RatelimitFeature'
import { RetryFeature } from './feature/retry/RetryFeature'
import { TestFeature } from './feature/test/TestFeature'
import { TimeoutFeature } from './feature/timeout/TimeoutFeature'



const FEATURE_CLASS: Record<string, typeof BaseFeature> = {
   ratelimit: RatelimitFeature,
 retry: RetryFeature,
 test: TestFeature,
 timeout: TimeoutFeature,

}


// Per-feature plugin DEFINITIONS (voxgig/plugin `Definition` values), from
// the model's active plugin groups. A feature that takes a `plugins` option
// (secrets over sekreto) reads its own entry; a feature with no plugins has
// none. Named imports above make each definition statically reachable, so
// an SDK carries exactly the plugin modules its model selects — the same
// leanness the old side-effect registry imports bought, without a registry.
const FEATURE_PLUGINS: Record<string, any[]> = {
  
}


class Config {

  makeFeature(this: any, fn: string) {
    const fc = FEATURE_CLASS[fn]
    const fi = new fc()
    // TODO: errors etc
    return fi
  }

  // False for a feature added at runtime via options.extend (station's
  // adopt path) - the constructor uses this to skip makeFeature for names
  // no generated class backs.
  hasFeature(this: any, fn: string) {
    return null != FEATURE_CLASS[fn]
  }


  main = {
    name: 'IamSmart',
        slug: "iam-smart",
    version: "0.0.1",
    target: "ts",

  }


  feature = {
     ratelimit:     {
      "options": {
        "active": false,
        "burst": 5,
        "rate": 5
      },
      "optspec": {
        "now": "`$FUNCTION`",
        "sleep": "`$FUNCTION`"
      },
      "strict": false,
      "transport": "wrap"
    },
 retry:     {
      "options": {
        "active": false,
        "factor": 2,
        "maxDelay": 2000,
        "minDelay": 50,
        "retries": 2,
        "statuses": [
          408,
          425,
          429,
          500,
          502,
          503,
          504
        ]
      },
      "optspec": {
        "jitter": "`$BOOLEAN`",
        "sleep": "`$FUNCTION`"
      },
      "strict": false,
      "transport": "wrap"
    },
 test:     {
      "options": {
        "active": false
      },
      "optspec": {
        "entity": "`$MAP`",
        "net": "`$MAP`"
      },
      "strict": false,
      "transport": "base"
    },
 timeout:     {
      "options": {
        "active": false,
        "ms": 30000
      },
      "optspec": {
        "clearTimer": "`$FUNCTION`",
        "setTimer": "`$FUNCTION`"
      },
      "strict": false,
      "transport": "wrap"
    },

  }


  options = {
    base: "https://www.digitalpolicy.gov.hk",

    headers: {
      "content-type": "application/json"
    },

    entity: {
      
        mobile_registration_point: {
        },
  
        registration_service_counter: {
        },
  
        self_registration_kiosk: {
        },
  
    }
  }


  entity = {
    "mobile_registration_point": {
      "fields": [
        {
          "name": "district",
          "short": "District where the mobile point operates",
          "type": "`$STRING`"
        },
        {
          "name": "id",
          "short": "Unique identifier for the mobile registration point",
          "type": "`$STRING`"
        },
        {
          "format": "double",
          "name": "latitude",
          "short": "Latitude coordinate",
          "type": "`$NUMBER`"
        },
        {
          "name": "location",
          "short": "Location description of the mobile point",
          "type": "`$STRING`"
        },
        {
          "name": "locationEn",
          "short": "English location description",
          "type": "`$STRING`"
        },
        {
          "name": "locationZh",
          "short": "Chinese location description",
          "type": "`$STRING`"
        },
        {
          "format": "double",
          "name": "longitude",
          "short": "Longitude coordinate",
          "type": "`$NUMBER`"
        },
        {
          "name": "name",
          "short": "Name of the mobile registration point location",
          "type": "`$STRING`"
        },
        {
          "name": "nameEn",
          "short": "English name of the mobile registration point",
          "type": "`$STRING`"
        },
        {
          "name": "nameZh",
          "short": "Chinese name of the mobile registration point",
          "type": "`$STRING`"
        },
        {
          "name": "region",
          "short": "Region (Hong Kong Island, Kowloon, New Territories)",
          "type": "`$STRING`"
        },
        {
          "name": "remarks",
          "short": "Additional remarks or notes",
          "type": "`$STRING`"
        },
        {
          "name": "schedule",
          "short": "Schedule of mobile registration point visits",
          "type": "`$ARRAY`"
        }
      ],
      "id": {
        "field": "id",
        "name": "id"
      },
      "name": "mobile_registration_point",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {},
              "kind": "http",
              "method": "GET",
              "orig": "/open_data/iam_smart/mobile-registration-points",
              "segments": [
                {
                  "lit": "open_data"
                },
                {
                  "lit": "iam_smart"
                },
                {
                  "lit": "mobile-registration-points"
                }
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "parts": [
                "open_data",
                "iam_smart",
                "mobile-registration-points"
              ]
            }
          ]
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "registration_service_counter": {
      "fields": [
        {
          "name": "address",
          "short": "Full address of the service counter",
          "type": "`$STRING`"
        },
        {
          "name": "addressEn",
          "short": "English address of the service counter",
          "type": "`$STRING`"
        },
        {
          "name": "addressZh",
          "short": "Chinese address of the service counter",
          "type": "`$STRING`"
        },
        {
          "name": "district",
          "short": "District where the service counter is located",
          "type": "`$STRING`"
        },
        {
          "name": "id",
          "short": "Unique identifier for the service counter",
          "type": "`$STRING`"
        },
        {
          "format": "double",
          "name": "latitude",
          "short": "Latitude coordinate",
          "type": "`$NUMBER`"
        },
        {
          "format": "double",
          "name": "longitude",
          "short": "Longitude coordinate",
          "type": "`$NUMBER`"
        },
        {
          "name": "name",
          "short": "Name of the service counter location",
          "type": "`$STRING`"
        },
        {
          "name": "nameEn",
          "short": "English name of the service counter location",
          "type": "`$STRING`"
        },
        {
          "name": "nameZh",
          "short": "Chinese name of the service counter location",
          "type": "`$STRING`"
        },
        {
          "name": "operatingHours",
          "short": "Operating hours of the service counter",
          "type": "`$STRING`"
        },
        {
          "name": "region",
          "short": "Region (Hong Kong Island, Kowloon, New Territories)",
          "type": "`$STRING`"
        },
        {
          "name": "remarks",
          "short": "Additional remarks or notes",
          "type": "`$STRING`"
        },
        {
          "name": "services",
          "short": "List of services available at this counter",
          "type": "`$ARRAY`"
        },
        {
          "name": "telephone",
          "short": "Contact telephone number",
          "type": "`$STRING`"
        }
      ],
      "id": {
        "field": "id",
        "name": "id"
      },
      "name": "registration_service_counter",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {},
              "kind": "http",
              "method": "GET",
              "orig": "/open_data/iam_smart/registration-service-counters",
              "segments": [
                {
                  "lit": "open_data"
                },
                {
                  "lit": "iam_smart"
                },
                {
                  "lit": "registration-service-counters"
                }
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "parts": [
                "open_data",
                "iam_smart",
                "registration-service-counters"
              ]
            }
          ]
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "self_registration_kiosk": {
      "fields": [
        {
          "name": "address",
          "short": "Full address of the kiosk",
          "type": "`$STRING`"
        },
        {
          "name": "addressEn",
          "short": "English address of the kiosk",
          "type": "`$STRING`"
        },
        {
          "name": "addressZh",
          "short": "Chinese address of the kiosk",
          "type": "`$STRING`"
        },
        {
          "name": "availability",
          "short": "Availability status of the kiosk",
          "type": "`$STRING`"
        },
        {
          "name": "district",
          "short": "District where the kiosk is located",
          "type": "`$STRING`"
        },
        {
          "name": "floor",
          "short": "Floor level where kiosk is located",
          "type": "`$STRING`"
        },
        {
          "name": "id",
          "short": "Unique identifier for the kiosk",
          "type": "`$STRING`"
        },
        {
          "format": "double",
          "name": "latitude",
          "short": "Latitude coordinate",
          "type": "`$NUMBER`"
        },
        {
          "format": "double",
          "name": "longitude",
          "short": "Longitude coordinate",
          "type": "`$NUMBER`"
        },
        {
          "name": "name",
          "short": "Name of the kiosk location",
          "type": "`$STRING`"
        },
        {
          "name": "nameEn",
          "short": "English name of the kiosk location",
          "type": "`$STRING`"
        },
        {
          "name": "nameZh",
          "short": "Chinese name of the kiosk location",
          "type": "`$STRING`"
        },
        {
          "name": "operatingHours",
          "short": "Operating hours of the kiosk",
          "type": "`$STRING`"
        },
        {
          "name": "region",
          "short": "Region (Hong Kong Island, Kowloon, New Territories)",
          "type": "`$STRING`"
        },
        {
          "name": "remarks",
          "short": "Additional remarks or notes",
          "type": "`$STRING`"
        }
      ],
      "id": {
        "field": "id",
        "name": "id"
      },
      "name": "self_registration_kiosk",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {},
              "kind": "http",
              "method": "GET",
              "orig": "/open_data/iam_smart/self-registration-kiosks",
              "segments": [
                {
                  "lit": "open_data"
                },
                {
                  "lit": "iam_smart"
                },
                {
                  "lit": "self-registration-kiosks"
                }
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "parts": [
                "open_data",
                "iam_smart",
                "self-registration-kiosks"
              ]
            }
          ]
        }
      },
      "relations": {
        "ancestors": []
      }
    }
  }
}


const config = new Config()

export {
  config,
  FEATURE_PLUGINS,
}

