#!/usr/bin/env python
"""Test the data module"""
import os
import unittest

import pytest

from omniversion.data.data import Data
from omniversion.file_info import FileInfo
from omniversion.package_info import PackageInfo
from omniversion.package_info.list.package_infos_list import PackageInfosList

test_file_path = os.path.join(os.path.dirname(__file__), "../vectors/test-env.txt")


class LocalConfigValuesTestCase(unittest.TestCase):
    def test_add_named_local_config_value(self):
        package_info = PackageInfo(name="test")
        package_infos = PackageInfosList([package_info])
        data = Data(file_infos=[FileInfo(data=package_infos, verb="list")])
        data.add_local_config_value(test_file_path, regex="^TEST=(?P<version>.*)$", name="package1")
        local_package_infos = PackageInfo(host="localhost", name="package1", version="1.0.0", pm="local file")
        self.assertListEqual([package_info, local_package_infos], data.items())

    def test_add_unnamed_local_config_value(self):
        package_info = PackageInfo(name="test")
        package_infos = PackageInfosList([package_info])
        data = Data(file_infos=[FileInfo(data=package_infos, verb="list")])
        data.add_local_config_value(test_file_path, regex="^(?P<name>.*)=(?P<version>.*)$", name="package1")
        local_package_infos = PackageInfo(host="localhost", name="package1", version="1.0.0", pm="local file")
        self.assertListEqual([package_info, local_package_infos], data.items())

    @staticmethod
    def test_invalid_regex():
        data = Data()
        with pytest.raises(IndexError):
            data.add_local_config_value(test_file_path, regex="^(.*)$", name="package1")
        with pytest.raises(IndexError):
            data.add_local_config_value(test_file_path, regex="^TEST=(?P<version>.*)$")


if __name__ == '__main__':
    unittest.main()
