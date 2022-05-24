def verify_expectation(hostname, packages, expectations):
    if hostname in expectations:
        if "count" in expectations[hostname]:
            assert len(packages) == expectations[hostname]["count"]
        if "packages" in expectations[hostname]:
            for package_name in expectations[hostname]["packages"]:
                assert packages.show(package_name=package_name)[0].current is not None
    else:
        assert len(packages) == 0
