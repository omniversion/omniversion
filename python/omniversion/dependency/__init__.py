#!/usr/bin/env python
"""Model classes for various types of dependency"""
from .audit import Advisory, Vulnerabilities
from .common import Dependencies, Dependency
from .match import VersionsMatch
from .outdated import AvailableUpdates
