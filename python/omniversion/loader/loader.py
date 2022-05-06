#!/usr/bin/env python
"""Helper for loading omniversion files"""
import os
from typing import Callable
import yaml

from omniversion.package_metadata import PackagesMetadataList
from omniversion.file_info import FileMetadata

AVAILABLE_VERBS = ["audit", "list", "refresh", "outdated", "version"]


def load_data(
        base_path: str,
        add_file: Callable[[FileMetadata], None],
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
            dir_path = os.path.join(host_path, package_manager)
            if package_managers is not None and package_manager not in package_managers:
                continue
            for verb in AVAILABLE_VERBS:
                if verbs is not None and verb not in verbs:
                    continue
                process_file(verb, host, dir_path, package_manager, add_file)


def process_file(
        verb: str,
        host: str,
        host_path: str,
        package_manager: str,
        add_file: Callable[[FileMetadata], None]
) -> None:
    """load the file data and hand `FileMetadata` object to the callback"""
    file_name = verb + ".omniversion.yaml"
    file_path = os.path.join(host_path, file_name)
    if os.path.exists(file_path):
        version, packages_data, time = extract_yaml_data(file_path)
        if packages_data is None:
            add_file(
                FileMetadata(
                    None, version, file_name, host, package_manager, verb, time, file_path
                )
            )
        else:
            for package_data in packages_data:
                package_data["package_manager"] = package_manager
                package_data["host"] = host
            add_file(
                FileMetadata(
                    PackagesMetadataList(packages_data),
                    version,
                    file_name,
                    host,
                    package_manager,
                    verb,
                    time,
                    file_path,
                )
            )


def extract_yaml_data(file_path: str) -> tuple[str | None, list[any] | None, float | None]:
    """load an omniversion file containing yaml data"""
    time = None
    try:
        time = os.stat(file_path).st_ctime
        with open(file_path, encoding="utf8") as file:
            file_data = yaml.safe_load(file)
            return file_data["version"], file_data["items"], time
    except TypeError:
        return None, None, time
    except yaml.YAMLError:
        return None, None, time
    except FileNotFoundError:
        return None, None, time
