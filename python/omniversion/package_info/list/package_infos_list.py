#!/usr/bin/env python
"""A list of dependencies, i.e. software packages"""
from collections import UserList
from itertools import groupby
from typing import Any

from dacite import from_dict

from ..package_info import PackageInfo
from omniversion.pretty import pretty


class PackageInfosList(UserList[PackageInfo]):
    """A list of software package mate data infos"""

    def __str__(self):
        """Human-readable description of each dependency"""
        num_items = len(self)
        if num_items > 0:
            table_items = [
                f'\t{(item.host or "").ljust(12)}'
                + f'\t{(item.version or "").ljust(20)}'
                + f'\t{(item.pm or "").ljust(12)}'
                for item in self
            ]
            return (
                    f'{num_items} version{"" if num_items == 1 else "s"} found\n'
                    + "\n".join(table_items)
            )
        return pretty.traffic_light("No versions found", "red")

    def overview(self):
        """Summary of dependency counts grouped by host"""
        sorted_dependencies = sorted(self, key=lambda package_info: package_info.host)
        grouped_dependencies = groupby(
            sorted_dependencies, lambda dependency: dependency.host
        )
        result = ""
        for host, items in grouped_dependencies:
            result += "\n  " + pretty.hostname(host) + "\n"
            result += "    " + pretty.dependency_count(len(list(items)))
        return result

    @staticmethod
    def from_list(list_data: list[dict[str, Any]]):
        """Create from a list of package infos"""
        return PackageInfosList([from_dict(data_class=PackageInfo, data=item) for item in list_data])
