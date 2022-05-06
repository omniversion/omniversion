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
