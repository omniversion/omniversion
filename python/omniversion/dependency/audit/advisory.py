#!/usr/bin/env python
from dataclasses import dataclass
from typing import Optional
from ...pretty import pretty


@dataclass
class Advisory:  # pylint: disable=too-many-instance-attributes
    access: Optional[str]
    cvss_score: Optional[float]
    identifier: Optional[int]
    overview: Optional[str]
    patched_versions: Optional[str]
    recommendation: Optional[str]
    references: Optional[str]
    severity: Optional[str]
    title: Optional[str]
    url: Optional[str]
    vulnerable_versions: Optional[str]

    def __str__(self):
        severity = pretty.severity(self.severity)
        version = pretty.bright_on_lightblack(self.patched_versions)
        return f"a known vulnerability of severity {severity} and should be updated to {version}"
