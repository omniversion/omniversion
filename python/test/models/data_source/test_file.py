import os
import time
import unittest

from omniversion.models import FileDataSource

VECTORS_DIR = os.path.join(os.path.dirname(__file__), "../../vectors")


class FileDataSourceTestCase(unittest.TestCase):
    def test_is_stale(self):
        self.assertFalse(FileDataSource(timestamp=time.time()).is_stale)
        self.assertTrue(FileDataSource(timestamp=time.time() - 60 * 60 * 2).is_stale)
        self.assertTrue(FileDataSource().is_stale)

    def test_has_data(self):
        self.assertTrue(FileDataSource(num_packages=3).has_data)
        self.assertFalse(FileDataSource(num_packages=0).has_data)
        self.assertFalse(FileDataSource(num_packages=None).has_data)
        self.assertFalse(FileDataSource().has_data)

    def test_load_invalid_yaml_file(self):
        file, packages = FileDataSource.load_data(os.path.join(VECTORS_DIR, "invalid_yaml.omniversion.yaml"), "", "",
                                                  "")
        self.assertIsNone(file.version)
        self.assertListEqual([], packages)
        self.assertTrue(file.timestamp > 0)

    def test_load_invalid_items(self):
        file, packages = FileDataSource.load_data(os.path.join(VECTORS_DIR, "invalid_items.omniversion.yaml"), "", "",
                                                  "")
        self.assertEqual("test", file.version)
        self.assertEqual(1, len(packages))
        self.assertTrue(file.timestamp > 0)

    def test_load_non_existent_file(self):
        file, packages = FileDataSource.load_data(os.path.join(VECTORS_DIR, "404.omniversion.yaml"), "", "", "")
        self.assertIsNone(file)
        self.assertListEqual([], packages)

    def test_process_empty_file(self):
        file, packages = FileDataSource.load_data(os.path.join(VECTORS_DIR, "empty_file.omniversion.yaml"), "", "", "")
        self.assertIsNotNone(file)
        self.assertListEqual([], packages)

    def test_process_empty_documents(self):
        file, packages = FileDataSource.load_data(os.path.join(VECTORS_DIR, "many_empty_documents.omniversion.yaml"), "", "", "")
        self.assertIsNotNone(file)
        self.assertListEqual([], packages)

    def test_process_valid_file(self):
        file, packages = FileDataSource.load_data(os.path.join(VECTORS_DIR, "test_host/test_pm/list.omniversion.yaml"),
                                                  "list", "test_host", "test_pm")
        self.assertIsNotNone(file)
        self.assertIsNotNone(packages)
        self.assertEqual(2, len(packages))
        self.assertEqual(2, file.num_packages)

    def test_load_multiple_documents_in_same_file(self):
        file, packages = FileDataSource.load_data(
            os.path.join(VECTORS_DIR, "test_host3/test_pm4/list.omniversion.yaml"),
            "list", "test_host", "test_pm")
        self.assertIsNotNone(file)
        self.assertIsNotNone(file.timestamp)
        self.assertEqual(3, len(packages))
        self.assertEqual(3, file.num_packages)


if __name__ == '__main__':
    unittest.main()
