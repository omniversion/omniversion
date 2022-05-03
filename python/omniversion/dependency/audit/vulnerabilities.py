#!/usr/bin/env python
from dataclasses import dataclass

from ..list import Dependencies
from ...pretty import pretty


@dataclass
class Vulnerabilities(Dependencies):
    def __str__(self):
        num_vulnerabilities = len(self.data)
        if num_vulnerabilities > 1:
            audit_items = [
                f'\t{(data_item.name or "").ljust(20)}'
                + f'\t{(data_item.host or "").ljust(12)}'
                + f'\t{(data_item.pm or "").ljust(12)}'
                + f'\t{data_item.advisories[0].severity or ""}'
                for data_item in self.data
            ]
            return pretty.traffic_light(
                f"{num_vulnerabilities} vulnerabilities found\n"
                + "\n".join(audit_items),
                "red",
            )
        if num_vulnerabilities == 1:
            return pretty.traffic_light("One vulnerability found", "red")
        return pretty.traffic_light("No vulnerabilities found")
