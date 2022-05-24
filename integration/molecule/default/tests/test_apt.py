from .shared import verify_expectation
from omniversion import Omniversion

EXPECTATION = {
    "brew-instance": {"count": 337, "packages": ["nmap", "mysql-common", "nginx"]},
    "go-instance": {"count": 337, "packages": ["apt"]},
    "node-instance": {"count": 337},
    "python2-instance": {"count": 358},
    "python3-instance": {"count": 365},
    "rbenv-instance": {"count": 337},
    "rvm-instance": {"count": 354},
}


def test_apt_packages(host):
    omniversion = Omniversion()
    hostname = host.backend.get_pytest_id().split("://")[1]
    apt_packages = omniversion.packages.filter(package_manager="apt", host=hostname, verb=["list", "version"])
    verify_expectation(hostname, apt_packages, EXPECTATION)
