-- IamSmart SDK error

local IamSmartError = {}
IamSmartError.__index = IamSmartError


function IamSmartError.new(code, msg, ctx)
  local self = setmetatable({}, IamSmartError)
  self.is_sdk_error = true
  self.sdk = "IamSmart"
  self.code = code or ""
  self.msg = msg or ""
  self.ctx = ctx
  self.result = nil
  self.spec = nil
  return self
end


function IamSmartError:error()
  return self.msg
end


function IamSmartError:__tostring()
  return self.msg
end


return IamSmartError
