
const envlocal = __dirname + '/../../../.env.local'
require('dotenv').config({ quiet: true, path: [envlocal] })

import Path from 'node:path'
import * as Fs from 'node:fs'

import { test, describe, afterEach } from 'node:test'
import assert from 'node:assert'


import { IamSmartSDK, BaseFeature, stdutil } from '../../..'

import {
  envOverride,
  liveDelay,
  makeCtrl,
  makeMatch,
  makeReqdata,
  makeStepData,
  makeValid,
  maybeSkipControl,
} from '../../utility'


describe('MobileRegistrationPointEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when IAMSMART_TEST_LIVE=TRUE.
  afterEach(liveDelay('IAMSMART_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = IamSmartSDK.test()
    const ent = testsdk.MobileRegistrationPoint()
    assert(null != ent)
  })


  test('basic', async (t) => {

    const live = 'TRUE' === process.env.IAM_SMART_TEST_LIVE
    for (const op of ['list']) {
      if (maybeSkipControl(t, 'entityOp', 'mobile_registration_point.' + op, live)) return
    }

    const setup = basicSetup()
    // The basic flow consumes synthetic IDs and field values from the
    // fixture (entity TestData.json). Those don't exist on the live API.
    // Skip live runs unless the user provided a real ENTID env override.
    if (setup.syntheticOnly) {
      t.skip('live entity test uses synthetic IDs from fixture — set IAM_SMART_TEST_MOBILE_REGISTRATION_POINT_ENTID JSON to run live')
      return
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select

    let mobile_registration_point_ref01_data = Object.values(setup.data.existing.mobile_registration_point)[0] as any

    // LIST
    const mobile_registration_point_ref01_ent = client.MobileRegistrationPoint()
    const mobile_registration_point_ref01_match: any = {}

    const mobile_registration_point_ref01_list = await mobile_registration_point_ref01_ent.list(mobile_registration_point_ref01_match)


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

  // Detect whether the user provided a real ENTID JSON via env var. The
  // basic flow consumes synthetic IDs from the fixture file; without an
  // override those synthetic IDs reach the live API and 4xx. Surface this
  // to the test so it can skip rather than fail.
  const idmapEnvVal = process.env['IAM_SMART_TEST_MOBILE_REGISTRATION_POINT_ENTID']
  const idmapOverridden = null != idmapEnvVal && idmapEnvVal.trim().startsWith('{')

  const env = envOverride({
    'IAM_SMART_TEST_MOBILE_REGISTRATION_POINT_ENTID': idmap,
    'IAM_SMART_TEST_LIVE': 'FALSE',
    'IAM_SMART_TEST_EXPLAIN': 'FALSE',
  })

  idmap = env['IAM_SMART_TEST_MOBILE_REGISTRATION_POINT_ENTID']

  const live = 'TRUE' === env.IAM_SMART_TEST_LIVE

  if (live) {
    client = new IamSmartSDK(merge([
      {
      },
      extra
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
    syntheticOnly: live && !idmapOverridden,
    now: Date.now(),
  }

  return setup
}
  
