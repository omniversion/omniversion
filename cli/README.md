[![Software License](https://img.shields.io/badge/license-AGPL--3.0-green.svg?style=for-the-badge)](https://github.com/omniversion/omniversion/LICENSE)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=for-the-badge)](http://godoc.org/github.com/omniversion/omniversion/cli)
[![Go Report Card](https://goreportcard.com/badge/github.com/omniversion/omniversion/cli?style=for-the-badge)](https://goreportcard.com/report/github.com/omniversion/omniversion/cli)

# omniversion/cli

<!--suppress HtmlDeprecatedAttribute -->
<div align="center">
    <img src="../docs/assets/omniversion_cli.png" width="128" height="86" alt="omniversion logo" />
    <h2 align="center">omniversion command line tool</h2>
    <br />
</div>

The omniversion command line tool translates the output of many different version managers into a single, coherent list.

### How to install

#### Via npm
```shell
npm install -g omniversion
```

#### Via homebrew
```shell
brew tap omniversion/tap && brew install omniversion
```

### How to use

#### `parse` subcommand

The `parse` subcommand understands many different types of package manager output, translating the available information into the universal [omniversion format](../docs/MODELS.md).
```shell
apt list --installed | omniversion parse apt
apt list --upgradeable | omniversion parse apt
apt list --upgradeable --all-versions | omniversion parse apt

npm audit | omniversion parse npm
npm audit --json | omniversion parse npm
npm list | omniversion parse npm
npm list --json | omniversion parse npm
npm list --parseable | omniversion parse npm
npm outdated | omniversion parse npm
npm outdated --json | omniversion parse npm
npm outdated --parseable | omniversion parse npm
cat package.json | omniversion parse npm
cat package-lock.json | omniversion parse npm
npm --versions | omniversion parse npm
npm --versions --json | omniversion parse npm

rvm list | omniversion parse rvm
rvm version | omniversion parse rvm

brew list | omniversion parse homebrew

gem list | omniversion parse rubygems
bundle-audit | omniversion parse rubygems
```

##### Stderr output

```shell
npm audit 2>&1 || true | omniversion parse npm

```

### How to uninstall

#### Via npm
```shell
npm uninstall -g omniversion
```

#### Via homebrew
```shell
brew uninstall omniversion
```

If you are unsure how you installed `omniversion/cli`, look for the `via` field in the output of the `version` subcommand:
```shell
omniversion version | grep via
```