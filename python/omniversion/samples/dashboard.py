from omniversion.data import Data
from omniversion import pretty


def show_dashboard(data: Data):
    print(pretty.header("Data sources"))
    print(data.files)

    print(pretty.header("Dependency count"))
    print(data.dependencies().overview())

    print(pretty.header("App versions"))
    print(data.dependencies(package_name="test1"))

    print(pretty.header("Version consistency"))
    print(data.match_versions(["test1"], "Test"))

    print(pretty.header("Available updates"))
    print(data.available_updates())

    print(pretty.header("Security audit"))
    print(data.vulnerabilities())
