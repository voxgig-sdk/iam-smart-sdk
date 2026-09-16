# IamSmart Ruby SDK



The Ruby SDK for the IamSmart API — an entity-oriented client using idiomatic Ruby conventions.

The SDK exposes the API as capitalised, semantic **Entities** — for example `client.MobileRegistrationPoint` — with named operations (`list`) instead of raw URL paths and query strings. Working with resources and verbs keeps call sites self-describing and reduces cognitive load.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
This package is not yet published to RubyGems. Install it from the
GitHub release tag (`rb/vX.Y.Z`):

- Releases: [https://github.com/voxgig-sdk/iam-smart-sdk/releases](https://github.com/voxgig-sdk/iam-smart-sdk/releases)


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```ruby
require_relative "IamSmart_sdk"

client = IamSmartSDK.new
```

### 2. List mobileregistrationpoint records

```ruby
begin
  # list returns an Array of MobileRegistrationPoint records — iterate directly.
  mobileregistrationpoints = client.MobileRegistrationPoint.list
  mobileregistrationpoints.each do |item|
    puts "#{item["id"]} #{item["district"]}"
  end
rescue => err
  warn "list failed: #{err}"
end
```


## Error handling

Entity operations raise on failure, so rescue them:

```ruby
begin
  selfregistrationkiosks = client.SelfRegistrationKiosk.list()
rescue => err
  warn "list failed: #{err}"
end
```

`direct` does **not** raise — it returns the result hash. Branch on
`ok`; on failure `status` holds the HTTP status (for error responses) and
`err` holds a transport error, so read both defensively:

```ruby
result = client.direct({
  "path" => "/api/resource/{id}",
  "method" => "GET",
  "params" => { "id" => "example_id" },
})

warn "request failed: #{result["err"] || "HTTP #{result["status"]}"}" unless result["ok"]
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```ruby
result = client.direct({
  "path" => "/api/resource/{id}",
  "method" => "GET",
  "params" => { "id" => "example" },
})

if result["ok"]
  puts result["status"]  # 200
  puts result["data"]    # response body
else
  # On an HTTP error status there is no err (only a transport failure sets
  # it), so fall back to the status code.
  warn(result["err"] || "HTTP #{result["status"]}")
end
```

### Prepare a request without sending it

```ruby
begin
  fetchdef = client.prepare({
    "path" => "/api/resource/{id}",
    "method" => "DELETE",
    "params" => { "id" => "example" },
  })
  puts fetchdef["url"]
  puts fetchdef["method"]
  puts fetchdef["headers"]
rescue => err
  warn "prepare failed: #{err}"
end
```

### Use test mode

Create a mock client for unit testing — no server required:

```ruby
client = IamSmartSDK.test

# Entity ops return the ENTITY (raises on error);
# call data_get for the mock record.
selfregistrationkiosk = client.SelfRegistrationKiosk.list()
puts selfregistrationkiosk
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```ruby
mock_fetch = ->(url, init) {
  return {
    "status" => 200,
    "statusText" => "OK",
    "headers" => {},
    "json" => ->() { { "id" => "mock01" } },
  }, nil
}

client = IamSmartSDK.new({
  "base" => "http://localhost:8080",
  "system" => {
    "fetch" => mock_fetch,
  },
})
```

### Run live tests

Create a `.env.local` file at the project root:

```
IAM_SMART_TEST_LIVE=TRUE
```

Then run:

```bash
cd rb && ruby -Itest -e "Dir['test/*_test.rb'].each { |f| require_relative f }"
```


## Reference

### IamSmartSDK

```ruby
require_relative "IamSmart_sdk"
client = IamSmartSDK.new(options)
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `base` | `String` | Base URL of the API server. |
| `prefix` | `String` | URL path prefix prepended to all requests. |
| `suffix` | `String` | URL path suffix appended to all requests. |
| `feature` | `Hash` | Feature activation flags. |
| `extend` | `Hash` | Additional Feature instances to load. |
| `system` | `Hash` | System overrides (e.g. custom `fetch` lambda). |

### test

```ruby
client = IamSmartSDK.test(testopts, sdkopts)
```

Creates a test-mode client with mock transport. Both arguments may be `nil`.

### IamSmartSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `options_map` | `() -> Hash` | Deep copy of current SDK options. |
| `get_utility` | `() -> Utility` | Copy of the SDK utility object. |
| `prepare` | `(fetchargs) -> Hash` | Build an HTTP request definition without sending. Raises on error. |
| `direct` | `(fetchargs) -> Hash` | Build and send an HTTP request. Returns a result hash (`result["ok"]`); does not raise. |
| `MobileRegistrationPoint` | `(data) -> MobileRegistrationPointEntity` | Create a MobileRegistrationPoint entity instance. |
| `RegistrationServiceCounter` | `(data) -> RegistrationServiceCounterEntity` | Create a RegistrationServiceCounter entity instance. |
| `SelfRegistrationKiosk` | `(data) -> SelfRegistrationKioskEntity` | Create a SelfRegistrationKiosk entity instance. |

### Entity interface

All entities share the same interface.

| Method | Signature | Description |
| --- | --- | --- |
| `list` | `(reqmatch = nil, ctrl) -> Array` | List entities matching the criteria (call with no argument to list all). Raises on error. |
| `data_get` | `() -> Hash` | Get entity data. |
| `data_set` | `(data)` | Set entity data. |
| `match_get` | `() -> Hash` | Get entity match criteria. |
| `match_set` | `(match)` | Set entity match criteria. |
| `make` | `() -> Entity` | Create a new instance with the same options. |
| `get_name` | `() -> String` | Return the entity name. |

### Result shape

Entity operations return the result data directly. On failure they
raise a `IamSmartError` (a `StandardError` subclass), so wrap
calls in `begin`/`rescue` where you need to handle errors.

The `direct` escape hatch is the exception: it never raises and instead
returns a result `Hash` with these keys:

| Key | Type | Description |
| --- | --- | --- |
| `ok` | `Boolean` | `true` if the HTTP status is 2xx. |
| `status` | `Integer` | HTTP status code. |
| `headers` | `Hash` | Response headers. |
| `data` | `any` | Parsed JSON response body. |
| `err` | `Error` | Present when `ok` is `false`. |

### Entities

#### MobileRegistrationPoint

| Field | Description |
| --- | --- |
| `district` | District where the mobile point operates |
| `id` | Unique identifier for the mobile registration point |
| `latitude` | Latitude coordinate |
| `location` | Location description of the mobile point |
| `locationEn` | English location description |
| `locationZh` | Chinese location description |
| `longitude` | Longitude coordinate |
| `name` | Name of the mobile registration point location |
| `nameEn` | English name of the mobile registration point |
| `nameZh` | Chinese name of the mobile registration point |
| `region` | Region (Hong Kong Island, Kowloon, New Territories) |
| `remarks` | Additional remarks or notes |
| `schedule` | Schedule of mobile registration point visits |

Operations: List.

API path: `/open_data/iam_smart/mobile-registration-points`

#### RegistrationServiceCounter

| Field | Description |
| --- | --- |
| `address` | Full address of the service counter |
| `addressEn` | English address of the service counter |
| `addressZh` | Chinese address of the service counter |
| `district` | District where the service counter is located |
| `id` | Unique identifier for the service counter |
| `latitude` | Latitude coordinate |
| `longitude` | Longitude coordinate |
| `name` | Name of the service counter location |
| `nameEn` | English name of the service counter location |
| `nameZh` | Chinese name of the service counter location |
| `operatingHours` | Operating hours of the service counter |
| `region` | Region (Hong Kong Island, Kowloon, New Territories) |
| `remarks` | Additional remarks or notes |
| `services` | List of services available at this counter |
| `telephone` | Contact telephone number |

Operations: List.

API path: `/open_data/iam_smart/registration-service-counters`

#### SelfRegistrationKiosk

| Field | Description |
| --- | --- |
| `address` | Full address of the kiosk |
| `addressEn` | English address of the kiosk |
| `addressZh` | Chinese address of the kiosk |
| `availability` | Availability status of the kiosk |
| `district` | District where the kiosk is located |
| `floor` | Floor level where kiosk is located |
| `id` | Unique identifier for the kiosk |
| `latitude` | Latitude coordinate |
| `longitude` | Longitude coordinate |
| `name` | Name of the kiosk location |
| `nameEn` | English name of the kiosk location |
| `nameZh` | Chinese name of the kiosk location |
| `operatingHours` | Operating hours of the kiosk |
| `region` | Region (Hong Kong Island, Kowloon, New Territories) |
| `remarks` | Additional remarks or notes |

Operations: List.

API path: `/open_data/iam_smart/self-registration-kiosks`



## Entities


### MobileRegistrationPoint

Create an instance: `mobile_registration_point = client.MobileRegistrationPoint`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `district` | `String` | District where the mobile point operates |
| `id` | `String` | Unique identifier for the mobile registration point |
| `latitude` | `Float` | Latitude coordinate |
| `location` | `String` | Location description of the mobile point |
| `locationEn` | `String` | English location description |
| `locationZh` | `String` | Chinese location description |
| `longitude` | `Float` | Longitude coordinate |
| `name` | `String` | Name of the mobile registration point location |
| `nameEn` | `String` | English name of the mobile registration point |
| `nameZh` | `String` | Chinese name of the mobile registration point |
| `region` | `String` | Region (Hong Kong Island, Kowloon, New Territories) |
| `remarks` | `String` | Additional remarks or notes |
| `schedule` | `Array` | Schedule of mobile registration point visits |

#### Example: List

```ruby
# list returns an Array of MobileRegistrationPoint records (raises on error).
mobile_registration_points = client.MobileRegistrationPoint.list
```


### RegistrationServiceCounter

Create an instance: `registration_service_counter = client.RegistrationServiceCounter`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `address` | `String` | Full address of the service counter |
| `addressEn` | `String` | English address of the service counter |
| `addressZh` | `String` | Chinese address of the service counter |
| `district` | `String` | District where the service counter is located |
| `id` | `String` | Unique identifier for the service counter |
| `latitude` | `Float` | Latitude coordinate |
| `longitude` | `Float` | Longitude coordinate |
| `name` | `String` | Name of the service counter location |
| `nameEn` | `String` | English name of the service counter location |
| `nameZh` | `String` | Chinese name of the service counter location |
| `operatingHours` | `String` | Operating hours of the service counter |
| `region` | `String` | Region (Hong Kong Island, Kowloon, New Territories) |
| `remarks` | `String` | Additional remarks or notes |
| `services` | `Array` | List of services available at this counter |
| `telephone` | `String` | Contact telephone number |

#### Example: List

```ruby
# list returns an Array of RegistrationServiceCounter records (raises on error).
registration_service_counters = client.RegistrationServiceCounter.list
```


### SelfRegistrationKiosk

Create an instance: `self_registration_kiosk = client.SelfRegistrationKiosk`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `address` | `String` | Full address of the kiosk |
| `addressEn` | `String` | English address of the kiosk |
| `addressZh` | `String` | Chinese address of the kiosk |
| `availability` | `String` | Availability status of the kiosk |
| `district` | `String` | District where the kiosk is located |
| `floor` | `String` | Floor level where kiosk is located |
| `id` | `String` | Unique identifier for the kiosk |
| `latitude` | `Float` | Latitude coordinate |
| `longitude` | `Float` | Longitude coordinate |
| `name` | `String` | Name of the kiosk location |
| `nameEn` | `String` | English name of the kiosk location |
| `nameZh` | `String` | Chinese name of the kiosk location |
| `operatingHours` | `String` | Operating hours of the kiosk |
| `region` | `String` | Region (Hong Kong Island, Kowloon, New Territories) |
| `remarks` | `String` | Additional remarks or notes |

#### Example: List

```ruby
# list returns an Array of SelfRegistrationKiosk records (raises on error).
self_registration_kiosks = client.SelfRegistrationKiosk.list
```

## Features

This SDK ships 4 optional features. Each is **inactive until you
switch it on**, so an SDK you have not configured behaves exactly as if none of
them existed — no retries, no cache, no logging, no measurable overhead.

Activate a feature by name in the client options, alongside the options shown
above:

| Feature | What it does |
|---|---|
| [`ratelimit`](#ratelimit) | Client-side rate limiting via a token bucket |
| [`retry`](#retry) | Automatic retry of transient failures with exponential backoff |
| [`test`](#test) | In-memory mock transport for testing without a live server |
| [`timeout`](#timeout) | Per-request timeout with transport abort |

> **Order matters for `ratelimit`, `retry`, `timeout`.** These wrap the
> transport, so each one wraps whatever is already installed: the order you
> activate them in IS the nesting order. Activating them as an ordered list
> rather than a map is what fixes that order.

### ratelimit

Client-side rate limiting via a token bucket.

| Option | Default |
|---|---|
| `active` | `false` |
| `burst` | `5` |
| `rate` | `5` |

Set `feature.ratelimit.active` to enable it, then override any of the options above.

`ratelimit` wraps the transport, so its position among the other
transport features decides what it sees. A feature activated later wraps one
activated earlier.

### retry

Automatic retry of transient failures with exponential backoff.

| Option | Default |
|---|---|
| `active` | `false` |
| `factor` | `2` |
| `maxDelay` | `2000` |
| `minDelay` | `50` |
| `retries` | `2` |
| `statuses` | `[408, 425, 429, 500, 502, 503, 504]` |

Set `feature.retry.active` to enable it, then override any of the options above.

`retry` wraps the transport, so its position among the other
transport features decides what it sees. A feature activated later wraps one
activated earlier.

### test

In-memory mock transport for testing without a live server.

| Option | Default |
|---|---|
| `active` | `false` |

Set `feature.test.active` to enable it, then override any of the options above.

### timeout

Per-request timeout with transport abort.

| Option | Default |
|---|---|
| `active` | `false` |
| `ms` | `30000` |

Set `feature.timeout.active` to enable it, then override any of the options above.

`timeout` wraps the transport, so its position among the other
transport features decides what it sees. A feature activated later wraps one
activated earlier.


## Advanced

> The sections above cover everyday use. The material below explains the
> SDK's internals — useful when extending it with custom features, but not
> needed for normal use.

### The operation pipeline

Every entity operation follows a six-stage pipeline. Each stage fires a
feature hook before executing:

```
PrePoint → PreSpec → PreRequest → PreResponse → PreResult → PreDone
```

- **PrePoint**: Resolves which API endpoint to call based on the
  operation name and entity configuration.
- **PreSpec**: Builds the HTTP spec — URL, method, headers, body —
  from the resolved point and the caller's parameters.
- **PreRequest**: Sends the HTTP request. Features can intercept here
  to replace the transport (as TestFeature does with mocks).
- **PreResponse**: Parses the raw HTTP response.
- **PreResult**: Extracts the business data from the parsed response.
- **PreDone**: Final stage before returning to the caller. Entity
  state (match, data) is updated here.

If any stage errors, the pipeline short-circuits and the error surfaces
to the caller — see [Error handling](#error-handling) for how that looks
in this language.

### Features and hooks

Features are the extension mechanism. A feature is a Ruby class
with hook methods named after pipeline stages (e.g. `PrePoint`,
`PreSpec`). Each method receives the context.

The SDK ships with built-in features:

- **RatelimitFeature**: Client-side rate limiting via a token bucket
- **RetryFeature**: Automatic retry of transient failures with exponential backoff
- **TestFeature**: In-memory mock transport for testing without a live server
- **TimeoutFeature**: Per-request timeout with transport abort

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as hashes

The Ruby SDK uses plain Ruby hashes throughout rather than typed
objects. This mirrors the dynamic nature of the API and keeps the
SDK flexible — no code generation is needed when the API schema
changes.

Use `Helpers.to_map()` to safely validate that a value is a hash.

### Module structure

```
rb/
├── IamSmart_sdk.rb       -- Main SDK module
├── config.rb                  -- Configuration
├── features.rb                -- Feature factory
├── core/                      -- Core types and context
├── entity/                    -- Entity implementations
├── feature/                   -- Built-in features (Base, Test, Log)
├── utility/                   -- Utility functions and struct library
└── test/                      -- Test suites
```

The main module (`IamSmart_sdk`) exports the SDK class
and test helper. Import entity or utility modules directly only
when needed.

### Entity state

Entity instances are stateful. After a successful `list`, the entity
stores the returned data and match criteria internally.

```ruby
selfregistrationkiosk = client.SelfRegistrationKiosk
selfregistrationkiosk.list()

# selfregistrationkiosk.data_get now returns the selfregistrationkiosk data from the last list
# selfregistrationkiosk.match_get returns the last match criteria
```

Call `make` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

`direct` gives full control over the HTTP request. Use it for
non-standard endpoints, bulk operations, or any path not modelled as
an entity. `prepare` builds the request without sending it — useful
for debugging or custom transport.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.
