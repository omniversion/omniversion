#!/usr/bin/env python
"""Test the data module"""
import os

from omniversion.data.data import Data
from omniversion.file_info import FileInfo
from omniversion.package_info import PackageInfo
from omniversion.package_info.list.package_infos_list import PackageInfosList


def test_summary_for_no_data():
    data = Data(file_infos=[])
    data_as_str = data.__str__()
    assert "No files loaded" in data_as_str


def test_summary_for_single_file():
    package_infos = PackageInfosList([PackageInfo(name="package", version="0.1.2", pm="test")])
    data = Data(file_infos=[FileInfo(data=package_infos, name="test")])
    data_as_str = data.__str__()
    assert "1 file loaded" in data_as_str


def test_summary_for_multiple_files():
    package_infos = PackageInfosList([PackageInfo(name="package", version="0.1.2", pm="test")])
    data = Data(file_infos=[FileInfo(data=package_infos, name="test"), FileInfo(data=package_infos, name="test")])
    data_as_str = data.__str__()
    assert "2 files loaded" in data_as_str


def test_load_files():
    base_path = os.path.join(os.path.dirname(__file__), "./../vectors")
    data = Data(base_path=base_path)
    data_as_str = data.__str__()
    assert "4 files loaded" in data_as_str
