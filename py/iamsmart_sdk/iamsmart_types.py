# Typed models for the IamSmart SDK.
#
# GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
# params (op.<name>.points[].args.params[]). Field/param types come from the
# canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
# @voxgig/apidef VALID_CANON). Do not edit by hand.
#
# These are TypedDicts, not dataclasses: the SDK ops return/accept plain dicts
# at runtime, and a TypedDict IS a dict shape, so the types match the runtime.
# Optional (req:false) keys are modelled as TypedDict key-optionality
# (total=False), split into a required base + total=False subclass when a type
# has both required and optional keys.

from __future__ import annotations

from typing import TypedDict, Any


class MobileRegistrationPoint(TypedDict, total=False):
    district: str
    id: str
    latitude: float
    location: str
    locationEn: str
    locationZh: str
    longitude: float
    name: str
    nameEn: str
    nameZh: str
    region: str
    remarks: str
    schedule: list


class MobileRegistrationPointListMatch(TypedDict, total=False):
    district: str
    id: str
    latitude: float
    location: str
    locationEn: str
    locationZh: str
    longitude: float
    name: str
    nameEn: str
    nameZh: str
    region: str
    remarks: str
    schedule: list


class RegistrationServiceCounter(TypedDict, total=False):
    address: str
    addressEn: str
    addressZh: str
    district: str
    id: str
    latitude: float
    longitude: float
    name: str
    nameEn: str
    nameZh: str
    operatingHours: str
    region: str
    remarks: str
    services: list
    telephone: str


class RegistrationServiceCounterListMatch(TypedDict, total=False):
    address: str
    addressEn: str
    addressZh: str
    district: str
    id: str
    latitude: float
    longitude: float
    name: str
    nameEn: str
    nameZh: str
    operatingHours: str
    region: str
    remarks: str
    services: list
    telephone: str


class SelfRegistrationKiosk(TypedDict, total=False):
    address: str
    addressEn: str
    addressZh: str
    availability: str
    district: str
    floor: str
    id: str
    latitude: float
    longitude: float
    name: str
    nameEn: str
    nameZh: str
    operatingHours: str
    region: str
    remarks: str


class SelfRegistrationKioskListMatch(TypedDict, total=False):
    address: str
    addressEn: str
    addressZh: str
    availability: str
    district: str
    floor: str
    id: str
    latitude: float
    longitude: float
    name: str
    nameEn: str
    nameZh: str
    operatingHours: str
    region: str
    remarks: str
