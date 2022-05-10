from omniversion import Omniversion, pretty


def show_dashboard(omniversion: Omniversion):
    print(pretty.header("Data sources"))
    print(omniversion.files)

    print(pretty.header("Dependency count"))
    print(omniversion.ls().summary())

    print(pretty.header("App versions"))
    print(omniversion.ls(package_name="test1"))

    print(pretty.header("Version consistency"))
    print(omniversion.show(["test1"], "Test"))

    print(pretty.header("Available updates"))
    print(omniversion.outdated())

    print(pretty.header("Security audit"))
    print(omniversion.audit())
