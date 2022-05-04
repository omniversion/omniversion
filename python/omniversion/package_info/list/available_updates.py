#!/usr/bin/env python
"""Result of a versions match, i.e. a list of installations with corresponding versions"""
from itertools import groupby

from .package_infos_list import PackageInfosList
from python.omniversion.pretty import pretty


class AvailableUpdates(PackageInfosList):
    """List of packages for which a newer version is available"""
    def __str__(self):
        """Human-readable description of the available update"""
        sorted_items = sorted(self.data, key=lambda item: item.host)
        grouped_items = groupby(sorted_items, lambda item: item.host)
        result = ""
        for host, items in grouped_items:
            result += "\n  " + pretty.hostname(host) + "\n"
            sorted_pms = sorted(items, key=lambda item: item.pm)
            grouped_pms = groupby(sorted_pms, lambda item: item.pm)
            for package_manager, items_for_pm in grouped_pms:
                result += "    " + pretty.package_manager(package_manager) + "\n"
                updates = [item for item in items_for_pm if item.version is not None]
                not_installed = [item for item in items_for_pm if item.version is None]
                if len(updates) == 0:
                    if len(not_installed) == 0:
                        result += "      " + pretty.traffic_light("up-to-date") + "\n"
                    else:
                        result += (
                            "      "
                            + pretty.traffic_light(
                                f"{len(not_installed)} dependencies not installed",
                                "amber",
                            )
                            + "\n"
                        )
                else:
                    result += (
                        "      "
                        + pretty.traffic_light("updates available", "red")
                        + "\n"
                    )
                    for item in updates:
                        result += (
                            "        "
                            + f"update for {pretty.black_on_white(item.name)} available:"
                            + f" {pretty.white(item.version)} -> {pretty.white(item.latest)}"
                            + "\n"
                        )
        return result
