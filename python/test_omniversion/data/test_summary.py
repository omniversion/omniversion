#!/usr/bin/env python
"""Test the data module"""
import os
import unittest

from omniversion import Omniversion
from omniversion.file_info import FileMetadata
from omniversion.package_metadata import PackageMetadata
from omniversion.package_metadata.list.packages_metadata_list import PackagesMetadataList

test_file_path = os.path.join(os.path.dirname(__file__), "../vectors/test-env.txt")


class SummaryTestCase(unittest.TestCase):
    def test_summary_for_no_data(self):
        data = Omniversion(file_infos=[])
        data_as_str = data.__str__()
        self.assertIn("No files loaded", data_as_str)

    def test_summary_for_single_file(self):
        package_infos = PackagesMetadataList([PackageMetadata(name="package", current="0.1.2", package_manager="test")])
        data = Omniversion(file_infos=[FileMetadata(data=package_infos, name="test")])
        data_as_str = data.__str__()
        self.assertIn("1 file loaded", data_as_str)

    def test_summary_for_multiple_files(self):
        package_infos = PackagesMetadataList([PackageMetadata(name="package", current="0.1.2", package_manager="test")])
        data = Omniversion(
            file_infos=[FileMetadata(data=package_infos, name="test"), FileMetadata(data=package_infos, name="test")])
        data_as_str = data.__str__()
        self.assertIn("2 files loaded", data_as_str)

    def test_load_files(self):
        base_path = "test_omniversion/vectors"
        data = Omniversion(base_path=base_path, hosts=["test_host"], package_managers=["test_pm"])
        data_as_str = data.__str__()
        self.assertIn("4 files loaded", data_as_str)

    def test_collect_hosts(self):
        data = Omniversion(file_infos=[FileMetadata(name="1", host="test1"),
                                       FileMetadata(name="2", host="test2"),
                                       FileMetadata(name="3", host="test1")])
        self.assertListEqual(["test1", "test2"], data.hosts())


if __name__ == '__main__':
    unittest.main()
