#!/usr/bin/env python
from typing import List
from dataclasses import dataclass
from itertools import groupby

from ..dependency import Dependency
from ...pretty import pretty


@dataclass
class Dependencies:
    data: List[Dependency]

    def __str__(self):
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
        sorted_dependencies = sorted(self.data, key=lambda dependency: dependency.host)
        grouped_dependencies = groupby(
            sorted_dependencies, lambda dependency: dependency.host
        )
        result = ""
        for host, items in grouped_dependencies:
            result += "\n  " + pretty.hostname(host) + "\n"
            result += "    " + pretty.dependency_count(len(list(items)))
        return result
