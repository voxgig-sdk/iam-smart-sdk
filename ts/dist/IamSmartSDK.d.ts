import { MobileRegistrationPointEntity } from './entity/MobileRegistrationPointEntity';
import { RegistrationServiceCounterEntity } from './entity/RegistrationServiceCounterEntity';
import { SelfRegistrationKioskEntity } from './entity/SelfRegistrationKioskEntity';
export type * from './IamSmartTypes';
import { inspect } from 'node:util';
import type { Context, Feature } from './types';
import { config } from './Config';
import { IamSmartEntityBase } from './IamSmartEntityBase';
import { Utility } from './utility/Utility';
import { BaseFeature } from './feature/base/BaseFeature';
declare const stdutil: Utility;
declare class IamSmartSDK {
    _mode: string;
    _options: any;
    _utility: Utility;
    _features: Feature[];
    _rootctx: Context;
    constructor(options?: any);
    options(): any;
    utility(): any;
    prepare(fetchargs?: any): Promise<any>;
    direct(fetchargs?: any): Promise<Error | {
        ok: boolean;
        status: number;
        headers: any;
        data: any;
        err?: undefined;
    } | {
        ok: boolean;
        err: any;
        status?: undefined;
        headers?: undefined;
        data?: undefined;
    }>;
    _rawRequest(fetchargs?: any): Promise<Error | {
        ok: boolean;
        status: number;
        headers: any;
        data: any;
        err?: undefined;
    } | {
        ok: boolean;
        err: any;
        status?: undefined;
        headers?: undefined;
        data?: undefined;
    }>;
    graphql(query: string, variables?: any, ctrl?: any): Promise<any>;
    MobileRegistrationPoint(entopts?: Record<string, any>): MobileRegistrationPointEntity;
    RegistrationServiceCounter(entopts?: Record<string, any>): RegistrationServiceCounterEntity;
    SelfRegistrationKiosk(entopts?: Record<string, any>): SelfRegistrationKioskEntity;
    static test(testoptsarg?: any, sdkoptsarg?: any): IamSmartSDK;
    tester(testopts?: any, sdkopts?: any): IamSmartSDK;
    toJSON(): {
        name: string;
    };
    toString(): string;
    [inspect.custom](): string;
}
declare const SDK: typeof IamSmartSDK;
export { stdutil, config, BaseFeature, IamSmartEntityBase, IamSmartSDK, SDK, };
