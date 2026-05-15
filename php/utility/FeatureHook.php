<?php
declare(strict_types=1);

// IamSmart SDK utility: feature_hook

class IamSmartFeatureHook
{
    public static function call(IamSmartContext $ctx, string $name): void
    {
        if (!$ctx->client) {
            return;
        }
        $features = $ctx->client->features ?? null;
        if (!$features) {
            return;
        }
        foreach ($features as $f) {
            if (method_exists($f, $name)) {
                $f->$name($ctx);
            }
        }
    }
}
