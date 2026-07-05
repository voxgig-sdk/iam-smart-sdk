# IamSmart TypeScript SDK Reference

Complete API reference for the IamSmart TypeScript SDK.


## IamSmartSDK

### Constructor

```ts
new IamSmartSDK(options?: object)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `object` | SDK configuration options. |
| `options.base` | `string` | Base URL for API requests. |
| `options.prefix` | `string` | URL prefix appended after base. |
| `options.suffix` | `string` | URL suffix appended after path. |
| `options.headers` | `object` | Custom headers for all requests. |
| `options.feature` | `object` | Feature configuration. |
| `options.system` | `object` | System overrides (e.g. custom fetch). |


### Static Methods

#### `IamSmartSDK.test(testopts?, sdkopts?)`

Create a test client with mock features active.

```ts
const client = IamSmartSDK.test()
```

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `testopts` | `object` | Test feature options. |
| `sdkopts` | `object` | Additional SDK options merged with test defaults. |

**Returns:** `IamSmartSDK` instance in test mode.


### Instance Methods

#### `MobileRegistrationPoint(data?: object)`

Create a new `MobileRegistrationPoint` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `MobileRegistrationPointEntity` instance.

#### `RegistrationServiceCounter(data?: object)`

Create a new `RegistrationServiceCounter` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `RegistrationServiceCounterEntity` instance.

#### `SelfRegistrationKiosk(data?: object)`

Create a new `SelfRegistrationKiosk` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `SelfRegistrationKioskEntity` instance.

#### `options()`

Return a deep copy of the current SDK options.

**Returns:** `object`

#### `utility()`

Return a copy of the SDK utility object.

**Returns:** `object`

#### `direct(fetchargs?: object)`

Make a direct HTTP request to any API endpoint.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs.path` | `string` | URL path with optional `{param}` placeholders. |
| `fetchargs.method` | `string` | HTTP method (default: `GET`). |
| `fetchargs.params` | `object` | Path parameter values for `{param}` substitution. |
| `fetchargs.query` | `object` | Query string parameters. |
| `fetchargs.headers` | `object` | Request headers (merged with defaults). |
| `fetchargs.body` | `any` | Request body (objects are JSON-serialized). |
| `fetchargs.ctrl` | `object` | Control options (e.g. `{ explain: true }`). |

**Returns:** `Promise<{ ok, status, headers, data } | Error>`

#### `prepare(fetchargs?: object)`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `direct()`.

**Returns:** `Promise<{ url, method, headers, body } | Error>`

#### `tester(testopts?, sdkopts?)`

Alias for `IamSmartSDK.test()`.

**Returns:** `IamSmartSDK` instance in test mode.


---

## MobileRegistrationPointEntity

```ts
const mobile_registration_point = client.MobileRegistrationPoint()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `district` | `string` | No |  |
| `id` | `string` | No |  |
| `latitude` | `number` | No |  |
| `location` | `string` | No |  |
| `location_en` | `string` | No |  |
| `location_zh` | `string` | No |  |
| `longitude` | `number` | No |  |
| `name` | `string` | No |  |
| `name_en` | `string` | No |  |
| `name_zh` | `string` | No |  |
| `region` | `string` | No |  |
| `remark` | `string` | No |  |
| `schedule` | `any[]` | No |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.MobileRegistrationPoint().list()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `MobileRegistrationPointEntity` instance with the same client and
options.

#### `client()`

Return the parent `IamSmartSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## RegistrationServiceCounterEntity

```ts
const registration_service_counter = client.RegistrationServiceCounter()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `address` | `string` | No |  |
| `address_en` | `string` | No |  |
| `address_zh` | `string` | No |  |
| `district` | `string` | No |  |
| `id` | `string` | No |  |
| `latitude` | `number` | No |  |
| `longitude` | `number` | No |  |
| `name` | `string` | No |  |
| `name_en` | `string` | No |  |
| `name_zh` | `string` | No |  |
| `operating_hour` | `string` | No |  |
| `region` | `string` | No |  |
| `remark` | `string` | No |  |
| `service` | `any[]` | No |  |
| `telephone` | `string` | No |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.RegistrationServiceCounter().list()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `RegistrationServiceCounterEntity` instance with the same client and
options.

#### `client()`

Return the parent `IamSmartSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## SelfRegistrationKioskEntity

```ts
const self_registration_kiosk = client.SelfRegistrationKiosk()
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
| `latitude` | `number` | No |  |
| `longitude` | `number` | No |  |
| `name` | `string` | No |  |
| `name_en` | `string` | No |  |
| `name_zh` | `string` | No |  |
| `operating_hour` | `string` | No |  |
| `region` | `string` | No |  |
| `remark` | `string` | No |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.SelfRegistrationKiosk().list()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `SelfRegistrationKioskEntity` instance with the same client and
options.

#### `client()`

Return the parent `IamSmartSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```ts
const client = new IamSmartSDK({
  feature: {
    test: { active: true },
  }
})
```

