package voxgigiamsmartsdk

import (
	"github.com/voxgig-sdk/iam-smart-sdk/go/core"
	"github.com/voxgig-sdk/iam-smart-sdk/go/entity"
	"github.com/voxgig-sdk/iam-smart-sdk/go/feature"
	_ "github.com/voxgig-sdk/iam-smart-sdk/go/utility"
)

// Type aliases preserve external API.
type IamSmartSDK = core.IamSmartSDK
type Context = core.Context
type Utility = core.Utility
type Feature = core.Feature
type Entity = core.Entity
type IamSmartEntity = core.IamSmartEntity
type FetcherFunc = core.FetcherFunc
type Spec = core.Spec
type Result = core.Result
type Response = core.Response
type Operation = core.Operation
type Control = core.Control
type IamSmartError = core.IamSmartError

// BaseFeature from feature package.
type BaseFeature = feature.BaseFeature

func init() {
	core.NewBaseFeatureFunc = func() core.Feature {
		return feature.NewBaseFeature()
	}
	core.NewTestFeatureFunc = func() core.Feature {
		return feature.NewTestFeature()
	}
	core.NewMobileRegistrationPointEntityFunc = func(client *core.IamSmartSDK, entopts map[string]any) core.IamSmartEntity {
		return entity.NewMobileRegistrationPointEntity(client, entopts)
	}
	core.NewRegistrationServiceCounterEntityFunc = func(client *core.IamSmartSDK, entopts map[string]any) core.IamSmartEntity {
		return entity.NewRegistrationServiceCounterEntity(client, entopts)
	}
	core.NewSelfRegistrationKioskEntityFunc = func(client *core.IamSmartSDK, entopts map[string]any) core.IamSmartEntity {
		return entity.NewSelfRegistrationKioskEntity(client, entopts)
	}
}

// Constructor re-exports.
var NewIamSmartSDK = core.NewIamSmartSDK
var TestSDK = core.TestSDK
var NewContext = core.NewContext
var NewSpec = core.NewSpec
var NewResult = core.NewResult
var NewResponse = core.NewResponse
var NewOperation = core.NewOperation
var MakeConfig = core.MakeConfig

// No-arg convenience constructors. Go has no default-argument syntax,
// so these aliases let callers write `sdk.New()` / `sdk.Test()`
// instead of `sdk.NewIamSmartSDK(nil)` / `sdk.TestSDK(nil, nil)`
// for the common no-options case.
func New() *IamSmartSDK  { return NewIamSmartSDK(nil) }
func Test() *IamSmartSDK { return TestSDK(nil, nil) }
var NewBaseFeature = feature.NewBaseFeature
var NewTestFeature = feature.NewTestFeature
