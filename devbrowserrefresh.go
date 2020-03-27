package cbutil

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"time"
)

func DevRefreshBrowserOnLoadAndTemplateChange(domain string) {
	if os.Getenv("CHROMEREFRESH") == "1" && runtime.GOOS == "darwin" {
		go func() {
			Sleep(time.Second)

			cmd := exec.Command(`/usr/bin/osascript`, `-e`, fmt.Sprintf(`tell application "Google Chrome"
	set window_list to every window # get the windows

	repeat with the_window in window_list # for every window
		set tab_list to every tab in the_window # get the tabs

		repeat with the_tab in tab_list # for every tab
			if the URL of the_tab contains "%s" then
				tell the_tab to reload
			end if
		end repeat
	end repeat
end tell`, domain))
			e := cmd.Run()
			if e != nil {
				log.Println(e.Error())
			}

			e = exec.Command("pkill", "-f", "fswatch").Start()
			if e != nil {
				log.Println(e.Error())
			}

			watch := exec.Command(`/bin/bash`, `-c`, fmt.Sprintf(`/usr/local/bin/fswatch -e ".*" -i "\\.gohtml$" -o . | xargs -n1 -I {} osascript -e 'tell application "Google Chrome"
	set window_list to every window # get the windows

	repeat with the_window in window_list # for every window
		set tab_list to every tab in the_window # get the tabs

		repeat with the_tab in tab_list # for every tab
			if the URL of the_tab contains "%s" then
				tell the_tab to reload
			end if
		end repeat
	end repeat
end tell'`, domain))
			e = watch.Start()
			if e != nil {
				log.Println(e.Error())
			}
		}()
	}
}
