# IamSmart Python SDK Reference

Complete API reference for the IamSmart Python SDK.


## IamSmartSDK

### Constructor

```python
from iam-smart_sdk import IamSmartSDK

client = IamSmartSDK(options)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `dict` | SDK configuration options. |
| `options["apikey"]` | `str` | API key for authentication. |
| `options["base"]` | `str` | Base URL for API requests. |
| `options["prefix"]` | `str` | URL prefix appended after base. |
| `options["suffix"]` | `str` | URL suffix appended after path. |
| `options["headers"]` | `dict` | Custom headers for all requests. |
| `options["feature"]` | `dict` | Feature configuration. |
| `options["system"]` | `dict` | System overrides (e.g. custom fetch). |


### Static Methods

#### `IamSmartSDK.test(testopts=None, sdkopts=None)`

Create a test client with mock features active. Both arguments may be `None`.

```python
client = IamSmartSDK.test()
```


### Instance Methods

#### `MobileRegistrationPoint(data=None)`

Create a new `MobileRegistrationPointEntity` instance. Pass `None` for no initial data.

#### `RegistrationServiceCounter(data=None)`

Create a new `RegistrationServiceCounterEntity` instance. Pass `None` for no initial data.

#### `SelfRegistrationKiosk(data=None)`

Create a new `SelfRegistrationKioskEntity` instance. Pass `None` for no initial data.

#### `options_map() -> dict`

Return a deep copy of the current SDK options.

#### `get_utility() -> Utility`

Return a copy of the SDK utility object.

#### `direct(fetchargs=None) -> tuple`

Make a direct HTTP request to any API endpoint. Returns `(result, err)`.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs["path"]` | `str` | URL path with optional `{param}` placeholders. |
| `fetchargs["method"]` | `str` | HTTP method (default: `"GET"`). |
| `fetchargs["params"]` | `dict` | Path parameter values. |
| `fetchargs["query"]` | `dict` | Query string parameters. |
| `fetchargs["headers"]` | `dict` | Request headers (merged with defaults). |
| `fetchargs["body"]` | `any` | Request body (dicts are JSON-serialized). |

**Returns:** `(result_dict, err)`

#### `prepare(fetchargs=None) -> tuple`

Prepare a fetch definition without sending. Returns `(fetchdef, err)`.


---

## MobileRegistrationPointEntity

```python
mobile_registration_point = client.MobileRegistrationPoint()
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

#### `list(reqmatch, ctrl=None) -> tuple`

List entities matching the given criteria. Returns an array.

```python
results, err = client.MobileRegistrationPoint().list({})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `MobileRegistrationPointEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## RegistrationServiceCounterEntity

```python
registration_service_counter = client.RegistrationServiceCounter()
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

#### `list(reqmatch, ctrl=None) -> tuple`

List entities matching the given criteria. Returns an array.

```python
results, err = client.RegistrationServiceCounter().list({})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `RegistrationServiceCounterEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## SelfRegistrationKioskEntity

```python
self_registration_kiosk = client.SelfRegistrationKiosk()
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

#### `list(reqmatch, ctrl=None) -> tuple`

List entities matching the given criteria. Returns an array.

```python
results, err = client.SelfRegistrationKiosk().list({})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `SelfRegistrationKioskEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```python
client = IamSmartSDK({
    "feature": {
        "test": {"active": True},
    },
})
```

