

import Path from 'node:path'
import * as Fs from 'node:fs'

import { test, describe, afterEach } from 'node:test'
import assert from 'node:assert'
import { createLiveTransport } from '../../live-runner'
import { runLiveEntity } from '../../live-entity'


import { IamSmartSDK, BaseFeature, stdutil } from '../../..'

import {
  envOverride,
  liveClientOptions,
  liveDelay,
  loadEnvLocal,
  makeCtrl,
  makeMatch,
  makeReqdata,
  makeStepData,
  makeValid,
  maybeSkipControl,
} from '../../utility'


// AFTER the imports on purpose: TypeScript hoists `import` above any
// statement in the emitted CommonJS, so a loader placed above them would
// run only after every imported module had already been evaluated - and
// anything reading process.env at module scope would miss these values.
loadEnvLocal(__dirname + '/../../../.env.local')


describe('MobileRegistrationPointEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when IAM_SMART_TEST_LIVE=TRUE.
  afterEach(liveDelay('IAM_SMART_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = IamSmartSDK.test()
    const ent = testsdk.MobileRegistrationPoint()
    assert(null != ent)
  })


  test('basic', async (t) => {

    const live = 'TRUE' === process.env.IAM_SMART_TEST_LIVE
    for (const op of ['list']) {
      if (!live && maybeSkipControl(t, 'entityOp', 'mobile_registration_point.' + op, live)) return
    }

    
    const setup = basicSetup()
    if (setup.live) {
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":[{"active":true,"name":"district","req":false,"short":"District where the mobile point operates","type":"`$STRING`","index$":0},{"active":true,"name":"id","req":false,"short":"Unique identifier for the mobile registration point","type":"`$STRING`","index$":1},{"active":true,"format":"double","name":"latitude","req":false,"short":"Latitude coordinate","type":"`$NUMBER`","index$":2},{"active":true,"name":"location","req":false,"short":"Location description of the mobile point","type":"`$STRING`","index$":3},{"active":true,"name":"locationEn","req":false,"short":"English location description","type":"`$STRING`","index$":4},{"active":true,"name":"locationZh","req":false,"short":"Chinese location description","type":"`$STRING`","index$":5},{"active":true,"format":"double","name":"longitude","req":false,"short":"Longitude coordinate","type":"`$NUMBER`","index$":6},{"active":true,"name":"name","req":false,"short":"Name of the mobile registration point location","type":"`$STRING`","index$":7},{"active":true,"name":"nameEn","req":false,"short":"English name of the mobile registration point","type":"`$STRING`","index$":8},{"active":true,"name":"nameZh","req":false,"short":"Chinese name of the mobile registration point","type":"`$STRING`","index$":9},{"active":true,"name":"region","req":false,"short":"Region (Hong Kong Island, Kowloon, New Territories)","type":"`$STRING`","index$":10},{"active":true,"name":"remarks","req":false,"short":"Additional remarks or notes","type":"`$STRING`","index$":11},{"active":true,"name":"schedule","req":false,"short":"Schedule of mobile registration point visits","type":"`$ARRAY`","index$":12}],"id":{"field":"id","name":"id"},"name":"mobile_registration_point","op":{"list":{"input":"data","name":"list","points":[{"active":true,"args":{},"contract":{"id":"GET /open_data/iam_smart/mobile-registration-points","json":"{\"operationId\":\"getMobileRegistrationPoints\",\"parameters\":[],\"protocol\":\"http\",\"responses\":{\"200\":{\"content\":{\"application/json\":{\"schema\":{\"items\":{\"description\":\"Information about a mobile registration point location\",\"properties\":{\"district\":{\"description\":\"District where the mobile point operates\",\"type\":\"string\"},\"id\":{\"description\":\"Unique identifier for the mobile registration point\",\"type\":\"string\"},\"latitude\":{\"description\":\"Latitude coordinate\",\"format\":\"double\",\"type\":\"number\"},\"location\":{\"description\":\"Location description of the mobile point\",\"type\":\"string\"},\"locationEn\":{\"description\":\"English location description\",\"type\":\"string\"},\"locationZh\":{\"description\":\"Chinese location description\",\"type\":\"string\"},\"longitude\":{\"description\":\"Longitude coordinate\",\"format\":\"double\",\"type\":\"number\"},\"name\":{\"description\":\"Name of the mobile registration point location\",\"type\":\"string\"},\"nameEn\":{\"description\":\"English name of the mobile registration point\",\"type\":\"string\"},\"nameZh\":{\"description\":\"Chinese name of the mobile registration point\",\"type\":\"string\"},\"region\":{\"description\":\"Region (Hong Kong Island, Kowloon, New Territories)\",\"enum\":[\"Hong Kong Island\",\"Kowloon\",\"New Territories\"],\"type\":\"string\"},\"remarks\":{\"description\":\"Additional remarks or notes\",\"type\":\"string\"},\"schedule\":{\"description\":\"Schedule of mobile registration point visits\",\"items\":{\"properties\":{\"date\":{\"description\":\"Date of service\",\"format\":\"date\",\"type\":\"string\"},\"endTime\":{\"description\":\"End time of service\",\"type\":\"string\"},\"startTime\":{\"description\":\"Start time of service\",\"type\":\"string\"}},\"type\":\"object\"},\"type\":\"array\"}},\"type\":\"object\"},\"type\":\"array\"}}},\"description\":\"Successful response containing mobile registration point locations\"},\"500\":{\"content\":{\"application/json\":{\"schema\":{\"description\":\"Error response object\",\"properties\":{\"code\":{\"description\":\"Error code\",\"type\":\"string\"},\"details\":{\"description\":\"Additional error details\",\"type\":\"string\"},\"message\":{\"description\":\"Error message describing what went wrong\",\"type\":\"string\"}},\"required\":[\"code\",\"message\"],\"type\":\"object\"}}},\"description\":\"Internal server error\"}},\"securitySource\":\"unspecified\"}","source":"openapi3","version":1},"kind":"http","method":"GET","orig":"/open_data/iam_smart/mobile-registration-points","segments":[{"lit":"open_data"},{"lit":"iam_smart"},{"lit":"mobile-registration-points"}],"select":{},"transform":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"list"}},"relations":{"ancestors":[]},"key$":"mobile_registration_point","name__orig":"mobile_registration_point","Name":"MobileRegistrationPoint","name_":"mobile_registration_point","name-":"mobile-registration-point","NAME":"MOBILE_REGISTRATION_POINT","index$":0}, {"active":true,"entity":"mobile_registration_point","key$":"BasicMobileRegistrationPointFlow","kind":"basic","name":"BasicMobileRegistrationPointFlow","param":{},"step":[{"active":true,"data":{},"input":{},"match":{},"op":"list","spec":[],"valid":[{"apply":"ItemExists","def":{"ref":"mobile_registration_point_ref01"}}],"index$":0}]}, 'MobileRegistrationPoint')
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select

    let mobile_registration_point_ref01_data = Object.values(setup.data.existing.mobile_registration_point)[0] as any

    // LIST
    const mobile_registration_point_ref01_ent = client.MobileRegistrationPoint()
    const mobile_registration_point_ref01_match: any = {}

    const mobile_registration_point_ref01_list = (await mobile_registration_point_ref01_ent.list(mobile_registration_point_ref01_match)).map((e: any) => e.data())


  })
})



