from .shared import verify_expectation
from omniversion import Omniversion

EXPECTATION = {
    "brew-instance": {"packages": ["npm", "node", "v8"]},
    "node-instance": {"packages": ["npm", "node", "v8", "pug", "test", "async.js"]},
}


def test_npm_packages(host):
    omniversion = Omniversion()
    hostname = host.backend.get_pytest_id().split("://")[1]
    npm_packages = omniversion.packages.filter(package_manager="npm", host=hostname, verb=["list", "version"])
    verify_expectation(hostname, npm_packages, EXPECTATION)
