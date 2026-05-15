package core

var UtilityRegistrar func(u *Utility)

var NewBaseFeatureFunc func() Feature

var NewTestFeatureFunc func() Feature

var NewMobileRegistrationPointEntityFunc func(client *IamSmartSDK, entopts map[string]any) IamSmartEntity

var NewRegistrationServiceCounterEntityFunc func(client *IamSmartSDK, entopts map[string]any) IamSmartEntity

var NewSelfRegistrationKioskEntityFunc func(client *IamSmartSDK, entopts map[string]any) IamSmartEntity

