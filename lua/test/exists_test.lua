-- IamSmart SDK exists test

local sdk = require("iam-smart_sdk")

describe("IamSmartSDK", function()
  it("should create test SDK", function()
    local testsdk = sdk.test(nil, nil)
    assert.is_not_nil(testsdk)
  end)
end)
