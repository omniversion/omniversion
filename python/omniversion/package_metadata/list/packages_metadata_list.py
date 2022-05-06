#!/usr/bin/env python
"""A list of dependencies, i.e. software packages"""
from collections import UserList

from ..package_metadata import PackageMetadata
from omniversion.pretty import pretty
from ...helpers import group_by_host


class PackagesMetadataList(UserList[PackageMetadata]):
    """A list of software package mate data infos"""

    def __str__(self):
        """Human-readable description of each dependency"""
        num_items = len(self)
        if num_items > 0:
            table_items = [
                f'\t{(item.host or "").ljust(12)}'
                + f'\t{(item.current or "").ljust(20)}'
                + f'\t{(item.package_manager or "").ljust(12)}'
                for item in self
            ]
            return (
                    f'{num_items} version{"" if num_items == 1 else "s"} found\n'
                    + "\n".join(table_items)
            )
        return pretty.traffic_light("No versions found", "red")

    def summary(self):
        """Summary of dependency counts grouped by host"""
        result = ""
        for host, items in group_by_host(self):
            result += "\n  " + pretty.hostname(host) + "\n"
            result += "    " + pretty.dependency_count(len(list(items)))
        return result
