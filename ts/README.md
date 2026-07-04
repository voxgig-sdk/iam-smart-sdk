# IamSmart TypeScript SDK



The TypeScript SDK for the IamSmart API — a type-safe, entity-oriented client with full async/await support.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
This package is not yet published to npm. Install it from the GitHub
release tag (`ts/vX.Y.Z`):

- Releases: [https://github.com/voxgig-sdk/iam-smart-sdk/releases](https://github.com/voxgig-sdk/iam-smart-sdk/releases)


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```ts
import { IamSmartSDK } from '@voxgig-sdk/iam-smart'

const client = new IamSmartSDK()
```

### 2. List mobileregistrationpoint records

`list()` resolves to an array of MobileRegistrationPoint objects — iterate it directly:

```ts
const mobileregistrationpoints = await client.MobileRegistrationPoint().list()

for (const mobileregistrationpoint of mobileregistrationpoints) {
  console.log(mobileregistrationpoint)
}
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```ts
const result = await client.direct({
  path: '/api/resource/{id}',
  method: 'GET',
  params: { id: 'example' },
})

if (result instanceof Error) {
  throw result
}
if (result.ok) {
  console.log(result.status)  // 200
  console.log(result.data)    // response body
}
```

### Prepare a request without sending it

```ts
const fetchdef = await client.prepare({
  path: '/api/resource/{id}',
  method: 'DELETE',
  params: { id: 'example' },
})

// Inspect before sending
console.log(fetchdef.url)
console.log(fetchdef.method)
console.log(fetchdef.headers)
```

### Use test mode

Create a mock client for unit testing — no server required:

```ts
const client = IamSmartSDK.test()

const mobileregistrationpoint = await client.MobileRegistrationPoint().load({ id: 'test01' })
// mobileregistrationpoint is a bare entity populated with mock response data
console.log(mobileregistrationpoint)
```

You can also use the instance method:

```ts
const client = new IamSmartSDK()
const testClient = client.tester()
```

### Retain entity state across calls

Entity instances remember their last match and data:

```ts
const entity = client.MobileRegistrationPoint()

// First call sets internal match
await entity.load({ id: 'example' })

// Subsequent calls reuse the stored match
const data = entity.data()
console.log(data.id) // 'example'
```

### Add custom middleware

Pass features via the `extend` option:

```ts
const logger = {
  hooks: {
    PreRequest: (ctx: any) => {
      console.log('Requesting:', ctx.spec.method, ctx.spec.path)
    },
    PreResponse: (ctx: any) => {
      console.log('Status:', ctx.out.request?.status)
    },
  },
}

const client = new IamSmartSDK({
  extend: [logger],
})
```

### Run live tests

Create a `.env.local` file at the project root:

```
IAM_SMART_TEST_LIVE=TRUE
```

Then run:

```bash
cd ts && npm test
```


## Reference

### IamSmartSDK

#### Constructor

```ts
new IamSmartSDK(options?: {
  base?: string
  prefix?: string
  suffix?: string
  feature?: Record<string, { active: boolean }>
  extend?: Feature[]
})
```

| Option | Type | Description |
| --- | --- | --- |
| `base` | `string` | Base URL of the API server. |
| `prefix` | `string` | URL path prefix prepended to all requests. |
| `suffix` | `string` | URL path suffix appended to all requests. |
| `feature` | `object` | Feature activation flags (e.g. `{ test: { active: true } }`). |
| `extend` | `Feature[]` | Additional feature instances to load. |

#### Methods

| Method | Returns | Description |
| --- | --- | --- |
| `options()` | `object` | Deep copy of current SDK options. |
| `utility()` | `Utility` | Deep copy of the SDK utility object. |
| `prepare(fetchargs?)` | `Promise<FetchDef>` | Build an HTTP request definition without sending it. |
| `direct(fetchargs?)` | `Promise<DirectResult>` | Build and send an HTTP request. |
| `MobileRegistrationPoint(data?)` | `MobileRegistrationPointEntity` | Create a MobileRegistrationPoint entity instance. |
| `RegistrationServiceCounter(data?)` | `RegistrationServiceCounterEntity` | Create a RegistrationServiceCounter entity instance. |
| `SelfRegistrationKiosk(data?)` | `SelfRegistrationKioskEntity` | Create a SelfRegistrationKiosk entity instance. |
| `tester(testopts?, sdkopts?)` | `IamSmartSDK` | Create a test-mode client instance. |

#### Static methods

| Method | Returns | Description |
| --- | --- | --- |
| `IamSmartSDK.test(testopts?, sdkopts?)` | `IamSmartSDK` | Create a test-mode client. |

### Entity interface

All entities share the same interface.

#### Methods

| Method | Signature | Description |
| --- | --- | --- |
| `load` | `load(reqmatch?, ctrl?): Promise<Entity>` | Load a single entity by match criteria. |
| `list` | `list(reqmatch?, ctrl?): Promise<Entity[]>` | List entities matching the criteria. |
| `create` | `create(reqdata?, ctrl?): Promise<Entity>` | Create a new entity. |
| `update` | `update(reqdata?, ctrl?): Promise<Entity>` | Update an existing entity. |
| `remove` | `remove(reqmatch?, ctrl?): Promise<void>` | Remove an entity. |
| `data` | `data(data?): any` | Get or set entity data. |
| `match` | `match(match?): any` | Get or set entity match criteria. |
| `make` | `make(): Entity` | Create a new instance with the same options. |
| `client` | `client(): IamSmartSDK` | Return the parent SDK client. |
| `entopts` | `entopts(): object` | Return a copy of the entity options. |

#### Return values

Entity operations resolve to the entity data directly — there is no
result envelope:

- `load`, `create` and `update` resolve to a single entity object.
- `list` resolves to an **array** of entity objects (iterate it directly;
  there is no `.data` and no `.ok`).
- `remove` resolves to `void`.

On a failed request these methods **throw**, so wrap calls in
`try`/`catch` to handle errors. Only `direct()` returns the result
envelope described below.

### DirectResult shape

The `direct()` method returns:

```ts
{
  ok: boolean
  status: number
  headers: object
  data: any
}
```

On error, `ok` is `false` and an `err` property contains the error.

### FetchDef shape

The `prepare()` method returns:

```ts
{
  url: string
  method: string
  headers: Record<string, string>
  body?: any
}
```

### Entities

#### MobileRegistrationPoint

| Field | Description |
| --- | --- |
| `district` |  |
| `id` |  |
| `latitude` |  |
| `location` |  |
| `location_en` |  |
| `location_zh` |  |
| `longitude` |  |
| `name` |  |
| `name_en` |  |
| `name_zh` |  |
| `region` |  |
| `remark` |  |
| `schedule` |  |

Operations: list.

API path: `/open_data/iam_smart/mobile-registration-points`

#### RegistrationServiceCounter

| Field | Description |
| --- | --- |
| `address` |  |
| `address_en` |  |
| `address_zh` |  |
| `district` |  |
| `id` |  |
| `latitude` |  |
| `longitude` |  |
| `name` |  |
| `name_en` |  |
| `name_zh` |  |
| `operating_hour` |  |
| `region` |  |
| `remark` |  |
| `service` |  |
| `telephone` |  |

Operations: list.

API path: `/open_data/iam_smart/registration-service-counters`

#### SelfRegistrationKiosk

| Field | Description |
| --- | --- |
| `address` |  |
| `address_en` |  |
| `address_zh` |  |
| `availability` |  |
| `district` |  |
| `floor` |  |
| `id` |  |
| `latitude` |  |
| `longitude` |  |
| `name` |  |
| `name_en` |  |
| `name_zh` |  |
| `operating_hour` |  |
| `region` |  |
| `remark` |  |

Operations: list.

API path: `/open_data/iam_smart/self-registration-kiosks`



## Entities


### MobileRegistrationPoint

Create an instance: `const mobile_registration_point = client.MobileRegistrationPoint()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `district` | ``$STRING`` |  |
| `id` | ``$STRING`` |  |
| `latitude` | ``$NUMBER`` |  |
| `location` | ``$STRING`` |  |
| `location_en` | ``$STRING`` |  |
| `location_zh` | ``$STRING`` |  |
| `longitude` | ``$NUMBER`` |  |
| `name` | ``$STRING`` |  |
| `name_en` | ``$STRING`` |  |
| `name_zh` | ``$STRING`` |  |
| `region` | ``$STRING`` |  |
| `remark` | ``$STRING`` |  |
| `schedule` | ``$ARRAY`` |  |

#### Example: List

```ts
const mobile_registration_points = await client.MobileRegistrationPoint().list()
```


### RegistrationServiceCounter

Create an instance: `const registration_service_counter = client.RegistrationServiceCounter()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `address` | ``$STRING`` |  |
| `address_en` | ``$STRING`` |  |
| `address_zh` | ``$STRING`` |  |
| `district` | ``$STRING`` |  |
| `id` | ``$STRING`` |  |
| `latitude` | ``$NUMBER`` |  |
| `longitude` | ``$NUMBER`` |  |
| `name` | ``$STRING`` |  |
| `name_en` | ``$STRING`` |  |
| `name_zh` | ``$STRING`` |  |
| `operating_hour` | ``$STRING`` |  |
| `region` | ``$STRING`` |  |
| `remark` | ``$STRING`` |  |
| `service` | ``$ARRAY`` |  |
| `telephone` | ``$STRING`` |  |

#### Example: List

```ts
const registration_service_counters = await client.RegistrationServiceCounter().list()
```


### SelfRegistrationKiosk

Create an instance: `const self_registration_kiosk = client.SelfRegistrationKiosk()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `address` | ``$STRING`` |  |
| `address_en` | ``$STRING`` |  |
| `address_zh` | ``$STRING`` |  |
| `availability` | ``$STRING`` |  |
| `district` | ``$STRING`` |  |
| `floor` | ``$STRING`` |  |
| `id` | ``$STRING`` |  |
| `latitude` | ``$NUMBER`` |  |
| `longitude` | ``$NUMBER`` |  |
| `name` | ``$STRING`` |  |
| `name_en` | ``$STRING`` |  |
| `name_zh` | ``$STRING`` |  |
| `operating_hour` | ``$STRING`` |  |
| `region` | ``$STRING`` |  |
| `remark` | ``$STRING`` |  |

#### Example: List

```ts
const self_registration_kiosks = await client.SelfRegistrationKiosk().list()
```


## Explanation

### The operation pipeline

Every entity operation (load, list, create, update, remove) follows a
six-stage pipeline. Each stage fires a feature hook before executing:

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

If any stage returns an error, the pipeline short-circuits and the
error is returned to the caller.

An unexpected exception triggers the `PreUnexpected` hook before
propagating.

### Features and hooks

Features are the extension mechanism. A feature is an object with a
`hooks` map. Each hook key is a pipeline stage name, and the value is
a function that receives the context.

The SDK ships with built-in features:

- **TestFeature**: In-memory mock transport for testing without a live server

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Module structure

```
iam-smart/
├── src/
│   ├── IamSmartSDK.ts        # Main SDK class
│   ├── entity/             # Entity implementations
│   ├── feature/            # Built-in features (Base, Test, Log)
│   └── utility/            # Utility functions
├── test/                   # Test suites
└── dist/                   # Compiled output
```

Import the SDK from the package root:

```ts
import { IamSmartSDK } from '@voxgig-sdk/iam-smart'
```

### Entity state

Entity instances are stateful. After a successful `load`, the entity
stores the returned data and match criteria internally. Subsequent
calls on the same instance can rely on this state.

```ts
const mobileregistrationpoint = client.MobileRegistrationPoint()
await mobileregistrationpoint.load({ id: "example_id" })

// mobileregistrationpoint.data() now returns the loaded mobileregistrationpoint data
// mobileregistrationpoint.match() returns { id: "example_id" }
```

Call `make()` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

The `direct` method gives full control over the HTTP request. Use it
for non-standard endpoints, bulk operations, or any path not modelled
as an entity. The `prepare` method is useful for debugging — it
shows exactly what `direct` would send.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.
