import { IamSmartEntityBase } from '../IamSmartEntityBase';
import type { IamSmartSDK } from '../IamSmartSDK';
import type { Control } from '../types';
import type { MobileRegistrationPoint, MobileRegistrationPointListMatch } from '../IamSmartTypes';
declare class MobileRegistrationPointEntity extends IamSmartEntityBase<MobileRegistrationPoint> {
    constructor(client: IamSmartSDK, entopts: any);
    make(this: MobileRegistrationPointEntity): MobileRegistrationPointEntity;
    list(this: any, reqmatch?: MobileRegistrationPointListMatch, ctrl?: Control): Promise<MobileRegistrationPointEntity[]>;
}
export { MobileRegistrationPointEntity };
