#!/usr/bin/env python
"""Root class used to load and extract all omniversion data"""
import os
import time
import re

from omniversion.package_metadata import AvailableUpdates, PackageMetadata, PackagesMetadataList, VersionsMatch, \
    Vulnerabilities
from omniversion.file_info import FileMetadata, FileInfosList

from omniversion.pretty import pretty
from omniversion.loader import load_data


class Data:
    """Root class used to load and extract all omniversion data"""
    files: FileInfosList

    def __init__(self,
                 base_path: str | None = None,
                 file_infos: FileInfosList | list[FileMetadata] | None = None,
                 package_managers: list[str] | None = None,
                 hosts: list[str] | None = None):
        """Initialize the root class"""
        self.files = FileInfosList() if file_infos is None else FileInfosList(file_infos)
        if base_path is not None:
            load_data(base_path, self.files.append, package_managers=package_managers, hosts=hosts)

    def __str__(self):
        """Human-readable summary of the data, counting loaded files"""
        return pretty.file_count(len(self.files))

    def hosts(self):
        """Deduplicated list of hosts for which files are present in the list"""
        return sorted(list({file.host for file in self.files}))

    def items(
            self,
            verb: str | list[str] = "list",
            host: str | None = None,
            package_manager: str | None = None,
            package_name: str | list[str] | None = None,
    ) -> list[PackageMetadata]:
        """List all dependencies matching the given criteria"""
        def file_condition(file: FileMetadata) -> bool:
            if file.data is None:
                return False
            if host is not None and file.host != host:
                return False
            if package_manager is not None and file.package_manager != package_manager:
                return False
            if isinstance(verb, list):
                return file.verb in verb
            return file.verb == verb

        files_with_dependencies_data: list[FileMetadata] = [
            file for file in self.files if file_condition(file)
        ]
        all_items = [
            item for file_info in files_with_dependencies_data for item in file_info.data
        ]

        def package_condition(package: PackageMetadata) -> bool:
            if package_name is None:
                return True
            if isinstance(package_name, list):
                return package.name in package_name
            return package.name == package_name

        return [item for item in all_items if package_condition(item)]

    def vulnerabilities(
            self,
            host: str | None = None,
            package_manager: str | None = None,
            package_name: str | list[str] | None = None,
    ):
        """List security vulnerabilities"""
        return Vulnerabilities(self.items("audit", host, package_manager, package_name))

    def list_packages(
            self,
            host: str | None = None,
            package_manager: str | None = None,
            package_name: str | list[str] | None = None,
    ):
        """List software packages"""
        return PackagesMetadataList(self.items(["list", "version"], host, package_manager, package_name))

    def available_updates(
            self,
            host: str | None = None,
            package_manager: str | None = None,
            package_name: str | list[str] | None = None,
    ):
        """List available updates"""
        return AvailableUpdates(self.items("outdated", host, package_manager, package_name))

    def match_versions(
            self, package_name: str | list[str], display_name: str | None = None
    ):
        """Match versions of all installations of a particular package"""
        return VersionsMatch(
            PackagesMetadataList(self.items(["list", "version"], package_name=package_name)),
            package_name,
            display_name,
        )

    def add_local_config_value(
            self, file_path, regex: str, name: str | None = None
    ):
        """Add dependency meta data from a local file"""
        absolute_file_path = os.path.realpath(file_path)
        with open(absolute_file_path, encoding="utf8") as file:
            try:
                matches = re.compile(regex).finditer(file.read())
                package_name = name
                for match in matches:
                    version = match.group("version")
                    if package_name is None:
                        package_name = match.group("name")
                    package_metadata = PackageMetadata(
                        host="localhost",
                        name=package_name,
                        package_manager="local file",
                        current=version,
                    )
                    file_name = os.path.basename(absolute_file_path)
                    file_info = FileMetadata(
                        PackagesMetadataList([package_metadata]),
                        version=None,
                        name=file_name,
                        host="localhost",
                        package_manager="local file",
                        verb="list",
                        time=time.time(),
                        path=file_path,
                    )
                    self.files.append(file_info)
            except IndexError:
                raise IndexError("Invalid regex. You need to provide a named group called `version` and either a name "
                                 "parameter or a named group called `name`")
