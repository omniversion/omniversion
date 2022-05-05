#!/usr/bin/env python
"""Result of a versions match, i.e. a list of installations with corresponding versions"""
from .packages_metadata_list import PackagesMetadataList
from omniversion.pretty import pretty
from ...helpers.group_by import group_by_host, group_by_pm


class AvailableUpdates(PackagesMetadataList):
    """List of packages for which a newer version is available"""

    def __str__(self):
        """Human-readable description of the available update"""
        result = ""
        for host, items_for_host in group_by_host(self):
            result += "\n  " + pretty.hostname(host) + "\n"
            for package_manager, items_for_pm in group_by_pm(items_for_host):
                result += "    " + pretty.package_manager(package_manager) + "\n"
                items = list(items_for_pm)
                updates = [item for item in items if item.current is not None]
                not_installed = [item for item in items if item.current is None]
                if len(not_installed) > 0:
                    result += (
                            "      "
                            + pretty.traffic_light(
                                f"{len(not_installed)} dependencies not installed", "amber",
                            )
                            + "\n"
                    )
                if len(updates) > 0:
                    result += (
                            "      "
                            + pretty.traffic_light("updates available", "red")
                            + "\n"
                    )
                    for item in updates:
                        result += (
                                "        "
                                + f"update for {pretty.black_on_white(item.name)} available:"
                                + f" {pretty.white(item.current)} -> {pretty.white(item.latest)}"
                                + "\n"
                        )
        return result