function basicSetup(extra?: any) {
  // TODO: fix test def options
  const options: any = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname, 
      '../../../../.sdk/test/entity/mobile_registration_point/MobileRegistrationPointTestData.json')

  // TODO: file ready util needed?
  const entityDataSource = Fs.readFileSync(entityDataFile).toString('utf8')

  // TODO: need a xlang JSON parse utility in voxgig/struct with better error msgs
  const entityData = JSON.parse(entityDataSource)

  options.entity = entityData.existing

  let client = IamSmartSDK.test(options, extra)
  const struct = client.utility().struct
  const merge = struct.merge
  const transform = struct.transform

  let idmap = transform(
    ['mobile_registration_point01','mobile_registration_point02','mobile_registration_point03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'IAM_SMART_TEST_MOBILE_REGISTRATION_POINT_ENTID': idmap,
    'IAM_SMART_TEST_LIVE': 'FALSE',
    'IAM_SMART_TEST_EXPLAIN': 'FALSE',
  })

  idmap = env['IAM_SMART_TEST_MOBILE_REGISTRATION_POINT_ENTID']

  const live = 'TRUE' === env.IAM_SMART_TEST_LIVE

  const transport = createLiveTransport()
  if (live) {
    const rawIds = process.env['IAM_SMART_TEST_MOBILE_REGISTRATION_POINT_ENTID']
    idmap = rawIds && rawIds.trim() ? JSON.parse(rawIds) : {}
    if (!idmap || Array.isArray(idmap) || typeof idmap !== 'object') {
      throw new Error('Live ENTID must be a JSON object')
    }
    client = new IamSmartSDK(merge([
      // FIRST, so the generated fields below win: sdk-test-control.json's
      // test.client.options adds to the live client, it does not redirect it.
      liveClientOptions(),
      {
      },
      // 'extra || {}', not a bare 'extra': struct.merge returns UNDEFINED when the
      // last entry is undefined, and basicSetup is normally called with no
      // argument at all - so a bare 'extra' silently discarded the apikey
      // and server values above and handed the SDK undefined. Harmless
      // while there was nothing in that object; not harmless now.
      extra || {},
      { system: { fetch: transport.fetch } }
    ]))
  }

  const setup = {
    idmap,
    env,
    options,
    client,
    struct,
    data: entityData,
    explain: 'TRUE' === env.IAM_SMART_TEST_EXPLAIN,
    live,
    transport,
    now: Date.now(),
  }

  return setup
}
  
