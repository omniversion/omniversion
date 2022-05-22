# Ansible Galaxy configuration options

There are no particular configuration options for [Ansible Galaxy](https://galaxy.ansible.com).

Specify a truthy value to activate, a falsy value to deactivate. If no value is
specified and `auto` is set to `true`, auto-detection will be used.

```yaml
var_omniversion:
  # enable Ansible Galaxy
  galaxy: true
```
