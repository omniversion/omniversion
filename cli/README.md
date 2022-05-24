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

The `parse` subcommand understands many different types of package manager output, translating it into the universal [omniversion format](../docs/MODELS.md).
```shell
ansible-galaxy list | omniversion parse galaxy
ansible-galaxy list -vvv | omniversion parse galaxy
ansible-galaxy --version | omniversion parse galaxy
cat requirements.yaml | omniversion parse galaxy

apt list --installed | omniversion parse apt
apt list --upgradable | omniversion parse apt
apt list --upgradable --all-versions | omniversion parse apt
cat apt_preferences | omniversion parse apt

brew list | omniversion parse brew

gem list | omniversion parse gem
bundle-audit | omniversion parse gem

go list | omniversion parse go
go list -m -json all | omniversion parse go
go go version | omniversion parse go
cat go.mod | omniversion parse go
cat go.sum | omniversion parse go

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

nvm list | omniversion parse nvm
nvm --version | omniversion parse nvm
cat .nvmrc | omniversion parse nvm

pip list | omniversion parse pip
pip list --format=json | omniversion parse pip
pip list --format=freeze | omniversion parse pip
pip freeze | omniversion parse pip
pip list --outdated | omniversion parse pip
pip list --outdated --format=json | omniversion parse pip
pip --version | omniversion parse pip

echo "v1.2.3" | omniversion parse raw --name="test"
echo 'test1=v1.2.3\ntest2=v2.3.4\ntest3=v3.4.5' | omniversion parse raw --regex="(?m)^(?P<name>\S*)=(?P<version>\S*)$"

rvm list | omniversion parse rvm
rvm --version | omniversion parse rvm
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