#!/usr/bin/env python
"""List of imported files, including meta data"""
from collections import UserList

from ..helpers import group_by_host, group_by_verb
from ..pretty import pretty

from .file_info import FileMetadata


class FileInfosList(UserList[FileMetadata]):
    """List of imported files, including meta data"""

    def hosts(self):
        """Deduplicated list of hosts for which files are present in the list"""
        return sorted(list({file.host for file in self}))

    def __str__(self):
        result = ""
        for host, package_infos_by_host in group_by_host(self):
            if host == "localhost":
                continue
            result += "\n  " + pretty.hostname(host)
            for verb, infos_by_verb in group_by_verb(package_infos_by_host):
                result += "\n    " + pretty.verb(verb)
                for file in infos_by_verb:
                    result += "\n      " + file.__str__()
            result += "\n"
        return result
