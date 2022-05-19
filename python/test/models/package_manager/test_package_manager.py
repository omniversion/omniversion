"""Test the data module"""
import unittest

from omniversion.models import PackagesMetadataList, PackageMetadata
from omniversion.models.package_manager import PackageManager


class PackageManagerTestCase(unittest.TestCase):
    def test_str(self):
        package_manager = PackageManager(identifier="test", packages=PackagesMetadataList([]))
        self.assertEqual("test", package_manager.identifier)
        self.assertEqual("package manager: test", package_manager.__str__())

    def test_list_from_packages(self):
        package_managers = PackageManager.list_from_packages(packages=PackagesMetadataList([
            PackageMetadata(name="test1", package_manager="pm1"),
            PackageMetadata(name="test2", package_manager="pm1"),
            PackageMetadata(name="test3", package_manager="pm2"),
        ]))
        self.assertEqual(2, len(package_managers))
        self.assertEqual("pm1", package_managers[0].identifier)
        self.assertEqual("package manager: pm1", package_managers[0].__str__())


if __name__ == '__main__':
    unittest.main()
