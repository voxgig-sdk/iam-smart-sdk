# IamSmart SDK exists test

require "minitest/autorun"
require_relative "../IamSmart_sdk"

class ExistsTest < Minitest::Test
  def test_create_test_sdk
    testsdk = IamSmartSDK.test(nil, nil)
    assert !testsdk.nil?
  end
end
