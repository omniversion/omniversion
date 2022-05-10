#!/usr/bin/env python
"""Test the data module"""
import os
import unittest

import pytest

from omniversion import Omniversion
from omniversion.file_info import FileMetadata
from omniversion.package_metadata import PackageMetadata
from omniversion.package_metadata.list.packages_metadata_list import PackagesMetadataList

test_file_path = os.path.join(os.path.dirname(__file__), "../vectors/test_env.txt")


class LocalConfigValuesTestCase(unittest.TestCase):
    def test_add_named_local_config_value(self):
        package_metadata = PackageMetadata(name="test")
        package_infos = PackagesMetadataList([package_metadata])
        data = Omniversion(file_infos=[FileMetadata(data=package_infos, verb="list")])
        data.add_local_config_value(test_file_path, regex="^TEST=(?P<version>.*)$", name="package1")
        local_package_infos = PackageMetadata(host="localhost", name="package1", current="1.0.0",
                                              package_manager="local file")
        self.assertListEqual([package_metadata, local_package_infos], data.items())

    def test_add_unnamed_local_config_value(self):
        package_metadata = PackageMetadata(name="test")
        package_infos = PackagesMetadataList([package_metadata])
        data = Omniversion(file_infos=[FileMetadata(data=package_infos, verb="list")])
        data.add_local_config_value(test_file_path, regex="^(?P<name>.*)=(?P<version>.*)$", name="package1")
        local_package_infos = PackageMetadata(host="localhost", name="package1", current="1.0.0",
                                              package_manager="local file")
        self.assertListEqual([package_metadata, local_package_infos], data.items())

    @staticmethod
    def test_invalid_regex():
        data = Omniversion()
        with pytest.raises(IndexError):
            data.add_local_config_value(test_file_path, regex="^(.*)$", name="package1")
        with pytest.raises(IndexError):
            data.add_local_config_value(test_file_path, regex="^TEST=(?P<version>.*)$")


if __name__ == '__main__':
    unittest.main()
