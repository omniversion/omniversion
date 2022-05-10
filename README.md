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

> ### 🚧 **Status: proof of concept**
> Feedback much appreciated


`omniversion` is a **dependency management toolbox** streamlining common maintenance tasks.

### Collection
[omniversion/ansible](ansible) collects versions and dependencies from servers orchestrated via Ansible.

### Aggregation
[omniversion/cli](cli) translates the output of many different version managers into a single, unified list.

### Analysis
[omniversion/python](python) adds convenience methods to create dashboards and reports in a few lines of code.

## Use cases
The `omniversion` tools are pretty flexible. [Here is how and why we use them](docs/WHY.md).

## Quick start

### Prerequisites
You should have `Ansible` installed and some hosts set up. You will also need `Python` and `npm` (or `homebrew`).

### Steps

1. Install the Ansible collection, the CLI and the Python module:
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

### Operating systems

This _proof of concept_ has been built and tested on a MacOS control node and linux hosts in mind, but we do aim to support linux control nodes as well.

Since Ansible does not support Windows control nodes, the same is currently true for `omniversion`. We might create a dedicated task runner to overcome this limitation at some point in the future, if there is demand.

### Package managers

| Name             | Supported    |
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
