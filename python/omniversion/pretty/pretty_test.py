#!/usr/bin/env python
from colorama import Fore, Back, Style

from . import pretty


def test_white_on_black():
    assert f"{Fore.BLACK}{Back.WHITE} test {Style.RESET_ALL}" == pretty.white_on_black("test")

