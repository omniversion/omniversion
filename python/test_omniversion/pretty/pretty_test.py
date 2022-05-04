#!/usr/bin/env python
"""Test the pretty module"""
import unittest
from colorama import Fore, Back, Style

from omniversion.pretty import pretty


class InitializationTestCase(unittest.TestCase):
    def test_black_on_white(self):
        """`black_on_white` shows """
        formatted_text = pretty.black_on_white("test")
        self.assertIn(Fore.BLACK, formatted_text)
        self.assertIn(Back.WHITE, formatted_text)
        self.assertIn("test", formatted_text)
        self.assertIn(Style.RESET_ALL, formatted_text)


if __name__ == '__main__':
    unittest.main()
