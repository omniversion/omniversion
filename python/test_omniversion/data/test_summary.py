#!/usr/bin/env python
"""Test the data module"""
import os
import unittest

from omniversion.data.data import Data
from omniversion.file_info import FileInfo
from omniversion.package_info import PackageInfo
from omniversion.package_info.list.package_infos_list import PackageInfosList

test_file_path = os.path.join(os.path.dirname(__file__), "../vectors/test-env.txt")


class InitializationTestCase(unittest.TestCase):
    def test_summary_for_no_data(self):
        data = Data(file_infos=[])
        data_as_str = data.__str__()
        self.assertIn("No files loaded", data_as_str)

    def test_summary_for_single_file(self):
        package_infos = PackageInfosList([PackageInfo(name="package", version="0.1.2", pm="test")])
        data = Data(file_infos=[FileInfo(data=package_infos, name="test")])
        data_as_str = data.__str__()
        self.assertIn("1 file loaded", data_as_str)

    def test_summary_for_multiple_files(self):
        package_infos = PackageInfosList([PackageInfo(name="package", version="0.1.2", pm="test")])
        data = Data(file_infos=[FileInfo(data=package_infos, name="test"), FileInfo(data=package_infos, name="test")])
        data_as_str = data.__str__()
        self.assertIn("2 files loaded", data_as_str)

    def test_load_files(self):
        base_path = "test_omniversion/vectors"
        data = Data(base_path=base_path)
        data_as_str = data.__str__()
        self.assertIn("5 files loaded", data_as_str)

    def test_collect_hosts(self):
        data = Data(file_infos=[FileInfo(name="1", host="test1"),
                                FileInfo(name="2", host="test2"),
                                FileInfo(name="3", host="test1")])
        self.assertListEqual(["test1", "test2"], data.hosts())


if __name__ == '__main__':
    unittest.main()
