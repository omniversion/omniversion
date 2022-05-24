from typing import Any

from .shared import verify_expectation
from omniversion import Omniversion

EXPECTATION: Any = {
    "brew-instance": {"count": 6, "packages": ["pip", "python", "setuptools"]},
    "node-instance": {"count": 16, "packages": ["pip", "python"]},
    "python2-instance": {"count": 39, "packages": ["pip", "python", "botocore", "PyYAML", "six"]},
    "python3-instance": {"count": 33, "packages": ["pip", "python", "botocore", "PyYAML", "six"]},
}


def test_pip_packages(host):
    omniversion = Omniversion()
    hostname = host.backend.get_pytest_id().split("://")[1]
    pip_packages = omniversion.packages.filter(package_manager="pip", host=hostname, verb=["list", "version"])
    verify_expectation(hostname, pip_packages, EXPECTATION)
