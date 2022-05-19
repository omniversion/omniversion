package apt

import (
	"github.com/omniversion/omniversion/cli/cmd/parse/shared"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseAptSimpleOutput(t *testing.T) {
	vector := `Listing... Done
autoconf/bionic,now 2.69-11 all [installed]
autotools-dev/bionic-updates,now 20180224.1 all [installed,automatic]
mde-netfilter/insiders-fast,bionic,now 100.69.32 amd64 [installed,upgradable to: 100.69.45]`

	result, err := parseAptOutput(vector)

	assert.Nil(t, err)
	assert.Equal(t, 3, len(result))

	item := result[0]
	assert.Equal(t, "autoconf", item.Name)
	assert.Equal(t, "2.69-11", item.Current)
	assert.Equal(t, "all", item.Architecture)
	assert.Equal(t, 2, len(item.Sources))
	assert.Equal(t, "bionic", item.Sources[0].Identifier)
	assert.Equal(t, "now", item.Sources[1].Identifier)

	item = result[1]
	assert.Equal(t, "autotools-dev", item.Name)
	assert.Equal(t, "20180224.1", item.Current)
	assert.Equal(t, "all", item.Architecture)
	assert.Equal(t, 2, len(item.Sources))
	assert.Equal(t, "bionic-updates", item.Sources[0].Identifier)
	assert.Equal(t, "now", item.Sources[1].Identifier)

	item = result[2]
	assert.Equal(t, "mde-netfilter", item.Name)
	assert.Equal(t, "100.69.32", item.Current)
	assert.Equal(t, "amd64", item.Architecture)
	assert.Equal(t, 3, len(item.Sources))
	assert.Equal(t, "insiders-fast", item.Sources[0].Identifier)
	assert.Equal(t, "bionic", item.Sources[1].Identifier)
	assert.Equal(t, "now", item.Sources[2].Identifier)
	assert.Equal(t, "100.69.45", item.Latest)
}

func TestParseAptNotInstalledOutput(t *testing.T) {
	vector := `Listing... Done
zssh/bionic 1.5c.debian.1-4 amd64`

	result, err := parseAptOutput(vector)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(result))

	item := result[0]
	assert.Empty(t, item.PackageManager)
	assert.Equal(t, "zssh", item.Name)
	assert.Equal(t, "", item.Current)
	assert.Equal(t, "1.5c.debian.1-4", item.Wanted)
	assert.Equal(t, "amd64", item.Architecture)
	assert.Equal(t, 1, len(item.Sources))
	assert.Equal(t, "bionic", item.Sources[0].Identifier)
}

func TestParseAptOutdatedOutput(t *testing.T) {
	vector := `Listing... Done
mde-netfilter/insiders-fast 100.69.45 amd64 [upgradable from: 100.69.32]
N: There are 4 additional versions. Please use the '-a' switch to see them.`

	previousInjectValue := shared.InjectPackageManager
	defer func() { shared.InjectPackageManager = previousInjectValue }()
	shared.InjectPackageManager = true
	result, err := parseAptOutput(vector)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(result))

	item := result[0]
	assert.Equal(t, "apt", item.PackageManager)
	assert.Equal(t, "mde-netfilter", item.Name)
	assert.Equal(t, "100.69.32", item.Current)
	assert.Equal(t, "100.69.45", item.Latest)
	assert.Equal(t, "100.69.32", item.Installations[0].Version)
	assert.Equal(t, "amd64", item.Architecture)
	assert.Equal(t, 1, len(item.Sources))
	assert.Equal(t, "insiders-fast", item.Sources[0].Identifier)
}

