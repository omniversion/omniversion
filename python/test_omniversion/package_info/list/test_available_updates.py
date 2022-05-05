import unittest

from omniversion import AvailableUpdates


class AvailableUpdatesTestCase(unittest.TestCase):
    def test_pretty_print(self):
        up_to_date = AvailableUpdates([])
        self.assertIn("up-to-date", up_to_date.__str__())

        self.assertIn("up-to-date", up_to_date.__str__())

        self.assertEqual(True, False)  # add assertion here


if __name__ == '__main__':
    unittest.main()
