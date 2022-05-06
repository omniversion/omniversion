import os
import unittest

from omniversion.file_info import FileMetadata
from omniversion.loader.loader import extract_yaml_data, process_file, load_data

invalid_yaml_file_path = os.path.join(os.path.dirname(__file__), "../vectors/invalid_yaml.txt")
empty_file_path = os.path.join(os.path.dirname(__file__), "../vectors/empty_file.txt")


class LoaderTestCase(unittest.TestCase):
    def test_load_invalid_yaml_file(self):
        version, file_data, time = extract_yaml_data(invalid_yaml_file_path)
        self.assertIsNone(version)
        self.assertIsNone(file_data)
        self.assertTrue(time > 0)

    def test_load_non_existent_file(self):
        version, file_data, time = extract_yaml_data("i_do_not_exist.dat")
        self.assertIsNone(version)
        self.assertIsNone(file_data)
        self.assertIsNone(time)

    def test_process_empty_file(self):
        files: [FileMetadata] = []
        process_file(verb="empty_file",
                     host="test",
                     host_path=os.path.join(os.path.dirname(__file__), "../vectors/test"),
                     package_manager="test",
                     add_file=files.append)
        self.assertEqual(1, len(files))
        self.assertIsNone(files[0].data)

    def test_load_data_should_ignore_host_and_pm_not_on_allow_list(self):
        files: [FileMetadata] = []
        load_data(base_path=os.path.join(os.path.dirname(__file__), "../vectors/"),
                  hosts=["test_host2"],
                  package_managers=["test_pm3"],
                  add_file=files.append)
        self.assertEqual(1, len(files))
        self.assertEqual("list.omniversion.yaml", files[0].name)

    def test_process_valid_file(self):
        files: [FileMetadata] = []
        process_file(verb="list",
                     host="test_host",
                     host_path=os.path.join(os.path.dirname(__file__), "../vectors/test_host/test_pm"),
                     package_manager="test_pm",
                     add_file=files.append)
        self.assertEqual(1, len(files))
        self.assertIsNotNone(1, files[0].data)
        self.assertEqual(2, len(files[0].data))


if __name__ == '__main__':
    unittest.main()
