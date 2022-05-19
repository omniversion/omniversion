"""Pretty-print some data in a human-readable string"""
from typing import List
from pprint import PrettyPrinter

from colorama import Fore, Back, Style

from omniversion import Omniversion
from omniversion.models import DataSources, Host, FileDataSource, ConfigDataSource, PackagesMetadataList, \
    AvailableUpdates, Vulnerabilities, VersionsMatch, PackageMetadata
from omniversion.models.package_manager import PackageManager


def header(text):
    """Text styled like a header"""
    return f"\n{Back.LIGHTBLACK_EX} {text} {Style.RESET_ALL}"


def traffic_light(text, status="green"):
    """Green, amber or red text with leading symbol"""
    if status == "green":
        return f"{Fore.GREEN}[✔]{Style.DIM} {text} {Style.RESET_ALL}"
    if status == "amber":
        return f"{Fore.YELLOW}[-]{Style.DIM} {text} {Style.RESET_ALL}"
    if status == "red":
        return f"{Fore.RED}[✘]{Style.DIM} {text} {Style.RESET_ALL}"
    return "-"


def multiline_indent(multiline_string: str, levels: int = 2):
    return " " * levels + ("\n" + " " * levels).join(multiline_string.splitlines())


class OmniversionPrettyPrinter(PrettyPrinter):
    def format(self, obj: object, context, depth: int, level: int) -> tuple[str, bool, bool]:
        def recursive_format(value: object):
            return self.format(value, context=context, depth=depth, level=level - 1)[0]

        if isinstance(obj, Omniversion):
            def host_summary(host, dependencies):
                return f"{recursive_format(host)}\n{multiline_indent(recursive_format(dependencies))}\n\n"

            return "\n".join([
                header('Data sources'),
                multiline_indent(recursive_format(obj.sources)),
                header('Dependency count'),
                *[multiline_indent(host_summary(host, host.packages.ls())) for host in obj.hosts],
                header('Available updates'),
                *[multiline_indent(host_summary(host, host.packages.outdated())) for host in obj.hosts],
                header('Security audit'),
                *[multiline_indent(host_summary(host, host.packages.audit())) for host in obj.hosts],
            ]), False, False

        if isinstance(obj, DataSources):
            def host_summary(host_info: tuple[str, List[FileDataSource], List[ConfigDataSource]]):
                sources = [recursive_format(file) for file in host_info[1]] + \
                          [recursive_format(config) for config in host_info[2]]
                formatted_sources = "\n".join(sources)
                return f"{Back.BLUE} {host_info[0]} {Style.RESET_ALL}\n{multiline_indent(formatted_sources)}\n"

            return "\n".join([host_summary(host_info) for host_info in obj.host_infos]), False, False

        if isinstance(obj, Host):
            return "\n".join([
                f"{Back.BLUE} {obj.name} {Style.RESET_ALL}",
            ]), False, False

        if isinstance(obj, FileDataSource):
            color = "green"
            if obj.is_stale or obj.num_packages == 0:
                color = "amber"
            if obj.num_packages is None:
                color = "red"
            return traffic_light(obj, color), False, False

        if isinstance(obj, ConfigDataSource):
            color = "green"
            if obj.num_packages == 0:
                color = "amber"
            return traffic_light(obj, color), False, False

        if isinstance(obj, PackagesMetadataList):
            num_dependencies = len(obj)
            pretty_deps = "dependenc" + ("y" if num_dependencies == 1 else "ies")
            return f"{num_dependencies} {Style.DIM}{pretty_deps}{Style.RESET_ALL}", False, False

        if isinstance(obj, AvailableUpdates):
            def update_description(update: PackageMetadata):
                if update.current is None:
                    latest_version_text = f" version {Fore.WHITE}{update.latest} {Style.RESET_ALL}" \
                        if update.latest else ""
                    return f"package {Fore.BLACK}{Back.WHITE} `{update.name}` {Style.RESET_ALL}{latest_version_text}" \
                           f" can be installed"
                return f"update for {Fore.BLACK}{Back.WHITE} `{update.name}` {Style.RESET_ALL} available: " \
                       f"{Fore.WHITE}{update.current}{Style.RESET_ALL} -> {Fore.WHITE}{update.latest}{Style.RESET_ALL}"

            def package_manager_summary(pm: PackageManager):
                updates_for_pm = AvailableUpdates(pm.packages.data)
                summary = traffic_light(updates_for_pm.__repr__(), "green" if len(updates_for_pm) == 0 else "red")
                return "\n".join([
                    f"{Back.CYAN} {pm.identifier} {Style.RESET_ALL}",
                    multiline_indent(summary),
                    *[multiline_indent(multiline_indent(update_description(update))) for update in updates_for_pm]
                ])

            package_managers = PackageManager.list_from_packages(obj.data)
            if len(package_managers) == 0:
                return traffic_light('no package managers found', 'amber'), False, False
            return "\n".join([package_manager_summary(pm) for pm in package_managers]) + "\n", False, False

        if isinstance(obj, Vulnerabilities):
            advisories = [[data_item, advisory] for data_item in obj for advisory in data_item.advisories]
            audit_items = [
                f'\t{(data_item.name or "").ljust(20)}'
                + f'\t{(data_item.host or "").ljust(12)}'
                + f'\t{(data_item.package_manager or "").ljust(12)}'
                + f'\t{advisory.severity or ""}'
                for [data_item, advisory] in advisories
            ]
            return traffic_light(
                f"{obj}\n"
                + "\n".join(audit_items),
                "green" if len(advisories) == 0 else "red"), True, False

        if isinstance(obj, VersionsMatch):
            table_items = [
                f'\t{(item.current or "").ljust(20)}'
                + f'\t{(item.host or "").ljust(12)}'
                + f'\t{(item.package_manager or "").ljust(12)}'
                + f'\t{item.name or ""}'
                for item in obj.data
            ]
            color = "green"
            if len(obj.versions) <= 1 or not obj.is_consistent or not obj.has_all_expected_items:
                color = "red"
            return traffic_light(obj.__str__() + "\n".join(table_items), color), False, False

        return super().format(obj, context, depth, level)


def pprint(obj: object, indent=1, width=80, depth=None, compact=False, stream=None):
    OmniversionPrettyPrinter(indent=indent, width=width, depth=depth, compact=compact, stream=stream) \
        .pprint(obj)


def pformat(obj: object, indent=1, width=80, depth=None, compact=False) -> str:
    return OmniversionPrettyPrinter(indent=indent, width=width, depth=depth, compact=compact).pformat(obj)
