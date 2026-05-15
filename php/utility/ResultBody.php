<?php
declare(strict_types=1);

// IamSmart SDK utility: result_body

class IamSmartResultBody
{
    public static function call(IamSmartContext $ctx): ?IamSmartResult
    {
        $response = $ctx->response;
        $result = $ctx->result;
        if ($result && $response && $response->json_func && $response->body) {
            $result->body = ($response->json_func)();
        }
        return $result;
    }
}
