import unittest

from omniversion.file_info import FileMetadata, FileInfosList
from omniversion.helpers import group_by_host


class GroupByCase(unittest.TestCase):
    def test_group_by_host(self):
        file_info1 = FileMetadata(host="test2", name="package1")
        file_info2 = FileMetadata(host="test1", name="package2")
        file_info3 = FileMetadata(host="test2", name="package3")
        file_infos = FileInfosList([file_info1, file_info2, file_info3])
        grouped_infos = [[host, list(host_items)] for host, host_items in group_by_host(file_infos)]
        self.assertEqual("test1", grouped_infos[0][0])
        self.assertEqual(1, len(grouped_infos[0][1]))
        self.assertEqual("test2", grouped_infos[1][0])
        self.assertEqual(2, len(grouped_infos[1][1]))

if __name__ == '__main__':
    unittest.main()
