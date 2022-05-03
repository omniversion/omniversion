#!/usr/bin/env python
"""List of imported files, including meta data"""
from dataclasses import dataclass
from itertools import groupby
from typing import List
from ..pretty import pretty

from .omniversion_file_info import OmniversionFileInfo


@dataclass
class OmniversionFileInfos:
    files: List[OmniversionFileInfo]

    def hosts(self):
        return list({file.host for file in self.files})

    def __str__(self):
        sorted_files = sorted(self.files, key=lambda file: file.host)
        grouped_files = groupby(sorted_files, lambda file: file.host)
        result = ""
        for host, files in grouped_files:
            if host == "localhost":
                continue
            result += "\n  " + pretty.hostname(host) + "\n"
            sorted_verbs = sorted(files, key=lambda file: file.verb)
            grouped_verbs = groupby(sorted_verbs, lambda file: file.verb)
            for verb, files_for_verb in grouped_verbs:
                result += "    " + pretty.verb(verb) + "\n"
                for file in files_for_verb:
                    result += "      " + file.__str__() + "\n"
        return result
