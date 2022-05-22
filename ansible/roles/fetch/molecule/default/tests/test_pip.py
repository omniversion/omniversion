from typing import Any

from omniversion import Omniversion

EXPECTATION: Any = {
    "rvm-instance": {"count": 0},
    "golang-instance": {"count": 0},
    "redmine-instance": {"count": 0},
    "ubuntu-ansible-instance": {"count": 17, "exist": ["pip", "python", "cryptography"]},
    "debian-instance": {"count": 17, "exist": ["pip", "python", "cryptography"]},
    "docker-php-instance": {"count": 0},
    "fedora-instance": {"count": 18, "exist": ["pip", "python", "cryptography"]},
    "python2-instance": {"count": 37, "exist": ["pip", "python", "cryptography"]},
    "guestbook-instance": {"count": 11, "exist": ["pip", "python", "cryptography"]},
    "node-instance": {"count": 16, "exist": ["pip", "python", "cryptography"]},
    "rails-demo-instance": {"count": 0}
}


def test_pip_packages(host):
    omniversion = Omniversion()
    hostname = host.backend.get_pytest_id().split("://")[1]
    pip_packages = omniversion.packages.filter(package_manager="pip", host=hostname, verb=["list", "version"])
    assert len(pip_packages) == EXPECTATION[hostname]["count"]
    if "exist" in EXPECTATION[hostname]:
        for package_name in EXPECTATION[hostname]["exist"]:
            assert pip_packages.show(package_name=package_name)[0].current is not None
