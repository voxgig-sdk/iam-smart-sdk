# IamSmart Lua SDK Reference

Complete API reference for the IamSmart Lua SDK.


## IamSmartSDK

### Constructor

```lua
local sdk = require("iam-smart_sdk")
local client = sdk.new(options)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `table` | SDK configuration options. |
| `options.apikey` | `string` | API key for authentication. |
| `options.base` | `string` | Base URL for API requests. |
| `options.prefix` | `string` | URL prefix appended after base. |
| `options.suffix` | `string` | URL suffix appended after path. |
| `options.headers` | `table` | Custom headers for all requests. |
| `options.feature` | `table` | Feature configuration. |
| `options.system` | `table` | System overrides (e.g. custom fetch). |


### Static Methods

#### `sdk.test(testopts, sdkopts)`

Create a test client with mock features active. Both arguments may be `nil`.

```lua
local client = sdk.test(nil, nil)
```


### Instance Methods

#### `MobileRegistrationPoint(data)`

Create a new `MobileRegistrationPoint` entity instance. Pass `nil` for no initial data.

#### `RegistrationServiceCounter(data)`

Create a new `RegistrationServiceCounter` entity instance. Pass `nil` for no initial data.

#### `SelfRegistrationKiosk(data)`

Create a new `SelfRegistrationKiosk` entity instance. Pass `nil` for no initial data.

#### `options_map() -> table`

Return a deep copy of the current SDK options.

#### `get_utility() -> Utility`

Return a copy of the SDK utility object.

#### `direct(fetchargs) -> table, err`

Make a direct HTTP request to any API endpoint.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs.path` | `string` | URL path with optional `{param}` placeholders. |
| `fetchargs.method` | `string` | HTTP method (default: `"GET"`). |
| `fetchargs.params` | `table` | Path parameter values for `{param}` substitution. |
| `fetchargs.query` | `table` | Query string parameters. |
| `fetchargs.headers` | `table` | Request headers (merged with defaults). |
| `fetchargs.body` | `any` | Request body (tables are JSON-serialized). |
| `fetchargs.ctrl` | `table` | Control options (e.g. `{ explain = true }`). |

**Returns:** `table, err`

#### `prepare(fetchargs) -> table, err`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `direct()`.

**Returns:** `table, err`


---

## MobileRegistrationPointEntity

```lua
local mobile_registration_point = client:MobileRegistrationPoint(nil)
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

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:MobileRegistrationPoint(nil):list(nil, nil)
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `MobileRegistrationPointEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## RegistrationServiceCounterEntity

```lua
local registration_service_counter = client:RegistrationServiceCounter(nil)
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

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:RegistrationServiceCounter(nil):list(nil, nil)
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `RegistrationServiceCounterEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## SelfRegistrationKioskEntity

```lua
local self_registration_kiosk = client:SelfRegistrationKiosk(nil)
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

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:SelfRegistrationKiosk(nil):list(nil, nil)
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `SelfRegistrationKioskEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```lua
local client = sdk.new({
  feature = {
    test = { active = true },
  },
})
```

