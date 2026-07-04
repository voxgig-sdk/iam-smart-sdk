# frozen_string_literal: true

# Typed models for the IamSmart SDK.
#
# GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
# params (op.<name>.points[].args.params[]). Member types come from the
# canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
# @voxgig/apidef VALID_CANON). Ruby types are unenforced; these YARD
# annotations document the shapes. Do not edit by hand.

# MobileRegistrationPoint entity data model.
#
# @!attribute [rw] district
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] latitude
#   @return [Float, nil]
#
# @!attribute [rw] location
#   @return [String, nil]
#
# @!attribute [rw] location_en
#   @return [String, nil]
#
# @!attribute [rw] location_zh
#   @return [String, nil]
#
# @!attribute [rw] longitude
#   @return [Float, nil]
#
# @!attribute [rw] name
#   @return [String, nil]
#
# @!attribute [rw] name_en
#   @return [String, nil]
#
# @!attribute [rw] name_zh
#   @return [String, nil]
#
# @!attribute [rw] region
#   @return [String, nil]
#
# @!attribute [rw] remark
#   @return [String, nil]
#
# @!attribute [rw] schedule
#   @return [Array, nil]
MobileRegistrationPoint = Struct.new(
  :district,
  :id,
  :latitude,
  :location,
  :location_en,
  :location_zh,
  :longitude,
  :name,
  :name_en,
  :name_zh,
  :region,
  :remark,
  :schedule,
  keyword_init: true
)

# Match filter for MobileRegistrationPoint#list (any subset of MobileRegistrationPoint fields).
#
# @!attribute [rw] district
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] latitude
#   @return [Float, nil]
#
# @!attribute [rw] location
#   @return [String, nil]
#
# @!attribute [rw] location_en
#   @return [String, nil]
#
# @!attribute [rw] location_zh
#   @return [String, nil]
#
# @!attribute [rw] longitude
#   @return [Float, nil]
#
# @!attribute [rw] name
#   @return [String, nil]
#
# @!attribute [rw] name_en
#   @return [String, nil]
#
# @!attribute [rw] name_zh
#   @return [String, nil]
#
# @!attribute [rw] region
#   @return [String, nil]
#
# @!attribute [rw] remark
#   @return [String, nil]
#
# @!attribute [rw] schedule
#   @return [Array, nil]
MobileRegistrationPointListMatch = Struct.new(
  :district,
  :id,
  :latitude,
  :location,
  :location_en,
  :location_zh,
  :longitude,
  :name,
  :name_en,
  :name_zh,
  :region,
  :remark,
  :schedule,
  keyword_init: true
)

# RegistrationServiceCounter entity data model.
#
# @!attribute [rw] address
#   @return [String, nil]
#
# @!attribute [rw] address_en
#   @return [String, nil]
#
# @!attribute [rw] address_zh
#   @return [String, nil]
#
# @!attribute [rw] district
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] latitude
#   @return [Float, nil]
#
# @!attribute [rw] longitude
#   @return [Float, nil]
#
# @!attribute [rw] name
#   @return [String, nil]
#
# @!attribute [rw] name_en
#   @return [String, nil]
#
# @!attribute [rw] name_zh
#   @return [String, nil]
#
# @!attribute [rw] operating_hour
#   @return [String, nil]
#
# @!attribute [rw] region
#   @return [String, nil]
#
# @!attribute [rw] remark
#   @return [String, nil]
#
# @!attribute [rw] service
#   @return [Array, nil]
#
# @!attribute [rw] telephone
#   @return [String, nil]
RegistrationServiceCounter = Struct.new(
  :address,
  :address_en,
  :address_zh,
  :district,
  :id,
  :latitude,
  :longitude,
  :name,
  :name_en,
  :name_zh,
  :operating_hour,
  :region,
  :remark,
  :service,
  :telephone,
  keyword_init: true
)

# Match filter for RegistrationServiceCounter#list (any subset of RegistrationServiceCounter fields).
#
# @!attribute [rw] address
#   @return [String, nil]
#
# @!attribute [rw] address_en
#   @return [String, nil]
#
# @!attribute [rw] address_zh
#   @return [String, nil]
#
# @!attribute [rw] district
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] latitude
#   @return [Float, nil]
#
# @!attribute [rw] longitude
#   @return [Float, nil]
#
# @!attribute [rw] name
#   @return [String, nil]
#
# @!attribute [rw] name_en
#   @return [String, nil]
#
# @!attribute [rw] name_zh
#   @return [String, nil]
#
# @!attribute [rw] operating_hour
#   @return [String, nil]
#
# @!attribute [rw] region
#   @return [String, nil]
#
# @!attribute [rw] remark
#   @return [String, nil]
#
# @!attribute [rw] service
#   @return [Array, nil]
#
# @!attribute [rw] telephone
#   @return [String, nil]
RegistrationServiceCounterListMatch = Struct.new(
  :address,
  :address_en,
  :address_zh,
  :district,
  :id,
  :latitude,
  :longitude,
  :name,
  :name_en,
  :name_zh,
  :operating_hour,
  :region,
  :remark,
  :service,
  :telephone,
  keyword_init: true
)

# SelfRegistrationKiosk entity data model.
#
# @!attribute [rw] address
#   @return [String, nil]
#
# @!attribute [rw] address_en
#   @return [String, nil]
#
# @!attribute [rw] address_zh
#   @return [String, nil]
#
# @!attribute [rw] availability
#   @return [String, nil]
#
# @!attribute [rw] district
#   @return [String, nil]
#
# @!attribute [rw] floor
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] latitude
#   @return [Float, nil]
#
# @!attribute [rw] longitude
#   @return [Float, nil]
#
# @!attribute [rw] name
#   @return [String, nil]
#
# @!attribute [rw] name_en
#   @return [String, nil]
#
# @!attribute [rw] name_zh
#   @return [String, nil]
#
# @!attribute [rw] operating_hour
#   @return [String, nil]
#
# @!attribute [rw] region
#   @return [String, nil]
#
# @!attribute [rw] remark
#   @return [String, nil]
SelfRegistrationKiosk = Struct.new(
  :address,
  :address_en,
  :address_zh,
  :availability,
  :district,
  :floor,
  :id,
  :latitude,
  :longitude,
  :name,
  :name_en,
  :name_zh,
  :operating_hour,
  :region,
  :remark,
  keyword_init: true
)

# Match filter for SelfRegistrationKiosk#list (any subset of SelfRegistrationKiosk fields).
#
# @!attribute [rw] address
#   @return [String, nil]
#
# @!attribute [rw] address_en
#   @return [String, nil]
#
# @!attribute [rw] address_zh
#   @return [String, nil]
#
# @!attribute [rw] availability
#   @return [String, nil]
#
# @!attribute [rw] district
#   @return [String, nil]
#
# @!attribute [rw] floor
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] latitude
#   @return [Float, nil]
#
# @!attribute [rw] longitude
#   @return [Float, nil]
#
# @!attribute [rw] name
#   @return [String, nil]
#
# @!attribute [rw] name_en
#   @return [String, nil]
#
# @!attribute [rw] name_zh
#   @return [String, nil]
#
# @!attribute [rw] operating_hour
#   @return [String, nil]
#
# @!attribute [rw] region
#   @return [String, nil]
#
# @!attribute [rw] remark
#   @return [String, nil]
SelfRegistrationKioskListMatch = Struct.new(
  :address,
  :address_en,
  :address_zh,
  :availability,
  :district,
  :floor,
  :id,
  :latitude,
  :longitude,
  :name,
  :name_en,
  :name_zh,
  :operating_hour,
  :region,
  :remark,
  keyword_init: true
)

