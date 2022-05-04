#!/usr/bin/env python
"""Test the data module"""
import os

from .data import Data
from ..file_info import FileInfo
from ..package_info import PackageInfo
from ..package_info.list.package_infos_list import PackageInfosList


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
    base_path = os.path.join(os.path.dirname(__file__), "./../test/data")
    data = Data(base_path=base_path)
    data_as_str = data.__str__()
    assert "2 files loaded" in data_as_str
