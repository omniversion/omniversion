package ansible

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseRequirementsYamlFile(t *testing.T) {
	vector := "# from galaxy\n- src: yatesr.timezone\n\n# from GitHub\n- src: https://github.com/bennojoy/nginx\n\n# from GitHub, overriding the name and specifying a specific tag\n- src: https://github.com/bennojoy/nginx\n  version: master\n  name: nginx_role\n\n# from a webserver, where the role is packaged in a tar.gz\n- src: https://some.webserver.example.com/files/master.tar.gz\n  name: http-role\n\n# from Bitbucket\n- src: git+http://bitbucket.org/willthames/git-ansible-galaxy\n  version: v1.4\n\n# from Bitbucket, alternative syntax and caveats\n- src: http://bitbucket.org/willthames/hg-ansible-galaxy\n  scm: hg\n\n# from GitLab or other git-based scm\n- src: git@gitlab.company.com:mygroup/ansible-base.git\n  scm: git\n  version: \"0.1\"  # quoted, so YAML doesn't parse this as a floating-point value\n"

	result, err := ParseRequirementsYamlFile(vector)
	assert.Nil(t, err)
	assert.Equal(t, 7, len(result))

	item := result[0]
	assert.Equal(t, "yatesr.timezone", item.Name)
	assert.Zero(t, len(item.Installations))

	item = result[2]
	assert.Equal(t, "https://github.com/bennojoy/nginx", item.Name)
	assert.Equal(t, "master", item.Wanted)

	item = result[4]
	assert.Equal(t, "git+http://bitbucket.org/willthames/git-ansible-galaxy", item.Name)
	assert.Equal(t, "1.4", item.Wanted)
}
