import { Context } from './Context';
declare class IamSmartError extends Error {
    isIamSmartError: boolean;
    sdk: string;
    code: string;
    ctx: Context;
    status: number;
    get notFound(): boolean;
    constructor(code: string, msg: string, ctx: Context);
}
export { IamSmartError };
