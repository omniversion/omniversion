# Shared tasks


## Parse

Parse package manager output using the `omniversion` CLI.

### Variables

#### Input
`var_omniversion_parse_input`: The stdout of the package manager command.
`var_omniversion_pm`: The package manager's identifier. This is required by `omniversion` to apply the appropriate parser.

#### Output
`var_omniversion_parse_result`: The result of the `omniversion` invocation. `var_omniversion_parse_result.stdout` contains the actual output.


## Save

Save the parsed result to a file in the `omniversion` results directory.

### Variables

#### Input
`var_omniversion_output_dir` (optional): The base path where the `omniversion` results will be stored, grouped by hostname and package manager.
`var_omniversion_verb`: The verb (such as `audit`, `list` etc.) that will be used to determine the file name.
`var_omniversion_save_input`: The actual data to be saved to the file.

#### Output
_None_