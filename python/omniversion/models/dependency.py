#!/usr/bin/env python
from dataclasses import dataclass
from dacite import from_dict
from typing import List, Optional
from .advisory import Advisory
from ..format import format


@dataclass
class Dependency:
    host: Optional[str]
    name: Optional[str]
    pm: Optional[str]
    version: Optional[str]

    advisories: Optional[List[Advisory]] = None
    architecture: Optional[str] = None
    author: Optional[str] = None
    description: Optional[str] = None
    homepage: Optional[str] = None
    identifier: Optional[str] = None
    latest: Optional[str] = None
    license: Optional[str] = None
    sources: Optional[List[str]] = None
    wanted: Optional[str] = None

    def __str__(self):
        if self.advisories is not None and len(self.advisories) > 0:
            return f'package {format.white_on_black(self.name)} has version {format.white(self.version)} with {self.advisories[0]}'
        return f'package {format.white_on_black(self.name)} has version {format.white(self.version)}'

    @staticmethod
    def from_list(list_data: List[str]):
        return list(map(lambda item: from_dict(data_class=Dependency, data=item), list_data))
