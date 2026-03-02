package install

import (
	"runtime"

	"github.com/gogf/gf/v2/os/gproc"

	"jettjia/igo-ddd-cli/util"
)

func Install() {
	goroot := runtime.GOROOT()

	src := ""
	dst := ""

	switch runtime.GOOS {
	case "windows":
		src = "./go-ddd-cli.exe"
		dst = `C:\Program Files` + src
		if goroot != "" && len(goroot) > 0 {
			dst = goroot + "/bin" + src
		}
		_, _ = util.Copy(src, dst)

	default:
		src = "go-ddd-cli"
		dst = `/usr/local/bin/` + src
		if goroot != "" && len(goroot) > 0 {
			dst = goroot + "/bin/" + src
		}
		_, _ = util.Copy(src, dst)
		_, _ = gproc.ShellExec("chmod -R 755 " + dst)
	}

	return
}
