#!/usr/bin/env python
import yaml
import time
import os
from dacite import from_dict
from typing import Callable, List, Optional

from .. import constants
from ..models.advisory import Advisory
from ..models.dependency import Dependency
from ..models.file_info import OmniversionFileInfo


def load_file(file_path: str):
    try:
        with open(file_path) as file:
            return yaml.safe_load(file), os.stat(file_path).st_ctime
    except yaml.YAMLError:
        return None
    except FileNotFoundError:
        return None


def load_data(base_path: str, add_file: Callable[[OmniversionFileInfo], None], hosts: Optional[List[str]] = None, pms: Optional[List[str]] = None, verbs: Optional[List[str]] = None):
    # we look for subdirectories containing data for a particular host
    host_dirs = [directory for directory in os.listdir(base_path) if os.path.isdir(os.path.join(base_path, directory))]
    for host in host_dirs:
        if hosts is not None and host not in hosts:
            continue
        host_path = os.path.join(base_path, host)
        pm_dirs = [directory for directory in os.listdir(host_path) if os.path.isdir(os.path.join(host_path, directory))]
        for pm in pm_dirs:
            if pms is not None and pm not in pms:
                continue
            available_verbs = ['audit', 'list', 'refresh', 'outdated', 'version']
            for verb in available_verbs:
                if verbs is not None and verb not in verbs:
                    continue
                file_name = verb + ".omniversion.yaml"
                file_path = os.path.join(host_path, pm, file_name)
                if os.path.exists(file_path):
                    file_data, time = load_file(file_path)
                    if file_data is None:
                        add_file(OmniversionFileInfo(None, file_name, host, pm, verb, time, file_path))
                    else:
                        for item in file_data:
                            item["pm"] = pm
                            item["host"] = host
                        add_file(OmniversionFileInfo(Dependency.from_list(file_data), file_name, host, pm, verb, time, file_path))
