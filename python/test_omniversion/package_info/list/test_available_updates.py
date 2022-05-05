import unittest

from omniversion import AvailableUpdates, PackageMetadata


class AvailableUpdatesTestCase(unittest.TestCase):
    def test_pretty_print_available_updates(self):
        not_installed = AvailableUpdates([
            PackageMetadata(package_manager="test_pm", name="package_name", host="test_host")
        ])
        self.assertIn("test_host", not_installed.__str__())
        self.assertIn("test_pm", not_installed.__str__())
        self.assertIn("not installed", not_installed.__str__())
        self.assertNotIn("updates available", not_installed.__str__())

        updates_available = AvailableUpdates([
            PackageMetadata(package_manager="test_pm", name="package_name", host="test_host", current="1.2.3")
        ])
        self.assertIn("test_host", updates_available.__str__())
        self.assertIn("test_pm", updates_available.__str__())
        self.assertNotIn("not installed", updates_available.__str__())
        self.assertIn("updates available", updates_available.__str__())


if __name__ == '__main__':
    unittest.main()
