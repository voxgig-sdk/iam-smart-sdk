<?php
declare(strict_types=1);

// IamSmart SDK utility: feature_add

class IamSmartFeatureAdd
{
    public static function call(IamSmartContext $ctx, mixed $f): void
    {
        $ctx->client->features[] = $f;
    }
}
