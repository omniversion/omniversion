from dataclasses import dataclass
from typing import List, Optional


@dataclass
class PackageSource:
    identifier: Optional[str] = None
    url: Optional[str] = None
    versions: Optional[List[str]] = None
