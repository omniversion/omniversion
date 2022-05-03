#!/usr/bin/env python
from typing import Dict, List, Optional, Union
from datetime import datetime
from dataclasses import dataclass
from .dependency import Dependency
from ..format import format


@dataclass
class VersionsMatch:
    items: List[Dependency]
    package_name: str
    display_name: str
    match: bool
    expected_num: Optional[int]

    def __init__(self, items: List[Dependency], package_name: str, display_name: Optional[str] = None, expected_num: Optional[int] = None):
        self.items = items
        self.package_name = package_name
        self.display_name = display_name or package_name
        self.expected_num = expected_num

    @property
    def is_consistent(self):
        if len(self.deduped_versions()) == 1:
            return True
        return False

    @property
    def has_all_expected_items(self):
        return self.expected_num is None or len(self.items) == self.expected_num

    def __str__(self):
        table_items = [f'\t{(item.version or "").ljust(20)}\t{(item.host or "").ljust(12)}\t{(item.pm or "").ljust(12)}\t{item.name or ""}' for item in self.items]
        if len(self.deduped_versions()) == 0:
            return format.traffic_light(f'No {self.display_name} versions found', "red")
        elif len(self.deduped_versions()) == 1:
            return format.traffic_light(f'Only one {self.display_name} version found\n' + "\n".join(table_items), "red")
        if self.is_consistent:
            if self.has_all_expected_items:
                return format.traffic_light(f'{self.display_name.capitalize()} versions match', "green")
            else:
                return format.traffic_light(f'Only {len(self.items)} installations of {self.display_name} found (expected {self.expected_num})', "amber")
        else:
            if self.has_all_expected_items:
                return format.traffic_light(f'{self.display_name.capitalize()} versions mismatch\n' + "\n".join(table_items), "red")
            else:
                return format.traffic_light(f'{self.display_name.capitalize()} versions mismatch (only {len(self.items)} of {self.expected_num} installations found)\n' + "\n".join(table_items), "red")

    def deduped_versions(self):
        return set([dependency.version for dependency in self.items])
