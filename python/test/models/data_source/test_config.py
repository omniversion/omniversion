import unittest

from omniversion.models import ConfigDataSource


class ConfigDataSourceTestCase(unittest.TestCase):
    def test_str(self):
        self.assertEqual("No packages found in config", ConfigDataSource(num_packages=0, host="localhost",
                                                                         file_path="/test", regex="").__str__())
        self.assertEqual("1 package found in config", ConfigDataSource(num_packages=1, host="localhost",
                                                                       file_path="/test", regex="").__str__())
        self.assertEqual("5 packages found in config", ConfigDataSource(num_packages=5, host="localhost",
                                                                        file_path="/test", regex="").__str__())


if __name__ == '__main__':
    unittest.main()
