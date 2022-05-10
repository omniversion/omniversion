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

## Why?

### The problem

Maintaining a modern software stack across deployments is hard work.

Given the manual effort needed just to keep dependencies up-to-date, avoid version conflicts and patch vulnerabilities, it is no surprise that these tasks are often neglected.

Version control systems, package managers and vulnerability scanners offer partial solutions, but they leave a lot to be desired:

* There is no obvious way of ensuring version consistency across package managers, e.g. when an `apt` package and an `npm` dependency need to be in sync.
* Multiple package managers need to be called in turn to answer simple questions like "Is there anything to patch?".
* Repeating this for each version currently deployed on a server is time-consuming and error-prone.
* Many software versions are not actually controlled by package managers, including - more often than not - the package managers themselves.
* Versions kept in configuration files are also frequently unmanaged, leading to hidden inconsistencies.
* Package managers differ greatly in their syntax, features, terminology and underlying model.

### The solution

For our daily maintenance work, we wanted a single dashboard containing all relevant information - across package managers, across servers and even across projects.

Given that we perform maintenance fairly frequently, we found that the added complexity of automation was worthwhile, saving time and preventing human error.

This is why we built the `omniversion` toolbox.

We make it available as Free Open Source Software in the hope that it might benefit other people as well.

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
