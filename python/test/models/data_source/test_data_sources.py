import os
import unittest

import pytest

from omniversion.models import PackageMetadata, FileDataSource, ConfigDataSource
from omniversion.models.data_source.data_sources import DataSources

TEST_ENV_TXT_PATH = os.path.join(os.path.dirname(__file__), "../../vectors/test_env.txt")


class DataSourcesTestCase(unittest.TestCase):
    def test_add_file(self):
        data_sources = DataSources()
        self.assertTrue(data_sources.add_file(
            file_path=os.path.join(os.path.dirname(__file__),
                                   "../../vectors/test_host/test_pm/list.omniversion.yaml"),
            verb="list",
            host="test_host",
            package_manager="test_pm"))
        self.assertEqual(1, len(data_sources.files))
        self.assertIsNotNone(data_sources.files[0].num_packages)
        self.assertEqual(2, data_sources.files[0].num_packages)

    def test_add_multiple_documents_in_same_file(self):
        data_sources = DataSources()
        self.assertTrue(data_sources.add_file(
            file_path=os.path.join(os.path.dirname(__file__),
                                   "../../vectors/test_host3/test_pm4/list.omniversion.yaml"),
            verb="list",
            host="test_host",
            package_manager="test_pm"))
        self.assertEqual(1, len(data_sources.files))
        self.assertIsNotNone(data_sources.files[0].version)
        self.assertIsNotNone(data_sources.files[0].timestamp)
        self.assertIsNotNone(data_sources.files[0].num_packages)
        self.assertEqual(3, data_sources.files[0].num_packages)

    def test_add_files(self):
        data_sources = DataSources()
        self.assertTrue(data_sources.add_files(base_path=os.path.join(os.path.dirname(__file__), "../../vectors")))
        self.assertGreater(len(data_sources.files), 4)
        self.assertIsNotNone(data_sources.files[0].version)
        self.assertIsNotNone(data_sources.files[0].timestamp)
        self.assertIsNotNone(data_sources.files[0].num_packages)
        self.assertEqual(2, data_sources.files[0].num_packages)

        self.assertIn("test_host", data_sources.hostnames)
        test_host_info = data_sources.host_infos[0]
        self.assertEqual("test_host", test_host_info[0])
        self.assertEqual(4, len(test_host_info[1]))
        self.assertEqual(0, len(test_host_info[2]))

        self.assertIn("test_pm", data_sources.package_manager_identifiers)

    def test_add_config(self):
        data_sources = DataSources()
        packages = data_sources.add_config(TEST_ENV_TXT_PATH, regex="^TEST=(?P<version>.*)$", name="package1")
        local_package_infos = PackageMetadata(host="localhost", name="package1", current="1.0.0",
                                              package_manager="file")
        self.assertListEqual([local_package_infos], packages)
        self.assertEqual(1, len(data_sources.configs))
        self.assertEqual(1, data_sources.configs[0].num_packages)

    def test_add_unnamed_config_value(self):
        data_sources = DataSources()
        packages = data_sources.add_config(TEST_ENV_TXT_PATH, regex="^(?P<name>.*)=(?P<version>.*)$")
        local_package_infos = PackageMetadata(host="localhost", name="TEST", current="1.0.0",
                                              package_manager="file")
        self.assertListEqual([local_package_infos], packages)
        self.assertEqual(1, len(data_sources.configs))
        self.assertEqual(1, data_sources.configs[0].num_packages)

    @staticmethod
    def test_invalid_regex():
        data_sources = DataSources()
        with pytest.raises(IndexError):
            data_sources.add_config(TEST_ENV_TXT_PATH, regex="^(.*)$", name="package1")
        with pytest.raises(IndexError):
            data_sources.add_config(TEST_ENV_TXT_PATH, regex="^TEST=(?P<version>.*)$")

    def test_create_with_files_and_config(self):
        data_sources = DataSources(files=[
            FileDataSource(path=os.path.join(os.path.dirname(__file__),
                                             "../../vectors/test_host/test_pm/list.omniversion.yaml"),
                           verb="list", package_manager="test_pm", host="test_host")
        ],
            configs=[
                ConfigDataSource(file_path=TEST_ENV_TXT_PATH, regex="^(?P<name>.*)=(?P<version>.*)$")
            ])
        self.assertEqual(3, len(data_sources.packages))


if __name__ == '__main__':
    unittest.main()
