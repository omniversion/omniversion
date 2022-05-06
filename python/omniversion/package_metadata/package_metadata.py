#!/usr/bin/env python
"""A generic dependency"""
from dataclasses import dataclass, field

from .advisory import Advisory
from .installed_package import InstalledPackage


@dataclass
class PackageMetadata:  # pylint: disable=too-many-instance-attributes
    """PackageMetadata contains all available information about a particular software package,
including existing installations, security advisories and all sorts of other metadata.

If multiple versions of the same package need to be tracked
for a single combination of package manager and host,
this should be done using a single PackageMetadata structure.
By contrast, otherwise identical packages made available by different package managers
are considered different, as their metadata and/or content *may* indeed differ.
We also distinguish between package metadata on different hosts,
as installations are different and available versions
will depend on the configured repositories, firewall settings etc.
"""

    name: str | None = None
    aliases: list[str] | None = field(default_factory=lambda: [])

    package_manager: str | None = None
    host: str | None = None
    installations: list[InstalledPackage] = field(default_factory=lambda: [])

    current: str | None = None
    default: str | None = None
    latest: str | None = None
    wanted: str | None = None
    locked: str | None = None

    architecture: str | None = None
    author: str | None = None
    description: str | None = None
    homepage: str | None = None
    license: str | None = None
    sources: list[str] = field(default_factory=lambda: [])
    advisories: list[Advisory] = field(default_factory=lambda: [])
