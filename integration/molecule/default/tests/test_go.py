from .shared import verify_expectation
from omniversion import Omniversion

EXPECTATION = {
    "go-instance": {"packages": ["go",
                                 "cloud.google.com/go/datastore",
                                 "github.com/sirupsen/logrus"]},
}


def test_go_packages(host):
    omniversion = Omniversion()
    hostname = host.backend.get_pytest_id().split("://")[1]
    go_packages = omniversion.packages.filter(package_manager="go", host=hostname, verb=["list", "version"])
    verify_expectation(hostname, go_packages, EXPECTATION)
