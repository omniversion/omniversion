<!--suppress HtmlDeprecatedAttribute -->
<div align="center">
    <img src="../docs/assets/Ansible.png" width="128" height="128" alt="omniversion logo" />
    <h1 align="center">omniversion Ansible collection</h1>
    <br />
</div>

The omniversion Ansible collection contains roles and playbooks to fetch versions and dependency information from
servers orchestrated via Ansible.

## Install

```shell
ansible-galaxy collection install layer9gmbh.omniversion
```

## Use

### Configure package managers (optional)

For each host, configure your package managers via the `var_omniversion` variable. It should be a dictionary with the
following values:

#### The `auto` option

The `auto` option is a boolean value that toggles auto-detection of available package managers. The default value
is `true`,
which will cause available package managers to be determined automatically on each host. Note that this will only pick
up global dependencies. For more fine-grained control, configure the package managers you need. Setting `auto`
to `false` might also speed up execution a little.

#### Example

```yaml
var_omniversion:
  # turn off auto-detection so `omniversion` will only attempt to use the package managers explicitly configured
  auto: false
  # ... more configuration here ...
```

#### The `npm` options

These package managers can be configured by providing a dictionaries with the following keys:

| Config option | Type                 | Optional | Default | Function                                                                                                                                     |
|---------------|----------------------|----------|---------|----------------------------------------------------------------------------------------------------------------------------------------------|
| `global`      | `bool`               | yes      | `true`  | Whether global dependencies should be collected as well.                                                                                     |
| `local`       | `[string]` or `null` | yes      | `null`  | Working directories from which local dependencies should be collected. Omit or set to `null` to deactivate collection of local dependencies. |

If you don't want to provide any options, you can also set a falsy/truthy value to (de-)activate global dependency collection.

#### Example

```yaml
var_omniversion:
  # only use `npm` and `rubygems`
  auto: false
  # collect both global `npm` dependencies and local dependencies from the specified directories,
  # which should each contain a `package.json` file
  npm:
    global: true
    local:
      - '/srv/foobar/current/frontend'
      - '/srv/foobar2/current/frontend'
```

#### The `apt`, `gem` and `rvm` options

These package managers have no particular configuration options. Specify a truthy value to activate, a falsy value to deactivate. If no value is
specified and `auto` is set to `true`, auto-detection will be used to determine installed package managers.

#### Examples

```yaml
var_omniversion:
  # only use `apt`
  auto: false
  apt: true
```

```yaml
var_omniversion:
  # use auto-detection, but exclude `rvm`
  auto: true
  rvm: false
```

#### The `custom` option

To collect versions not controlled by a package manager, provide an array of options with the following keys:

| Config option | Type       | Optional | Default | Function                                                                           |
|---------------|------------|----------|---------|------------------------------------------------------------------------------------|
| `command`     | `string`   | no       |         | Shell command that outputs a string, which will be written to the `version` field. |
| `dir`         | `string`   | yes      | "/"     | Working directory in which the command should be executed, if any.                 |
| `name`        | `string`   | no       |         | The name of the dependency. It will be written to the `name` field.                |

#### Example

```yaml
var_omniversion:
  custom:
    # execute this command to fetch a version string and write it to the results file under the specified name
    - command: "test-app --version | sed 's/Test app version: //'"
      dir: test/app
      name: test-app
```

### Configure the output directory (optional)

`omniversion/ansible` stores the results of its runs on the control node. Use the `var_omniversion_output_dir` variable
to control the directory. This should be the same for all hosts. `omniversion/ansible` will automatically create a
subdirectory for each host.

The default value is `/tmp/omniversion`.

```yaml
var_omniversion_output_dir: ~/Documents/testortestington/foobar/tmp/omniversion
```

### Collect data

The `omniversion/ansible` collection contains both roles and playbooks. You can either include one of the roles in a
playbook or simply run a playbook directly:

```shell
ansible-playbook layer9gmbh.omniversion.all
```

The following are available as both roles and playbooks of the same name:

| Role/playbook name                                                    | Description                                                                                                                                                 |
|-----------------------------------------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------|
| [`layer9gmbh.omniversion.all`](ansible/roles/all/README.md)           | Collects all available data by executing all other `omniversion` roles in turn.                                                                             |
| [`layer9gmbh.omniversion.audit`](ansible/roles/audit/README.md)       | Collects security notices by running the `audit` (or equivalent) command on each host (only applicable for package managers that offer this functionality). |
| [`layer9gmbh.omniversion.list`](ansible/roles/list/README.md)         | Collects currently installed dependencies by running the `list` (or equivalent) command on each host.                                                       |
| [`layer9gmbh.omniversion.outdated`](ansible/roles/outdated/README.md) | Collects available updates by running the `outdated` (or equivalent) command on each host.                                                                  |
| [`layer9gmbh.omniversion.refresh`](ansible/roles/refresh/README.md)   | Updates the package manager caches by running the `update` (or equivalent) command on each host.                                                            |
| [`layer9gmbh.omniversion.version`](ansible/roles/version/README.md)   | Collects versions not controlled by any package manager (including versions of package managers themselves) on each host.                                   |

In addition, two internally used roles do not have an associated playbook, as you usually don't need to call them individually:
* [`layer9gmbh.omniversion.check`](ansible/roles/check/README.md) verifies that the `omniversion` CLI is installed on the control node.
* [`layer9gmbh.omniversion.autodetect`](ansible/roles/autodetect/README.md) determines which package managers are available on each host.

## Update

```shell
ansible-galaxy collection install layer9gmbh.omniversion --force
```

## Uninstall

`ansible-galaxy` offers no uninstall command, but you can simply delete the `layer9gmbh.omniversion` folder in your
Ansible collections directory.
