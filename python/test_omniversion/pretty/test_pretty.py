#!/usr/bin/env python
"""Test the pretty module"""
import unittest
from colorama import Fore, Back, Style

from omniversion.pretty import pretty


class PrettyPrintTestCase(unittest.TestCase):
    def test_black_on_white(self):
        formatted_text = pretty.black_on_white("test")
        self.assertIn(Fore.BLACK, formatted_text)
        self.assertIn(Back.WHITE, formatted_text)
        self.assertIn("test", formatted_text)
        self.assertIn(Style.RESET_ALL, formatted_text)

    def test_bright_on_lightblack(self):
        formatted_text = pretty.bright_on_lightblack("test")
        self.assertIn(Style.BRIGHT, formatted_text)
        self.assertIn(Back.LIGHTBLACK_EX, formatted_text)
        self.assertIn("test", formatted_text)
        self.assertIn(Style.RESET_ALL, formatted_text)

    def test_white(self):
        formatted_text = pretty.white("test")
        self.assertIn(Fore.WHITE, formatted_text)
        self.assertIn("test", formatted_text)
        self.assertIn(Style.RESET_ALL, formatted_text)

    def test_cyan(self):
        formatted_text = pretty.cyan("test")
        self.assertIn(Fore.CYAN, formatted_text)
        self.assertIn("test", formatted_text)
        self.assertIn(Style.RESET_ALL, formatted_text)

    def test_severity(self):
        formatted_text = pretty.severity("critical")
        self.assertIn(Style.BRIGHT, formatted_text)
        self.assertIn("critical", formatted_text)
        self.assertIn(Style.RESET_ALL, formatted_text)

        self.assertIn(Back.RED, pretty.severity("critical"))
        self.assertIn(Back.RED, pretty.severity("high"))
        self.assertIn(Back.YELLOW, pretty.severity("medium"))
        self.assertIn(Back.YELLOW, pretty.severity("moderate"))
        self.assertIn(Back.BLUE, pretty.severity("low"))

        self.assertNotIn(Style.BRIGHT, pretty.severity("unknown"))

    def test_traffic_light(self):
        self.assertIn(Fore.RED, pretty.traffic_light("test", "red"))
        self.assertIn(Fore.YELLOW, pretty.traffic_light("test", "amber"))
        self.assertIn(Fore.GREEN, pretty.traffic_light("test"))
        self.assertEqual("-", pretty.traffic_light("test", "unknown"))

    def test_header(self):
        self.assertIn(Back.LIGHTBLACK_EX, pretty.header("test"))

    def test_hostname(self):
        self.assertIn(Back.BLUE, pretty.hostname("test"))

    def test_verb(self):
        self.assertIn(Back.GREEN, pretty.verb("test"))

    def test_package_manager(self):
        self.assertIn(Back.CYAN, pretty.package_manager("test"))

    def test_dependency_count(self):
        self.assertIn(Style.DIM, pretty.dependency_count(3, "test"))
        self.assertIn("test dependencies", pretty.dependency_count(3, "test"))

    def test_updates_count(self):
        self.assertIn(Style.DIM, pretty.updates_count(3, "test"))
        self.assertIn("test updates available", pretty.updates_count(3, "test"))

    def test_file_count(self):
        self.assertIn("3 files loaded", pretty.file_count(3))

    def test_green(self):
        self.assertIn(Fore.GREEN, pretty.green("test"))


if __name__ == '__main__':
    unittest.main()
