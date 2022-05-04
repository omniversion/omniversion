#!/usr/bin/env python
"""A vulnerability notice"""
from dataclasses import dataclass
from python.omniversion.pretty import pretty


@dataclass
class Advisory:  # pylint: disable=too-many-instance-attributes
    """A vulnerability notice"""
    access: str | None
    cvss_score: float | None
    identifier: int | None
    overview: str | None
    patched_versions: str | None
    recommendation: str | None
    references: str | None
    severity: str | None
    title: str | None
    url: str | None
    vulnerable_versions: str | None

    def __str__(self):
        """Human-readable description of the security advisory"""
        severity = pretty.severity(self.severity)
        version = pretty.bright_on_lightblack(self.patched_versions)
        return f"a known vulnerability of severity {severity} and should be updated to {version}"
