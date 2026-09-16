# IamSmart SDK configuration


# The sekreto plugin DEFINITIONS the model selected per feature, imported
# above by name from the modules the catalogue's active `plugin.def`
# entries declare. Handed to each feature (secrets builds its Sekreto
# with them): a provider kind not listed here is unknown to that SDK.
FEATURE_PLUGINS = {
}


_shared_config = None


def shared_config():
    """Return the process-wide config, built once on first use.

    The SDK reads the config on every request and never writes to it, so one
    instance is shared by every client rather than rebuilt per client.

    The returned dict is shared: treat it as read-only. Callers that need to
    mutate should use make_config, which always returns a fresh copy.
    """
    global _shared_config
    if _shared_config is None:
        _shared_config = make_config()
    return _shared_config


def make_config():
    """Build a fresh, fully materialised config dict.

    Every call rebuilds the whole structure, so prefer shared_config unless
    you need a private copy you intend to mutate.
    """
    return {
        "main": {
            "name": "IamSmart",
            "slug": "iam-smart",
            "version": "0.0.1",
            "target": "py",
        },
        "feature": {
            "ratelimit": {
        "options": {
          "active": False,
          "burst": 5,
          "rate": 5,
        },
        "optspec": {
          "now": "`$FUNCTION`",
          "sleep": "`$FUNCTION`",
        },
        "strict": False,
        "transport": "wrap",
      },
            "retry": {
        "options": {
          "active": False,
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
            504,
          ],
        },
        "optspec": {
          "jitter": "`$BOOLEAN`",
          "sleep": "`$FUNCTION`",
        },
        "strict": False,
        "transport": "wrap",
      },
            "test": {
        "options": {
          "active": False,
        },
        "optspec": {
          "entity": "`$MAP`",
          "net": "`$MAP`",
        },
        "strict": False,
        "transport": "base",
      },
            "timeout": {
        "options": {
          "active": False,
          "ms": 30000,
        },
        "optspec": {
          "clearTimer": "`$FUNCTION`",
          "setTimer": "`$FUNCTION`",
        },
        "strict": False,
        "transport": "wrap",
      },
        },
        "options": {
            "base": "https://www.digitalpolicy.gov.hk",
            "headers": {
        "content-type": "application/json",
      },
            "entity": {
                "mobile_registration_point": {},
                "registration_service_counter": {},
                "self_registration_kiosk": {},
            },
        },
        "entity": {
      "mobile_registration_point": {
        "fields": [
          {
            "name": "district",
            "short": "District where the mobile point operates",
            "type": "`$STRING`",
          },
          {
            "name": "id",
            "short": "Unique identifier for the mobile registration point",
            "type": "`$STRING`",
          },
          {
            "format": "double",
            "name": "latitude",
            "short": "Latitude coordinate",
            "type": "`$NUMBER`",
          },
          {
            "name": "location",
            "short": "Location description of the mobile point",
            "type": "`$STRING`",
          },
          {
            "name": "locationEn",
            "short": "English location description",
            "type": "`$STRING`",
          },
          {
            "name": "locationZh",
            "short": "Chinese location description",
            "type": "`$STRING`",
          },
          {
            "format": "double",
            "name": "longitude",
            "short": "Longitude coordinate",
            "type": "`$NUMBER`",
          },
          {
            "name": "name",
            "short": "Name of the mobile registration point location",
            "type": "`$STRING`",
          },
          {
            "name": "nameEn",
            "short": "English name of the mobile registration point",
            "type": "`$STRING`",
          },
          {
            "name": "nameZh",
            "short": "Chinese name of the mobile registration point",
            "type": "`$STRING`",
          },
          {
            "name": "region",
            "short": "Region (Hong Kong Island, Kowloon, New Territories)",
            "type": "`$STRING`",
          },
          {
            "name": "remarks",
            "short": "Additional remarks or notes",
            "type": "`$STRING`",
          },
          {
            "name": "schedule",
            "short": "Schedule of mobile registration point visits",
            "type": "`$ARRAY`",
          },
        ],
        "id": {
          "field": "id",
          "name": "id",
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
                    "lit": "open_data",
                  },
                  {
                    "lit": "iam_smart",
                  },
                  {
                    "lit": "mobile-registration-points",
                  },
                ],
                "select": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "parts": [
                  "open_data",
                  "iam_smart",
                  "mobile-registration-points",
                ],
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "registration_service_counter": {
        "fields": [
          {
            "name": "address",
            "short": "Full address of the service counter",
            "type": "`$STRING`",
          },
          {
            "name": "addressEn",
            "short": "English address of the service counter",
            "type": "`$STRING`",
          },
          {
            "name": "addressZh",
            "short": "Chinese address of the service counter",
            "type": "`$STRING`",
          },
          {
            "name": "district",
            "short": "District where the service counter is located",
            "type": "`$STRING`",
          },
          {
            "name": "id",
            "short": "Unique identifier for the service counter",
            "type": "`$STRING`",
          },
          {
            "format": "double",
            "name": "latitude",
            "short": "Latitude coordinate",
            "type": "`$NUMBER`",
          },
          {
            "format": "double",
            "name": "longitude",
            "short": "Longitude coordinate",
            "type": "`$NUMBER`",
          },
          {
            "name": "name",
            "short": "Name of the service counter location",
            "type": "`$STRING`",
          },
          {
            "name": "nameEn",
            "short": "English name of the service counter location",
            "type": "`$STRING`",
          },
          {
            "name": "nameZh",
            "short": "Chinese name of the service counter location",
            "type": "`$STRING`",
          },
          {
            "name": "operatingHours",
            "short": "Operating hours of the service counter",
            "type": "`$STRING`",
          },
          {
            "name": "region",
            "short": "Region (Hong Kong Island, Kowloon, New Territories)",
            "type": "`$STRING`",
          },
          {
            "name": "remarks",
            "short": "Additional remarks or notes",
            "type": "`$STRING`",
          },
          {
            "name": "services",
            "short": "List of services available at this counter",
            "type": "`$ARRAY`",
          },
          {
            "name": "telephone",
            "short": "Contact telephone number",
            "type": "`$STRING`",
          },
        ],
        "id": {
          "field": "id",
          "name": "id",
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
                    "lit": "open_data",
                  },
                  {
                    "lit": "iam_smart",
                  },
                  {
                    "lit": "registration-service-counters",
                  },
                ],
                "select": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "parts": [
                  "open_data",
                  "iam_smart",
                  "registration-service-counters",
                ],
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "self_registration_kiosk": {
        "fields": [
          {
            "name": "address",
            "short": "Full address of the kiosk",
            "type": "`$STRING`",
          },
          {
            "name": "addressEn",
            "short": "English address of the kiosk",
            "type": "`$STRING`",
          },
          {
            "name": "addressZh",
            "short": "Chinese address of the kiosk",
            "type": "`$STRING`",
          },
          {
            "name": "availability",
            "short": "Availability status of the kiosk",
            "type": "`$STRING`",
          },
          {
            "name": "district",
            "short": "District where the kiosk is located",
            "type": "`$STRING`",
          },
          {
            "name": "floor",
            "short": "Floor level where kiosk is located",
            "type": "`$STRING`",
          },
          {
            "name": "id",
            "short": "Unique identifier for the kiosk",
            "type": "`$STRING`",
          },
          {
            "format": "double",
            "name": "latitude",
            "short": "Latitude coordinate",
            "type": "`$NUMBER`",
          },
          {
            "format": "double",
            "name": "longitude",
            "short": "Longitude coordinate",
            "type": "`$NUMBER`",
          },
          {
            "name": "name",
            "short": "Name of the kiosk location",
            "type": "`$STRING`",
          },
          {
            "name": "nameEn",
            "short": "English name of the kiosk location",
            "type": "`$STRING`",
          },
          {
            "name": "nameZh",
            "short": "Chinese name of the kiosk location",
            "type": "`$STRING`",
          },
          {
            "name": "operatingHours",
            "short": "Operating hours of the kiosk",
            "type": "`$STRING`",
          },
          {
            "name": "region",
            "short": "Region (Hong Kong Island, Kowloon, New Territories)",
            "type": "`$STRING`",
          },
          {
            "name": "remarks",
            "short": "Additional remarks or notes",
            "type": "`$STRING`",
          },
        ],
        "id": {
          "field": "id",
          "name": "id",
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
                    "lit": "open_data",
                  },
                  {
                    "lit": "iam_smart",
                  },
                  {
                    "lit": "self-registration-kiosks",
                  },
                ],
                "select": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "parts": [
                  "open_data",
                  "iam_smart",
                  "self-registration-kiosks",
                ],
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
    },
    }
