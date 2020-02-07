package cbutil

import (
	"log"
	"os"
	"os/exec"
	"runtime"
	"time"
)

func DevRefreshBrowserOnLoadAndTemplateChange() {
	if os.Getenv("CHROMEREFRESH") == "1" && runtime.GOOS == "darwin" {
		go func() {
			Sleep(time.Second)

			cmd := exec.Command(`/usr/bin/osascript`, `-e`, `tell application "Google Chrome" to reload (tabs of its first window whose URL contains "localhost")`)
			e := cmd.Run()
			if e != nil {
				log.Println(e.Error())
			}

			watch := exec.Command(`/bin/bash`, `-c`, `/usr/local/bin/fswatch -e ".*" -i "\\.gohtml$" -o . | xargs -n1 -I {} osascript -e 'tell application "Google Chrome" to reload (tabs of its first window whose URL contains "localhost")'`)
			e = watch.Start()
			if e != nil {
				log.Println(e.Error())
			}
		}()
	}
}
