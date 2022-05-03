#!/usr/bin/env python
from dataclasses import dataclass
from itertools import groupby
from typing import List, Optional
from ..format import format

from .file_info import OmniversionFileInfo


@dataclass
class OmniversionFileInfos:
    files: List[OmniversionFileInfo]

    def hosts(self):
        return list(set([file.host for file in self.files]))

    def __str__(self):
        sorted_files = sorted(self.files, key=lambda file: file.host)
        grouped_files = groupby(sorted_files, lambda file: file.host)
        result = ""
        for host, files in grouped_files:
            if host == "localhost":
                continue
            result += "\n  " + format.hostname(host) + "\n"
            sorted_verbs = sorted(files, key=lambda file: file.verb)
            grouped_verbs = groupby(sorted_verbs, lambda file: file.verb)
            for verb, files_for_verb in grouped_verbs:
                result += "    " + format.verb(verb) + "\n"
                for file in files_for_verb:
                    result += "      " + file.__str__() + "\n"
        return result
