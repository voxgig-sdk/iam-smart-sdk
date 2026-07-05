# IamSmart PHP SDK Reference

Complete API reference for the IamSmart PHP SDK.


## IamSmartSDK

### Constructor

```php
require_once __DIR__ . '/iamsmart_sdk.php';

$client = new IamSmartSDK($options);
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `$options` | `array` | SDK configuration options. |
| `$options["base"]` | `string` | Base URL for API requests. |
| `$options["prefix"]` | `string` | URL prefix appended after base. |
| `$options["suffix"]` | `string` | URL suffix appended after path. |
| `$options["headers"]` | `array` | Custom headers for all requests. |
| `$options["feature"]` | `array` | Feature configuration. |
| `$options["system"]` | `array` | System overrides (e.g. custom fetch). |


### Static Methods

#### `IamSmartSDK::test($testopts = null, $sdkopts = null)`

Create a test client with mock features active. Both arguments may be `null`.

```php
$client = IamSmartSDK::test();
```


### Instance Methods

#### `MobileRegistrationPoint($data = null)`

Create a new `MobileRegistrationPointEntity` instance. Pass `null` for no initial data.

#### `RegistrationServiceCounter($data = null)`

Create a new `RegistrationServiceCounterEntity` instance. Pass `null` for no initial data.

#### `SelfRegistrationKiosk($data = null)`

Create a new `SelfRegistrationKioskEntity` instance. Pass `null` for no initial data.

#### `options_map(): array`

Return a deep copy of the current SDK options.

#### `get_utility(): IamSmartUtility`

Return a copy of the SDK utility object.

#### `direct(array $fetchargs = []): array`

Make a direct HTTP request to any API endpoint. This is the raw-HTTP escape
hatch: it does **not** throw. It returns a result array
`["ok" => bool, "status" => int, "headers" => array, "data" => mixed]`, or
`["ok" => false, "err" => \Exception]` on failure. Branch on `$result["ok"]`.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `$fetchargs["path"]` | `string` | URL path with optional `{param}` placeholders. |
| `$fetchargs["method"]` | `string` | HTTP method (default: `"GET"`). |
| `$fetchargs["params"]` | `array` | Path parameter values for `{param}` substitution. |
| `$fetchargs["query"]` | `array` | Query string parameters. |
| `$fetchargs["headers"]` | `array` | Request headers (merged with defaults). |
| `$fetchargs["body"]` | `mixed` | Request body (arrays are JSON-serialized). |
| `$fetchargs["ctrl"]` | `array` | Control options. |

**Returns:** `array` — the result dict (see above); never throws.

#### `prepare(array $fetchargs = []): mixed`

Prepare a fetch definition without sending the request. Returns the
`$fetchdef` array. Throws on error.


---

## MobileRegistrationPointEntity

```php
$mobile_registration_point = $client->MobileRegistrationPoint();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `district` | `string` | No |  |
| `id` | `string` | No |  |
| `latitude` | `float` | No |  |
| `location` | `string` | No |  |
| `location_en` | `string` | No |  |
| `location_zh` | `string` | No |  |
| `longitude` | `float` | No |  |
| `name` | `string` | No |  |
| `name_en` | `string` | No |  |
| `name_zh` | `string` | No |  |
| `region` | `string` | No |  |
| `remark` | `string` | No |  |
| `schedule` | `array` | No |  |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->MobileRegistrationPoint()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): MobileRegistrationPointEntity`

Create a new `MobileRegistrationPointEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## RegistrationServiceCounterEntity

```php
$registration_service_counter = $client->RegistrationServiceCounter();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `address` | `string` | No |  |
| `address_en` | `string` | No |  |
| `address_zh` | `string` | No |  |
| `district` | `string` | No |  |
| `id` | `string` | No |  |
| `latitude` | `float` | No |  |
| `longitude` | `float` | No |  |
| `name` | `string` | No |  |
| `name_en` | `string` | No |  |
| `name_zh` | `string` | No |  |
| `operating_hour` | `string` | No |  |
| `region` | `string` | No |  |
| `remark` | `string` | No |  |
| `service` | `array` | No |  |
| `telephone` | `string` | No |  |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->RegistrationServiceCounter()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): RegistrationServiceCounterEntity`

Create a new `RegistrationServiceCounterEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## SelfRegistrationKioskEntity

```php
$self_registration_kiosk = $client->SelfRegistrationKiosk();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `address` | `string` | No |  |
| `address_en` | `string` | No |  |
| `address_zh` | `string` | No |  |
| `availability` | `string` | No |  |
| `district` | `string` | No |  |
| `floor` | `string` | No |  |
| `id` | `string` | No |  |
| `latitude` | `float` | No |  |
| `longitude` | `float` | No |  |
| `name` | `string` | No |  |
| `name_en` | `string` | No |  |
| `name_zh` | `string` | No |  |
| `operating_hour` | `string` | No |  |
| `region` | `string` | No |  |
| `remark` | `string` | No |  |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->SelfRegistrationKiosk()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): SelfRegistrationKioskEntity`

Create a new `SelfRegistrationKioskEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```php
$client = new IamSmartSDK([
  "feature" => [
    "test" => ["active" => true],
  ],
]);
```

