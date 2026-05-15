# IamSmart SDK



Available for [Golang](go/) and [Lua](lua/) and [PHP](php/) and [Python](py/) and [Ruby](rb/) and [TypeScript](ts/).


## Entities

The API exposes 3 entities:

| Entity | Description | API path |
| --- | --- | --- |
| **MobileRegistrationPoint** |  | `/open_data/iam_smart/mobile-registration-points` |
| **RegistrationServiceCounter** |  | `/open_data/iam_smart/registration-service-counters` |
| **SelfRegistrationKiosk** |  | `/open_data/iam_smart/self-registration-kiosks` |

Each entity supports the following operations where available: **load**, **list**, **create**,
**update**, and **remove**.


## Architecture

### Entity-operation model

Every SDK call follows the same pipeline:

1. **Point** — resolve the API endpoint from the operation definition.
2. **Spec** — build the HTTP specification (URL, method, headers, body).
3. **Request** — send the HTTP request.
4. **Response** — receive and parse the response.
5. **Result** — extract the result data for the caller.

At each stage a feature hook fires (e.g. `PrePoint`, `PreSpec`,
`PreRequest`), allowing features to inspect or modify the pipeline.

### Features

Features are hook-based middleware that extend SDK behaviour.

| Feature | Purpose |
| --- | --- |
| **TestFeature** | In-memory mock transport for testing without a live server |

You can add custom features by passing them in the `extend` option at
construction time.

### Direct and Prepare

For endpoints not covered by the entity model, use the low-level methods:

- **`direct(fetchargs)`** — build and send an HTTP request in one step.
- **`prepare(fetchargs)`** — build the request without sending it.

Both accept a map with `path`, `method`, `params`, `query`, `headers`,
and `body`.


## Quick start

### Golang

```go
import sdk "github.com/voxgig-sdk/iam-smart-sdk"

client := sdk.NewIamSmartSDK(map[string]any{
    "apikey": os.Getenv("IAM-SMART_APIKEY"),
})

// List all mobileregistrationpoints
mobileregistrationpoints, err := client.MobileRegistrationPoint(nil).List(nil, nil)
```

### Lua

```lua
local sdk = require("iam-smart_sdk")

local client = sdk.new({
  apikey = os.getenv("IAM-SMART_APIKEY"),
})

-- List all mobileregistrationpoints
local mobileregistrationpoints, err = client:MobileRegistrationPoint(nil):list(nil, nil)
```

### PHP

```php
<?php
require_once 'iamsmart_sdk.php';

$client = new IamSmartSDK([
    "apikey" => getenv("IAM-SMART_APIKEY"),
]);

// List all mobileregistrationpoints
[$mobileregistrationpoints, $err] = $client->MobileRegistrationPoint(null)->list(null, null);
```

### Python

```python
import os
from iamsmart_sdk import IamSmartSDK

client = IamSmartSDK({
    "apikey": os.environ.get("IAM-SMART_APIKEY"),
})

# List all mobileregistrationpoints
mobileregistrationpoints, err = client.MobileRegistrationPoint(None).list(None, None)
```

### Ruby

```ruby
require_relative "IamSmart_sdk"

client = IamSmartSDK.new({
  "apikey" => ENV["IAM-SMART_APIKEY"],
})

# List all mobileregistrationpoints
mobileregistrationpoints, err = client.MobileRegistrationPoint(nil).list(nil, nil)
```

### TypeScript

```ts
import { IamSmartSDK } from 'iam-smart'

const client = new IamSmartSDK({
  apikey: process.env.IAM-SMART_APIKEY,
})

// List all mobileregistrationpoints
const mobileregistrationpoints = await client.MobileRegistrationPoint().list()
```


## Testing

Both SDKs provide a test mode that replaces the HTTP transport with an
in-memory mock, so tests run without a network connection.

### Golang

```go
client := sdk.TestSDK(nil, nil)
result, err := client.MobileRegistrationPoint(nil).Load(
    map[string]any{"id": "test01"}, nil,
)
```

### Lua

```lua
local client = sdk.test(nil, nil)
local result, err = client:MobileRegistrationPoint(nil):load(
  { id = "test01" }, nil
)
```

### PHP

```php
$client = IamSmartSDK::test(null, null);
[$result, $err] = $client->MobileRegistrationPoint(null)->load(
    ["id" => "test01"], null
);
```

### Python

```python
client = IamSmartSDK.test(None, None)
result, err = client.MobileRegistrationPoint(None).load(
    {"id": "test01"}, None
)
```

### Ruby

```ruby
client = IamSmartSDK.test(nil, nil)
result, err = client.MobileRegistrationPoint(nil).load(
  { "id" => "test01" }, nil
)
```

### TypeScript

```ts
const client = IamSmartSDK.test()
const result = await client.MobileRegistrationPoint().load({ id: 'test01' })
// result.ok === true, result.data contains mock data
```


## How-to guides

### Make a direct API call

When the entity interface does not cover an endpoint, use `direct`:

**Go:**
```go
result, err := client.Direct(map[string]any{
    "path":   "/api/resource/{id}",
    "method": "GET",
    "params": map[string]any{"id": "example"},
})
```

**Lua:**
```lua
local result, err = client:direct({
  path = "/api/resource/{id}",
  method = "GET",
  params = { id = "example" },
})
```

**PHP:**
```php
[$result, $err] = $client->direct([
    "path" => "/api/resource/{id}",
    "method" => "GET",
    "params" => ["id" => "example"],
]);
```

**Python:**
```python
result, err = client.direct({
    "path": "/api/resource/{id}",
    "method": "GET",
    "params": {"id": "example"},
})
```

**Ruby:**
```ruby
result, err = client.direct({
  "path" => "/api/resource/{id}",
  "method" => "GET",
  "params" => { "id" => "example" },
})
```

**TypeScript:**
```ts
const result = await client.direct({
  path: '/api/resource/{id}',
  method: 'GET',
  params: { id: 'example' },
})
console.log(result.data)
```


## Language-specific documentation

- [Golang SDK](go/README.md)
- [Lua SDK](lua/README.md)
- [PHP SDK](php/README.md)
- [Python SDK](py/README.md)
- [Ruby SDK](rb/README.md)
- [TypeScript SDK](ts/README.md)

