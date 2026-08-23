
import { BaseFeature } from './feature/base/BaseFeature'
import { TestFeature } from './feature/test/TestFeature'



const FEATURE_CLASS: Record<string, typeof BaseFeature> = {
   test: TestFeature,

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
     test:     {
      "options": {
        "active": false
      }
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
              "parts": [
                "open_data",
                "iam_smart",
                "mobile-registration-points"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
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
          "name": "latitude",
          "short": "Latitude coordinate",
          "type": "`$NUMBER`"
        },
        {
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
              "parts": [
                "open_data",
                "iam_smart",
                "registration-service-counters"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
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
          "name": "latitude",
          "short": "Latitude coordinate",
          "type": "`$NUMBER`"
        },
        {
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
              "parts": [
                "open_data",
                "iam_smart",
                "self-registration-kiosks"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
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
  config
}

