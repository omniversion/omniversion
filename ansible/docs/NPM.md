# Node package manager config options

## Key: `npm`

[Node package manager](https://npmjs.org) can be configured by providing a dictionary/map with the following keys and values:

| Key      | Value type           | Optional | Default | Purpose                                                                                                                                      |
|----------|----------------------|----------|---------|----------------------------------------------------------------------------------------------------------------------------------------------|
| `global` | `bool`               | yes      | `true`  | (De-)activate fetching of global dependencies.                                                                                               |
| `local`  | `[string]` or `null` | yes      | `null`  | Working directories from which local dependencies should be collected. Omit or set to `null` to deactivate collection of local dependencies. |

If you don't want to provide any options, you can also set a falsy/truthy value to (de-)activate global dependency collection.

#### Example

```yaml
var_omniversion:
  # only use npm
  auto: false
  # collect both global npm dependencies and local dependencies from the specified directories,
  # which should each contain a `package.json` file
  npm:
    global: true
    local:
      - '/srv/foobar/current/frontend'
      - '/srv/foobar2/current/frontend'
```
