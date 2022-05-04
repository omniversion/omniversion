#!/usr/bin/env python
"""A generic dependency"""
from dataclasses import dataclass, field

from omniversion.package_info.advisory import Advisory
from omniversion.pretty import pretty


@dataclass
class PackageInfo:  # pylint: disable=too-many-instance-attributes
    """A generic dependency"""
    host: str | None = None
    name: str | None = None
    pm: str | None = None  # pylint: disable=invalid-name
    version: str | None = None

    advisories: list[Advisory] = field(default_factory=lambda: [])
    architecture: str | None = None
    author: str | None = None
    description: str | None = None
    homepage: str | None = None
    id: str | None = None  # pylint: disable=invalid-name
    latest: str | None = None
    license: str | None = None
    sources: list[str] = field(default_factory=lambda: [])
    wanted: str | None = None

    def __str__(self):
        """Human-readable description of the dependency"""
        name = pretty.black_on_white(self.name)
        version = pretty.white(self.version)
        if self.advisories is not None and len(self.advisories) > 0:
            return f"package {name} has version {version} with {self.advisories[0]}"
        return f"package {name} has version {version}"
