#!/usr/bin/env python
from enum import Enum

# data that is more than 1h out of date is considered stale
STALE_DATA_THRESHOLD_IN_SECONDS = 60 * 60

PackageManager = Enum("PackageManager", "apt npm rubygems rvm")
