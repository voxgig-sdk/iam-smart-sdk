// Typed models for the IamSmart SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.

export interface MobileRegistrationPoint {
  district?: string
  id?: string
  latitude?: number
  location?: string
  locationEn?: string
  locationZh?: string
  longitude?: number
  name?: string
  nameEn?: string
  nameZh?: string
  region?: string
  remarks?: string
  schedule?: any[]
}

export interface MobileRegistrationPointListMatch {
  district?: string
  id?: string
  latitude?: number
  location?: string
  locationEn?: string
  locationZh?: string
  longitude?: number
  name?: string
  nameEn?: string
  nameZh?: string
  region?: string
  remarks?: string
  schedule?: any[]
}

export interface RegistrationServiceCounter {
  address?: string
  addressEn?: string
  addressZh?: string
  district?: string
  id?: string
  latitude?: number
  longitude?: number
  name?: string
  nameEn?: string
  nameZh?: string
  operatingHours?: string
  region?: string
  remarks?: string
  services?: any[]
  telephone?: string
}

export interface RegistrationServiceCounterListMatch {
  address?: string
  addressEn?: string
  addressZh?: string
  district?: string
  id?: string
  latitude?: number
  longitude?: number
  name?: string
  nameEn?: string
  nameZh?: string
  operatingHours?: string
  region?: string
  remarks?: string
  services?: any[]
  telephone?: string
}

export interface SelfRegistrationKiosk {
  address?: string
  addressEn?: string
  addressZh?: string
  availability?: string
  district?: string
  floor?: string
  id?: string
  latitude?: number
  longitude?: number
  name?: string
  nameEn?: string
  nameZh?: string
  operatingHours?: string
  region?: string
  remarks?: string
}

export interface SelfRegistrationKioskListMatch {
  address?: string
  addressEn?: string
  addressZh?: string
  availability?: string
  district?: string
  floor?: string
  id?: string
  latitude?: number
  longitude?: number
  name?: string
  nameEn?: string
  nameZh?: string
  operatingHours?: string
  region?: string
  remarks?: string
}

