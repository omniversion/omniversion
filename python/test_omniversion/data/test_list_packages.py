#!/usr/bin/env python
"""Test the data module"""
import unittest

from omniversion import Advisory
from omniversion.data.data import Data
from omniversion.file_info import FileInfo
from omniversion.package_info import PackageInfo
from omniversion.package_info.list.package_infos_list import PackageInfosList


class ListPackagesTestCase(unittest.TestCase):
    def test_list_vulnerabilities(self):
        package_info = PackageInfo(name="match", advisories=[Advisory()])
        mismatch_package_info = PackageInfo(name="mismatch", advisories=[Advisory()])
        package_infos = PackageInfosList([package_info])
        mismatch_package_infos = PackageInfosList([mismatch_package_info])
        data = Data(file_infos=[
            FileInfo(name="1"),
            FileInfo(name="2", data=package_infos, host="host1", package_manager="pm1", verb="audit"),
            FileInfo(name="3", data=mismatch_package_infos, host="mismatch", package_manager="pm1", verb="audit"),
            FileInfo(name="4", data=mismatch_package_infos, host="host1", package_manager="mismatch", verb="audit"),
            FileInfo(name="5", data=mismatch_package_infos, host="host1", package_manager="pm1", verb="list")
        ])
        self.assertListEqual([package_info], data.vulnerabilities(host="host1", package_manager="pm1").data)

    def test_list_specific_packages(self):
        package_info1 = PackageInfo(name="test1")
        package_info2 = PackageInfo(name="test2")
        package_info3 = PackageInfo(name="test3")
        package_infos = PackageInfosList([package_info1, package_info2, package_info3])
        data = Data(file_infos=[FileInfo(data=package_infos, verb="list")])
        self.assertListEqual([package_info1], data.items(package_name="test1"))
        self.assertListEqual([package_info1, package_info2], data.items(package_name=["test1", "test2"]))

    def test_list_multiple_verbs(self):
        package_info = PackageInfo(name="test")
        package_infos = PackageInfosList([package_info])
        data = Data(file_infos=[FileInfo(data=package_infos, verb="audit"),
                                FileInfo(data=package_infos, verb="list"),
                                FileInfo(data=package_infos, verb="outdated")])
        self.assertListEqual([package_info, package_info], data.items(verb=["audit", "list"]))

    def test_package_infos(self):
        package_info = PackageInfo(name="test")
        package_infos = PackageInfosList([package_info])
        data = Data(file_infos=[FileInfo(data=package_infos, verb="audit"),
                                FileInfo(data=package_infos, verb="list"),
                                FileInfo(data=package_infos, verb="outdated"),
                                FileInfo(data=package_infos, verb="version")])
        self.assertListEqual([package_info, package_info], data.list_packages().data)

    def test_available_updates(self):
        package_info = PackageInfo(name="test")
        outdated_package_info = PackageInfo(name="outdated")
        package_infos = PackageInfosList([package_info])
        outdated_package_infos = PackageInfosList([outdated_package_info])
        data = Data(file_infos=[FileInfo(data=package_infos, verb="audit"),
                                FileInfo(data=package_infos, verb="list"),
                                FileInfo(data=outdated_package_infos, verb="outdated"),
                                FileInfo(data=package_infos, verb="version")])
        self.assertListEqual([outdated_package_info], data.available_updates().data)

    def test_match_versions(self):
        package_info1 = PackageInfo(name="test", version="1.2.3")
        package_info2 = PackageInfo(name="test2", version="2.3.4")
        package_info3 = PackageInfo(name="test", version="3.4.5")
        package_info4 = PackageInfo(name="test3", version="1.0.0")
        package_infos1 = PackageInfosList([package_info1])
        package_infos2 = PackageInfosList([package_info2])
        package_infos3 = PackageInfosList([package_info3])
        package_infos4 = PackageInfosList([package_info4])
        data = Data(file_infos=[FileInfo(data=package_infos1, verb="list"),
                                FileInfo(data=package_infos2, verb="list"),
                                FileInfo(data=package_infos3, verb="version"),
                                FileInfo(data=package_infos4, verb="list")])
        self.assertListEqual([package_info1, package_info2, package_info3],
                             data.match_versions(["test", "test2"]).data)


if __name__ == '__main__':
    unittest.main()
