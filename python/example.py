import os

from omniversion.data import Data
from omniversion import pretty

data = Data(os.path.join(os.path.dirname(__file__), "./test_omniversion/vectors"))

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
