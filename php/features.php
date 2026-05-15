<?php
declare(strict_types=1);

// IamSmart SDK feature factory

require_once __DIR__ . '/feature/BaseFeature.php';
require_once __DIR__ . '/feature/TestFeature.php';


class IamSmartFeatures
{
    public static function make_feature(string $name)
    {
        switch ($name) {
            case "base":
                return new IamSmartBaseFeature();
            case "test":
                return new IamSmartTestFeature();
            default:
                return new IamSmartBaseFeature();
        }
    }
}
