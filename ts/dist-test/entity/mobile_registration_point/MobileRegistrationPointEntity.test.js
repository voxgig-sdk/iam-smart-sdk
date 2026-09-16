"use strict";
var __createBinding = (this && this.__createBinding) || (Object.create ? (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    var desc = Object.getOwnPropertyDescriptor(m, k);
    if (!desc || ("get" in desc ? !m.__esModule : desc.writable || desc.configurable)) {
      desc = { enumerable: true, get: function() { return m[k]; } };
    }
    Object.defineProperty(o, k2, desc);
}) : (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    o[k2] = m[k];
}));
var __setModuleDefault = (this && this.__setModuleDefault) || (Object.create ? (function(o, v) {
    Object.defineProperty(o, "default", { enumerable: true, value: v });
}) : function(o, v) {
    o["default"] = v;
});
var __importStar = (this && this.__importStar) || (function () {
    var ownKeys = function(o) {
        ownKeys = Object.getOwnPropertyNames || function (o) {
            var ar = [];
            for (var k in o) if (Object.prototype.hasOwnProperty.call(o, k)) ar[ar.length] = k;
            return ar;
        };
        return ownKeys(o);
    };
    return function (mod) {
        if (mod && mod.__esModule) return mod;
        var result = {};
        if (mod != null) for (var k = ownKeys(mod), i = 0; i < k.length; i++) if (k[i] !== "default") __createBinding(result, mod, k[i]);
        __setModuleDefault(result, mod);
        return result;
    };
})();
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const node_path_1 = __importDefault(require("node:path"));
const Fs = __importStar(require("node:fs"));
const node_test_1 = require("node:test");
const node_assert_1 = __importDefault(require("node:assert"));
const live_runner_1 = require("../../live-runner");
const live_entity_1 = require("../../live-entity");
const __1 = require("../../..");
const utility_1 = require("../../utility");
// AFTER the imports on purpose: TypeScript hoists `import` above any
// statement in the emitted CommonJS, so a loader placed above them would
// run only after every imported module had already been evaluated - and
// anything reading process.env at module scope would miss these values.
(0, utility_1.loadEnvLocal)(__dirname + '/../../../.env.local');
(0, node_test_1.describe)('MobileRegistrationPointEntity', async () => {
    // Per-test live pacing. Delay is read from sdk-test-control.json's
    // `test.live.delayMs`; only sleeps when IAM_SMART_TEST_LIVE=TRUE.
    (0, node_test_1.afterEach)((0, utility_1.liveDelay)('IAM_SMART_TEST_LIVE'));
    (0, node_test_1.test)('instance', async () => {
        const testsdk = __1.IamSmartSDK.test();
        const ent = testsdk.MobileRegistrationPoint();
        (0, node_assert_1.default)(null != ent);
    });
    (0, node_test_1.test)('basic', async (t) => {
        const live = 'TRUE' === process.env.IAM_SMART_TEST_LIVE;
        for (const op of ['list']) {
            if (!live && (0, utility_1.maybeSkipControl)(t, 'entityOp', 'mobile_registration_point.' + op, live))
                return;
        }
        const setup = basicSetup();
        if (setup.live) {
            return (0, live_entity_1.runLiveEntity)(setup, { "active": true, "alias": { "field": {} }, "fields": [{ "active": true, "name": "district", "req": false, "short": "District where the mobile point operates", "type": "`$STRING`", "index$": 0 }, { "active": true, "name": "id", "req": false, "short": "Unique identifier for the mobile registration point", "type": "`$STRING`", "index$": 1 }, { "active": true, "format": "double", "name": "latitude", "req": false, "short": "Latitude coordinate", "type": "`$NUMBER`", "index$": 2 }, { "active": true, "name": "location", "req": false, "short": "Location description of the mobile point", "type": "`$STRING`", "index$": 3 }, { "active": true, "name": "locationEn", "req": false, "short": "English location description", "type": "`$STRING`", "index$": 4 }, { "active": true, "name": "locationZh", "req": false, "short": "Chinese location description", "type": "`$STRING`", "index$": 5 }, { "active": true, "format": "double", "name": "longitude", "req": false, "short": "Longitude coordinate", "type": "`$NUMBER`", "index$": 6 }, { "active": true, "name": "name", "req": false, "short": "Name of the mobile registration point location", "type": "`$STRING`", "index$": 7 }, { "active": true, "name": "nameEn", "req": false, "short": "English name of the mobile registration point", "type": "`$STRING`", "index$": 8 }, { "active": true, "name": "nameZh", "req": false, "short": "Chinese name of the mobile registration point", "type": "`$STRING`", "index$": 9 }, { "active": true, "name": "region", "req": false, "short": "Region (Hong Kong Island, Kowloon, New Territories)", "type": "`$STRING`", "index$": 10 }, { "active": true, "name": "remarks", "req": false, "short": "Additional remarks or notes", "type": "`$STRING`", "index$": 11 }, { "active": true, "name": "schedule", "req": false, "short": "Schedule of mobile registration point visits", "type": "`$ARRAY`", "index$": 12 }], "id": { "field": "id", "name": "id" }, "name": "mobile_registration_point", "op": { "list": { "input": "data", "name": "list", "points": [{ "active": true, "args": {}, "contract": { "id": "GET /open_data/iam_smart/mobile-registration-points", "json": "{\"operationId\":\"getMobileRegistrationPoints\",\"parameters\":[],\"protocol\":\"http\",\"responses\":{\"200\":{\"content\":{\"application/json\":{\"schema\":{\"items\":{\"description\":\"Information about a mobile registration point location\",\"properties\":{\"district\":{\"description\":\"District where the mobile point operates\",\"type\":\"string\"},\"id\":{\"description\":\"Unique identifier for the mobile registration point\",\"type\":\"string\"},\"latitude\":{\"description\":\"Latitude coordinate\",\"format\":\"double\",\"type\":\"number\"},\"location\":{\"description\":\"Location description of the mobile point\",\"type\":\"string\"},\"locationEn\":{\"description\":\"English location description\",\"type\":\"string\"},\"locationZh\":{\"description\":\"Chinese location description\",\"type\":\"string\"},\"longitude\":{\"description\":\"Longitude coordinate\",\"format\":\"double\",\"type\":\"number\"},\"name\":{\"description\":\"Name of the mobile registration point location\",\"type\":\"string\"},\"nameEn\":{\"description\":\"English name of the mobile registration point\",\"type\":\"string\"},\"nameZh\":{\"description\":\"Chinese name of the mobile registration point\",\"type\":\"string\"},\"region\":{\"description\":\"Region (Hong Kong Island, Kowloon, New Territories)\",\"enum\":[\"Hong Kong Island\",\"Kowloon\",\"New Territories\"],\"type\":\"string\"},\"remarks\":{\"description\":\"Additional remarks or notes\",\"type\":\"string\"},\"schedule\":{\"description\":\"Schedule of mobile registration point visits\",\"items\":{\"properties\":{\"date\":{\"description\":\"Date of service\",\"format\":\"date\",\"type\":\"string\"},\"endTime\":{\"description\":\"End time of service\",\"type\":\"string\"},\"startTime\":{\"description\":\"Start time of service\",\"type\":\"string\"}},\"type\":\"object\"},\"type\":\"array\"}},\"type\":\"object\"},\"type\":\"array\"}}},\"description\":\"Successful response containing mobile registration point locations\"},\"500\":{\"content\":{\"application/json\":{\"schema\":{\"description\":\"Error response object\",\"properties\":{\"code\":{\"description\":\"Error code\",\"type\":\"string\"},\"details\":{\"description\":\"Additional error details\",\"type\":\"string\"},\"message\":{\"description\":\"Error message describing what went wrong\",\"type\":\"string\"}},\"required\":[\"code\",\"message\"],\"type\":\"object\"}}},\"description\":\"Internal server error\"}},\"securitySource\":\"unspecified\"}", "source": "openapi3", "version": 1 }, "kind": "http", "method": "GET", "orig": "/open_data/iam_smart/mobile-registration-points", "segments": [{ "lit": "open_data" }, { "lit": "iam_smart" }, { "lit": "mobile-registration-points" }], "select": {}, "transform": { "req": "`reqdata`", "res": "`body`" }, "index$": 0 }], "key$": "list" } }, "relations": { "ancestors": [] }, "key$": "mobile_registration_point", "name__orig": "mobile_registration_point", "Name": "MobileRegistrationPoint", "name_": "mobile_registration_point", "name-": "mobile-registration-point", "NAME": "MOBILE_REGISTRATION_POINT", "index$": 0 }, { "active": true, "entity": "mobile_registration_point", "key$": "BasicMobileRegistrationPointFlow", "kind": "basic", "name": "BasicMobileRegistrationPointFlow", "param": {}, "step": [{ "active": true, "data": {}, "input": {}, "match": {}, "op": "list", "spec": [], "valid": [{ "apply": "ItemExists", "def": { "ref": "mobile_registration_point_ref01" } }], "index$": 0 }] }, 'MobileRegistrationPoint');
        }
        const client = setup.client;
        const struct = setup.struct;
        const isempty = struct.isempty;
        const select = struct.select;
        let mobile_registration_point_ref01_data = Object.values(setup.data.existing.mobile_registration_point)[0];
        // LIST
        const mobile_registration_point_ref01_ent = client.MobileRegistrationPoint();
        const mobile_registration_point_ref01_match = {};
        const mobile_registration_point_ref01_list = (await mobile_registration_point_ref01_ent.list(mobile_registration_point_ref01_match)).map((e) => e.data());
    });
});
function basicSetup(extra) {
    // TODO: fix test def options
    const options = {}; // null
    // TODO: needs test utility to resolve path
    const entityDataFile = node_path_1.default.resolve(__dirname, '../../../../.sdk/test/entity/mobile_registration_point/MobileRegistrationPointTestData.json');
    // TODO: file ready util needed?
    const entityDataSource = Fs.readFileSync(entityDataFile).toString('utf8');
    // TODO: need a xlang JSON parse utility in voxgig/struct with better error msgs
    const entityData = JSON.parse(entityDataSource);
    options.entity = entityData.existing;
    let client = __1.IamSmartSDK.test(options, extra);
    const struct = client.utility().struct;
    const merge = struct.merge;
    const transform = struct.transform;
    let idmap = transform(['mobile_registration_point01', 'mobile_registration_point02', 'mobile_registration_point03'], {
        '`$PACK`': ['', {
                '`$KEY`': '`$COPY`',
                '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
            }]
    });
    const env = (0, utility_1.envOverride)({
        'IAM_SMART_TEST_MOBILE_REGISTRATION_POINT_ENTID': idmap,
        'IAM_SMART_TEST_LIVE': 'FALSE',
        'IAM_SMART_TEST_EXPLAIN': 'FALSE',
    });
    idmap = env['IAM_SMART_TEST_MOBILE_REGISTRATION_POINT_ENTID'];
    const live = 'TRUE' === env.IAM_SMART_TEST_LIVE;
    const transport = (0, live_runner_1.createLiveTransport)();
    if (live) {
        const rawIds = process.env['IAM_SMART_TEST_MOBILE_REGISTRATION_POINT_ENTID'];
        idmap = rawIds && rawIds.trim() ? JSON.parse(rawIds) : {};
        if (!idmap || Array.isArray(idmap) || typeof idmap !== 'object') {
            throw new Error('Live ENTID must be a JSON object');
        }
        client = new __1.IamSmartSDK(merge([
            // FIRST, so the generated fields below win: sdk-test-control.json's
            // test.client.options adds to the live client, it does not redirect it.
            (0, utility_1.liveClientOptions)(),
            {},
            // 'extra || {}', not a bare 'extra': struct.merge returns UNDEFINED when the
            // last entry is undefined, and basicSetup is normally called with no
            // argument at all - so a bare 'extra' silently discarded the apikey
            // and server values above and handed the SDK undefined. Harmless
            // while there was nothing in that object; not harmless now.
            extra || {},
            { system: { fetch: transport.fetch } }
        ]));
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
    };
    return setup;
}
//# sourceMappingURL=MobileRegistrationPointEntity.test.js.map