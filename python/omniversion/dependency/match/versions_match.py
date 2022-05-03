#!/usr/bin/env python
from typing import List, Optional
from dataclasses import dataclass
from ..dependency import Dependency
from ..list.dependencies import Dependencies
from ...pretty import pretty


@dataclass
class VersionsMatch(Dependencies):
    package_name: str
    display_name: str
    match: bool
    expected_num: Optional[int]

    def __init__(
            self,
            data: List[Dependency],
            package_name: str,
            display_name: Optional[str] = None,
            expected_num: Optional[int] = None,
    ):
        super().__init__(data)
        self.package_name = package_name
        self.display_name = display_name or package_name
        self.expected_num = expected_num

    @property
    def is_consistent(self):
        if len(self.deduplicated_versions()) == 1:
            return True
        return False

    @property
    def has_all_expected_items(self):
        return self.expected_num is None or len(self.data) == self.expected_num

    def __str__(self):
        table_items = [
            f'\t{(item.version or "").ljust(20)}'
            + f'\t{(item.host or "").ljust(12)}'
            + f'\t{(item.pm or "").ljust(12)}'
            + f'\t{item.name or ""}'
            for item in self.data
        ]
        if len(self.deduplicated_versions()) == 0:
            return pretty.traffic_light(f"No {self.display_name} versions found", "red")
        if len(self.deduplicated_versions()) == 1:
            return pretty.traffic_light(
                f"Only one {self.display_name} version found\n"
                + "\n".join(table_items),
                "red",
            )
        if self.is_consistent:
            if self.has_all_expected_items:
                return pretty.traffic_light(
                    f"{self.display_name.capitalize()} versions match", "green"
                )
            return pretty.traffic_light(
                f"Only {len(self.data)} installations of {self.display_name} found"
                + f" (expected {self.expected_num})",
                "amber",
            )
        if self.has_all_expected_items:
            return pretty.traffic_light(
                f"{self.display_name.capitalize()} versions mismatch\n"
                + "\n".join(table_items),
                "red",
            )
        return pretty.traffic_light(
            f"{self.display_name.capitalize()} versions mismatch"
            + f" (only {len(self.data)} of {self.expected_num} installations found)\n"
            + "\n".join(table_items),
            "red",
        )

    def deduplicated_versions(self):
        return {dependency.version for dependency in self.data}
