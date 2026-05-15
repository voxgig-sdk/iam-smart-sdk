
import { test, describe } from 'node:test'
import { equal } from 'node:assert'


import { IamSmartSDK } from '..'


describe('exists', async () => {

  test('test-mode', async () => {
    const testsdk = await IamSmartSDK.test()
    equal(null !== testsdk, true)
  })

})
