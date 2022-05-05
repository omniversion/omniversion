#!/usr/bin/env python
"""Result of a versions match, i.e. a list of installations with corresponding versions"""
from .package_infos_list import PackageInfosList
from omniversion.pretty import pretty
from ...helpers.group_by import group_by_host, group_by_pm


class AvailableUpdates(PackageInfosList):
    """List of packages for which a newer version is available"""
    def __str__(self):
        """Human-readable description of the available update"""
        result = ""
        for host, items_for_host in group_by_host(self):
            result += "\n  " + pretty.hostname(host) + "\n"
            for package_manager, items_for_pm in group_by_pm(items_for_host):
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
