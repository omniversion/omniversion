#!/usr/bin/env python
"""Test the data module"""
import unittest

from omniversion import Advisory
from omniversion.data.data import Data
from omniversion.file_info import FileInfo
from omniversion.package_metadata import PackageMetadata
from omniversion.package_metadata.list.packages_metadata_list import PackagesMetadataList


class ListPackagesTestCase(unittest.TestCase):
    def test_list_vulnerabilities(self):
        package_metadata = PackageMetadata(name="match", advisories=[Advisory()])
        mismatch_package_info = PackageMetadata(name="mismatch", advisories=[Advisory()])
        package_infos = PackagesMetadataList([package_metadata])
        mismatch_package_infos = PackagesMetadataList([mismatch_package_info])
        data = Data(file_infos=[
            FileInfo(),
            FileInfo(data=package_infos, host="host1", package_manager="pm1", verb="audit"),
            FileInfo(data=mismatch_package_infos, host="mismatch", package_manager="pm1", verb="audit"),
            FileInfo(data=mismatch_package_infos, host="host1", package_manager="mismatch", verb="audit"),
            FileInfo(data=mismatch_package_infos, host="host1", package_manager="pm1", verb="list")
        ])
        self.assertListEqual([package_metadata], data.vulnerabilities(host="host1", package_manager="pm1").data)

    def test_list_specific_packages(self):
        package_info1 = PackageMetadata(name="test1")
        package_info2 = PackageMetadata(name="test2")
        package_info3 = PackageMetadata(name="test3")
        package_infos = PackagesMetadataList([package_info1, package_info2, package_info3])
        data = Data(file_infos=[FileInfo(data=package_infos, verb="list")])
        self.assertListEqual([package_info1], data.items(package_name="test1"))
        self.assertListEqual([package_info1, package_info2], data.items(package_name=["test1", "test2"]))

    def test_list_multiple_verbs(self):
        package_metadata = PackageMetadata(name="test")
        package_infos = PackagesMetadataList([package_metadata])
        data = Data(file_infos=[FileInfo(data=package_infos, verb="audit"),
                                FileInfo(data=package_infos, verb="list"),
                                FileInfo(data=package_infos, verb="outdated")])
        self.assertListEqual([package_metadata, package_metadata], data.items(verb=["audit", "list"]))

    def test_package_infos(self):
        package_metadata = PackageMetadata(name="test")
        package_infos = PackagesMetadataList([package_metadata])
        data = Data(file_infos=[FileInfo(data=package_infos, verb="audit"),
                                FileInfo(data=package_infos, verb="list"),
                                FileInfo(data=package_infos, verb="outdated"),
                                FileInfo(data=package_infos, verb="version")])
        self.assertListEqual([package_metadata, package_metadata], data.list_packages().data)

    def test_available_updates(self):
        package_metadata = PackageMetadata(name="test")
        outdated_package_info = PackageMetadata(name="outdated")
        package_infos = PackagesMetadataList([package_metadata])
        outdated_package_infos = PackagesMetadataList([outdated_package_info])
        data = Data(file_infos=[FileInfo(data=package_infos, verb="audit"),
                                FileInfo(data=package_infos, verb="list"),
                                FileInfo(data=outdated_package_infos, verb="outdated"),
                                FileInfo(data=package_infos, verb="version")])
        self.assertListEqual([outdated_package_info], data.available_updates().data)

    def test_match_versions(self):
        package_info1 = PackageMetadata(name="test", current="1.2.3")
        package_info2 = PackageMetadata(name="test2", current="2.3.4")
        package_info3 = PackageMetadata(name="test", current="3.4.5")
        package_info4 = PackageMetadata(name="test3", current="1.0.0")
        package_infos1 = PackagesMetadataList([package_info1])
        package_infos2 = PackagesMetadataList([package_info2])
        package_infos3 = PackagesMetadataList([package_info3])
        package_infos4 = PackagesMetadataList([package_info4])
        data = Data(file_infos=[FileInfo(data=package_infos1, verb="list"),
                                FileInfo(data=package_infos2, verb="list"),
                                FileInfo(data=package_infos3, verb="version"),
                                FileInfo(data=package_infos4, verb="list")])
        self.assertListEqual([package_info1, package_info2, package_info3],
                             data.match_versions(["test", "test2"]).data)


if __name__ == '__main__':
    unittest.main()
