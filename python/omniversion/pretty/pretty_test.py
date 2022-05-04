#!/usr/bin/env python
"""Test the pretty module"""
from colorama import Fore, Back, Style

from . import pretty


def test_black_on_white():
    """`black_on_white` shows """
    formatted_text = pretty.black_on_white("test")
    assert Fore.BLACK in formatted_text
    assert Back.WHITE in formatted_text
    assert "test" in formatted_text
    assert Style.RESET_ALL in formatted_text
