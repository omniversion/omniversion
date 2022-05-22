# Auto-detection options

## Key: `auto`

The `auto` option is a boolean value that toggles auto-detection of available package managers. The default value is `true`, which will cause available package managers to be determined automatically on each host. Note that this will only pick up global dependencies. For more fine-grained control, configure the package managers you need.

Setting `auto` to `false` might also speed up execution a little.

##### Example

```yaml
var_omniversion:
  # turn off auto-detection so `omniversion` will only attempt to use the package managers explicitly configured
  auto: false
  # ... additional configuration ...
```
