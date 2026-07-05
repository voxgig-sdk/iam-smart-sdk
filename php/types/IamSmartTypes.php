<?php
declare(strict_types=1);

// Typed models for the IamSmart SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.
//
// These are documentation-grade value objects (PHP 8 typed properties),
// registered on the composer classmap autoload. The SDK boundary exchanges
// assoc-arrays; these classes name the shapes for tooling and typed callers.

/** MobileRegistrationPoint entity data model. */
class MobileRegistrationPoint
{
    public ?string $district = null;
    public ?string $id = null;
    public ?float $latitude = null;
    public ?string $location = null;
    public ?string $location_en = null;
    public ?string $location_zh = null;
    public ?float $longitude = null;
    public ?string $name = null;
    public ?string $name_en = null;
    public ?string $name_zh = null;
    public ?string $region = null;
    public ?string $remark = null;
    public ?array $schedule = null;
}

/** Request payload for MobileRegistrationPoint#list. */
class MobileRegistrationPointListMatch
{
    public ?string $district = null;
    public ?string $id = null;
    public ?float $latitude = null;
    public ?string $location = null;
    public ?string $location_en = null;
    public ?string $location_zh = null;
    public ?float $longitude = null;
    public ?string $name = null;
    public ?string $name_en = null;
    public ?string $name_zh = null;
    public ?string $region = null;
    public ?string $remark = null;
    public ?array $schedule = null;
}

/** RegistrationServiceCounter entity data model. */
class RegistrationServiceCounter
{
    public ?string $address = null;
    public ?string $address_en = null;
    public ?string $address_zh = null;
    public ?string $district = null;
    public ?string $id = null;
    public ?float $latitude = null;
    public ?float $longitude = null;
    public ?string $name = null;
    public ?string $name_en = null;
    public ?string $name_zh = null;
    public ?string $operating_hour = null;
    public ?string $region = null;
    public ?string $remark = null;
    public ?array $service = null;
    public ?string $telephone = null;
}

/** Request payload for RegistrationServiceCounter#list. */
class RegistrationServiceCounterListMatch
{
    public ?string $address = null;
    public ?string $address_en = null;
    public ?string $address_zh = null;
    public ?string $district = null;
    public ?string $id = null;
    public ?float $latitude = null;
    public ?float $longitude = null;
    public ?string $name = null;
    public ?string $name_en = null;
    public ?string $name_zh = null;
    public ?string $operating_hour = null;
    public ?string $region = null;
    public ?string $remark = null;
    public ?array $service = null;
    public ?string $telephone = null;
}

/** SelfRegistrationKiosk entity data model. */
class SelfRegistrationKiosk
{
    public ?string $address = null;
    public ?string $address_en = null;
    public ?string $address_zh = null;
    public ?string $availability = null;
    public ?string $district = null;
    public ?string $floor = null;
    public ?string $id = null;
    public ?float $latitude = null;
    public ?float $longitude = null;
    public ?string $name = null;
    public ?string $name_en = null;
    public ?string $name_zh = null;
    public ?string $operating_hour = null;
    public ?string $region = null;
    public ?string $remark = null;
}

/** Request payload for SelfRegistrationKiosk#list. */
class SelfRegistrationKioskListMatch
{
    public ?string $address = null;
    public ?string $address_en = null;
    public ?string $address_zh = null;
    public ?string $availability = null;
    public ?string $district = null;
    public ?string $floor = null;
    public ?string $id = null;
    public ?float $latitude = null;
    public ?float $longitude = null;
    public ?string $name = null;
    public ?string $name_en = null;
    public ?string $name_zh = null;
    public ?string $operating_hour = null;
    public ?string $region = null;
    public ?string $remark = null;
}

