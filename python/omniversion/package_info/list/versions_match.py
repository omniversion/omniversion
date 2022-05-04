#!/usr/bin/env python
"""List of installations of one particular package"""
from .package_infos_list import PackageInfosList
from omniversion.pretty import pretty


class VersionsMatch(PackageInfosList):
    """List of installations of one particular package"""
    package_name: str
    display_name: str
    match: bool
    expected_num: int | None

    def __init__(
            self,
            data: PackageInfosList,
            package_name: str,
            display_name: str | None = None,
            expected_num: int | None = None,
    ):
        """Initialize the dependency match"""
        super().__init__(data)
        self.package_name = package_name
        self.display_name = display_name or package_name
        self.expected_num = expected_num

    @property
    def is_consistent(self):
        """Do all versions match?"""
        if len(self.deduplicated_versions()) == 1:
            return True
        return False

    @property
    def has_all_expected_items(self):
        """Does the number of dependencies match the expected value?"""
        return self.expected_num is None or len(self.data) == self.expected_num

    def __str__(self):
        """Human-readable description of the version match"""
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
        """All dependency versions with duplicates removed"""
        return {dependency.version for dependency in self.data}
