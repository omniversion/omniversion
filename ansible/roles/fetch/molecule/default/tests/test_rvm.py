from omniversion import Omniversion

EXPECTATION = {
    "rvm-instance": {"count": 2},
    "golang-instance": {"count": 0},
    "redmine-instance": {"count": 0},
    "ubuntu-ansible-instance": {"count": 0},
    "debian-instance": {"count": 0},
    "docker-php-instance": {"count": 0},
    "fedora-instance": {"count": 0},
    "python2-instance": {"count": 0},
    "guestbook-instance": {"count": 0},
    "node-instance": {"count": 0},
    "rails-demo-instance": {"count": 0}
}


def test_rvm_packages(host):
    omniversion = Omniversion()
    hostname = host.backend.get_pytest_id().split("://")[1]
    rvm_packages = omniversion.packages.filter(package_manager="rvm", host=hostname, verb=["list", "version"])
    assert len(rvm_packages) == EXPECTATION[hostname]["count"]
