# IamSmart Python SDK Reference

Complete API reference for the IamSmart Python SDK.


## IamSmartSDK

### Constructor

```python
from iamsmart_sdk import IamSmartSDK

client = IamSmartSDK(options)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `dict` | SDK configuration options. |
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

#### `direct(fetchargs=None) -> dict`

Make a direct HTTP request to any API endpoint. Returns a result `dict` with `ok`, `status`, `headers`, and `data` (or `err` on failure). This escape hatch never raises — branch on `result["ok"]`.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs["path"]` | `str` | URL path with optional `{param}` placeholders. |
| `fetchargs["method"]` | `str` | HTTP method (default: `"GET"`). |
| `fetchargs["params"]` | `dict` | Path parameter values. |
| `fetchargs["query"]` | `dict` | Query string parameters. |
| `fetchargs["headers"]` | `dict` | Request headers (merged with defaults). |
| `fetchargs["body"]` | `any` | Request body (dicts are JSON-serialized). |

**Returns:** `result_dict`

#### `prepare(fetchargs=None) -> dict`

Prepare a fetch definition without sending. Returns the `fetchdef` and raises on error.


---

## MobileRegistrationPointEntity

```python
mobile_registration_point = client.MobileRegistrationPoint()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `district` | `str` | No |  |
| `id` | `str` | No |  |
| `latitude` | `float` | No |  |
| `location` | `str` | No |  |
| `location_en` | `str` | No |  |
| `location_zh` | `str` | No |  |
| `longitude` | `float` | No |  |
| `name` | `str` | No |  |
| `name_en` | `str` | No |  |
| `name_zh` | `str` | No |  |
| `region` | `str` | No |  |
| `remark` | `str` | No |  |
| `schedule` | `list` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.MobileRegistrationPoint().list()
for mobile_registration_point in results:
    print(mobile_registration_point)
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
| `address` | `str` | No |  |
| `address_en` | `str` | No |  |
| `address_zh` | `str` | No |  |
| `district` | `str` | No |  |
| `id` | `str` | No |  |
| `latitude` | `float` | No |  |
| `longitude` | `float` | No |  |
| `name` | `str` | No |  |
| `name_en` | `str` | No |  |
| `name_zh` | `str` | No |  |
| `operating_hour` | `str` | No |  |
| `region` | `str` | No |  |
| `remark` | `str` | No |  |
| `service` | `list` | No |  |
| `telephone` | `str` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.RegistrationServiceCounter().list()
for registration_service_counter in results:
    print(registration_service_counter)
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
| `address` | `str` | No |  |
| `address_en` | `str` | No |  |
| `address_zh` | `str` | No |  |
| `availability` | `str` | No |  |
| `district` | `str` | No |  |
| `floor` | `str` | No |  |
| `id` | `str` | No |  |
| `latitude` | `float` | No |  |
| `longitude` | `float` | No |  |
| `name` | `str` | No |  |
| `name_en` | `str` | No |  |
| `name_zh` | `str` | No |  |
| `operating_hour` | `str` | No |  |
| `region` | `str` | No |  |
| `remark` | `str` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.SelfRegistrationKiosk().list()
for self_registration_kiosk in results:
    print(self_registration_kiosk)
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

