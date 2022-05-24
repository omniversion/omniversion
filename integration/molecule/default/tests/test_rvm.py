from typing import Any

from .shared import verify_expectation
from omniversion import Omniversion

EXPECTATION: Any = {
    "rvm-instance": {"count": 2, "packages": ["ruby", "rvm"]},
}


def test_rvm_packages(host):
    omniversion = Omniversion()
    hostname = host.backend.get_pytest_id().split("://")[1]
    rvm_packages = omniversion.packages.filter(package_manager="rvm", host=hostname, verb=["list", "version"])
    verify_expectation(hostname, rvm_packages, EXPECTATION)
