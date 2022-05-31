from .shared import verify_expectation
from omniversion import Omniversion

EXPECTATION = {
    "rbenv-instance": {"packages": ["nokogiri", "rails-api"]},
    "rvm-instance": {"packages": ["nokogiri", "rails-api"]},
}


def test_gem_packages(host):
    omniversion = Omniversion()
    hostname = host.backend.get_pytest_id().split("://")[1]
    gem_packages = omniversion.packages.filter(package_manager="gem", host=hostname, verb=["list", "version"])
    verify_expectation(hostname, gem_packages, EXPECTATION)
