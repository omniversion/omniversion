#!/usr/bin/env python
"""A list of dependencies, i.e. software packages"""
from dataclasses import dataclass
from itertools import groupby
from dacite import from_dict

from python.omniversion.dependency.common.dependency import Dependency
from ...pretty import pretty


@dataclass
class Dependencies:
    """A list of dependencies, i.e. software packages"""
    data: list[Dependency]

    def __str__(self):
        """Human-readable description of each dependency"""
        num_items = len(self.data)
        if num_items > 0:
            table_items = [
                f'\t{(item.host or "").ljust(12)}'
                + f'\t{(item.version or "").ljust(20)}'
                + f'\t{(item.pm or "").ljust(12)}'
                for item in self.data
            ]
            return (
                    f'{num_items} version{"" if num_items == 1 else "s"} found\n'
                    + "\n".join(table_items)
            )
        return pretty.traffic_light("No versions found", "red")

    def overview(self):
        """Summary of dependency counts grouped by host"""
        sorted_dependencies = sorted(self.data, key=lambda dependency: dependency.host)
        grouped_dependencies = groupby(
            sorted_dependencies, lambda dependency: dependency.host
        )
        result = ""
        for host, items in grouped_dependencies:
            result += "\n  " + pretty.hostname(host) + "\n"
            result += "    " + pretty.dependency_count(len(list(items)))
        return result

    @staticmethod
    def from_list(list_data: list[str]):
        """Create a list of dependencies from """
        return Dependencies(list(
            map(lambda item: from_dict(data_class=Dependency, data=item), list_data)
        ))
