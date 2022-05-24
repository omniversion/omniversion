# Advanced packaging tool configuration options

## Key: `apt`

There are no particular configuration options for [Advanced packaging tool](https://salsa.debian.org/apt-team/apt).

Specify a truthy value to activate, a falsy value to deactivate. If no value is
specified and `auto` is set to `true`, auto-detection will be used.

```yaml
var_omniversion:
  # enable Aptitude
  apt: true
```
