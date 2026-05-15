# IamSmart SDK utility registration
require_relative '../core/utility_type'
require_relative 'clean'
require_relative 'done'
require_relative 'make_error'
require_relative 'feature_add'
require_relative 'feature_hook'
require_relative 'feature_init'
require_relative 'fetcher'
require_relative 'make_fetch_def'
require_relative 'make_context'
require_relative 'make_options'
require_relative 'make_request'
require_relative 'make_response'
require_relative 'make_result'
require_relative 'make_point'
require_relative 'make_spec'
require_relative 'make_url'
require_relative 'param'
require_relative 'prepare_auth'
require_relative 'prepare_body'
require_relative 'prepare_headers'
require_relative 'prepare_method'
require_relative 'prepare_params'
require_relative 'prepare_path'
require_relative 'prepare_query'
require_relative 'result_basic'
require_relative 'result_body'
require_relative 'result_headers'
require_relative 'transform_request'
require_relative 'transform_response'

IamSmartUtility.registrar = ->(u) {
  u.clean = IamSmartUtilities::Clean
  u.done = IamSmartUtilities::Done
  u.make_error = IamSmartUtilities::MakeError
  u.feature_add = IamSmartUtilities::FeatureAdd
  u.feature_hook = IamSmartUtilities::FeatureHook
  u.feature_init = IamSmartUtilities::FeatureInit
  u.fetcher = IamSmartUtilities::Fetcher
  u.make_fetch_def = IamSmartUtilities::MakeFetchDef
  u.make_context = IamSmartUtilities::MakeContext
  u.make_options = IamSmartUtilities::MakeOptions
  u.make_request = IamSmartUtilities::MakeRequest
  u.make_response = IamSmartUtilities::MakeResponse
  u.make_result = IamSmartUtilities::MakeResult
  u.make_point = IamSmartUtilities::MakePoint
  u.make_spec = IamSmartUtilities::MakeSpec
  u.make_url = IamSmartUtilities::MakeUrl
  u.param = IamSmartUtilities::Param
  u.prepare_auth = IamSmartUtilities::PrepareAuth
  u.prepare_body = IamSmartUtilities::PrepareBody
  u.prepare_headers = IamSmartUtilities::PrepareHeaders
  u.prepare_method = IamSmartUtilities::PrepareMethod
  u.prepare_params = IamSmartUtilities::PrepareParams
  u.prepare_path = IamSmartUtilities::PreparePath
  u.prepare_query = IamSmartUtilities::PrepareQuery
  u.result_basic = IamSmartUtilities::ResultBasic
  u.result_body = IamSmartUtilities::ResultBody
  u.result_headers = IamSmartUtilities::ResultHeaders
  u.transform_request = IamSmartUtilities::TransformRequest
  u.transform_response = IamSmartUtilities::TransformResponse
}
