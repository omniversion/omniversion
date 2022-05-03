#!/usr/bin/env python
from typing import Dict, List, Optional, Union
from datetime import datetime
from dataclasses import dataclass
from itertools import groupby

from .dependency import Dependency
from ..format import format


@dataclass
class AvailableUpdates:
    items: List[Dependency]

    def __str__(self):
        sorted_items = sorted(self.items, key=lambda item: item.host)
        grouped_items = groupby(sorted_items, lambda item: item.host)
        result = ""
        for host, items in grouped_items:
            result += "\n  " + format.hostname(host) + "\n"
            sorted_pms = sorted(items, key=lambda item: item.pm)
            grouped_pms = groupby(sorted_pms, lambda item: item.pm)
            for pm, items_for_pm in grouped_pms:
                result += "    " + format.pm(pm) + "\n"
                updates = [item for item in items_for_pm if item.version is not None]
                not_installed = [item for item in items_for_pm if item.version is None]
                if len(updates) == 0:
                    if len(not_installed) == 0:
                        result += "      " + format.traffic_light(f'up-to-date') + "\n"
                    else:
                        result += "      " + format.traffic_light(f'{len(not_installed)} dependencies not installed', "amber") + "\n"
                else:
                    result += "      " + format.traffic_light(f'updates available', "red") + "\n"
                    for item in updates:
                        result += "        " + f'update for {format.white_on_black(item.name)} available: {format.white(item.version)} -> {format.white(item.latest)}' + "\n"
        return result
