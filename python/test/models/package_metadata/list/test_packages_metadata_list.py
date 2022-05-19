import unittest

from omniversion.models import PackageMetadata, PackagesMetadataList
from omniversion.models.package_metadata.package_source import PackageSource


class PackagesMetadataListTestCase(unittest.TestCase):
    def test_packages_metadata_list_can_be_initialized(self):
        package_metadata = PackageMetadata(sources=[PackageSource(identifier='test1'), PackageSource(identifier='test2')])
        packages_metadata_list = PackagesMetadataList([package_metadata])
        self.assertEqual(1, len(packages_metadata_list))

    def test_package_metadata_list_instance(self):
        packages_metadata_list = PackagesMetadataList([])
        self.assertTrue(isinstance(packages_metadata_list, PackagesMetadataList))

    def test_str(self):
        self.assertEqual("0 packages", PackagesMetadataList([
        ]).__str__())
        self.assertEqual("1 package", PackagesMetadataList([
            PackageMetadata(),
        ]).__str__())
        self.assertEqual("3 packages", PackagesMetadataList([
            PackageMetadata(),
            PackageMetadata(),
            PackageMetadata(),
        ]).__str__())


if __name__ == '__main__':
    unittest.main()
