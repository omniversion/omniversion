package ansible

import (
	"github.com/omniversion/omniversion/cli/cmd/parse/shared"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseListOutput(t *testing.T) {
	vector := "ansible-galaxy [core 2.12.5]\n  config file = None\n  configured module search path = ['/Users/testor/.ansible/plugins/modules', '/usr/share/ansible/plugins/modules']\n  ansible python module location = /usr/local/lib/python3.8/site-packages/ansible\n  ansible collection location = /Users/testor/.ansible/collections:/usr/share/ansible/collections\n  executable location = /usr/local/bin/ansible-galaxy\n  python version = 3.8.13 (default, Mar 16 2022, 20:38:07) [Clang 13.0.0 (clang-1300.0.29.30)]\n  jinja version = 2.11.2\n  libyaml = True\nNo config file found; using defaults\n# /Users/testor/.ansible/roles\nOpened /Users/testor/.ansible/galaxy_token\n- atosatto.minio, v1.1.0\n- rvm.ruby, v2.1.2\n- ansistrano.deploy, 3.8.0\n- geerlingguy.mysql, 3.3.0\n- ansistrano.rollback, 3.1.0\n- geerlingguy.certbot, 3.0.3\n- elastic.elasticsearch, v7.12.0\n[WARNING]: - the configured path /usr/share/ansible/roles does not exist.\n[WARNING]: - the configured path /etc/ansible/roles does not exist.\n"

	previousInjectValue := shared.InjectPackageManager
	defer func() { shared.InjectPackageManager = previousInjectValue }()
	shared.InjectPackageManager = true
	result, err := ParseListOutput(vector)
	assert.Nil(t, err)
	assert.Equal(t, 7, len(result))

	item := result[0]
	assert.Equal(t, "atosatto.minio", item.Name)
	assert.Equal(t, []string{"minio"}, item.Aliases)
	assert.Equal(t, "ansible", item.PackageManager)
	assert.Equal(t, "1.1.0", item.Current)
	assert.Equal(t, 1, len(item.Installations))
	assert.Equal(t, "1.1.0", item.Installations[0].Version)

	item = result[1]
	assert.Equal(t, "rvm.ruby", item.Name)
	assert.Equal(t, []string{"ruby"}, item.Aliases)
	assert.Equal(t, "ansible", item.PackageManager)
	assert.Equal(t, "2.1.2", item.Current)
	assert.Equal(t, 1, len(item.Installations))
	assert.Equal(t, "2.1.2", item.Installations[0].Version)
}
