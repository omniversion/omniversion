from .shared import verify_expectation
from omniversion import Omniversion

EXPECTATION = {
    "go-instance": {"count": 5, "packages": ["go", "viper", "errors", "logrus"]},
}


def test_go_packages(host):
    omniversion = Omniversion()
    hostname = host.backend.get_pytest_id().split("://")[1]
    apt_packages = omniversion.packages.filter(package_manager="go", host=hostname, verb=["list", "version"])
    verify_expectation(hostname, apt_packages, EXPECTATION)
