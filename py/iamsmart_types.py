# Typed models for the IamSmart SDK.
#
# GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
# params (op.<name>.points[].args.params[]). Field/param types come from the
# canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
# @voxgig/apidef VALID_CANON). Do not edit by hand.

from __future__ import annotations

from dataclasses import dataclass
from typing import Optional, Any


@dataclass
class MobileRegistrationPoint:
    district: Optional[str] = None
    id: Optional[str] = None
    latitude: Optional[float] = None
    location: Optional[str] = None
    location_en: Optional[str] = None
    location_zh: Optional[str] = None
    longitude: Optional[float] = None
    name: Optional[str] = None
    name_en: Optional[str] = None
    name_zh: Optional[str] = None
    region: Optional[str] = None
    remark: Optional[str] = None
    schedule: Optional[list] = None


@dataclass
class MobileRegistrationPointListMatch:
    district: Optional[str] = None
    id: Optional[str] = None
    latitude: Optional[float] = None
    location: Optional[str] = None
    location_en: Optional[str] = None
    location_zh: Optional[str] = None
    longitude: Optional[float] = None
    name: Optional[str] = None
    name_en: Optional[str] = None
    name_zh: Optional[str] = None
    region: Optional[str] = None
    remark: Optional[str] = None
    schedule: Optional[list] = None


@dataclass
class RegistrationServiceCounter:
    address: Optional[str] = None
    address_en: Optional[str] = None
    address_zh: Optional[str] = None
    district: Optional[str] = None
    id: Optional[str] = None
    latitude: Optional[float] = None
    longitude: Optional[float] = None
    name: Optional[str] = None
    name_en: Optional[str] = None
    name_zh: Optional[str] = None
    operating_hour: Optional[str] = None
    region: Optional[str] = None
    remark: Optional[str] = None
    service: Optional[list] = None
    telephone: Optional[str] = None


@dataclass
class RegistrationServiceCounterListMatch:
    address: Optional[str] = None
    address_en: Optional[str] = None
    address_zh: Optional[str] = None
    district: Optional[str] = None
    id: Optional[str] = None
    latitude: Optional[float] = None
    longitude: Optional[float] = None
    name: Optional[str] = None
    name_en: Optional[str] = None
    name_zh: Optional[str] = None
    operating_hour: Optional[str] = None
    region: Optional[str] = None
    remark: Optional[str] = None
    service: Optional[list] = None
    telephone: Optional[str] = None


@dataclass
class SelfRegistrationKiosk:
    address: Optional[str] = None
    address_en: Optional[str] = None
    address_zh: Optional[str] = None
    availability: Optional[str] = None
    district: Optional[str] = None
    floor: Optional[str] = None
    id: Optional[str] = None
    latitude: Optional[float] = None
    longitude: Optional[float] = None
    name: Optional[str] = None
    name_en: Optional[str] = None
    name_zh: Optional[str] = None
    operating_hour: Optional[str] = None
    region: Optional[str] = None
    remark: Optional[str] = None


@dataclass
class SelfRegistrationKioskListMatch:
    address: Optional[str] = None
    address_en: Optional[str] = None
    address_zh: Optional[str] = None
    availability: Optional[str] = None
    district: Optional[str] = None
    floor: Optional[str] = None
    id: Optional[str] = None
    latitude: Optional[float] = None
    longitude: Optional[float] = None
    name: Optional[str] = None
    name_en: Optional[str] = None
    name_zh: Optional[str] = None
    operating_hour: Optional[str] = None
    region: Optional[str] = None
    remark: Optional[str] = None

