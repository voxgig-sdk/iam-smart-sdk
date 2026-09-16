# IamSmart Ruby SDK Reference

Complete API reference for the IamSmart Ruby SDK.


## IamSmartSDK

### Constructor

```ruby
require_relative 'IamSmart_sdk'

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
| `district` | `String` | No | District where the mobile point operates |
| `id` | `String` | No | Unique identifier for the mobile registration point |
| `latitude` | `Float` | No | Latitude coordinate |
| `location` | `String` | No | Location description of the mobile point |
| `locationEn` | `String` | No | English location description |
| `locationZh` | `String` | No | Chinese location description |
| `longitude` | `Float` | No | Longitude coordinate |
| `name` | `String` | No | Name of the mobile registration point location |
| `nameEn` | `String` | No | English name of the mobile registration point |
| `nameZh` | `String` | No | Chinese name of the mobile registration point |
| `region` | `String` | No | Region (Hong Kong Island, Kowloon, New Territories) |
| `remarks` | `String` | No | Additional remarks or notes |
| `schedule` | `Array` | No | Schedule of mobile registration point visits |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.MobileRegistrationPoint.list
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
| `address` | `String` | No | Full address of the service counter |
| `addressEn` | `String` | No | English address of the service counter |
| `addressZh` | `String` | No | Chinese address of the service counter |
| `district` | `String` | No | District where the service counter is located |
| `id` | `String` | No | Unique identifier for the service counter |
| `latitude` | `Float` | No | Latitude coordinate |
| `longitude` | `Float` | No | Longitude coordinate |
| `name` | `String` | No | Name of the service counter location |
| `nameEn` | `String` | No | English name of the service counter location |
| `nameZh` | `String` | No | Chinese name of the service counter location |
| `operatingHours` | `String` | No | Operating hours of the service counter |
| `region` | `String` | No | Region (Hong Kong Island, Kowloon, New Territories) |
| `remarks` | `String` | No | Additional remarks or notes |
| `services` | `Array` | No | List of services available at this counter |
| `telephone` | `String` | No | Contact telephone number |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.RegistrationServiceCounter.list
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
| `address` | `String` | No | Full address of the kiosk |
| `addressEn` | `String` | No | English address of the kiosk |
| `addressZh` | `String` | No | Chinese address of the kiosk |
| `availability` | `String` | No | Availability status of the kiosk |
| `district` | `String` | No | District where the kiosk is located |
| `floor` | `String` | No | Floor level where kiosk is located |
| `id` | `String` | No | Unique identifier for the kiosk |
| `latitude` | `Float` | No | Latitude coordinate |
| `longitude` | `Float` | No | Longitude coordinate |
| `name` | `String` | No | Name of the kiosk location |
| `nameEn` | `String` | No | English name of the kiosk location |
| `nameZh` | `String` | No | Chinese name of the kiosk location |
| `operatingHours` | `String` | No | Operating hours of the kiosk |
| `region` | `String` | No | Region (Hong Kong Island, Kowloon, New Territories) |
| `remarks` | `String` | No | Additional remarks or notes |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.SelfRegistrationKiosk.list
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
| `ratelimit` | 0.0.1 | Client-side rate limiting via a token bucket |
| `retry` | 0.0.1 | Automatic retry of transient failures with exponential backoff |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |
| `timeout` | 0.0.1 | Per-request timeout with transport abort |


Features are activated via the `feature` option:

```ruby
client = IamSmartSDK.new({
  "feature" => {
    "ratelimit" => { "active" => true },
    "retry" => { "active" => true },
    "test" => { "active" => true },
    "timeout" => { "active" => true },
  },
})
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

