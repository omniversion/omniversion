"""Test the data module"""
import unittest

from omniversion.models import InstalledPackage


class InstalledPackageTestCase(unittest.TestCase):
    def test_str(self):
        self.assertEqual("version `?` in location `?`", InstalledPackage().__str__())
        self.assertEqual("version `?` in location `test_location`", InstalledPackage(location="test_location").__str__())
        self.assertEqual("version `test_version` in location `?`", InstalledPackage(version="test_version").__str__())
        self.assertEqual("version `test_version` in location `test_location`", InstalledPackage(location="test_location", version="test_version").__str__())


if __name__ == '__main__':
    unittest.main()
