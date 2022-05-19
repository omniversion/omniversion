"""Test the data module"""
import unittest

from omniversion.models import Advisory, PackageMetadata, PackagesMetadataList

PACKAGE_1 = PackageMetadata(name="test1")
PACKAGE_2 = PackageMetadata(name="test2")
PACKAGE_3 = PackageMetadata(name="test3")
PACKAGE_4 = PackageMetadata(name="test4")


class ListPackagesTestCase(unittest.TestCase):
    def test_audit(self):
        packages = PackagesMetadataList([
            PackageMetadata(host="host1", package_manager="pm1", verb="audit", advisories=[Advisory()]),
            PackageMetadata(host="mismatch", package_manager="pm1", verb="audit", advisories=[Advisory()]),
            PackageMetadata(host="host1", package_manager="mismatch", verb="audit", advisories=[Advisory()]),
            PackageMetadata(host="host1", package_manager="pm1", verb="list"),
        ])
        self.assertEqual([packages[0]], packages.audit(host="host1", package_manager="pm1"))

    def test_specific_packages(self):
        data = PackagesMetadataList([
            PACKAGE_1,
            PACKAGE_2,
            PACKAGE_3,
        ])
        self.assertListEqual([PACKAGE_1], data.filter(package_name="test1").data)
        self.assertListEqual([PACKAGE_1, PACKAGE_2], data.filter(package_name=["test1", "test2"]).data)

    def test_multiple_verbs(self):
        packages = PackagesMetadataList([
            PackageMetadata(host="host1", package_manager="pm1", verb="audit"),
            PackageMetadata(host="host1", package_manager="pm1", verb="list"),
            PackageMetadata(host="host1", package_manager="pm1", verb="outdated"),
            ])
        self.assertListEqual([packages[0], packages[1]], packages.filter(verb=["audit", "list"]).data)

    def test_package_infos(self):
        packages = PackagesMetadataList([
            PackageMetadata(host="host1", package_manager="pm1", verb="audit"),
            PackageMetadata(host="host1", package_manager="pm1", verb="list"),
            PackageMetadata(host="host1", package_manager="pm1", verb="outdated"),
            PackageMetadata(host="host1", package_manager="pm1", verb="version"),
            ])
        self.assertListEqual([packages[1], packages[3]], packages.ls().data)

    def test_available_updates(self):
        packages = PackagesMetadataList([
            PackageMetadata(host="host1", package_manager="pm1", verb="audit"),
            PackageMetadata(host="host1", package_manager="pm1", verb="list"),
            PackageMetadata(host="host1", package_manager="pm1", verb="outdated"),
            PackageMetadata(host="host1", package_manager="pm1", verb="version"),
        ])
        self.assertListEqual([packages[2]], packages.outdated().data)

    def test_match_versions(self):
        packages = PackagesMetadataList([
            PackageMetadata(name="test", current="1.2.3", verb="list"),
            PackageMetadata(name="test2", current="2.3.4", verb="list"),
            PackageMetadata(name="test", current="3.4.5", verb="version"),
            PackageMetadata(name="test3", current="1.0.0", verb="list"),
        ])
        self.assertListEqual([packages[0], packages[1], packages[2]],
                             packages.show(["test", "test2"]).data)


if __name__ == '__main__':
    unittest.main()
