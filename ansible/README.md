<!--suppress HtmlDeprecatedAttribute -->
<div align="center">
    <img src="../docs/assets/Ansible.png" width="128" height="128" alt="omniversion logo" />
    <h1 align="center">omniversion Ansible collection</h1>
    <br />
</div>

The _omniversion Ansible collection_ contains roles and playbooks to fetch versions and dependency information from
servers orchestrated via Ansible.

## Install

```shell
ansible-galaxy collection install layer9gmbh.omniversion
```

## Use

Simply run the playbook [`layer9gmbh.omniversion.fetch`](ansible/playbooks/fetch/README.md) in your project's directory (or wherever Ansible can pick up your host definitions):

```shell
ansible-playbook layer9gmbh.omniversion.fetch
```

Package managers installed on each host will be detected automatically and all **global package** information will be fetched.

To fetch **local dependencies** as well, you will need to add some configuration so that `omniversion` knows which directories to look in.

### Configuration 

#### Package managers

Configure your package managers via the `var_omniversion` variable. It is a dictionary/map with the following values:

##### Package manager configuration options

| Key      | Value                                                                                   | 
|----------|-----------------------------------------------------------------------------------------|
| `apt`    | [Advanced packaging tool options](./docs/APT.md)                                        |
| `auto`   | [auto-detection options](./docs/AUTO.md)                                                |
| `galaxy` | [Ansible Galaxy options](./docs/GALAXY.md)                                              |
| `custom` | [custom command fetcher options](./docs/CUSTOM.md) _(fetch unmanaged package versions)_ |
| `file`   | [file fetcher options](./docs/FILE.md) _(fetch versions from configuration files)_      |
| `go`     | [go modules options](./docs/GO.md)                                                      |
| `brew`   | [homebrew options](./docs/BREW.md)                                                      |
| `npm`    | [Node package manager options](./docs/NPM.md)                                           |
| `nvm`    | [Node version manager options](./docs/NVM.md)                                           |
| `pip`    | [package installer for Python options](./docs/PIP.md)                                   |
| `gem`    | [rubygems options](./docs/GEM.md)                                                       |
| `rvm`    | [Ruby version manager options](./docs/RVM.md)                                           |

##### Example

```yaml
var_omniversion:
  
  # turn off auto-detection so omniversion will only attempt to use the package managers explicitly configured
  auto: false
  
  # enable Aptitude
  apt: true
  
  # enable npm with local dependencies
  npm:
    global: true
    local:
      - '/srv/foobar/current/frontend'
      - '/srv/foobar2/current/frontend'
  
  # enable rvm
  rvm: true
```


#### Output directory

The results of all `omniversion/ansible` runs are stored in temp files on the localhost (i.e. the control node). Use the `var_omniversion_output_dir` variable to control the directory. The default value is `/tmp/omniversion`.

This should be the same for all hosts. `omniversion/ansible` will automatically create subdirectories and store files separated by host, package manager and role, so that the `--tag` and `--limit` Ansible options can be used without overwriting unrelated data.

##### Example

```yaml
var_omniversion_output_dir: ~/Documents/testortestington/foobar/tmp/omniversion
```

### Playbooks and roles

The `omniversion/ansible` collection contains both roles and playbooks. You can either include one of the roles in a playbook or simply run a playbook directly.

The following are available as both roles and playbooks of the same name:

| Role/playbook name                                                    | Description                                                                                                                                                |
|-----------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------|
| [`layer9gmbh.omniversion.audit`](ansible/roles/audit/README.md)       | Fetches security notices by running the `audit` (or equivalent) command on each host (only applicable for package managers that offer this functionality). |
| [`layer9gmbh.omniversion.fetch`](ansible/roles/fetch/README.md)       | Fetches all available data by executing all other `omniversion` roles in turn.                                                                             |
| [`layer9gmbh.omniversion.list`](ansible/roles/list/README.md)         | Fetches currently installed dependencies by running the `list` (or equivalent) command on each host.                                                       |
| [`layer9gmbh.omniversion.outdated`](ansible/roles/outdated/README.md) | Fetches available updates by running the `outdated` (or equivalent) command on each host.                                                                  |
| [`layer9gmbh.omniversion.refresh`](ansible/roles/refresh/README.md)   | Updates the package manager caches by running the `update` (or equivalent) command on each host.                                                           |
| [`layer9gmbh.omniversion.version`](ansible/roles/version/README.md)   | Fetches versions not controlled by any package manager (including versions of package managers themselves) on each host.                                   |

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
