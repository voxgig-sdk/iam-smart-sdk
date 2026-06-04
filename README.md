# IamSmart SDK

Locate Hong Kong iAM Smart registration service counters, mobile teams, and self-registration kiosks

> TypeScript, Python, PHP, Golang, Ruby, Lua SDKs, a CLI, an interactive REPL, and an MCP server for AI agents — all generated from one OpenAPI spec by [@voxgig/sdkgen](https://github.com/voxgig/sdkgen).

## About iAM Smart API

This SDK wraps the open dataset of "iAM Smart" Registration Locations published by the [Digital Policy Office](https://www.digitalpolicy.gov.hk) of the Hong Kong SAR Government. iAM Smart is Hong Kong's digital identity service, and the dataset lists where the public can register or get help registering for it.

What you get from the API:
- Registration Service Counters — fixed in-person counters where staff can assist with iAM Smart sign-up
- Mobile Registration Points — scheduled mobile team visits at community locations
- Self-Registration Kiosks — unattended kiosks for self-service registration

Data is served as static JSON files hosted under `digitalpolicy.gov.hk/open_data/iam_smart/`. The files moved to this domain on 25 Jul 2024; the dataset is also catalogued on [data.gov.hk](https://data.gov.hk/en-data/dataset/hk-dpo-dpo_hp-iam-smart-registration-locations) where a field-level data dictionary PDF is linked.

Operational notes: no authentication is required, no rate limits are documented, and CORS is not enabled on the JSON endpoints — fetch from a server-side context if you hit browser CORS errors.

## Try it

**TypeScript**
```bash
npm install iam-smart
```

**Python**
```bash
pip install iam-smart-sdk
```

**PHP**
```bash
composer require voxgig/iam-smart-sdk
```

**Golang**
```bash
go get github.com/voxgig-sdk/iam-smart-sdk/go
```

**Ruby**
```bash
gem install iam-smart-sdk
```

**Lua**
```bash
luarocks install iam-smart-sdk
```

## 30-second quickstart

### TypeScript

```ts
import { IamSmartSDK } from 'iam-smart'

const client = new IamSmartSDK({})

// List all mobileregistrationpoints
const mobileregistrationpoints = await client.MobileRegistrationPoint().list()
```

See the [TypeScript README](ts/README.md) for the
full guide, or scroll down for the same example in other languages.

## What's in the box

| Surface | Use it for | Path |
| --- | --- | --- |
| **SDK** (TypeScript, Python, PHP, Golang, Ruby, Lua) | App integration | `ts/` `py/` `php/` `go/` `rb/` `lua/` |
| **CLI** | Scripts, CI, ops, one-off API calls | `go-cli/` |
| **MCP server** | AI agents (Claude, Cursor, Cline) | `go-mcp/` |

## Use it from an AI agent (MCP)

The generated MCP server exposes every operation in this SDK as an
[MCP](https://modelcontextprotocol.io) tool that Claude, Cursor or Cline
can call directly. Build and register it:

```bash
cd go-mcp && go build -o iam-smart-mcp .
```

Then add it to your agent's MCP config (Claude Desktop, Cursor, etc.):

```json
{
  "mcpServers": {
    "iam-smart": {
      "command": "/abs/path/to/iam-smart-mcp"
    }
  }
}
```

## Entities

The API exposes 3 entities:

| Entity | Description | API path |
| --- | --- | --- |
| **MobileRegistrationPoint** | Scheduled mobile team registration sessions held at community venues across Hong Kong, served from `/open_data/iam_smart/iAM-Smart-RegistrationMobileTeamService.json`. | `/open_data/iam_smart/mobile-registration-points` |
| **RegistrationServiceCounter** | Staffed in-person counters where the public can register for iAM Smart, served from `/open_data/iam_smart/iAM-Smart-RegistrationServiceCounter.json`. | `/open_data/iam_smart/registration-service-counters` |
| **SelfRegistrationKiosk** | Unattended kiosks that allow members of the public to self-register for iAM Smart, served from `/open_data/iam_smart/iAM-Smart-SelfRegistrationKiosk.json`. | `/open_data/iam_smart/self-registration-kiosks` |

Each entity supports the following operations where available: **load**,
**list**, **create**, **update**, and **remove**.

## Quickstart in other languages

### Python

```python
from iamsmart_sdk import IamSmartSDK

client = IamSmartSDK({})

# List all mobileregistrationpoints
mobileregistrationpoints, err = client.MobileRegistrationPoint(None).list(None, None)
```

### PHP

```php
<?php
require_once 'iamsmart_sdk.php';

$client = new IamSmartSDK([]);

// List all mobileregistrationpoints
[$mobileregistrationpoints, $err] = $client->MobileRegistrationPoint(null)->list(null, null);
```

### Golang

```go
import sdk "github.com/voxgig-sdk/iam-smart-sdk/go"

client := sdk.NewIamSmartSDK(map[string]any{})

// List all mobileregistrationpoints
mobileregistrationpoints, err := client.MobileRegistrationPoint(nil).List(nil, nil)
```

### Ruby

```ruby
require_relative "IamSmart_sdk"

client = IamSmartSDK.new({})

# List all mobileregistrationpoints
mobileregistrationpoints, err = client.MobileRegistrationPoint(nil).list(nil, nil)
```

### Lua

```lua
local sdk = require("iam-smart_sdk")

local client = sdk.new({})

-- List all mobileregistrationpoints
local mobileregistrationpoints, err = client:MobileRegistrationPoint(nil):list(nil, nil)
```

## Unit testing in offline mode

Every SDK ships a test mode that swaps the HTTP transport for an
in-memory mock, so unit tests run offline.

### TypeScript

```ts
const client = IamSmartSDK.test()
const result = await client.MobileRegistrationPoint().load({ id: 'test01' })
// result.ok === true, result.data contains mock data
```

### Python

```python
client = IamSmartSDK.test(None, None)
result, err = client.MobileRegistrationPoint(None).load(
    {"id": "test01"}, None
)
```

### PHP

```php
$client = IamSmartSDK::test(null, null);
[$result, $err] = $client->MobileRegistrationPoint(null)->load(
    ["id" => "test01"], null
);
```

### Golang

```go
client := sdk.TestSDK(nil, nil)
result, err := client.MobileRegistrationPoint(nil).Load(
    map[string]any{"id": "test01"}, nil,
)
```

### Ruby

```ruby
client = IamSmartSDK.test(nil, nil)
result, err = client.MobileRegistrationPoint(nil).load(
  { "id" => "test01" }, nil
)
```

### Lua

```lua
local client = sdk.test(nil, nil)
local result, err = client:MobileRegistrationPoint(nil):load(
  { id = "test01" }, nil
)
```

## How it works

Every SDK call runs the same five-stage pipeline:

1. **Point** — resolve the API endpoint from the operation definition.
2. **Spec** — build the HTTP specification (URL, method, headers, body).
3. **Request** — send the HTTP request.
4. **Response** — receive and parse the response.
5. **Result** — extract the result data for the caller.

A feature hook fires at each stage (e.g. `PrePoint`, `PreSpec`,
`PreRequest`), so features can inspect or modify the pipeline without
forking the SDK.

### Features

| Feature | Purpose |
| --- | --- |
| **TestFeature** | In-memory mock transport for testing without a live server |

Pass custom features via the `extend` option at construction time.

### Direct and Prepare

For endpoints the entity model doesn't cover, use the low-level methods:

- **`direct(fetchargs)`** — build and send an HTTP request in one step.
- **`prepare(fetchargs)`** — build the request without sending it.

Both accept a map with `path`, `method`, `params`, `query`,
`headers`, and `body`. See the [How-to guides](#how-to-guides) below.

## How-to guides

### Make a direct API call

When the entity interface does not cover an endpoint, use `direct`:

**TypeScript:**
```ts
const result = await client.direct({
  path: '/api/resource/{id}',
  method: 'GET',
  params: { id: 'example' },
})
console.log(result.data)
```

**Python:**
```python
result, err = client.direct({
    "path": "/api/resource/{id}",
    "method": "GET",
    "params": {"id": "example"},
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

**Go:**
```go
result, err := client.Direct(map[string]any{
    "path":   "/api/resource/{id}",
    "method": "GET",
    "params": map[string]any{"id": "example"},
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

**Lua:**
```lua
local result, err = client:direct({
  path = "/api/resource/{id}",
  method = "GET",
  params = { id = "example" },
})
```

## Per-language documentation

- [TypeScript](ts/README.md)
- [Python](py/README.md)
- [PHP](php/README.md)
- [Golang](go/README.md)
- [Ruby](rb/README.md)
- [Lua](lua/README.md)

## Using the iAM Smart API

- Upstream: [https://www.digitalpolicy.gov.hk](https://www.digitalpolicy.gov.hk)
- API docs: [https://data.gov.hk/en-data/dataset/hk-dpo-dpo_hp-iam-smart-registration-locations](https://data.gov.hk/en-data/dataset/hk-dpo-dpo_hp-iam-smart-registration-locations)

- Published as open data by the Hong Kong Digital Policy Office (DPO) via data.gov.hk
- Use is governed by the data.gov.hk Terms and Conditions
- No explicit attribution requirement is stated on the dataset page; check the data.gov.hk terms before redistribution
- Dataset is updated "as and when necessary"

---

Generated from the iAM Smart API OpenAPI spec by [@voxgig/sdkgen](https://github.com/voxgig/sdkgen).
