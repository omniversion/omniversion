from omniversion import Omniversion

EXPECTATION = {
    "brew-instance": {"count": 337},
    "go-instance": {"count": 1},
    "node-instance": {"count": 1},
    "python2-instance": {"count": 1},
    "python3-instance": {"count": 1},
    "rbenv-ansible-instance": {"count": 1},
    "rvm-instance": {"count": 1},
}


def test_rvm_packages(host):
    omniversion = Omniversion()
    hostname = host.backend.get_pytest_id().split("://")[1]
    apt_packages = omniversion.packages.filter(package_manager="apt", host=hostname, verb=["list", "version"])
    assert len(apt_packages) == EXPECTATION[hostname]["count"]
