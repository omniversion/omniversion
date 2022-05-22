# Custom command fetcher options

## Key: `custom`

To collect versions not controlled by a package manager, provide an array of options with the following keys:

| Config option | Type       | Optional | Default | Function                                                                           |
|---------------|------------|----------|---------|------------------------------------------------------------------------------------|
| `command`     | `string`   | no       | *N/A*   | Shell command that outputs a string, which will be written to the `version` field. |
| `dir`         | `string`   | yes      | "/"     | Working directory in which the command should be executed, if any.                 |
| `name`        | `string`   | no       | *N/A*   | The name of the dependency. It will be written to the `name` field.                |

## Example

```yaml
var_omniversion:
  custom:
    # execute this command to fetch a version string and write it to the results file under the specified name
    - command: "test-app --version | sed 's/Test app version: //'"
      dir: test/app
      name: test-app
```
