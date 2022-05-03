#!/usr/bin/env python
from typing import Dict, List, Optional, Union
from ..format import format
from dataclasses import dataclass
from .dependency import Dependency


@dataclass
class Vulnerabilities:
    data: List[Dependency]

    def __init__(self, data: List[Dependency]):
        self.data = data

    def __str__(self):
        num_vulnerabilities = len(self.data)
        if num_vulnerabilities > 1:
            audit_items = [f'\t{(data_item.name or "").ljust(20)}\t{(data_item.host or "").ljust(12)}\t{(data_item.pm or "").ljust(12)}\t{data_item.advisories[0].severity or ""}' for data_item in self.data]
            return format.traffic_light(f'{num_vulnerabilities} vulnerabilities found\n' + "\n".join(audit_items), "red")
        elif num_vulnerabilities == 1:
            return format.traffic_light("One vulnerability found", "red")
        else:
            return format.traffic_light("No vulnerabilities found")
