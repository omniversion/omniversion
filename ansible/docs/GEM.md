# Rubygems configuration options

## Key: `gem`

[rubygems](https://rubygems.org) can be configured by providing a dictionary/map with the following keys and values:

| Key     | Value type           | Optional | Default | Purpose                               |
|---------|----------------------|----------|---------|---------------------------------------|
| `dir`   | `string`             | yes      | `/`     | Working directory to fetch gems from. |

If you don't want to provide any options, you can also set a falsy/truthy value to (de-)activate global dependency collection.

```yaml
var_omniversion:
  # enable rubygems
  gem: true
```