func TestParseAptOutdatedOutput_DebianInstance(t *testing.T) {
	vector := "Listing... Done\nadduser/now 3.118 all [installed,local]\napt/now 1.8.2.3 amd64 [installed,local]\nbase-files/now 10.3+deb10u12 amd64 [installed,local]\nbase-passwd/now 3.5.46 amd64 [installed,local]\nbash/now 5.0-4 amd64 [installed,local]\nbinutils-common/now 2.31.1-16 amd64 [installed,local]\nbinutils-x86-64-linux-gnu/now 2.31.1-16 amd64 [installed,local]\nbinutils/now 2.31.1-16 amd64 [installed,local]\nbsdutils/now 1:2.33.1-0.1 amd64 [installed,local]\nbuild-essential/now 12.6 amd64 [installed,local]\nbzip2/now 1.0.6-9.2~deb10u1 amd64 [installed,local]\nca-certificates/now 20200601~deb10u2 all [installed,local]\ncoreutils/now 8.30-3 amd64 [installed,local]\ncpp-8/now 8.3.0-6 amd64 [installed,local]\ncpp/now 4:8.3.0-1 amd64 [installed,local]\ndash/now 0.5.10.2-5 amd64 [installed,local]\ndebconf/now 1.5.71+deb10u1 all [installed,local]\ndebian-archive-keyring/now 2019.1+deb10u1 all [installed,local]\ndebianutils/now 4.8.6.1 amd64 [installed,local]\ndh-python/now 3.20190308 all [installed,local]\ndiffutils/now 1:3.7-3 amd64 [installed,local]\ndmsetup/now 2:1.02.155-3 amd64 [installed,local]\ndpkg-dev/now 1.19.7 all [installed,local]\ndpkg/now 1.19.7 amd64 [installed,local]\ne2fsprogs/now 1.44.5-1+deb10u3 amd64 [installed,local]\nfdisk/now 2.33.1-0.1 amd64 [installed,local]\nfindutils/now 4.6.0+git+20190209-2 amd64 [installed,local]\ng++-8/now 8.3.0-6 amd64 [installed,local]\ng++/now 4:8.3.0-1 amd64 [installed,local]\ngcc-8-base/now 8.3.0-6 amd64 [installed,local]\ngcc-8/now 8.3.0-6 amd64 [installed,local]\ngcc/now 4:8.3.0-1 amd64 [installed,local]\ngpgv/now 2.2.12-1+deb10u1 amd64 [installed,local]\ngrep/now 3.3-1 amd64 [installed,local]\ngzip/now 1.9-3+deb10u1 amd64 [installed,local]\nhostname/now 3.21 amd64 [installed,local]\ninit-system-helpers/now 1.56+nmu1 all [installed,local]\niproute2/now 4.20.0-2+deb10u1 amd64 [installed,local]\niputils-ping/now 3:20180629-2+deb10u2 amd64 [installed,local]\nlibacl1/now 2.2.53-4 amd64 [installed,local]\nlibapparmor1/now 2.13.2-10 amd64 [installed,local]\nlibapt-inst2.0/now 1.8.2.3 amd64 [installed,local]\nlibapt-pkg5.0/now 1.8.2.3 amd64 [installed,local]\nlibargon2-1/now 0~20171227-0.2 amd64 [installed,local]\nlibasan5/now 8.3.0-6 amd64 [installed,local]\nlibatomic1/now 8.3.0-6 amd64 [installed,local]\nlibattr1/now 1:2.4.48-4 amd64 [installed,local]\nlibaudit-common/now 1:2.8.4-3 all [installed,local]\nlibaudit1/now 1:2.8.4-3 amd64 [installed,local]\nlibbinutils/now 2.31.1-16 amd64 [installed,local]\nlibblkid1/now 2.33.1-0.1 amd64 [installed,local]\nlibbz2-1.0/now 1.0.6-9.2~deb10u1 amd64 [installed,local]\nlibc-bin/now 2.28-10+deb10u1 amd64 [installed,local]\nlibc-dev-bin/now 2.28-10+deb10u1 amd64 [installed,local]\nlibc6-dev/now 2.28-10+deb10u1 amd64 [installed,local]\nlibc6/now 2.28-10+deb10u1 amd64 [installed,local]\nlibcap-ng0/now 0.7.9-2 amd64 [installed,local]\nlibcap2-bin/now 1:2.25-2 amd64 [installed,local]\nlibcap2/now 1:2.25-2 amd64 [installed,local]\nlibcc1-0/now 8.3.0-6 amd64 [installed,local]\nlibcom-err2/now 1.44.5-1+deb10u3 amd64 [installed,local]\nlibcryptsetup12/now 2:2.1.0-5+deb10u2 amd64 [installed,local]\nlibdb5.3/now 5.3.28+dfsg1-0.5 amd64 [installed,local]\nlibdebconfclient0/now 0.249 amd64 [installed,local]\nlibdevmapper1.02.1/now 2:1.02.155-3 amd64 [installed,local]\nlibdpkg-perl/now 1.19.7 all [installed,local]\nlibelf1/now 0.176-1.1 amd64 [installed,local]\nlibexpat1-dev/now 2.2.6-2+deb10u4 amd64 [installed,local]\nlibexpat1/now 2.2.6-2+deb10u4 amd64 [installed,local]\nlibext2fs2/now 1.44.5-1+deb10u3 amd64 [installed,local]\nlibfdisk1/now 2.33.1-0.1 amd64 [installed,local]\nlibffi-dev/now 3.2.1-9 amd64 [installed,local]\nlibffi6/now 3.2.1-9 amd64 [installed,local]\nlibgcc-8-dev/now 8.3.0-6 amd64 [installed,local]\nlibgcc1/now 1:8.3.0-6 amd64 [installed,local]\nlibgcrypt20/now 1.8.4-5+deb10u1 amd64 [installed,local]\nlibgdbm-compat4/now 1.18.1-4 amd64 [installed,local]\nlibgdbm6/now 1.18.1-4 amd64 [installed,local]\nlibgmp10/now 2:6.1.2+dfsg-4+deb10u1 amd64 [installed,local]\nlibgnutls30/now 3.6.7-4+deb10u7 amd64 [installed,local]\nlibgomp1/now 8.3.0-6 amd64 [installed,local]\nlibgpg-error0/now 1.35-1 amd64 [installed,local]\nlibhogweed4/now 3.4.1-1+deb10u1 amd64 [installed,local]\nlibidn11/now 1.33-2.2 amd64 [installed,local]\nlibidn2-0/now 2.0.5-1+deb10u1 amd64 [installed,local]\nlibip4tc0/now 1.8.2-4 amd64 [installed,local]\nlibisl19/now 0.20-2 amd64 [installed,local]\nlibitm1/now 8.3.0-6 amd64 [installed,local]\nlibjson-c3/now 0.12.1+ds-2+deb10u1 amd64 [installed,local]\nlibkmod2/now 26-1 amd64 [installed,local]\nliblsan0/now 8.3.0-6 amd64 [installed,local]\nliblz4-1/now 1.8.3-1+deb10u1 amd64 [installed,local]\nliblzma5/now 5.2.4-1+deb10u1 amd64 [installed,local]\nlibmnl0/now 1.0.4-2 amd64 [installed,local]\nlibmount1/now 2.33.1-0.1 amd64 [installed,local]\nlibmpc3/now 1.1.0-1 amd64 [installed,local]\nlibmpdec2/now 2.4.2-2 amd64 [installed,local]\nlibmpfr6/now 4.0.2-1 amd64 [installed,local]\nlibmpx2/now 8.3.0-6 amd64 [installed,local]\nlibncurses6/now 6.1+20181013-2+deb10u2 amd64 [installed,local]\nlibncursesw6/now 6.1+20181013-2+deb10u2 amd64 [installed,local]\nlibnettle6/now 3.4.1-1+deb10u1 amd64 [installed,local]\nlibp11-kit0/now 0.23.15-2+deb10u1 amd64 [installed,local]\nlibpam-modules-bin/now 1.3.1-5 amd64 [installed,local]\nlibpam-modules/now 1.3.1-5 amd64 [installed,local]\nlibpam-runtime/now 1.3.1-5 all [installed,local]\nlibpam0g/now 1.3.1-5 amd64 [installed,local]\nlibpcre2-8-0/now 10.32-5 amd64 [installed,local]\nlibpcre3/now 2:8.39-12 amd64 [installed,local]\nlibperl5.28/now 5.28.1-6+deb10u1 amd64 [installed,local]\nlibprocps7/now 2:3.3.15-2 amd64 [installed,local]\nlibpsl5/now 0.20.2-2 amd64 [installed,local]\nlibpython3-dev/now 3.7.3-1 amd64 [installed,local]\nlibpython3-stdlib/now 3.7.3-1 amd64 [installed,local]\nlibpython3.7-dev/now 3.7.3-2+deb10u3 amd64 [installed,local]\nlibpython3.7-minimal/now 3.7.3-2+deb10u3 amd64 [installed,local]\nlibpython3.7-stdlib/now 3.7.3-2+deb10u3 amd64 [installed,local]\nlibpython3.7/now 3.7.3-2+deb10u3 amd64 [installed,local]\nlibquadmath0/now 8.3.0-6 amd64 [installed,local]\nlibreadline7/now 7.0-5 amd64 [installed,local]\nlibseccomp2/now 2.3.3-4 amd64 [installed,local]\nlibselinux1/now 2.8-1+b1 amd64 [installed,local]\nlibsemanage-common/now 2.8-2 all [installed,local]\nlibsemanage1/now 2.8-2 amd64 [installed,local]\nlibsepol1/now 2.8-1 amd64 [installed,local]\nlibsmartcols1/now 2.33.1-0.1 amd64 [installed,local]\nlibsqlite3-0/now 3.27.2-3+deb10u1 amd64 [installed,local]\nlibss2/now 1.44.5-1+deb10u3 amd64 [installed,local]\nlibssl-dev/now 1.1.1n-0+deb10u1 amd64 [installed,local]\nlibssl1.1/now 1.1.1n-0+deb10u1 amd64 [installed,local]\nlibstdc++-8-dev/now 8.3.0-6 amd64 [installed,local]\nlibstdc++6/now 8.3.0-6 amd64 [installed,local]\nlibsystemd0/now 241-7~deb10u8 amd64 [installed,local]\nlibtasn1-6/now 4.13-3 amd64 [installed,local]\nlibtinfo6/now 6.1+20181013-2+deb10u2 amd64 [installed,local]\nlibtsan0/now 8.3.0-6 amd64 [installed,local]\nlibubsan1/now 8.3.0-6 amd64 [installed,local]\nlibudev1/now 241-7~deb10u8 amd64 [installed,local]\nlibunistring2/now 0.9.10-1 amd64 [installed,local]\nlibuuid1/now 2.33.1-0.1 amd64 [installed,local]\nlibxtables12/now 1.8.2-4 amd64 [installed,local]\nlibzstd1/now 1.3.8+dfsg-3+deb10u2 amd64 [installed,local]\nlinux-libc-dev/now 4.19.235-1 amd64 [installed,local]\nlogin/now 1:4.5-1.1 amd64 [installed,local]\nlsb-base/now 10.2019051400 all [installed,local]\nmake/now 4.2.1-1.2 amd64 [installed,local]\nmawk/now 1.3.3-17+b3 amd64 [installed,local]\nmime-support/now 3.62 all [installed,local]\nmount/now 2.33.1-0.1 amd64 [installed,local]\nncurses-base/now 6.1+20181013-2+deb10u2 all [installed,local]\nncurses-bin/now 6.1+20181013-2+deb10u2 amd64 [installed,local]\nopenssl/now 1.1.1n-0+deb10u1 amd64 [installed,local]\npasswd/now 1:4.5-1.1 amd64 [installed,local]\npatch/now 2.7.6-3+deb10u1 amd64 [installed,local]\nperl-base/now 5.28.1-6+deb10u1 amd64 [installed,local]\nperl-modules-5.28/now 5.28.1-6+deb10u1 all [installed,local]\nperl/now 5.28.1-6+deb10u1 amd64 [installed,local]\nprocps/now 2:3.3.15-2 amd64 [installed,local]\npython-apt-common/now 1.8.4.3 all [installed,local]\npython-pip-whl/now 18.1-5 all [installed,local]\npython3-apt/now 1.8.4.3 amd64 [installed,local]\npython3-dev/now 3.7.3-1 amd64 [installed,local]\npython3-distutils/now 3.7.3-1 all [installed,local]\npython3-lib2to3/now 3.7.3-1 all [installed,local]\npython3-minimal/now 3.7.3-1 amd64 [installed,local]\npython3-pip/now 18.1-5 all [installed,local]\npython3-pkg-resources/now 40.8.0-1 all [installed,local]\npython3-setuptools/now 40.8.0-1 all [installed,local]\npython3-wheel/now 0.32.3-2 all [installed,local]\npython3.7-dev/now 3.7.3-2+deb10u3 amd64 [installed,local]\npython3.7-minimal/now 3.7.3-2+deb10u3 amd64 [installed,local]\npython3.7/now 3.7.3-2+deb10u3 amd64 [installed,local]\npython3/now 3.7.3-1 amd64 [installed,local]\nreadline-common/now 7.0-5 all [installed,local]\nsed/now 4.7-1 amd64 [installed,local]\nsudo/now 1.8.27-1+deb10u3 amd64 [installed,local]\nsystemd-sysv/now 241-7~deb10u8 amd64 [installed,local]\nsystemd/now 241-7~deb10u8 amd64 [installed,local]\nsysvinit-utils/now 2.93-8 amd64 [installed,local]\ntar/now 1.30+dfsg-6 amd64 [installed,local]\ntzdata/now 2021a-0+deb10u4 all [installed,local]\nutil-linux/now 2.33.1-0.1 amd64 [installed,local]\nwget/now 1.20.1-1.1 amd64 [installed,local]\nxz-utils/now 5.2.4-1+deb10u1 amd64 [installed,local]\nzlib1g/now 1:1.2.11.dfsg-1+deb10u1 amd64 [installed,local]\n"

	previousInjectValue := shared.InjectPackageManager
	defer func() { shared.InjectPackageManager = previousInjectValue }()
	shared.InjectPackageManager = true
	result, err := parseAptOutput(vector)

	assert.Nil(t, err)
	assert.Equal(t, 185, len(result))
}
