# IamSmart Ruby SDK Reference

Complete API reference for the IamSmart Ruby SDK.


## IamSmartSDK

### Constructor

```ruby
require_relative 'iam-smart_sdk'

client = IamSmartSDK.new(options)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `Hash` | SDK configuration options. |
| `options["base"]` | `String` | Base URL for API requests. |
| `options["prefix"]` | `String` | URL prefix appended after base. |
| `options["suffix"]` | `String` | URL suffix appended after path. |
| `options["headers"]` | `Hash` | Custom headers for all requests. |
| `options["feature"]` | `Hash` | Feature configuration. |
| `options["system"]` | `Hash` | System overrides (e.g. custom fetch). |


### Static Methods

#### `IamSmartSDK.test(testopts = nil, sdkopts = nil)`

Create a test client with mock features active. Both arguments may be `nil`.

```ruby
client = IamSmartSDK.test
```


### Instance Methods

#### `MobileRegistrationPoint(data = nil)`

Create a new `MobileRegistrationPoint` entity instance. Pass `nil` for no initial data.

#### `RegistrationServiceCounter(data = nil)`

Create a new `RegistrationServiceCounter` entity instance. Pass `nil` for no initial data.

#### `SelfRegistrationKiosk(data = nil)`

Create a new `SelfRegistrationKiosk` entity instance. Pass `nil` for no initial data.

#### `options_map -> Hash`

Return a deep copy of the current SDK options.

#### `get_utility -> Utility`

Return a copy of the SDK utility object.

#### `direct(fetchargs = {}) -> Hash`

Make a direct HTTP request to any API endpoint. Returns a result hash
(`{ "ok" => ..., "status" => ..., "data" => ..., "err" => ... }`); it
does not raise — inspect `result["ok"]`.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs["path"]` | `String` | URL path with optional `{param}` placeholders. |
| `fetchargs["method"]` | `String` | HTTP method (default: `"GET"`). |
| `fetchargs["params"]` | `Hash` | Path parameter values for `{param}` substitution. |
| `fetchargs["query"]` | `Hash` | Query string parameters. |
| `fetchargs["headers"]` | `Hash` | Request headers (merged with defaults). |
| `fetchargs["body"]` | `any` | Request body (hashes are JSON-serialized). |
| `fetchargs["ctrl"]` | `Hash` | Control options (e.g. `{ "explain" => true }`). |

**Returns:** `Hash`

#### `prepare(fetchargs = {}) -> Hash`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `direct()`. Raises on error.

**Returns:** `Hash` (the fetch definition; raises on error)


---

## MobileRegistrationPointEntity

```ruby
mobile_registration_point = client.MobileRegistrationPoint
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `district` | ``$STRING`` | No |  |
| `id` | ``$STRING`` | No |  |
| `latitude` | ``$NUMBER`` | No |  |
| `location` | ``$STRING`` | No |  |
| `location_en` | ``$STRING`` | No |  |
| `location_zh` | ``$STRING`` | No |  |
| `longitude` | ``$NUMBER`` | No |  |
| `name` | ``$STRING`` | No |  |
| `name_en` | ``$STRING`` | No |  |
| `name_zh` | ``$STRING`` | No |  |
| `region` | ``$STRING`` | No |  |
| `remark` | ``$STRING`` | No |  |
| `schedule` | ``$ARRAY`` | No |  |

### Operations

#### `list(reqmatch, ctrl = nil) -> Array`

List entities matching the given criteria. Returns an array. Raises on error.

```ruby
results = client.MobileRegistrationPoint.list(nil)
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `MobileRegistrationPointEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## RegistrationServiceCounterEntity

```ruby
registration_service_counter = client.RegistrationServiceCounter
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `address` | ``$STRING`` | No |  |
| `address_en` | ``$STRING`` | No |  |
| `address_zh` | ``$STRING`` | No |  |
| `district` | ``$STRING`` | No |  |
| `id` | ``$STRING`` | No |  |
| `latitude` | ``$NUMBER`` | No |  |
| `longitude` | ``$NUMBER`` | No |  |
| `name` | ``$STRING`` | No |  |
| `name_en` | ``$STRING`` | No |  |
| `name_zh` | ``$STRING`` | No |  |
| `operating_hour` | ``$STRING`` | No |  |
| `region` | ``$STRING`` | No |  |
| `remark` | ``$STRING`` | No |  |
| `service` | ``$ARRAY`` | No |  |
| `telephone` | ``$STRING`` | No |  |

### Operations

#### `list(reqmatch, ctrl = nil) -> Array`

List entities matching the given criteria. Returns an array. Raises on error.

```ruby
results = client.RegistrationServiceCounter.list(nil)
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `RegistrationServiceCounterEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## SelfRegistrationKioskEntity

```ruby
self_registration_kiosk = client.SelfRegistrationKiosk
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `address` | ``$STRING`` | No |  |
| `address_en` | ``$STRING`` | No |  |
| `address_zh` | ``$STRING`` | No |  |
| `availability` | ``$STRING`` | No |  |
| `district` | ``$STRING`` | No |  |
| `floor` | ``$STRING`` | No |  |
| `id` | ``$STRING`` | No |  |
| `latitude` | ``$NUMBER`` | No |  |
| `longitude` | ``$NUMBER`` | No |  |
| `name` | ``$STRING`` | No |  |
| `name_en` | ``$STRING`` | No |  |
| `name_zh` | ``$STRING`` | No |  |
| `operating_hour` | ``$STRING`` | No |  |
| `region` | ``$STRING`` | No |  |
| `remark` | ``$STRING`` | No |  |

### Operations

#### `list(reqmatch, ctrl = nil) -> Array`

List entities matching the given criteria. Returns an array. Raises on error.

```ruby
results = client.SelfRegistrationKiosk.list(nil)
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `SelfRegistrationKioskEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```ruby
client = IamSmartSDK.new({
  "feature" => {
    "test" => { "active" => true },
  },
})
```

