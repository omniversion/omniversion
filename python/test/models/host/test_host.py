"""Test the data module"""
import unittest

from omniversion.models import PackagesMetadataList, PackageMetadata, Host


class HostTestCase(unittest.TestCase):
    def test_list_from_packages(self):
        packages = PackagesMetadataList([
            PackageMetadata(name="test1", host="host1"),
            PackageMetadata(name="test2", host="host1"),
            PackageMetadata(name="test3", host="host3"),
            PackageMetadata(name="test4", host="host2"),
            PackageMetadata(name="test5", host="localhost"),
            PackageMetadata(name="test5", host="host1"),
        ])
        hosts = Host.list_from_packages(packages)
        self.assertEqual(4, len(hosts))
        self.assertEqual("host1", hosts[0].name)
        self.assertEqual("host: host1", hosts[0].__str__())


if __name__ == '__main__':
    unittest.main()
