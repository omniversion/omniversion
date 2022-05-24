from .shared import verify_expectation
from omniversion import Omniversion

EXPECTATION = {
    "brew-instance": {"count": 3, "packages": ["brew", "git", "node"]},
}


def test_brew_packages(host):
    omniversion = Omniversion()
    hostname = host.backend.get_pytest_id().split("://")[1]
    apt_packages = omniversion.packages.filter(package_manager="brew", host=hostname, verb=["list", "version"])
    verify_expectation(hostname, apt_packages, EXPECTATION)
