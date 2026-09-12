import { IamSmartEntityBase } from '../IamSmartEntityBase';
import type { IamSmartSDK } from '../IamSmartSDK';
import type { Control } from '../types';
import type { SelfRegistrationKiosk, SelfRegistrationKioskListMatch } from '../IamSmartTypes';
declare class SelfRegistrationKioskEntity extends IamSmartEntityBase<SelfRegistrationKiosk> {
    constructor(client: IamSmartSDK, entopts: any);
    make(this: SelfRegistrationKioskEntity): SelfRegistrationKioskEntity;
    list(this: any, reqmatch?: SelfRegistrationKioskListMatch, ctrl?: Control): Promise<SelfRegistrationKioskEntity[]>;
}
export { SelfRegistrationKioskEntity };
