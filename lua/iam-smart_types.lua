-- Typed models for the IamSmart SDK (LuaLS annotations).
--
-- GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
-- params (op.<name>.points[].args.params[]). Field/param types come from the
-- canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
-- @voxgig/apidef VALID_CANON). Annotations only — no runtime effect. Do not
-- edit by hand.

---@class MobileRegistrationPoint
---@field district? string
---@field id? string
---@field latitude? number
---@field location? string
---@field location_en? string
---@field location_zh? string
---@field longitude? number
---@field name? string
---@field name_en? string
---@field name_zh? string
---@field region? string
---@field remark? string
---@field schedule? table

---@class MobileRegistrationPointListMatch

---@class RegistrationServiceCounter
---@field address? string
---@field address_en? string
---@field address_zh? string
---@field district? string
---@field id? string
---@field latitude? number
---@field longitude? number
---@field name? string
---@field name_en? string
---@field name_zh? string
---@field operating_hour? string
---@field region? string
---@field remark? string
---@field service? table
---@field telephone? string

---@class RegistrationServiceCounterListMatch

---@class SelfRegistrationKiosk
---@field address? string
---@field address_en? string
---@field address_zh? string
---@field availability? string
---@field district? string
---@field floor? string
---@field id? string
---@field latitude? number
---@field longitude? number
---@field name? string
---@field name_en? string
---@field name_zh? string
---@field operating_hour? string
---@field region? string
---@field remark? string

---@class SelfRegistrationKioskListMatch

local M = {}

return M
