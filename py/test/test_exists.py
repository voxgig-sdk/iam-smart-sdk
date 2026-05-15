# ProjectName SDK exists test

import pytest
from iamsmart_sdk import IamSmartSDK


class TestExists:

    def test_should_create_test_sdk(self):
        testsdk = IamSmartSDK.test(None, None)
        assert testsdk is not None
