#!/usr/bin/env python
from colorama import Fore, Back, Style, init
from typing import Optional


def white_on_black(text: str):
    return f'{Fore.BLACK}{Back.WHITE} {text} {Style.RESET_ALL}'


def bright_on_lightblack(text: str):
    return f'{Fore.BRIGHT}{Back.LIGHTBLACK_EX} {text} {Style.RESET_ALL}'


def white(text: str):
    return f'{Fore.WHITE} {text} {Style.RESET_ALL}'


def cyan(text: str):
    return f'{Fore.CYAN} {text} {Style.RESET_ALL}'


def severity(severity: str):
    if severity == "critical":
        return f'{Fore.BRIGHT}{Back.RED} {text} {Style.RESET_ALL}'
    if severity == "high":
        return f'{Fore.BRIGHT}{Back.RED} {text} {Style.RESET_ALL}'
    if severity == "moderate":
        return f'{Fore.BRIGHT}{Back.YELLOW} {text} {Style.RESET_ALL}'
    if severity == "low":
        return f'{Fore.BRIGHT}{Back.BLUE} {text} {Style.RESET_ALL}'
    return severity


def traffic_light(text, status = "green"):
    if status == "green":
        return f'{Fore.GREEN}[✔]{Style.DIM} {text} {Style.RESET_ALL}'
    if status == "amber":
        return f'{Fore.YELLOW}[-]{Style.DIM} {text} {Style.RESET_ALL}'
    if status == "red":
        return f'{Fore.RED}[✘]{Style.DIM} {text} {Style.RESET_ALL}'
    return "-"


def header(text):
    return f'\n{Back.LIGHTBLACK_EX} {text} {Style.RESET_ALL}'


def hostname(text):
    return f'{Back.BLUE} {text} {Style.RESET_ALL}'


def verb(text):
    return f'{Back.GREEN} {text} {Style.RESET_ALL}'


def pm(text):
    return f'{Back.CYAN} {text} {Style.RESET_ALL}'


def dependency_count(num_dependencies: int, pm: Optional[str] = None):
    return f'{num_dependencies} {Style.DIM}{(pm + " ") if pm else ""}dependenc{"y" if num_dependencies == 1 else "ies" }{Style.RESET_ALL}\n'


def updates_count(num_dependencies: int, pm: str):
    return f'{num_dependencies if num_dependencies > 0 else "No"} {Style.DIM}{pm} update{"" if num_dependencies == 1 else "s" } available{Style.RESET_ALL}\n'


def file_count(num_items: int):
    return f'{num_items if num_items > 0 else "No"} file{"" if num_items == 1 else "s"} loaded'


def green(text: str):
    return f'{Fore.GREEN} {text} {Style.RESET_ALL}'
