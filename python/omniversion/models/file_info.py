#!/usr/bin/env python
import time
from datetime import datetime
from dataclasses import dataclass
from typing import Dict, List, Optional, Union

from ..format import format
from ..constants import STALE_DATA_THRESHOLD_IN_SECONDS
from .dependency import Dependency


@dataclass
class OmniversionFileInfo:
    data: Optional[List[Dependency]]
    name: str
    host: str
    pm: str
    verb: str
    time: datetime
    path: str

    @property
    def is_stale(self):
        return time.time() > STALE_DATA_THRESHOLD_IN_SECONDS + self.time

    @property
    def has_data(self):
        return self.data is not None

    @property
    def num_entries(self):
        if self.data is None:
            return 0
        return len(self.data)

    def __str__(self):
        entries_text = f'1 entry' if self.num_entries == 1 else f'{self.num_entries} entries'
        if self.is_stale:
            return format.traffic_light(f'Stale data loaded for {self.pm} ({entries_text})', "amber")
        else:
            if self.has_data:
                return format.traffic_light(f'No entries found for {self.pm}', "amber")
            else:
                return format.traffic_light(f'Recent data loaded for {self.pm} ({entries_text})')
