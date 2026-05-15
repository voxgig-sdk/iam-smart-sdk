# IamSmart SDK feature factory

require_relative 'feature/base_feature'
require_relative 'feature/test_feature'


module IamSmartFeatures
  def self.make_feature(name)
    case name
    when "base"
      IamSmartBaseFeature.new
    when "test"
      IamSmartTestFeature.new
    else
      IamSmartBaseFeature.new
    end
  end
end
