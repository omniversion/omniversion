[![Release](https://img.shields.io/github/v/release/omniversion/omniversion-cli.svg?style=for-the-badge)](https://github.com/omniversion/omniversion-cli/releases/latest)
[![GitHub Workflow Status](https://img.shields.io/github/workflow/status/omniversion/omniversion-cli/Upload%20code%20coverage?style=for-the-badge)](https://github.com/omniversion/omniversion-cli/actions?query=workflow%3A%22Upload%20code%20coverage%22)
[![Codecov branch](https://img.shields.io/codecov/c/github/omniversion/omniversion/main.svg?style=for-the-badge&token=X126WJ5IU4)](https://codecov.io/gh/omniversion/omniversion)
[![Software License](https://img.shields.io/badge/license-AGPL--3.0-green.svg?style=for-the-badge)](/LICENSE)


> ### 🚧 **Status: in active development**
> [Feedback and suggestions](https://github.com/omniversion/omniversion/discussions/1) much appreciated


<!--suppress HtmlDeprecatedAttribute -->
<div align="center">
    <img src="docs/assets/omniversion.png" width="128" height="86" alt="omniversion logo" />
    <h1 align="center">omniversion</h1>
    <br />
</div>

`omniversion` is a **dependency management toolbox**.

Some or all of its tools might be useful if you need to:

* keep servers up-to-date
* deal with multiple package managers
* identify conflicts across package managers and/or servers
* show all information with a single command in one spot
* include unmanaged or unpinned packages (e.g. manual installs, `apt install nginx`, `nvm install node`)
* keep `qa` and `prod` _exactly_ in sync
* patch vulnerabilities as soon as they are reported
* develop your own toolchain for server maintenance

## The tools

### Collection

[omniversion/ansible](ansible) collects versions and dependencies from servers orchestrated via Ansible.

### Aggregation

[omniversion/cli](cli) translates the output of many different version managers into a single, unified list.

### Analysis

[omniversion/python](python) adds convenience methods to create dashboards and reports in a few lines of code.

## Quick start

### Prerequisites

* `Ansible` >= 2.8
* `npm` (or `homebrew`)
* `Python` >= 3.8

### Steps

1.  **Install** the Ansible collection, the CLI and the Python module:
    ```shell
    ansible-galaxy collection install layer9gmbh.omniversion
    npm install -g omniversion
    pip install omniversion
    ```


2.  **Fetch** some data by running the Ansible playbook
    ```shell
    ansible-playbook layer9gmbh.omniversion.fetch
    ```
    in a directory where Ansible can find your host definitions.


3.  **Display** the data on the sample website dashboard:
    ```shell
    python3 -m omniversion.dashboard.website
    ```
    or in the terminal:
    ```shell
    python3 -m omniversion.dashboard.terminal
    ```

### What if I don't use Ansible?

At the moment, Ansible is the only option to fetch version data automatically. We might add a custom task runner to the toolbox at some point in the (probably distant) future.

In the meantime, you can fetch package manager output using shell scripts or any other method, feeding it to the `omniversion/cli` tool to get a single, comprehensive dependency list in a consistent format.

Feel free to suggest additional features and integrations in the [feedback section](https://github.com/omniversion/omniversion/discussions/1).

## Documentation

* [Ansible documentation](ansible/README.md)
* [CLI documentation](https://pkg.go.dev/github.com/omniversion/omniversion/cli)
* [Python documentation](https://omniversion.github.io/omniversion/python/omniversion/)

## Why?

[Why we created omniversion](docs/WHY.md)


## Get in touch

- [<img alt="GitHub Discussions" src="https://icongr.am/octicons/heart-fill.svg?color=808080&amp;size=10"/> GitHub Discussions](https://github.com/omniversion/omniversion/discussions/1):
  provide feedback and participate in discussions
- [<img alt="GitHub Issues" src="https://icongr.am/octicons/mark-github.svg?color=808080&amp;size=10"/> GitHub Issues](https://github.com/omniversion/omniversion/issues):
  report a bug or request a feature

## Supported platforms

### Operating systems

`omniversion` is being built and tested mostly on macOS control nodes with linux hosts, but we do aim to support linux control
nodes as well.

Like Ansible, we not support Windows control nodes, but this might change in the future, if there is demand.

### Package managers

| Name             | Supported  |
|------------------|------------|
| `ansible-galaxy` | ✅ yes      |
| `apt`            | ✅ yes      |
| `brew`           | ✅ yes      |
| `Composer`       | 🕙 planned |
| `go mod`         | ✅ yes      |
| `maven`          | 🕙 planned |
| `npm`            | ✅ yes      |
| `nvm`            | ✅ yes      |
| `pip`            | ✅ yes      |
| `rubygems`       | ✅ yes      |
| `rvm`            | ✅ yes      |
| `yarn`           | 🕙 planned |
