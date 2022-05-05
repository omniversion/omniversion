import unittest

from omniversion import PackagesMetadataList, PackageMetadata


class PackageInfosListTestCase(unittest.TestCase):
    def test_packages_metadata_list_can_be_initialized(self):
        package_metadata = PackageMetadata(sources=['test1', 'test2'])
        packages_metadata_list = PackagesMetadataList([package_metadata])
        self.assertEqual(1, len(packages_metadata_list))


if __name__ == '__main__':
    unittest.main()
