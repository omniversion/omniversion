from .shared import verify_expectation
from omniversion import Omniversion

EXPECTATION = {
    "brew-instance": {"packages": ["git", "node"]},
}


def test_brew_packages(host):
    omniversion = Omniversion()
    hostname = host.backend.get_pytest_id().split("://")[1]
    brew_packages = omniversion.packages.filter(package_manager="brew", host=hostname, verb=["list", "version"])
    verify_expectation(hostname, brew_packages, EXPECTATION)
