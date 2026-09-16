# IamSmart SDK feature factory

require_relative 'feature/base_feature'
require_relative 'feature/ratelimit_feature'
require_relative 'feature/retry_feature'
require_relative 'feature/test_feature'
require_relative 'feature/timeout_feature'


module IamSmartFeatures
  def self.make_feature(name)
    case name
    when "base"
      IamSmartBaseFeature.new
    when "ratelimit"
      IamSmartRatelimitFeature.new
    when "retry"
      IamSmartRetryFeature.new
    when "test"
      IamSmartTestFeature.new
    when "timeout"
      IamSmartTimeoutFeature.new
    else
      IamSmartBaseFeature.new
    end
  end
end
