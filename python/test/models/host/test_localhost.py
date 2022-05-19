"""Test the data module"""
import unittest

from omniversion.models import Localhost, PackagesMetadataList


class LocalhostTestCase(unittest.TestCase):
    def test_str(self):
        localhost = Localhost(packages=PackagesMetadataList([]))
        self.assertEqual("localhost", localhost.name)
        self.assertEqual("localhost", localhost.__str__())


if __name__ == '__main__':
    unittest.main()
