[![Release](https://img.shields.io/github/v/release/omniversion/omniversion-cli.svg?style=for-the-badge)](https://github.com/omniversion/omniversion-cli/releases/latest)
[![GitHub Workflow Status](https://img.shields.io/github/workflow/status/omniversion/omniversion-cli/Upload%20code%20coverage?style=for-the-badge)](https://github.com/omniversion/omniversion-cli/actions?query=workflow%3A%22Upload%20code%20coverage%22)
[![Codecov branch](https://img.shields.io/codecov/c/github/omniversion/omniversion/main.svg?style=for-the-badge&token=X126WJ5IU4)](https://codecov.io/gh/omniversion/omniversion)
[![Software License](https://img.shields.io/badge/license-AGPL--3.0-green.svg?style=for-the-badge)](/LICENSE)


<!--suppress HtmlDeprecatedAttribute -->
<div align="center">
    <img src="docs/assets/omniversion.png" width="128" height="86" alt="omniversion logo" />
    <h1 align="center">omniversion</h1>
    <br />
</div>

## Dependency management toolbox

`omniversion` simplifies dependency management in complex projects by streamlining common tasks:

### Collection
[omniversion/ansible](ansible) collects versions and dependencies from servers orchestrated via Ansible.

### Aggregation
[omniversion/cli](cli) translates the output of many different version managers into a single, coherent list.

### Analysis
[omniversion/python](python) adds convenience methods to create dashboards and reports in a few lines of code.

## Why?

Maintaining a code base with a 

* Every self-respecting programming language has its own package manager, for some reason.
* Package managers differ greatly in their syntax, features, terminology and underlying model.
* Many versions are not actually controlled by package managers. More often than not, this includes the package manager itself.
* Versions checked into source control 
* Running commands on remote servers is time-consuming and error-prone.
* 


## Quick start

Prerequisites: You should have Ansible, npm and Python installed.

1. Install the Ansible collection, the CLI and the Python module on your machine:
    ```shell
    ansible-galaxy collection install layer9gmbh.omniversion
    npm install -g omniversion
    pip install omniversion
    ```


2. Run the `all` Ansible playbook
    ```shell
    ansible-playbook layer9gmbh.omniversion.all
    ```
    in a directory where Ansible can find your host definitions.


3. Run the Python sample dashboard:
    ```shell
    python3 -m omniversion.samples.dashboard
    ```

## Support

Mac & linux

| Package manager  | Supported    |
|------------------|--------------|
| `ansible-galaxy` | 🕙 planned   |
| `apt`            | ✅ yes        |
| `brew`           | ✅ yes        |
| `Composer`       | 🕙 planned   |
| `maven`          | 🕙 planned   |
| `npm`            | ✅ yes        |
| `pip`            | 🕙 planned   |
| `rubygems`       | ✅ yes        |
| `rvm`            | ✅ yes        |
| `yarn`           | 🕙 planned   |
