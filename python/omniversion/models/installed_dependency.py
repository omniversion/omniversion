#!/usr/bin/env python
from pyfields import field, autoclass


@autoclass
class InstalledDependency:
    _flat_properties = ["location", "version"]

    def __init__(self):
        for prop in self._flat_properties:
            if data.has_key(prop):
                setattr(self, prop, data[prop])
