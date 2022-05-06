#!/usr/bin/env python
"""A vulnerability notice"""
from dataclasses import dataclass


@dataclass
class Advisory:  # pylint: disable=too-many-instance-attributes
    """A vulnerability notice"""
    access: str | None = None
    cvss_score: float | None = None
    identifier: int | None = None
    overview: str | None = None
    patched_versions: str | None = None
    recommendation: str | None = None
    references: str | None = None
    severity: str | None = None
    title: str | None = None
    url: str | None = None
    vulnerable_versions: str | None = None
