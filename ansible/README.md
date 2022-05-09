<!--suppress HtmlDeprecatedAttribute -->
<div align="center">
    <img src="../docs/assets/Ansible.png" width="128" height="128" alt="omniversion logo" />
    <h1 align="center">omniversion Ansible collection</h1>
    <br />
</div>

The omniversion Ansible collection contains roles and playbooks to fetch versions and dependency information from servers orchestrated via Ansible.

## Install

```shell
ansible-galaxy collection install layer9gmbh.omniversion
```

## Use

### Configure package managers (optional)

For each host, configure the package managers to be used by setting the `var_omniversion` variable to a dictionary. Each key should be a package manager identifier, each value a (possibly empty) array of configuration options.

If `var_omniversion` is not specified (or `auto` is set to `true`), omniversion will detect available package managers automatically, but only collect global packages.

```yaml
var_omniversion:
  auto: true
  apt: true
  custom:
    - command: "test-app --version | sed 's/Test app version: //'"
      name: test-app
  npm:
    - dir: '/srv/foobar/current/frontend'
    - global: true
  rubygems:
    - dir: '/srv/foobar/current/backend'
  rvm: true
```

| Package manager | Available options                                                                                                                                                                                |
|-----------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| **auto**        | _Special value to configure auto-discovery of package managers._<br/>**No options**. Use any truthy value to enable global package collection for all package managers not otherwise configured. |
| **apt**         | **No options**. Use any truthy value to activate.                                                                                                                                                |
| **custom**      | _Special value for packages not covered by a package manager._<br/>`command`: A shell command that outputs a version.<br/>`name`: The name to be used for the package.                           |
| **npm**         | `dir`: The working directory to be used for local package collection.<br/>`global`: Set to `true` in any array entry to collect global packages as well.                                         |
| **rubygems**    | `dir`: The working directory to be used for local package collection.<br/>`global`: Set to `true` in any array entry to collect global packages as well.                                         |
| **rvm**         | **No options**. Use any truthy value to activate.                                                                                                                                                |


### Configure the output directory (optional)

`omniversion/ansible` stores the results of its runs on the control node. Use the `var_omniversion_output_dir` variable to control the directory. This should be the same for all hosts. `omniversion/ansible` will automatically create a subdirectory for each host.

The default value is `/tmp/omniversion`.

```yaml
var_omniversion_output_dir: ~/Documents/testortestington/foobar/tmp/omniversion
```

### Collect data

Either include one of the `omniversion/ansible` roles in a playbook or simply run an `omniversion` playbook directly in the directory in which your Ansible hosts are defined:

```shell
ansible-playbook layer9gmbh.omniversion.all
```

The following roles and playbooks are available:

| Role/playbook name                | Description                                                                                                                                                 |
|-----------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `layer9gmbh.omniversion.all`      | Collects all available data by executing all other `omniversion` roles in turn.                                                                             |
| `layer9gmbh.omniversion.audit`    | Collects security notices by running the `audit` (or equivalent) command on each host (only applicable for package managers that offer this functionality). |
| `layer9gmbh.omniversion.list`     | Collects currently installed dependencies by running the `list` (or equivalent) command on each host.                                                       |
| `layer9gmbh.omniversion.outdated` | Collects available updates by running the `outdated` (or equivalent) command on each host.                                                                  |
| `layer9gmbh.omniversion.refresh`  | Updates the package manager caches by running the `update` (or equivalent) command on each host.                                                            |
| `layer9gmbh.omniversion.version`  | Collects versions not controlled by any package manager (including versions of package managers themselves) on each host.                                   |


## Update

```shell
ansible-galaxy collection install layer9gmbh.omniversion --force
```

## Uninstall

`ansible-galaxy` offers no uninstall command, but you can simply delete the `layer9gmbh.omniversion` folder in your Ansible collections directory.
