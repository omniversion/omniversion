#!/usr/bin/env python
"""Helper for loading omniversion files"""
import os
from typing import Callable
import yaml

from omniversion.package_info import PackageInfosList
from omniversion.file_info import FileInfo

AVAILABLE_VERBS = ["audit", "list", "refresh", "outdated", "version"]


def load_file(file_path: str) -> tuple[any, float]:
    """load an omniversion file containing yaml data"""
    try:
        with open(file_path, encoding="utf8") as file:
            return yaml.safe_load(file), os.stat(file_path).st_ctime
    except yaml.YAMLError:
        return None, 0
    except FileNotFoundError:
        return None, 0


def load_data(
        base_path: str,
        add_file: Callable[[FileInfo], None],
        hosts: list[str] | None = None,
        package_managers: list[str] | None = None,
        verbs: list[str] | None = None,
) -> None:
    """load all omniversion files in the base path"""
    # we look for subdirectories containing data for a particular host
    for host in [
        directory
        for directory in os.listdir(base_path)
        if os.path.isdir(os.path.join(base_path, directory))
    ]:
        if hosts is not None and host not in hosts:
            continue
        host_path = os.path.join(base_path, host)
        package_manager_dirs = [
            directory
            for directory in os.listdir(host_path)
            if os.path.isdir(os.path.join(host_path, directory))
        ]
        for package_manager in package_manager_dirs:
            if package_managers is not None and package_manager not in package_managers:
                continue
            for verb in AVAILABLE_VERBS:
                if verbs is not None and verb not in verbs:
                    continue
                process_file(verb, host, host_path, package_manager, add_file)


def process_file(
        verb: str,
        host: str,
        host_path: str,
        package_manager: str,
        add_file: Callable[[FileInfo], None]
) -> None:
    """load the file data and hand an `FileInfo` object to the callback"""
    file_name = verb + ".omniversion.yaml"
    file_path = os.path.join(host_path, package_manager, file_name)
    if os.path.exists(file_path):
        file_data, time = load_file(file_path)
        if file_data is None:
            add_file(
                FileInfo(
                    None, file_name, host, package_manager, verb, time, file_path
                )
            )
        else:
            for item in file_data:
                item["pm"] = package_manager
                item["host"] = host
            add_file(
                FileInfo(
                    PackageInfosList.from_list(file_data),
                    file_name,
                    host,
                    package_manager,
                    verb,
                    time,
                    file_path,
                )
            )
