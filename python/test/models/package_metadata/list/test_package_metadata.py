"""Test the data module"""
import unittest

from omniversion.models import PackageMetadata


class PackageMetadataTestCase(unittest.TestCase):
    def test_str(self):
        self.assertEqual("unknown package name", PackageMetadata().__str__())
        self.assertEqual("unknown package name", PackageMetadata(current="1.0.0").__str__())
        self.assertEqual("package `test`", PackageMetadata(name="test").__str__())
        self.assertEqual("package `test`@1.0.0 installed", PackageMetadata(name="test", current="1.0.0").__str__())
        self.assertEqual("package `test`@2.0.5 wanted", PackageMetadata(name="test", wanted="2.0.5").__str__())
        self.assertEqual("package `test` via `test_pm`", PackageMetadata(name="test",
                                                                         package_manager="test_pm").__str__())
        self.assertEqual("package `test`@1.2.3+test wanted via `test_pm`",
                         PackageMetadata(name="test", wanted="1.2.3+test", package_manager="test_pm").__str__())


if __name__ == '__main__':
    unittest.main()
