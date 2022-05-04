#!/usr/bin/env python
"""A generic dependency"""
from dataclasses import dataclass

from omniversion.package_info.advisory import Advisory
from python.omniversion.pretty import pretty


@dataclass
class PackageInfo:  # pylint: disable=too-many-instance-attributes
    """A generic dependency"""
    host: str | None
    name: str | None
    pm: str | None  # pylint: disable=invalid-name
    version: str | None

    advisories: list[Advisory] | None = None
    architecture: str | None = None
    author: str | None = None
    description: str | None = None
    homepage: str | None = None
    id: str | None = None  # pylint: disable=invalid-name
    latest: str | None = None
    license: str | None = None
    sources: list[str] | None = None
    wanted: str | None = None

    def __str__(self):
        """Human-readable description of the dependency"""
        name = pretty.black_on_white(self.name)
        version = pretty.white(self.version)
        if self.advisories is not None and len(self.advisories) > 0:
            return f"package {name} has version {version} with {self.advisories[0]}"
        return f"package {name} has version {version}"
