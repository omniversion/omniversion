"""Test the data module"""
import unittest

from omniversion.models import Advisory


class AdvisoryTestCase(unittest.TestCase):
    def test_str(self):
        self.assertEqual("known vulnerability", Advisory().__str__())
        self.assertEqual("known vulnerability (severity `high`)", Advisory(severity="high").__str__())


if __name__ == '__main__':
    unittest.main()
