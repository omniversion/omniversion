#!/usr/bin/env python
import os
import time
import re

from typing import List, Optional, Union

import __main__


from ..dependency import AvailableUpdates, Dependency, Dependencies, VersionsMatch, Vulnerabilities
from ..omniversion_file import OmniversionFileInfo, OmniversionFileInfos

from ..pretty import pretty
from ..loader import load_data


class OmniversionData:
    _files: List[OmniversionFileInfo]

    def __init__(self, base_path: str):
        self._files = []
        load_data(base_path, self._files.append)

    def __str__(self):
        return pretty.file_count(len(self._files))

    def file_infos(self):
        return OmniversionFileInfos(self._files)

    def hosts(self):
        return list({file.host for file in self._files})

    def items(
            self,
            verb: Union[str, List[str]] = "list",
            host: Optional[str] = None,
            package_manager: Optional[str] = None,
            package_name: Optional[Union[str, List[str]]] = None,
    ):
        def file_condition(file):
            if file.data is None:
                return False
            if host is not None and file.host != host:
                return False
            if package_manager is not None and file.package_manager != package_manager:
                return False
            if isinstance(verb, list):
                return file.verb in verb
            return file.verb == verb

        files_with_dependencies_data = [
            file.data for file in self._files if file_condition(file)
        ]
        all_items = [
            item for file_data in files_with_dependencies_data for item in file_data
        ]

        def package_condition(package):
            if package_name is None:
                return True
            if isinstance(package_name, list):
                return package.name in package_name
            return package.name == package_name

        return [item for item in all_items if package_condition(item)]

    def vulnerabilities(
            self,
            host: Optional[str] = None,
            package_manager: Optional[str] = None,
            package_name: Optional[Union[str, List[str]]] = None,
    ):
        return Vulnerabilities(self.items("audit", host, package_manager, package_name))

    def dependencies(
            self,
            host: Optional[str] = None,
            package_manager: Optional[str] = None,
            package_name: Optional[Union[str, List[str]]] = None,
    ):
        return Dependencies(self.items(["list", "version"], host, package_manager, package_name))

    def available_updates(
            self,
            host: Optional[str] = None,
            package_manager: Optional[str] = None,
            package_name: Optional[Union[str, List[str]]] = None,
    ):
        return AvailableUpdates(self.items("outdated", host, package_manager, package_name))

    def match_versions(
            self, package_name: Union[str, List[str]], display_name: Optional[str] = None
    ):
        return VersionsMatch(
            self.items(["list", "version"], package_name=package_name),
            package_name,
            display_name,
        )

    def add_local_config_value(
            self, file_path, regex: str, package: Optional[str] = None
    ):
        script_path = os.path.dirname(os.path.abspath(__main__.__file__))
        absolute_file_path = os.path.realpath(os.path.join(script_path, file_path))
        with open(absolute_file_path, encoding="utf8") as file:
            matches = re.compile(regex).finditer(file.read())
            for match in matches:
                version = match.group("version")
                package_name = package
                if package_name is None:
                    package_name = match.group("name")
                dependency = Dependency(
                    host="localhost",
                    name=package_name,
                    pm="local file",
                    version=version,
                )
                file_name = os.path.basename(absolute_file_path)
                file_info = OmniversionFileInfo(
                    [dependency],
                    file_name,
                    "localhost",
                    "local file",
                    "list",
                    time.time(),
                    file_path,
                )
                self._files.append(file_info)
