<?php
declare(strict_types=1);

// IamSmart SDK utility registration

require_once __DIR__ . '/../core/UtilityType.php';
require_once __DIR__ . '/Clean.php';
require_once __DIR__ . '/Done.php';
require_once __DIR__ . '/MakeError.php';
require_once __DIR__ . '/FeatureAdd.php';
require_once __DIR__ . '/FeatureHook.php';
require_once __DIR__ . '/FeatureInit.php';
require_once __DIR__ . '/Fetcher.php';
require_once __DIR__ . '/MakeFetchDef.php';
require_once __DIR__ . '/MakeContext.php';
require_once __DIR__ . '/MakeOptions.php';
require_once __DIR__ . '/MakeRequest.php';
require_once __DIR__ . '/MakeResponse.php';
require_once __DIR__ . '/MakeResult.php';
require_once __DIR__ . '/MakePoint.php';
require_once __DIR__ . '/MakeSpec.php';
require_once __DIR__ . '/MakeUrl.php';
require_once __DIR__ . '/Param.php';
require_once __DIR__ . '/PrepareAuth.php';
require_once __DIR__ . '/PrepareBody.php';
require_once __DIR__ . '/PrepareHeaders.php';
require_once __DIR__ . '/PrepareMethod.php';
require_once __DIR__ . '/PrepareParams.php';
require_once __DIR__ . '/PreparePath.php';
require_once __DIR__ . '/PrepareQuery.php';
require_once __DIR__ . '/ResultBasic.php';
require_once __DIR__ . '/ResultBody.php';
require_once __DIR__ . '/ResultHeaders.php';
require_once __DIR__ . '/TransformRequest.php';
require_once __DIR__ . '/TransformResponse.php';

IamSmartUtility::setRegistrar(function (IamSmartUtility $u): void {
    $u->clean = [IamSmartClean::class, 'call'];
    $u->done = [IamSmartDone::class, 'call'];
    $u->make_error = [IamSmartMakeError::class, 'call'];
    $u->feature_add = [IamSmartFeatureAdd::class, 'call'];
    $u->feature_hook = [IamSmartFeatureHook::class, 'call'];
    $u->feature_init = [IamSmartFeatureInit::class, 'call'];
    $u->fetcher = [IamSmartFetcher::class, 'call'];
    $u->make_fetch_def = [IamSmartMakeFetchDef::class, 'call'];
    $u->make_context = [IamSmartMakeContext::class, 'call'];
    $u->make_options = [IamSmartMakeOptions::class, 'call'];
    $u->make_request = [IamSmartMakeRequest::class, 'call'];
    $u->make_response = [IamSmartMakeResponse::class, 'call'];
    $u->make_result = [IamSmartMakeResult::class, 'call'];
    $u->make_point = [IamSmartMakePoint::class, 'call'];
    $u->make_spec = [IamSmartMakeSpec::class, 'call'];
    $u->make_url = [IamSmartMakeUrl::class, 'call'];
    $u->param = [IamSmartParam::class, 'call'];
    $u->prepare_auth = [IamSmartPrepareAuth::class, 'call'];
    $u->prepare_body = [IamSmartPrepareBody::class, 'call'];
    $u->prepare_headers = [IamSmartPrepareHeaders::class, 'call'];
    $u->prepare_method = [IamSmartPrepareMethod::class, 'call'];
    $u->prepare_params = [IamSmartPrepareParams::class, 'call'];
    $u->prepare_path = [IamSmartPreparePath::class, 'call'];
    $u->prepare_query = [IamSmartPrepareQuery::class, 'call'];
    $u->result_basic = [IamSmartResultBasic::class, 'call'];
    $u->result_body = [IamSmartResultBody::class, 'call'];
    $u->result_headers = [IamSmartResultHeaders::class, 'call'];
    $u->transform_request = [IamSmartTransformRequest::class, 'call'];
    $u->transform_response = [IamSmartTransformResponse::class, 'call'];
});
