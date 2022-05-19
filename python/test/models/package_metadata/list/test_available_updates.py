"""Test the data module"""
import unittest

from omniversion.models import AvailableUpdates, PackageMetadata


class AvailableUpdatesTestCase(unittest.TestCase):
    def test_str(self):
        self.assertEqual("All packages up-to-date", AvailableUpdates([]).__str__())
        self.assertEqual("2 outdated packages, 1 package not installed", AvailableUpdates([
            PackageMetadata(name="test", current="1.0.2"),
            PackageMetadata(name="test1", current="1.0.2", latest="2.0.0"),
            PackageMetadata(name="test2", wanted="1.0.2", latest="3.0.0")
        ]).__str__())
        self.assertEqual("1 outdated package", AvailableUpdates([
            PackageMetadata(name="test", current="1.0.2", latest="2.0.0"),
        ]).__str__())
        self.assertEqual("2 outdated packages, 3 packages not installed", AvailableUpdates([
            PackageMetadata(name="test", current="1.0.2"),
            PackageMetadata(name="test1", current="1.0.2", latest="2.0.0"),
            PackageMetadata(name="test2", wanted="1.0.2", latest="3.0.0"),
            PackageMetadata(name="test3", wanted="1.0.2"),
            PackageMetadata(name="test4", wanted="1.0.2", latest="3.0.0")
        ]).__str__())


if __name__ == '__main__':
    unittest.main()
