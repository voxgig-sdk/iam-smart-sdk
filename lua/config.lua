-- IamSmart SDK configuration

-- Build a fresh, fully materialised config table. Every call rebuilds the
-- whole structure, so prefer require("config_shared") unless you need a
-- private copy you intend to mutate.
local function make_config()
  return {
    main = {
      name = "IamSmart",
    },
    feature = {
      ["test"] = {
        ["options"] = {
          ["active"] = false,
        },
      },
    },
    options = {
      base = "https://www.digitalpolicy.gov.hk",
      headers = {
        ["content-type"] = "application/json",
      },
      entity = {
        ["mobile_registration_point"] = {},
        ["registration_service_counter"] = {},
        ["self_registration_kiosk"] = {},
      },
    },
    entity = {
      ["mobile_registration_point"] = {
        ["fields"] = {
          {
            ["name"] = "district",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "id",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "latitude",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "location",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "locationEn",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "locationZh",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "longitude",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "name",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "nameEn",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "nameZh",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "region",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "remarks",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "schedule",
            ["type"] = "`$ARRAY`",
          },
        },
        ["name"] = "mobile_registration_point",
        ["op"] = {
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/open_data/iam_smart/mobile-registration-points",
                ["parts"] = {
                  "open_data",
                  "iam_smart",
                  "mobile-registration-points",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["registration_service_counter"] = {
        ["fields"] = {
          {
            ["name"] = "address",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "addressEn",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "addressZh",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "district",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "id",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "latitude",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "longitude",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "name",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "nameEn",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "nameZh",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "operatingHours",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "region",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "remarks",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "services",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "telephone",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "registration_service_counter",
        ["op"] = {
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/open_data/iam_smart/registration-service-counters",
                ["parts"] = {
                  "open_data",
                  "iam_smart",
                  "registration-service-counters",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["self_registration_kiosk"] = {
        ["fields"] = {
          {
            ["name"] = "address",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "addressEn",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "addressZh",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "availability",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "district",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "floor",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "id",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "latitude",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "longitude",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "name",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "nameEn",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "nameZh",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "operatingHours",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "region",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "remarks",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "self_registration_kiosk",
        ["op"] = {
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/open_data/iam_smart/self-registration-kiosks",
                ["parts"] = {
                  "open_data",
                  "iam_smart",
                  "self-registration-kiosks",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
    },
  }
end


local function make_feature(name)
  local features = require("features")
  local factory = features[name]
  if factory ~= nil then
    return factory()
  end
  return features.base()
end


-- Attach make_feature to the SDK class
local function setup_sdk(SDK)
  SDK._make_feature = make_feature
end


return make_config
