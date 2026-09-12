import { IamSmartEntityBase } from '../IamSmartEntityBase';
import type { IamSmartSDK } from '../IamSmartSDK';
import type { Control } from '../types';
import type { RegistrationServiceCounter, RegistrationServiceCounterListMatch } from '../IamSmartTypes';
declare class RegistrationServiceCounterEntity extends IamSmartEntityBase<RegistrationServiceCounter> {
    constructor(client: IamSmartSDK, entopts: any);
    make(this: RegistrationServiceCounterEntity): RegistrationServiceCounterEntity;
    list(this: any, reqmatch?: RegistrationServiceCounterListMatch, ctrl?: Control): Promise<RegistrationServiceCounterEntity[]>;
}
export { RegistrationServiceCounterEntity };
