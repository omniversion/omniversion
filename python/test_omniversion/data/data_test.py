#!/usr/bin/env python
"""Test the data module"""
import unittest

from omniversion import Advisory
from omniversion.data.data import Data
from omniversion.file_info import FileInfo
from omniversion.package_info import PackageInfo
from omniversion.package_info.list.package_infos_list import PackageInfosList


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
        self.assertIn("4 files loaded", data_as_str)

    def test_collect_hosts(self):
        data = Data(file_infos=[FileInfo(name="1", host="test1"),
                                FileInfo(name="2", host="test2"),
                                FileInfo(name="3", host="test1")])
        self.assertListEqual(["test1", "test2"], data.hosts())

    def test_list_vulnerabilities(self):
        package_info = PackageInfo(name="test", advisories=[Advisory()])
        package_infos = PackageInfosList([package_info])
        data = Data(file_infos=[
            FileInfo(name="1"),
            FileInfo(name="2", data=package_infos, host="test1", package_manager="test2", verb="audit"),
            FileInfo(name="3", data=package_infos, host="wrong", package_manager="test2", verb="audit"),
            FileInfo(name="4", data=package_infos, host="test1", package_manager="wrong", verb="audit"),
            FileInfo(name="5", data=package_infos, host="test1", package_manager="test2", verb="list")
        ])
        self.assertEqual([package_info], data.vulnerabilities(host="test1", package_manager="test2"))

    def test_list_specific_packages(self):
        package_info1 = PackageInfo(name="test1")
        package_info2 = PackageInfo(name="test2")
        package_info3 = PackageInfo(name="test3")
        package_infos = PackageInfosList([package_info1, package_info2, package_info3])
        data = Data(file_infos=[FileInfo(data=package_infos, verb="list")])
        self.assertListEqual([package_info1], data.items(package_name="test1"))
        self.assertListEqual([package_info1, package_info2], data.items(package_name=["test1", "test2"]))


if __name__ == '__main__':
    unittest.main()
