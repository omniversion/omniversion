# Go modules configuration options

There are no particular configuration options for [go modules](https://go.dev/ref/mod).

Specify a truthy value to activate, a falsy value to deactivate. If no value is
specified and `auto` is set to `true`, auto-detection will be used.

```yaml
var_omniversion:
  # enable go modules
  go: true
```
