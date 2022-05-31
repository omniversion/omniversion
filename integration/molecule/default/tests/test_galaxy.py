from .shared import verify_expectation
from omniversion import Omniversion

EXPECTATION = {
    "ansible-instance": {"count": 3, "packages": ["ansible", "geerlingguy.nginx", "k8s_manifests"]},
}


def test_galaxy_packages(host):
    omniversion = Omniversion()
    hostname = host.backend.get_pytest_id().split("://")[1]
    galaxy_packages = omniversion.packages.filter(package_manager="galaxy", host=hostname, verb=["list", "version"])
    verify_expectation(hostname, galaxy_packages, EXPECTATION)
