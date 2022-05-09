<!--suppress HtmlDeprecatedAttribute -->
<div align="center">
    <img src="../docs/assets/Ansible.png" width="128" height="128" alt="omniversion logo" />
    <h2 align="center">omniversion Ansible collection</h2>
    <br />
</div>

The omniversion Ansible collection contains roles and playbooks to help you fetch versions and dependency information from servers orchestrated via Ansible.

### How to install

```shell
ansible-galaxy collection install layer9gmbh.omniversion
```

### How to use

#### Configure the output directory



#### Configure package managers


#### Collect data

```shell
ansible-playbook layer9gmbh.omniversion.all
```

### Update

```shell
ansible-galaxy collection install layer9gmbh.omniversion --force
```