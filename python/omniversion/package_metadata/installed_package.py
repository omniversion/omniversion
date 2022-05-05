from dataclasses import dataclass


@dataclass
class InstalledPackage:
    """InstalledPackage is an installation of a software package
        in a particular location on a particular host.
    """

    location: str | None = None
    version: str | None = None
