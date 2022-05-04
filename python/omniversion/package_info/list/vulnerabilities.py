#!/usr/bin/env python
"""A list of dependencies, optionally containing vulnerability notices"""
from .package_infos_list import PackageInfosList
from omniversion.pretty import pretty


class Vulnerabilities(PackageInfosList):
    """A list of dependencies, optionally containing vulnerability notices"""
    def __str__(self):
        """Human-readable description of the security advisories for each dependency"""
        advisories = [[data_item, advisory] for data_item in self for advisory in data_item.advisories]
        if len(advisories) > 1:
            audit_items = [
                f'\t{(data_item.name or "").ljust(20)}'
                + f'\t{(data_item.host or "").ljust(12)}'
                + f'\t{(data_item.pm or "").ljust(12)}'
                + f'\t{advisory.severity or ""}'
                for [data_item, advisory] in advisories
            ]
            return pretty.traffic_light(
                f"{len(advisories)} vulnerabilities found\n"
                + "\n".join(audit_items),
                "red",
            )
        if len(advisories) == 1:
            return pretty.traffic_light("One vulnerability found", "red")
        return pretty.traffic_light("No vulnerabilities found")
