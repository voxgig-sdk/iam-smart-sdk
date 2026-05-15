# IamSmart SDK utility: make_context
require_relative '../core/context'
module IamSmartUtilities
  MakeContext = ->(ctxmap, basectx) {
    IamSmartContext.new(ctxmap, basectx)
  }
end
