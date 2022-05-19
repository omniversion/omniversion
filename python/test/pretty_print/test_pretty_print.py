"""Test the pretty module"""
import os
import time
import unittest

from colorama import Fore, Back
from wheel.util import StringIO

from omniversion import Omniversion
from omniversion.models import FileDataSource, PackageMetadata, PackagesMetadataList, VersionsMatch, Vulnerabilities, \
    Advisory, AvailableUpdates, DataSources, ConfigDataSource
from omniversion.pretty_print import pformat, pprint
from omniversion.pretty_print.pretty_print import header, traffic_light


class PrettyPrintTestCase(unittest.TestCase):

    def test_traffic_light(self):
        self.assertIn(Fore.RED, traffic_light("test", "red"))
        self.assertIn(Fore.YELLOW, traffic_light("test", "amber"))
        self.assertIn(Fore.GREEN, traffic_light("test"))
        self.assertEqual("-", traffic_light("test", "unknown"))

    def test_header(self):
        self.assertIn(Back.LIGHTBLACK_EX, header("test"))

    def test_file_data_source(self):
        self.assertIn("Recent `test` data loaded for `test`", pformat(
            FileDataSource(num_packages=1, verb="test", package_manager="test", timestamp=time.time())))
        self.assertIn("Stale `test` data loaded for `test`", pformat(
            FileDataSource(num_packages=1, verb="test", package_manager="test", timestamp=time.time() - 60 * 60 * 2)))

        self.assertIn("2 packages", pformat(FileDataSource(num_packages=2)))
        self.assertIn("1 package", pformat(FileDataSource(num_packages=1)))

        self.assertIn("No `test` data", pformat(FileDataSource(num_packages=0, verb="test", timestamp=time.time())))
        self.assertIn("No valid `test` data", pformat(
            FileDataSource(num_packages=None, verb="test", timestamp=time.time())))
        self.assertIn("No valid `test` data", pformat(FileDataSource(verb="test")))

    def test_pretty_print_packages_metadata_list(self):
        packages_metadata_list = PackagesMetadataList([
            PackageMetadata(current="2.3.4"),
            PackageMetadata(current="1.2.3")
        ])
        self.assertIn("2", pformat(packages_metadata_list))

        packages_metadata_list = PackagesMetadataList()
        self.assertIn("0", pformat(packages_metadata_list))

    def test_pretty_print_versions_match(self):
        self.assertIn("No `test` versions found", pformat(VersionsMatch(data=[], display_name="test")))

        self.assertIn("Only one `test` version found", pformat(VersionsMatch(
            data=[PackageMetadata(name="test", current="1.0.0")], display_name="test")))

        self.assertIn("Versions mismatch", pformat(VersionsMatch(
            data=[
                PackageMetadata(name="test", current="1.0.0"),
                PackageMetadata(name="test", current="2.0.0"),
            ])))

        self.assertIn("only 2 of 3 found", pformat(VersionsMatch(
            expected_num=3,
            data=[
                PackageMetadata(name="test", current="1.0.0"),
                PackageMetadata(name="test", current="2.0.0"),
            ])))

        self.assertIn("Versions match", pformat(VersionsMatch(
            data=[
                PackageMetadata(name="test", current="1.0.0"),
                PackageMetadata(name="test", current="1.0.0"),
            ])))

        self.assertIn("only 2 of 3 found", pformat(VersionsMatch(
            expected_num=3,
            data=[
                PackageMetadata(name="test", current="1.0.0"),
                PackageMetadata(name="test", current="1.0.0"),
            ])))

    def test_pretty_print_vulnerabilities(self):
        self.assertIn("No vulnerabilities found", pformat(Vulnerabilities([])))

        vulnerabilities = Vulnerabilities([
            PackageMetadata(package_manager="test_pm",
                            name="package_name",
                            host="test_host",
                            advisories=[
                                Advisory(severity="low")
                            ])
        ])
        formatted_result = pformat(vulnerabilities)
        self.assertIn("test_host", formatted_result)
        self.assertIn("test_pm", formatted_result)
        self.assertIn("package_name", formatted_result)
        self.assertIn("One vulnerability found", formatted_result)

        vulnerabilities = Vulnerabilities([
            PackageMetadata(package_manager="test_pm",
                            name="package_name",
                            host="test_host",
                            advisories=[
                                Advisory(severity="critical"),
                                Advisory(severity="low")
                            ])
        ])
        formatted_result = pformat(vulnerabilities)
        self.assertIn("test_host", formatted_result)
        self.assertIn("test_pm", formatted_result)
        self.assertIn("package_name", formatted_result)
        self.assertIn("2 vulnerabilities found", formatted_result)

    def test_pretty_print_available_updates(self):
        self.assertIn("no package managers found", pformat(AvailableUpdates([
            PackageMetadata(name="package_name", host="test_host")
        ])))

        not_installed = AvailableUpdates([
            PackageMetadata(package_manager="test_pm", name="package_name", host="test_host")
        ])
        self.assertIn("test_pm", pformat(not_installed))
        self.assertIn("not installed", pformat(not_installed))
        self.assertNotIn("available", pformat(not_installed))

        updates_available = AvailableUpdates([
            PackageMetadata(package_manager="test_pm", name="package_name", host="test_host", current="1.2.3",
                            latest="2.5.4")
        ])
        self.assertIn("test_pm", pformat(updates_available))
        self.assertNotIn("not installed", pformat(updates_available))
        self.assertIn("available", pformat(updates_available))

    def test_pretty_print_omniversion(self):
        data_sources = DataSources()
        data_sources.files = [FileDataSource(num_packages=1, name="1", host="test1", verb="test"),
                              FileDataSource(name="2", host="test2", verb="test"),
                              FileDataSource(name="3", host="test3", verb="test")]
        data_sources.configs = [ConfigDataSource(host="test4", regex="test", file_path="/test"),
                                ConfigDataSource(host="test5", regex="test", file_path="/test")]
        omniversion = Omniversion(base_path=None, data_sources=data_sources)
        pretty_output = pformat(omniversion)
        self.assertIn("Stale `test` data loaded", pretty_output)

    def test_pretty_print_omniversion_test_vectors(self):
        omniversion = Omniversion(base_path=os.path.join(os.path.dirname(__file__), "../vectors"))
        self.assertIn("Data sources", pformat(omniversion))
        self.assertIn("`audit` data loaded for `test_pm` (2 packages)", pformat(omniversion))

    def test_pretty_config_data_source(self):
        config_data_source = ConfigDataSource(num_packages=0, host="test4", regex="test", file_path="/test")
        self.assertIn("No packages", pformat(config_data_source))

    def test_default(self):
        self.assertEqual("{'test': 'object'}", pformat({"test": "object"}))

    def test_pprint(self):
        stream = StringIO()
        pprint("test", stream=stream)
        self.assertIn("test", stream.getvalue())


if __name__ == '__main__':
    unittest.main()
