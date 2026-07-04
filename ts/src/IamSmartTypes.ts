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
  location_en?: string
  location_zh?: string
  longitude?: number
  name?: string
  name_en?: string
  name_zh?: string
  region?: string
  remark?: string
  schedule?: any[]
}

export type MobileRegistrationPointListMatch = Partial<MobileRegistrationPoint>

export interface RegistrationServiceCounter {
  address?: string
  address_en?: string
  address_zh?: string
  district?: string
  id?: string
  latitude?: number
  longitude?: number
  name?: string
  name_en?: string
  name_zh?: string
  operating_hour?: string
  region?: string
  remark?: string
  service?: any[]
  telephone?: string
}

export type RegistrationServiceCounterListMatch = Partial<RegistrationServiceCounter>

export interface SelfRegistrationKiosk {
  address?: string
  address_en?: string
  address_zh?: string
  availability?: string
  district?: string
  floor?: string
  id?: string
  latitude?: number
  longitude?: number
  name?: string
  name_en?: string
  name_zh?: string
  operating_hour?: string
  region?: string
  remark?: string
}

export type SelfRegistrationKioskListMatch = Partial<SelfRegistrationKiosk>

