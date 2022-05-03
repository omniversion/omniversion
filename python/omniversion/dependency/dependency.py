#!/usr/bin/env python
from dataclasses import dataclass
from typing import List, Optional
from dacite import from_dict

from .audit.advisory import Advisory
from ..pretty import pretty


@dataclass
class Dependency:  # pylint: disable=too-many-instance-attributes
    host: Optional[str]
    name: Optional[str]
    pm: Optional[str]  # pylint: disable=invalid-name
    version: Optional[str]

    advisories: Optional[List[Advisory]] = None
    architecture: Optional[str] = None
    author: Optional[str] = None
    description: Optional[str] = None
    homepage: Optional[str] = None
    id: Optional[str] = None  # pylint: disable=invalid-name
    latest: Optional[str] = None
    license: Optional[str] = None
    sources: Optional[List[str]] = None
    wanted: Optional[str] = None

    def __str__(self):
        name = pretty.white_on_black(self.name)
        version = pretty.white(self.version)
        if self.advisories is not None and len(self.advisories) > 0:
            return f"package {name} has version {version} with {self.advisories[0]}"
        return f"package {name} has version {version}"

    @staticmethod
    def from_list(list_data: List[str]):
        return list(
            map(lambda item: from_dict(data_class=Dependency, data=item), list_data)
        )
