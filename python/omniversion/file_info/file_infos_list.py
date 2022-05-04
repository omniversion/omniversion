#!/usr/bin/env python
"""List of imported files, including meta data"""
from collections import UserList
from dataclasses import dataclass

from ..helpers.group_by import group_by_host, group_by_verb
from ..pretty import pretty

from .file_info import FileInfo


@dataclass
class FileInfosList(UserList[FileInfo]):
    """List of imported files, including meta data"""

    def hosts(self):
        """Deduplicated list of hosts for which files are present in the list"""
        return list({file.host for file in self})

    def __str__(self):
        result = ""
        for host, package_infos_by_host in group_by_host(self):
            if host == "localhost":
                continue
            result += "\n  " + pretty.hostname(host) + "\n"
            for verb, infos_by_verb in group_by_verb(package_infos_by_host):
                result += "    " + pretty.verb(verb) + "\n"
                for file in infos_by_verb:
                    result += "      " + file.__str__() + "\n"
        return result
