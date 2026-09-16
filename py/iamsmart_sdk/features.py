# IamSmart SDK feature factory

from iamsmart_sdk.feature.base_feature import IamSmartBaseFeature
from iamsmart_sdk.feature.ratelimit_feature import IamSmartRatelimitFeature
from iamsmart_sdk.feature.retry_feature import IamSmartRetryFeature
from iamsmart_sdk.feature.test_feature import IamSmartTestFeature
from iamsmart_sdk.feature.timeout_feature import IamSmartTimeoutFeature


_FEATURES = {
    "base": lambda: IamSmartBaseFeature(),
    "ratelimit": lambda: IamSmartRatelimitFeature(),
    "retry": lambda: IamSmartRetryFeature(),
    "test": lambda: IamSmartTestFeature(),
    "timeout": lambda: IamSmartTimeoutFeature(),
}


def _make_feature(name):
    factory = _FEATURES.get(name)
    if factory is not None:
        return factory()
    return _FEATURES["base"]()


# True when this SDK was generated with the named feature class - the
# constructor's tolerance for extend-carried features reads this (an
# active name with no generated class must not become a BaseFeature
# stray when an extend instance carries it).
def _has_feature(name):
    return name in _FEATURES
