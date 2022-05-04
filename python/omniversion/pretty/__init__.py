#!/usr/bin/env python
"""Pretty-print some data in a human-readable string"""
from colorama import init

from .pretty import black_on_white, bright_on_lightblack, white, cyan, severity, traffic_light, header, hostname, \
    verb, package_manager, dependency_count, updates_count, file_count, green

init(strip=None, convert=None, wrap=False)
