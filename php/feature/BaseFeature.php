<?php
declare(strict_types=1);

// IamSmart SDK base feature

class IamSmartBaseFeature
{
    public string $version;
    public string $name;
    public bool $active;

    public function __construct()
    {
        $this->version = '0.0.1';
        $this->name = 'base';
        $this->active = true;
    }

    public function get_version(): string { return $this->version; }
    public function get_name(): string { return $this->name; }
    public function get_active(): bool { return $this->active; }

    public function init(IamSmartContext $ctx, array $options): void {}
    public function PostConstruct(IamSmartContext $ctx): void {}
    public function PostConstructEntity(IamSmartContext $ctx): void {}
    public function SetData(IamSmartContext $ctx): void {}
    public function GetData(IamSmartContext $ctx): void {}
    public function GetMatch(IamSmartContext $ctx): void {}
    public function SetMatch(IamSmartContext $ctx): void {}
    public function PrePoint(IamSmartContext $ctx): void {}
    public function PreSpec(IamSmartContext $ctx): void {}
    public function PreRequest(IamSmartContext $ctx): void {}
    public function PreResponse(IamSmartContext $ctx): void {}
    public function PreResult(IamSmartContext $ctx): void {}
    public function PreDone(IamSmartContext $ctx): void {}
    public function PreUnexpected(IamSmartContext $ctx): void {}
}
