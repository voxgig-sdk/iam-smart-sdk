# IamSmart SDK feature factory

from feature.base_feature import IamSmartBaseFeature
from feature.test_feature import IamSmartTestFeature


def _make_feature(name):
    features = {
        "base": lambda: IamSmartBaseFeature(),
        "test": lambda: IamSmartTestFeature(),
    }
    factory = features.get(name)
    if factory is not None:
        return factory()
    return features["base"]()
