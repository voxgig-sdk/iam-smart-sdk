<?php
declare(strict_types=1);

// IamSmart SDK utility: result_headers

class IamSmartResultHeaders
{
    public static function call(IamSmartContext $ctx): ?IamSmartResult
    {
        $response = $ctx->response;
        $result = $ctx->result;
        if ($result) {
            if ($response && is_array($response->headers)) {
                $result->headers = $response->headers;
            } else {
                $result->headers = [];
            }
        }
        return $result;
    }
}
