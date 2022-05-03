#!/usr/bin/env python
from dataclasses import dataclass
from dacite import from_dict
from typing import List, Optional
from ..format import format


@dataclass
class Advisory:
    access: Optional[str]
    cvss_score: Optional[float]
    id: Optional[int]
    overview: Optional[str]
    patched_versions: Optional[str]
    recommendation: Optional[str]
    references: Optional[str]
    severity: Optional[str]
    title: Optional[str]
    url: Optional[str]
    vulnerable_versions: Optional[str]

    def __str__(self):
        return f'a known vulnerability of severity {format.severity(self.severity)} and should be updated to {format.bright_on_lightblack(self.patched_versions)}'
