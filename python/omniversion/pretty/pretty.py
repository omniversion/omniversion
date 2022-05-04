#!/usr/bin/env python
"""Pretty-print some data in a human-readable string"""
from colorama import Fore, Back, Style


def black_on_white(text: str):
    """White text on black background"""
    return f"{Fore.BLACK}{Back.WHITE} {text} {Style.RESET_ALL}"


def bright_on_lightblack(text: str):
    """Bright text on dark gray background"""
    return f"{Style.BRIGHT}{Back.LIGHTBLACK_EX} {text} {Style.RESET_ALL}"


def white(text: str):
    """White text"""
    return f"{Fore.WHITE} {text} {Style.RESET_ALL}"


def cyan(text: str):
    """Cyan text"""
    return f"{Fore.CYAN} {text} {Style.RESET_ALL}"


def severity(text: str):
    """Format a vulnerability severity level"""
    if text == "critical":
        return f"{Style.BRIGHT}{Back.RED} {text} {Style.RESET_ALL}"
    if text == "high":
        return f"{Style.BRIGHT}{Back.RED} {text} {Style.RESET_ALL}"
    if text == "moderate":
        return f"{Style.BRIGHT}{Back.YELLOW} {text} {Style.RESET_ALL}"
    if text == "low":
        return f"{Style.BRIGHT}{Back.BLUE} {text} {Style.RESET_ALL}"
    return text


def traffic_light(text, status="green"):
    """Green, amber or red text with leading symbol"""
    if status == "green":
        return f"{Fore.GREEN}[✔]{Style.DIM} {text} {Style.RESET_ALL}"
    if status == "amber":
        return f"{Fore.YELLOW}[-]{Style.DIM} {text} {Style.RESET_ALL}"
    if status == "red":
        return f"{Fore.RED}[✘]{Style.DIM} {text} {Style.RESET_ALL}"
    return "-"


def header(text):
    """Text styled like a header"""
    return f"\n{Back.LIGHTBLACK_EX} {text} {Style.RESET_ALL}"


def hostname(text):
    """Consistent styling for hostnames"""
    return f"{Back.BLUE} {text} {Style.RESET_ALL}"


def verb(text):
    """Consistent styling for verbs (e.g. `list`, `audit`, etc.)"""
    return f"{Back.GREEN} {text} {Style.RESET_ALL}"


def package_manager(text):
    """Consistent styling for package manager names"""
    return f"{Back.CYAN} {text} {Style.RESET_ALL}"


def dependency_count(num_dependencies: int, package_manager_name: str | None = None):
    """Format the number of dependencies for a particular package manager"""
    pretty_pm = (package_manager_name + " ") if package_manager_name else ""
    pretty_deps = "dependenc" + ("y" if num_dependencies == 1 else "ies" )
    return f'{num_dependencies} {Style.DIM}{pretty_pm}{pretty_deps}{Style.RESET_ALL}\n'


def updates_count(num_dependencies: int, package_manager_name: str):
    """Format the number of available updates for a particular package manager"""
    pretty_num = num_dependencies if num_dependencies > 0 else "No"
    pretty_updates = "update" + ("" if num_dependencies == 1 else "s") + " available"
    return f'{pretty_num} {Style.DIM}{package_manager_name} {pretty_updates}{Style.RESET_ALL}\n'


def file_count(num_items: int):
    """Format the number of files loaded"""
    return f'{num_items if num_items > 0 else "No"} file{"" if num_items == 1 else "s"} loaded'


def green(text: str):
    """Green text"""
    return f"{Fore.GREEN} {text} {Style.RESET_ALL}"
