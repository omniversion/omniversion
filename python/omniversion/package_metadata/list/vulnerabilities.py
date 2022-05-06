#!/usr/bin/env python
"""A list of dependencies, optionally containing vulnerability notices"""
from .packages_metadata_list import PackagesMetadataList
from omniversion.pretty import pretty


class Vulnerabilities(PackagesMetadataList):
    """A list of dependencies, optionally containing vulnerability notices"""

    def __str__(self):
        """Human-readable description of the security advisories for each dependency"""
        advisories = [[data_item, advisory] for data_item in self for advisory in data_item.advisories]
        if len(advisories) == 0:
            return pretty.traffic_light("No vulnerabilities found")
        audit_items = [
            f'\t{(data_item.name or "").ljust(20)}'
            + f'\t{(data_item.host or "").ljust(12)}'
            + f'\t{(data_item.package_manager or "").ljust(12)}'
            + f'\t{advisory.severity or ""}'
            for [data_item, advisory] in advisories
        ]
        summary = f"{len(advisories)} vulnerabilities found" if len(advisories) > 1 else "One vulnerability found"
        return pretty.traffic_light(
            f"{summary}\n"
            + "\n".join(audit_items),
            "red")
