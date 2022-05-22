# Package installer for Python config options

## Key: `pip`

[Package installer for Python](https://pypi.org/project/pip/) can be configured by providing a dictionary/map with the following keys and values:

| Config option | Type                 | Optional | Default | Purpose                                                                                                                                                |
|---------------|----------------------|----------|---------|--------------------------------------------------------------------------------------------------------------------------------------------------------|
| `global`      | `bool`               | yes      | `true`  | Toggle fetching of global dependencies.                                                                                                                |
| `venv`        | `[string]` or `null` | yes      | `null`  | Paths to virtual environments from which local dependencies should be collected. Omit or set to `null` to deactivate collection of local dependencies. |

If you don't want to provide any options, you can also set a falsy/truthy value to (de-)activate global dependency collection.

#### Example

```yaml
var_omniversion:
  # collect both global pip dependencies and local dependencies from two virtual environments created using venv
  npm:
    global: true
    venv:
      - '/srv/foobar/current/env'
      - '/srv/foobar2/current/env'
```
