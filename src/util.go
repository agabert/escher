package main

import (
	"crypto/sha256"
	"fmt"
	"github.com/Sirupsen/logrus"
	"io"
	"os"
	"strings"
)

var embargo = []string{"/space/escher/trunk/",
	"/Library/",
	"/usr/local/Cellar",
	"/usr/local/lib/python",
	"screenshot",
	"Screenshot",
	"ScreenShot",
	"Screen Shot",
	"TerrariaClone/textures/",
	"/usr/lib/",
	"/usr/share/",
	"/usr/NX/share/",
	"/System/Library",
	"/.title.png",
	"/.thumb",
	"/Applications/Numbers.app/",
	"/midonet-manager/",
	"/.vscode/",
	"/.vscode-insiders/",
	"/mastodon/",
	"/github.com/agabert/",
	"/github.com/midonet/",
	"/doc/assets/",
	"/bower_components/",
	"/public/assets/",
	"/node_modules/",
	"/Keynote/",
	".key/Data/",
	".key/preview-micro.jpg",
	".key/preview-web.jpg",
	".key/preview.jpg",
	".graffle/",
	"/assets/",
	"/FullEMP/",
	"/Samples2/",
	"/Current/allpixels/",
	"/Snippets/",
	"/Setup/diablo3/",
	"/resources/media/version",
	"/resources/media/face/",
	"/resources/derivatives",
	"/resources/proxies/derivatives/",
	"/resources/renders/",
	"/www.dah4ever.de/",
	"/wp-content/",
	"/wp-includes/",
	"/wp-admin/",
	"/midonet-docs/",
	"/public/books/",
	"/projects/inovex/",
	"/home/agabert/.eclipse/",
	"/workspace/trunk/",
	"/MSTR_64bit/",
	"/Active_Mining/",
	"/Berlitz/",
	"/TSS/",
	"/OLD_SVN_",
	"/hardened-toolchain-development",
	"/presentations/gentoo-hardened-fosdem07",
	"/sspx_paper/",
	"/ssxp_paper/",
	"/FH-Trier_Studium/",
	"/var/www/",
	"/joinus_update/",
	"/ftblite/",
	"/.presentations/",
	"/www.midokura.de/",
	"_Life.Altis/",
	"/steam/Steam/steamapps/",
	"/files/virtualserver_1/",
	"/MW/page/",
	"/DZCP-",
	"/html.port81/",
	"/html.OLD001/",
	"/var/www/dah4ever/",
	"DHS/Server/",
	"/Arma 3 - Other Profiles/",
	"/projects/midokura/github/",
	"/static/dashboard/img/",
	"/static/horizon/",
	"/openstack_dashboard/",
	"github.com_OLD_REPOS",
	"/var/lib/latpak/",
	"/var/lib/app-info/",
	"/thumbnails/",
	"/.thumbnails/",
	"/.thumbnail/",
	"/.openshot_qt/thumbnail/",
	"/dashboard/docs/design/",
	"/screenshots/20",
	"/gitlab.app40.de/agabert/",
	"/.gradle/",
	"/Pingendo/New Site/assets/",
	"/Books/dspguide/",
	"/.cache/",
	"/.config/",
	"/.local/",
	"/.Trash/",
	"/var/lib/docker/aufs/diff",
	"/var/lib/docker/overlay2",
	"/var/lib/app-info/",
	"agabert/Books/",
	"/google-chrome/Default/Extensions/",
	"/.tile.png",
	"/elegant/",
	"/github.com/apereo/",
	"/private/var/root/Library/Caches/",
	"/opt/vagrant/",
	"github.com/openshift/",
	"github.com/cirosantilli/",
	"github.com/hashicorp/",
	"/kops/vendor/",
	"/Applications/",
	"/sandworm/buckets/",
	".mindnode/QuickLook/",
	".mindnode/resources/",
	"/opt/Rocket.Chat/programs/",
	"/acs-engine/docs/images/",
	"/cert-manager/vendor/",
	"/dashboard/docks/desin/mockups/",
	"/gitlab.xoreaxeax.de/",
	"/gitlab.benningen.xoreaxeax.de/",
	"/gitlab.inovex.de/",
	"/lib/firmware/",
	"/tank001/space/backups/storage001/space3/public/videos/",
	"/jpegvideocomplement_",
	"/teamspeak/sound/default",
}

func filterPath(filePath string) bool {
	i := 0

	for range embargo {
		if strings.Contains(strings.ToLower(filePath), strings.ToLower(embargo[i])) {
			logrus.Debug("filtered: ", filePath, "(", embargo[i], ")")
			return true
		}
		i++
	}

	return false
}

func filterExtension(extension string) bool {
	switch strings.ToLower(extension) {
	case
		"tga",
		"tif",
		"tiff",
		"gif",
		"png",
		"jpg",
		"jpeg",
		"bmp",
		"xcf",
		"psd",
		"pcx",
		"rle",
		"mp3",
		"wav",
		"ogg",
		"avi",
		"flv",
		"mov",
		"mpg",
		"mpeg",
		"mp4":
		return false
	}
	return true
}

func getFileChecksum(fPath string) string {
	fHandle, err := os.Open(fPath)

	if err != nil {
		logrus.Fatal(err)
	}

	defer fHandle.Close()

	hasher := sha256.New()

	if _, err := io.Copy(hasher, fHandle); err != nil {
		logrus.Fatal(err)
	}

	return fmt.Sprintf("%x", hasher.Sum(nil))
}
