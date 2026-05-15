<?php
declare(strict_types=1);

// IamSmart SDK utility: make_context

require_once __DIR__ . '/../core/Context.php';

class IamSmartMakeContext
{
    public static function call(array $ctxmap, ?IamSmartContext $basectx): IamSmartContext
    {
        return new IamSmartContext($ctxmap, $basectx);
    }
}
