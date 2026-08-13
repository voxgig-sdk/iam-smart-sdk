# IamSmart SDK utility: make_context

from iamsmart_sdk.core.context import IamSmartContext


def make_context_util(ctxmap, basectx):
    return IamSmartContext(ctxmap, basectx)
