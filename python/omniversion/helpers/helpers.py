from typing import Any, Optional, Union, List


def apply_filters(items: List[Any],
                  host: Optional[Union[str, List[str]]] = None,
                  package_manager: Optional[str] = None,
                  package_name: Optional[Union[str, List[str]]] = None,
                  verb: Optional[Union[str, List[str]]] = None,
                  ):
    def matches(value: str, filter_values: Union[str, List[str]]) -> bool:
        if isinstance(filter_values, list):
            return value in filter_values
        return value == filter_values

    def filter_condition(item: Any) -> bool:
        if host is not None and not matches(item.host, host):
            return False
        if package_manager is not None and not matches(item.package_manager, package_manager):
            return False
        if package_name is not None and not matches(item.package_name, package_name):
            return False
        if verb is not None and not matches(item.verb, verb):
            return False
        return True

    return [item for item in items if filter_condition(item)]
