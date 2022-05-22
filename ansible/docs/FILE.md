# File fetcher options

## Key: `file`

To collect versions from arbitrary files on your hosts (including localhost), you can provide an **array** of dictionaries/maps, each with the following keys and values:

| Key      | Value type | Optional | Default | Purpose                                                                                                                                                                                                                                                 |
|----------|------------|----------|---------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `path`   | `string`   | no       | _N/A_   | Absolute path to the file.                                                                                                                                                                                                                              |
| `regex`  | `string`   | yes      | `null`  | Golang-flavored regular expression containing a named group called `version` and optionally a named group called `name` to extract the package data from the file. Multiple matches are supported. Set to `null` if the file contains only the version. |
| `name`   | `string`   | yes      | `null`  | Name of the package, if not contained in the file.                                                                                                                                                                                                      |
| `parser` | `string`   | yesy     | `null`  | d                                                                                                                                                                                                                                                       |

#### Example

```yaml
var_omniversion:
  file:
    # extract each line as `<name>=<version>`
    - path: /srv/foobar/frontend/.env
      regex: "^(?P<name>\S+)=(?P<version>\S+)$"
    # use the entire file contents as the version string
    - path: /srv/foobar/frontend/.nvmrc
      name: node
```
