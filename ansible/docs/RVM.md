# Ruby version manager configuration options

## Key: `rvm`


[Ruby version manager](https://rvm.io) can be configured by providing a dictionary/map with the following keys and values:

| Key      | Value type           | Optional | Default                     | Purpose                        |
|----------|----------------------|----------|-----------------------------|--------------------------------|
| `gemset` | `[string]`           | yes      | output of `rvm gemset list` | List of gemsets to be fetched. |

There are no particular configuration options for [Ruby version manager](https://rvm.io).

Specify a truthy value to activate, a falsy value to deactivate. If no value is
specified and `auto` is set to `true`, auto-detection will be used.

```yaml
var_omniversion:
  # enable Ruby version manager
  rvm: true
  gemset:
    - test1
    - test2
```
