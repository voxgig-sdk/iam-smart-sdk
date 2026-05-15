<?php
declare(strict_types=1);

// IamSmart SDK utility: prepare_body

class IamSmartPrepareBody
{
    public static function call(IamSmartContext $ctx): mixed
    {
        if ($ctx->op->input === 'data') {
            return ($ctx->utility->transform_request)($ctx);
        }
        return null;
    }
}
