from collections import UserList
from typing import Callable, TypeAlias
from itertools import groupby

from omniversion.file_info import FileInfo
from omniversion.package_info import PackageInfo

InfoType: TypeAlias = PackageInfo | FileInfo
InfosListType: TypeAlias = UserList[PackageInfo] | UserList[FileInfo]


def group_by(data: InfosListType, extractor: Callable[[InfoType], str]) -> list[str, list[InfoType]]:
    sorted_data = sorted(data, key=extractor)
    return list(groupby(sorted_data, extractor))


def group_by_host(data: InfosListType) -> [str, str, list[InfoType]]:
    return group_by(data, lambda item: item.host)


def group_by_pm(data: InfosListType) -> [str, str, list[InfoType]]:
    return group_by(data, lambda item: item.host)


def group_by_verb(data: InfosListType) -> [str, str, list[InfoType]]:
    return group_by(data, lambda item: item.verb)


def group_by_host_and_pm(data: InfosListType) -> [str, str, list[InfoType]]:
    return [[host, pm, item] for host, host_items in group_by_host(data) for pm, item in group_by_pm(host_items)]


def group_by_host_and_verb(data: InfosListType) -> [str, str, list[InfoType]]:
    return [[host, verb, item] for host, host_items in group_by_host(data) for verb, item in group_by_verb(host_items)]
