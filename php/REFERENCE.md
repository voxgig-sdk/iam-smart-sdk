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
| `district` | `string` | No | District where the mobile point operates |
| `id` | `string` | No | Unique identifier for the mobile registration point |
| `latitude` | `float` | No | Latitude coordinate |
| `location` | `string` | No | Location description of the mobile point |
| `locationEn` | `string` | No | English location description |
| `locationZh` | `string` | No | Chinese location description |
| `longitude` | `float` | No | Longitude coordinate |
| `name` | `string` | No | Name of the mobile registration point location |
| `nameEn` | `string` | No | English name of the mobile registration point |
| `nameZh` | `string` | No | Chinese name of the mobile registration point |
| `region` | `string` | No | Region (Hong Kong Island, Kowloon, New Territories) |
| `remarks` | `string` | No | Additional remarks or notes |
| `schedule` | `array` | No | Schedule of mobile registration point visits |

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
| `address` | `string` | No | Full address of the service counter |
| `addressEn` | `string` | No | English address of the service counter |
| `addressZh` | `string` | No | Chinese address of the service counter |
| `district` | `string` | No | District where the service counter is located |
| `id` | `string` | No | Unique identifier for the service counter |
| `latitude` | `float` | No | Latitude coordinate |
| `longitude` | `float` | No | Longitude coordinate |
| `name` | `string` | No | Name of the service counter location |
| `nameEn` | `string` | No | English name of the service counter location |
| `nameZh` | `string` | No | Chinese name of the service counter location |
| `operatingHours` | `string` | No | Operating hours of the service counter |
| `region` | `string` | No | Region (Hong Kong Island, Kowloon, New Territories) |
| `remarks` | `string` | No | Additional remarks or notes |
| `services` | `array` | No | List of services available at this counter |
| `telephone` | `string` | No | Contact telephone number |

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
| `address` | `string` | No | Full address of the kiosk |
| `addressEn` | `string` | No | English address of the kiosk |
| `addressZh` | `string` | No | Chinese address of the kiosk |
| `availability` | `string` | No | Availability status of the kiosk |
| `district` | `string` | No | District where the kiosk is located |
| `floor` | `string` | No | Floor level where kiosk is located |
| `id` | `string` | No | Unique identifier for the kiosk |
| `latitude` | `float` | No | Latitude coordinate |
| `longitude` | `float` | No | Longitude coordinate |
| `name` | `string` | No | Name of the kiosk location |
| `nameEn` | `string` | No | English name of the kiosk location |
| `nameZh` | `string` | No | Chinese name of the kiosk location |
| `operatingHours` | `string` | No | Operating hours of the kiosk |
| `region` | `string` | No | Region (Hong Kong Island, Kowloon, New Territories) |
| `remarks` | `string` | No | Additional remarks or notes |

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
| `ratelimit` | 0.0.1 | Client-side rate limiting via a token bucket |
| `retry` | 0.0.1 | Automatic retry of transient failures with exponential backoff |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |
| `timeout` | 0.0.1 | Per-request timeout with transport abort |


Features are activated via the `feature` option:

```php
$client = new IamSmartSDK([
  "feature" => [
    "ratelimit" => ["active" => true],
    "retry" => ["active" => true],
    "test" => ["active" => true],
    "timeout" => ["active" => true],
  ],
]);
```


### Configuring features

Each feature is inactive until switched on, and an SDK with no feature
configured does no feature work at all. Every option below keeps its default
unless you name it.

The array form of \`feature\` is significant: several features wrap the
transport, and the order you list them in is the order they nest.

#### Ordering

`ratelimit`, `retry`, `timeout` wrap the transport. Each
wraps whatever is already installed, so **activation order is nesting order**:
a feature activated later sits OUTSIDE one activated earlier, and sees the call
first.

That decides behaviour, not just sequence: a feature that short-circuits the
call, such as a cache serving a hit, stops every feature nested inside it from
ever seeing that call.

`test` attach to pipeline hooks
rather than the transport, so their order does not affect what they observe.

#### `ratelimit`

Client-side rate limiting via a token bucket.

**Configuration**

| Option | Default |
|---|---|
| `active` | `false` |
| `burst` | `5` |
| `rate` | `5` |

| Option | Type |
|---|---|
| `now` | function |
| `sleep` | function |

These take no default: the feature behaves one way when you supply them and
another when you do not.

**Usage**

Set `feature.ratelimit.active` to true in the client options, and override any option above in the same entry. Every option keeps
its default unless you name it.

**Considerations**

- Wraps the transport: its place in the activation order decides what it
  sees. See [Ordering](#ordering) above.
- Inactive by default: leaving it out costs nothing at runtime.

#### `retry`

Automatic retry of transient failures with exponential backoff.

**Configuration**

| Option | Default |
|---|---|
| `active` | `false` |
| `factor` | `2` |
| `maxDelay` | `2000` |
| `minDelay` | `50` |
| `retries` | `2` |
| `statuses` | `[408, 425, 429, 500, 502, 503, 504]` |

| Option | Type |
|---|---|
| `jitter` | boolean |
| `sleep` | function |

These take no default: the feature behaves one way when you supply them and
another when you do not.

**Usage**

Set `feature.retry.active` to true in the client options, and override any option above in the same entry. Every option keeps
its default unless you name it.

**Considerations**

- Wraps the transport: its place in the activation order decides what it
  sees. See [Ordering](#ordering) above.
- Inactive by default: leaving it out costs nothing at runtime.

#### `test`

In-memory mock transport for testing without a live server.

**Configuration**

| Option | Default |
|---|---|
| `active` | `false` |

| Option | Type |
|---|---|
| `entity` | map |
| `net` | map |

These take no default: the feature behaves one way when you supply them and
another when you do not.

**Usage**

Set `feature.test.active` to true in the client options, and override any option above in the same entry. Every option keeps
its default unless you name it.

**Considerations**

- Attaches to pipeline hooks, not the transport, so activation order does
  not change what it observes.
- Installs the BASE transport that the wrapping features wrap, so it must be
  activated before them.
- Inactive by default: leaving it out costs nothing at runtime.

#### `timeout`

Per-request timeout with transport abort.

**Configuration**

| Option | Default |
|---|---|
| `active` | `false` |
| `ms` | `30000` |

| Option | Type |
|---|---|
| `clearTimer` | function |
| `setTimer` | function |

These take no default: the feature behaves one way when you supply them and
another when you do not.

**Usage**

Set `feature.timeout.active` to true in the client options, and override any option above in the same entry. Every option keeps
its default unless you name it.

**Considerations**

- Wraps the transport: its place in the activation order decides what it
  sees. See [Ordering](#ordering) above.
- Inactive by default: leaving it out costs nothing at runtime.

