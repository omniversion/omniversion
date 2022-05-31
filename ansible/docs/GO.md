# Go modules configuration options

[go modules](https://go.dev/ref/mod) can be configured by providing a dictionary/map with the following keys and values:

| Key      | Value type | Optional | Default | Purpose                                                   |
|----------|------------|----------|---------|-----------------------------------------------------------|
| `system` | `boolean`  | yes      | `true`  | Whether system library packages should be fetched.        |
| `local`  | `[string]` | yes      | `[]`    | List of directories from which modules should be fetched. |




Specify a truthy value to activate, a falsy value to deactivate. If no value is
specified and `auto` is set to `true`, auto-detection will be used.

```yaml
var_omniversion:
  # enable go modules
  go:
    # omit system libraries
    system: false
    local:
      # these directories should each contain a `go.mod` file
      - /some/dir
      - /some/other/dir
```
