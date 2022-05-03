#!/usr/bin/env python
"""Common types used by other modules"""
from enum import Enum

PackageManager = Enum("PackageManager", "apt npm rubygems rvm")
