import unittest

from omniversion import PackageMetadata, Vulnerabilities, Advisory


class VulnerabilitiesTestCase(unittest.TestCase):
    def test_pretty_print_vulnerabilities(self):
        self.assertIn("No vulnerabilities found", Vulnerabilities([]).__str__())

        vulnerabilities = Vulnerabilities([
            PackageMetadata(package_manager="test_pm",
                            name="package_name",
                            host="test_host",
                            advisories=[
                                Advisory(severity="low")
                            ])
        ])
        self.assertIn("test_host", vulnerabilities.__str__())
        self.assertIn("test_pm", vulnerabilities.__str__())
        self.assertIn("package_name", vulnerabilities.__str__())
        self.assertIn("One vulnerability found", vulnerabilities.__str__())

        vulnerabilities = Vulnerabilities([
            PackageMetadata(package_manager="test_pm",
                            name="package_name",
                            host="test_host",
                            advisories=[
                                Advisory(severity="critical"),
                                Advisory(severity="low")
                            ])
        ])
        self.assertIn("test_host", vulnerabilities.__str__())
        self.assertIn("test_pm", vulnerabilities.__str__())
        self.assertIn("package_name", vulnerabilities.__str__())
        self.assertIn("2 vulnerabilities found", vulnerabilities.__str__())


if __name__ == '__main__':
    unittest.main()
