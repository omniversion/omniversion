from typing import Callable
from itertools import groupby


def group_by(data: any, extractor: Callable[[any], str]):
    sorted_data = sorted(data, key=extractor)
    return groupby(sorted_data, extractor)


def group_by_host(data: any) -> [str, str, list[any]]:
    return group_by(data, lambda item: item.host)


def group_by_pm(data: any) -> [str, str, list[any]]:
    return group_by(data, lambda item: item.package_manager)


def group_by_verb(data: any) -> [str, str, list[any]]:
    return group_by(data, lambda item: item.verb)


def group_by_host_and_pm(data: any) -> [str, str, list[any]]:
    return [[host, pm, item] for host, host_items in group_by_host(data) for pm, item in group_by_pm(host_items)]


def group_by_host_and_verb(data: any) -> [str, str, list[any]]:
    return [[host, verb, item] for host, host_items in group_by_host(data) for verb, item in group_by_verb(host_items)]
