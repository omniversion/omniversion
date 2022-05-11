package homebrew

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseHomebrewListOutput(t *testing.T) {
	vector := "==> Formulae\nallure                  fribidi                 icu4c                   libksba                 libvorbis               nfpm                    protobuf                tcl-tk\nansible                 gcc                     imath                   libmng                  libvpx                  nspr                    python@3.10             terraform\naom                     gd                      ioctl                   libmpc                  libx11                  nss                     python@3.8              tesseract\nassimp                  gdbm                    isl                     libnghttp2              libxau                  omniversion             python@3.9              theora\nautoconf                gdk-pixbuf              jasper                  libogg                  libxcb                  open-mpi                qemu                    unbound\nautomake                gettext                 jpeg                    libpng                  libxdmcp                opencore-amr            qt                      vault\nbdw-gc                  giflib                  jpeg-xl                 libpthread-stubs        libxext                 openexr                 rav1e                   vde\nbrotli                  git                     lame                    librist                 libxrender              openjdk                 readline                wakeonlan\nca-certificates         glib                    leptonica               librsvg                 libyaml                 openjpeg                redis                   webp\ncairo                   gmp                     libarchive              libsamplerate           litestream              openssl@1.1             rsync                   x264\ncjson                   gnu-sed                 libass                  libslirp                little-cms2             opus                    rtmpdump                x265\ncmake                   gnutls                  libassuan               libsndfile              lz4                     p11-kit                 rubberband              xorgproto\ncmocka                  go                      libavif                 libsodium               lzo                     pango                   sdl2                    xvid\ncoreutils               gobject-introspection   libb2                   libsoxr                 m4                      pcre                    six                     xxhash\ndav1d                   goreleaser              libbluray               libssh                  mbedtls                 pcre2                   snappy                  xz\ndbus                    graphite2               libcbor                 libtasn1                md4c                    pdf2svg                 socat                   yq\ndouble-conversion       graphviz                libdvdcss               libtiff                 mpdecimal               pinentry-mac            speex                   yubikey-agent\nffmpeg                  gts                     libevent                libtool                 mpfr                    pixman                  sqlite                  zeromq\nflac                    guile                   libffi                  libunistring            mysql                   pkg-config              srt                     zimg\nfontconfig              harfbuzz                libfido2                libusb                  ncurses                 podman                  sshpass                 zlib\nfreetype                hunspell                libgpg-error            libvidstab              netpbm                  poppler                 svu                     zstd\nfrei0r                  hwloc                   libidn2                 libvmaf                 nettle                  popt                    swig\n\n==> Casks\ndb-browser-for-sqlite   secretive               sekey                   xquartz\n"
	result, err := parseHomebrewOutput(vector)

	assert.Nil(t, err)
	assert.Equal(t, 157, len(result))

	item := result[0]
	assert.Equal(t, "allure", item.Name)
	assert.Equal(t, "", item.Current)
	assert.Equal(t, 1, len(item.Installations))
	assert.Equal(t, "", item.Installations[0].Version)

	item = result[1]
	assert.Equal(t, "fribidi", item.Name)
	assert.Equal(t, "", item.Current)
	assert.Equal(t, 1, len(item.Installations))
	assert.Equal(t, "", item.Installations[0].Version)
}

