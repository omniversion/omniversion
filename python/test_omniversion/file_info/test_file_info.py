import time
import unittest

from omniversion import PackagesMetadataList, PackageMetadata
from omniversion.file_info import FileInfo


class FileInfoTestCase(unittest.TestCase):
    def test_is_stale(self):
        self.assertFalse(FileInfo(time=time.time()).is_stale)
        self.assertTrue(FileInfo(time=time.time() - 60 * 60 * 2).is_stale)
        self.assertTrue(FileInfo().is_stale)

    def test_has_data(self):
        package_metadata = PackageMetadata(name="test")
        self.assertTrue(FileInfo(data=PackagesMetadataList([package_metadata])).has_data)
        self.assertFalse(FileInfo(data=PackagesMetadataList([])).has_data)
        self.assertFalse(FileInfo().has_data)

    def test_num_entries(self):
        package_metadata = PackageMetadata(name="test")
        self.assertEqual(1, FileInfo(data=PackagesMetadataList([package_metadata])).num_entries)
        self.assertEqual(0, FileInfo(data=PackagesMetadataList([])).num_entries)
        self.assertEqual(0, FileInfo().num_entries)

    def test_pretty_print(self):
        package_metadata = PackageMetadata(name="test")
        single_package = PackagesMetadataList([package_metadata])
        two_packages = PackagesMetadataList([package_metadata, package_metadata])

        self.assertIn("Recent data loaded", FileInfo(data=single_package, time=time.time()).__str__())
        self.assertIn("Stale data loaded", FileInfo(data=single_package, time=time.time() - 60 * 60 * 2).__str__())

        self.assertIn("2 entries", FileInfo(data=two_packages).__str__())
        self.assertIn("1 entry", FileInfo(data=single_package).__str__())

        self.assertIn("No entries", FileInfo(data=PackagesMetadataList([]), time=time.time()).__str__())
        self.assertIn("0 entries", FileInfo(data=PackagesMetadataList([])).__str__())
        self.assertIn("0 entries", FileInfo().__str__())


if __name__ == '__main__':
    unittest.main()
