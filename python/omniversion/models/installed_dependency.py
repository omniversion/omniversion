#!/usr/bin/env python
from pyfields import field, autoclass

@autoclass
class InstalledDependency:
    _flat_properties = ["location", "version"]
    def __init__(self, properties):
        for property in self._flat_properties:
            if data.has_key(property):
                setattr(self, property, data[property])
