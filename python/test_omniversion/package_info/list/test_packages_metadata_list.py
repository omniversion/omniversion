import unittest

from omniversion import PackagesMetadataList, PackageMetadata


class PackagesMetadataListTestCase(unittest.TestCase):
    def test_packages_metadata_list_can_be_initialized(self):
        package_metadata = PackageMetadata(sources=['test1', 'test2'])
        packages_metadata_list = PackagesMetadataList([package_metadata])
        self.assertEqual(1, len(packages_metadata_list))

    def test_pretty_print_packages_metadata_list(self):
        packages_metadata_list = PackagesMetadataList([
            PackageMetadata(current="2.3.4"),
            PackageMetadata(current="1.2.3")
        ])
        self.assertIn("2 versions found", packages_metadata_list.__str__())

        packages_metadata_list = PackagesMetadataList()
        self.assertIn("No versions found", packages_metadata_list.__str__())

        packages_metadata_list = PackagesMetadataList([
            PackageMetadata(host="test_host", package_manager="test_pm", current="1.2.3")
        ])
        self.assertIn("test_host", packages_metadata_list.summary())


if __name__ == '__main__':
    unittest.main()
