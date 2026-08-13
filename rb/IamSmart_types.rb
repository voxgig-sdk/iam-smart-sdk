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
# @!attribute [rw] locationEn
#   @return [String, nil]
#
# @!attribute [rw] locationZh
#   @return [String, nil]
#
# @!attribute [rw] longitude
#   @return [Float, nil]
#
# @!attribute [rw] name
#   @return [String, nil]
#
# @!attribute [rw] nameEn
#   @return [String, nil]
#
# @!attribute [rw] nameZh
#   @return [String, nil]
#
# @!attribute [rw] region
#   @return [String, nil]
#
# @!attribute [rw] remarks
#   @return [String, nil]
#
# @!attribute [rw] schedule
#   @return [Array, nil]
MobileRegistrationPoint = Struct.new(
  :district,
  :id,
  :latitude,
  :location,
  :locationEn,
  :locationZh,
  :longitude,
  :name,
  :nameEn,
  :nameZh,
  :region,
  :remarks,
  :schedule,
  keyword_init: true
)

# Request payload for MobileRegistrationPoint#list.
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
# @!attribute [rw] locationEn
#   @return [String, nil]
#
# @!attribute [rw] locationZh
#   @return [String, nil]
#
# @!attribute [rw] longitude
#   @return [Float, nil]
#
# @!attribute [rw] name
#   @return [String, nil]
#
# @!attribute [rw] nameEn
#   @return [String, nil]
#
# @!attribute [rw] nameZh
#   @return [String, nil]
#
# @!attribute [rw] region
#   @return [String, nil]
#
# @!attribute [rw] remarks
#   @return [String, nil]
#
# @!attribute [rw] schedule
#   @return [Array, nil]
MobileRegistrationPointListMatch = Struct.new(
  :district,
  :id,
  :latitude,
  :location,
  :locationEn,
  :locationZh,
  :longitude,
  :name,
  :nameEn,
  :nameZh,
  :region,
  :remarks,
  :schedule,
  keyword_init: true
)

# RegistrationServiceCounter entity data model.
#
# @!attribute [rw] address
#   @return [String, nil]
#
# @!attribute [rw] addressEn
#   @return [String, nil]
#
# @!attribute [rw] addressZh
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
# @!attribute [rw] nameEn
#   @return [String, nil]
#
# @!attribute [rw] nameZh
#   @return [String, nil]
#
# @!attribute [rw] operatingHours
#   @return [String, nil]
#
# @!attribute [rw] region
#   @return [String, nil]
#
# @!attribute [rw] remarks
#   @return [String, nil]
#
# @!attribute [rw] services
#   @return [Array, nil]
#
# @!attribute [rw] telephone
#   @return [String, nil]
RegistrationServiceCounter = Struct.new(
  :address,
  :addressEn,
  :addressZh,
  :district,
  :id,
  :latitude,
  :longitude,
  :name,
  :nameEn,
  :nameZh,
  :operatingHours,
  :region,
  :remarks,
  :services,
  :telephone,
  keyword_init: true
)

# Request payload for RegistrationServiceCounter#list.
#
# @!attribute [rw] address
#   @return [String, nil]
#
# @!attribute [rw] addressEn
#   @return [String, nil]
#
# @!attribute [rw] addressZh
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
# @!attribute [rw] nameEn
#   @return [String, nil]
#
# @!attribute [rw] nameZh
#   @return [String, nil]
#
# @!attribute [rw] operatingHours
#   @return [String, nil]
#
# @!attribute [rw] region
#   @return [String, nil]
#
# @!attribute [rw] remarks
#   @return [String, nil]
#
# @!attribute [rw] services
#   @return [Array, nil]
#
# @!attribute [rw] telephone
#   @return [String, nil]
RegistrationServiceCounterListMatch = Struct.new(
  :address,
  :addressEn,
  :addressZh,
  :district,
  :id,
  :latitude,
  :longitude,
  :name,
  :nameEn,
  :nameZh,
  :operatingHours,
  :region,
  :remarks,
  :services,
  :telephone,
  keyword_init: true
)

# SelfRegistrationKiosk entity data model.
#
# @!attribute [rw] address
#   @return [String, nil]
#
# @!attribute [rw] addressEn
#   @return [String, nil]
#
# @!attribute [rw] addressZh
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
# @!attribute [rw] nameEn
#   @return [String, nil]
#
# @!attribute [rw] nameZh
#   @return [String, nil]
#
# @!attribute [rw] operatingHours
#   @return [String, nil]
#
# @!attribute [rw] region
#   @return [String, nil]
#
# @!attribute [rw] remarks
#   @return [String, nil]
SelfRegistrationKiosk = Struct.new(
  :address,
  :addressEn,
  :addressZh,
  :availability,
  :district,
  :floor,
  :id,
  :latitude,
  :longitude,
  :name,
  :nameEn,
  :nameZh,
  :operatingHours,
  :region,
  :remarks,
  keyword_init: true
)

# Request payload for SelfRegistrationKiosk#list.
#
# @!attribute [rw] address
#   @return [String, nil]
#
# @!attribute [rw] addressEn
#   @return [String, nil]
#
# @!attribute [rw] addressZh
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
# @!attribute [rw] nameEn
#   @return [String, nil]
#
# @!attribute [rw] nameZh
#   @return [String, nil]
#
# @!attribute [rw] operatingHours
#   @return [String, nil]
#
# @!attribute [rw] region
#   @return [String, nil]
#
# @!attribute [rw] remarks
#   @return [String, nil]
SelfRegistrationKioskListMatch = Struct.new(
  :address,
  :addressEn,
  :addressZh,
  :availability,
  :district,
  :floor,
  :id,
  :latitude,
  :longitude,
  :name,
  :nameEn,
  :nameZh,
  :operatingHours,
  :region,
  :remarks,
  keyword_init: true
)

