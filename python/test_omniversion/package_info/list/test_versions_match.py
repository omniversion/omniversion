import unittest

from omniversion import PackageMetadata, VersionsMatch, PackagesMetadataList


class VersionsMatchTestCase(unittest.TestCase):
    def test_pretty_print_versions_match(self):
        self.assertIn("No test versions found", VersionsMatch(
            data=PackagesMetadataList([]),
            package_name="test").__str__())

        self.assertIn("Only one test version found", VersionsMatch(
            data=PackagesMetadataList([PackageMetadata(name="test", current="1.0.0")]),
            package_name="test").__str__())

        self.assertIn("versions mismatch", VersionsMatch(
            data=PackagesMetadataList([
                PackageMetadata(name="test", current="1.0.0"),
                PackageMetadata(name="test", current="2.0.0"),
            ]),
            package_name="test").__str__())

        self.assertIn("only 2 of 3 installations found", VersionsMatch(
            expected_num=3,
            data=PackagesMetadataList([
                PackageMetadata(name="test", current="1.0.0"),
                PackageMetadata(name="test", current="2.0.0"),
            ]),
            package_name="test").__str__())

        self.assertIn("versions match", VersionsMatch(
            data=PackagesMetadataList([
                PackageMetadata(name="test", current="1.0.0"),
                PackageMetadata(name="test", current="1.0.0"),
            ]),
            package_name="test").__str__())

        self.assertIn("2 installations of test found (expected 3)", VersionsMatch(
            expected_num=3,
            data=PackagesMetadataList([
                PackageMetadata(name="test", current="1.0.0"),
                PackageMetadata(name="test", current="1.0.0"),
            ]),
            package_name="test").__str__())


if __name__ == '__main__':
    unittest.main()
