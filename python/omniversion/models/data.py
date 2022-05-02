#!/usr/bin/env python
import os
import time
import re
from dacite import from_dict
from typing import Dict, List, Optional, Union
from datetime import datetime
import __main__

from .advisory import Advisory
from .available_updates import AvailableUpdates
from .dependency import Dependency
from .dependencies import Dependencies
from .versions_match import VersionsMatch
from .vulnerabilities import Vulnerabilities
from .file_info import OmniversionFileInfo
from .file_infos import OmniversionFileInfos

from ..constants import PackageManager
from ..format import format
from ..loader import load_data, load_file


class OmniversionData:
    _files: List[OmniversionFileInfo]

    def __init__(self, base_path: str):
        self._files = []
        load_data(base_path, lambda file: self._files.append(file))

    def __str__(self):
        return format.file_count(len(self.files))

    def file_infos(self):
        return OmniversionFileInfos(self._files)

    def hosts(self):
        all_hosts = [file.host for file in self._files]
        return list(set(all_hosts))

    def items(self, verb: Union[str, List[str]] = "list", host: Optional[str] = None, pm: Optional[str] = None, package_name: Optional[Union[str, List[str]]] = None):
        def file_condition(file):
            if file.data is None:
                return False
            if host is not None and file.host != host:
                return False
            if pm is not None and file.pm != pm:
                return False
            if type(verb) is list:
                return file.verb in verb
            return file.verb == verb
        files_with_dependencies_data = [file.data for file in self._files if file_condition(file)]
        all_items = [item for file_data in files_with_dependencies_data for item in file_data]
        def package_condition(package):
            if package_name is None:
                return True
            if type(package_name) is list:
                return package.name in package_name
            return package.name == package_name
        return [item for item in all_items if package_condition(item)]

    def vulnerabilities(self, host: Optional[str] = None, pm: Optional[str] = None, package_name: Optional[Union[str, List[str]]] = None):
        return Vulnerabilities(self.items("audit", host, pm, package_name))

    def dependencies(self, host: Optional[str] = None, pm: Optional[str] = None, package_name: Optional[Union[str, List[str]]] = None):
        return Dependencies(self.items(["list", "version"], host, pm, package_name))

    def available_updates(self, host: Optional[str] = None, pm: Optional[str] = None, package_name: Optional[Union[str, List[str]]] = None):
        return AvailableUpdates(self.items("outdated", host, pm, package_name))

    def match_versions(self, package_name: Union[str, List[str]], display_name: Optional[str] = None):
        return VersionsMatch(self.items(["list", "version"], package_name=package_name), package_name, display_name)

    def add_local_source(self, file_path: str, pm: PackageManager):
        list_command = f'omniversion run list {pm}'
        list_result = os.system(command)
        try:
             with open(file_path) as file:
                 return yaml.safe_load(file), os.stat(file_path).st_ctime
        except yaml.YAMLError:
            return None
        except FileNotFoundError:
            return None

    def add_local_config_value(self, file_path, regex: str, package: Optional[str] = None):
        compiled_regex = re.compile(regex)
        script_path = os.path.dirname(os.path.abspath(__main__.__file__))
        absolute_file_path = os.path.realpath(os.path.join(script_path, file_path))
        with open(absolute_file_path) as file:
            content = file.read()
            matches = compiled_regex.finditer(content)
            for match in matches:
                version = match.group('version')
                package_name = package
                if package_name is None:
                    package_name = match.group('name')
                dependency = Dependency(host="localhost", name=package_name, pm="local file", version=version)
                file_name = os.path.basename(absolute_file_path)
                update_time = time.time()
                file_info = OmniversionFileInfo([dependency], file_name, "localhost", "local file", "list", update_time, file_path)
                self._files.append(file_info)
