# IamSmart SDK utility: make_context

from core.context import IamSmartContext


def make_context_util(ctxmap, basectx):
    return IamSmartContext(ctxmap, basectx)
