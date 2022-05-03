#!/usr/bin/env python
"""An imported file including meta data"""
import time
from dataclasses import dataclass
from typing import List, Optional

from ..dependency.dependency import Dependency
from ..pretty import pretty

STALENESS_THRESHOLD_IN_SECS: int = 60 * 60

@dataclass
class OmniversionFileInfo:
    """An imported file including meta data"""
    data: Optional[List[Dependency]]
    name: str
    host: str
    package_manager: str
    verb: str
    time: float
    path: str

    @property
    def is_stale(self):
        """Data that was fetched a long time ago is considered stale"""
        return time.time() > STALENESS_THRESHOLD_IN_SECS + self.time

    @property
    def has_data(self):
        """Does the file contain any parseable data at all?"""
        return self.data is not None

    @property
    def num_entries(self):
        """Number of entries contained in the file"""
        if self.data is None:
            return 0
        return len(self.data)

    def __str__(self):
        """Pretty string representation describing the file"""
        entries_text = (
            "1 entry" if self.num_entries == 1 else f"{self.num_entries} entries"
        )
        if self.is_stale:
            return pretty.traffic_light(
                f"Stale data loaded for {self.package_manager} ({entries_text})", "amber"
            )
        if self.has_data:
            return pretty.traffic_light(f"No entries found for {self.package_manager}", "amber")
        return pretty.traffic_light(
            f"Recent data loaded for {self.package_manager} ({entries_text})"
        )
