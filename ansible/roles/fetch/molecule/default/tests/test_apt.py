from omniversion import Omniversion

EXPECTATION = {
    "rvm-instance": {"count": 409},
    "golang-instance": {"count": 202},
    "redmine-instance": {"count": 236},
    "ubuntu-ansible-instance": {"count": 203},
    "debian-instance": {"count": 185},
    "docker-php-instance": {"count": 269},
    "fedora-instance": {"count": 0},
    "python2-instance": {"count": 443},
    "guestbook-instance": {"count": 0},
    "node-instance": {"count": 518},
    "rails-demo-instance": {"count": 388}
}


def test_rvm_packages(host):
    omniversion = Omniversion()
    hostname = host.backend.get_pytest_id().split("://")[1]
    apt_packages = omniversion.packages.filter(package_manager="apt", host=hostname, verb=["list", "version"])
    assert len(apt_packages) == EXPECTATION[hostname]["count"]
