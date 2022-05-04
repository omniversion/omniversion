import unittest

from omniversion import PackageInfosList, PackageInfo


class InitializationTestCase(unittest.TestCase):
    def test_package_infos_list_can_be_initialized(self):
        package_info = PackageInfo(sources=['test1', 'test2'])
        package_infos_list = PackageInfosList([package_info])
        self.assertEqual(1, len(package_infos_list))


if __name__ == '__main__':
    unittest.main()
