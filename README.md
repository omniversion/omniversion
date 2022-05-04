[![Release](https://img.shields.io/github/v/release/omniversion/omniversion-cli.svg?style=for-the-badge)](https://github.com/omniversion/omniversion-cli/releases/latest)
[![GitHub Workflow Status](https://img.shields.io/github/workflow/status/omniversion/omniversion-cli/Upload%20code%20coverage?style=for-the-badge)](https://github.com/omniversion/omniversion-cli/actions?query=workflow%3A%22Upload%20code%20coverage%22)
[![Codecov branch](https://img.shields.io/codecov/c/github/omniversion/omniversion/main.svg?style=for-the-badge&token=X126WJ5IU4)](https://codecov.io/gh/omniversion/omniversion)
[![Software License](https://img.shields.io/badge/license-AGPL--3.0-green.svg?style=for-the-badge)](/LICENSE)

# omniversion

<!--suppress HtmlDeprecatedAttribute -->
<div align="center">
    <img src="docs/assets/omniversion.png" width="256" height="172" alt="omniversion logo" />
    <h2 align="center">dependency management toolbox</h2>
    <br />
</div>

`omniversion` simplifies dependency management in complex projects by streamlining the following steps:

### Collection
[omniversion/ansible](ansible) collects version and dependency information from servers managed via Ansible.

### Aggregation
[omniversion/cli](cli) collates the output of different version managers into a single, coherent list.

### Analysis
[omniversion/python](python) adds convenience methods for analysing the dependency lists to create dashboards, alerts and much more in a few lines of code.