func TestParseHomebrewListVersionsOutput(t *testing.T) {
	vector := "allure 2.13.5\nansible 5.7.0\naom 3.3.0\nassimp 5.2.3\nautoconf 2.71\nautomake 1.16.5\nbdw-gc 8.0.6\nbrotli 1.0.9\nca-certificates 2022-04-26\ncairo 1.16.0_5\ncjson 1.7.15\ncmake 3.18.3\ncmocka 1.1.5\ncoreutils 9.0_1\ndav1d 0.9.2\ndbus 1.14.0\ndouble-conversion 3.2.0\nffmpeg 5.0.1\nflac 1.3.4 1.3.3\nfontconfig 2.14.0 2.13.1\nfreetype 2.12.0 2.11.1\nfrei0r 1.7.0 1.8.0\nfribidi 1.0.10 1.0.12\ngcc 10.2.0\ngd 2.3.3_2\ngdbm 1.23\ngdk-pixbuf 2.42.8\ngettext 0.21\ngiflib 5.2.1\ngit 2.30.0\nglib 2.72.1\ngmp 6.2.1_1\ngnu-sed 4.8\ngnutls 3.7.4\ngo 1.17.5\ngobject-introspection 1.72.0\ngoreleaser 0.155.0\ngraphite2 1.3.14\ngraphviz 3.0.0\ngts 0.7.6_2\nguile 3.0.8\nharfbuzz 4.2.1\nhunspell 1.7.0_2\nhwloc 2.3.0\nicu4c 69.1 70.1\nimath 3.1.5 3.1.3\nioctl 1.1.3\nisl 0.22.1\njasper 3.0.3\njpeg 9e\njpeg-xl 0.6.1\nlame 3.100\nleptonica 1.82.0\nlibarchive 3.6.1\nlibass 0.15.2\nlibassuan 2.5.5\nlibavif 0.10.1\nlibb2 0.98.1\nlibbluray 1.3.1 1.2.0\nlibcbor 0.9.0\nlibdvdcss 1.4.2\nlibevent 2.1.12\nlibffi 3.4.2\nlibfido2 1.10.0\nlibgpg-error 1.38 1.39 1.41 1.44\nlibidn2 2.3.2\nlibksba 1.6.0\nlibmng 2.0.3\nlibmpc 1.2.0\nlibnghttp2 1.47.0\nlibogg 1.3.4 1.3.5\nlibpng 1.6.37\nlibpthread-stubs 0.4\nlibrist 0.2.7\nlibrsvg 2.54.1\nlibsamplerate 0.1.9_1\nlibslirp 4.7.0\nlibsndfile 1.1.0 1.0.30\nlibsodium 1.0.18_1\nlibsoxr 0.1.3\nlibssh 0.9.6\nlibtasn1 4.18.0\nlibtiff 4.3.0\nlibtool 2.4.7\nlibunistring 1.0\nlibusb 1.0.26\nlibvidstab 1.1.0\nlibvmaf 2.3.1 2.3.0_1\nlibvorbis 1.3.7\nlibvpx 1.11.0 1.9.0\nlibx11 1.7.5\nlibxau 1.0.9\nlibxcb 1.14_2\nlibxdmcp 1.1.3\nlibxext 1.3.4\nlibxrender 0.9.10\nlibyaml 0.2.5\nlitestream 0.3.2\nlittle-cms2 2.13.1\nlz4 1.9.3\nlzo 2.10\nm4 1.4.19\nmbedtls 3.1.0\nmd4c 0.4.8\nmpdecimal 2.5.1\nmpfr 4.1.0\nmysql 8.0.29\nncurses 6.3\nnetpbm 10.86.32_1\nnettle 3.7.3\nnfpm 2.2.2\nnspr 4.33 4.29\nnss 3.58 3.78\nomniversion 0.43.3\nopen-mpi 4.0.5\nopencore-amr 0.1.5\nopenexr 3.1.5 3.1.3\nopenjdk 17.0.1_1\nopenjpeg 2.4.0\nopenssl@1.1 1.1.1n\nopus 1.3.1\np11-kit 0.24.1\npango 1.50.6\npcre 8.45\npcre2 10.36 10.40\npdf2svg 0.2.3_6\npinentry-mac 1.1.1.1\npixman 0.40.0\npkg-config 0.29.2_3\npodman 4.0.3\npoppler 22.04.0\npopt 1.18\nprotobuf 3.12.4 3.13.0 3.19.2 3.19.4\npython@3.10 3.10.4\npython@3.8 3.8.13\npython@3.9 3.9.12\nqemu 6.2.0_1\nqt 5.15.1 6.2.3_1\nrav1e 0.5.1 0.3.4\nreadline 8.1.2\nredis 6.0.8\nrsync 3.2.3\nrtmpdump 2.4+20151223_1\nrubberband 1.9.0 2.0.2\nsdl2 2.0.12_1 2.0.22\nsix 1.16.0_2\nsnappy 1.1.9\nsocat 1.7.4.3\nspeex 1.2.0\nsqlite 3.38.3\nsrt 1.4.2 1.4.4\nsshpass 1.06\nsvu 1.8.0\nswig 4.0.2\ntcl-tk 8.6.12_1\nterraform 0.14.7\ntesseract 5.1.0\ntheora 1.1.1\nunbound 1.15.0\nvault 1.6.3\nvde 2.3.2_1\nwakeonlan 0.41\nwebp 1.2.2\nx264 r3011 r3060\nx265 3.4 3.5\nxorgproto 2022.1\nxvid 1.3.7\nxxhash 0.8.0\nxz 5.2.5\nyq 4.7.1\nyubikey-agent 0.1.5\nzeromq 4.3.4\nzimg 3.0.4\nzlib 1.2.11\nzstd 1.5.2 1.5.1 1.4.5\ndb-browser-for-sqlite 3.12.1\nsecretive 1.0.2\nsekey 0.1\nxquartz 2.8.1\n"

	result, err := parseHomebrewOutput(vector)

	assert.Nil(t, err)
	assert.Equal(t, 179, len(result))

	item := result[0]
	assert.Equal(t, "allure", item.Name)
	assert.Equal(t, "2.13.5", item.Current)
	assert.Equal(t, 1, len(item.Installations))
	assert.Equal(t, "2.13.5", item.Installations[0].Version)

	item = result[1]
	assert.Equal(t, "ansible", item.Name)
	assert.Equal(t, "5.7.0", item.Current)
	assert.Equal(t, 1, len(item.Installations))
	assert.Equal(t, "5.7.0", item.Installations[0].Version)
}
