package cmds

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/spf13/cobra"
)

const (
	ARGUSRC    = "argusrc"
	rtspScheme = "rtsp"
)

var (
	rtspURL   string
	targetDir string
)

func init() {
	RootCmd.PersistentFlags().StringVarP(&rtspURL, "url", "u", "rtsp://admin:@localhost:554", "Specify RTSP url and port")
	RootCmd.PersistentFlags().StringVarP(&targetDir, "dir", "o", "/tmp", "Target directory to save files")
}

var RootCmd = &cobra.Command{
	Use:   ARGUSRC,
	Short: ARGUSRC + " is a tool to invoke Argus rate control",
}

func normalizeURL(rtspURL string) (u *url.URL, err error) {
	// Fix scheme
	if strings.Index(rtspURL, "//") == 0 {
		rtspURL = rtspScheme + ":" + rtspURL
	}
	if strings.Index(rtspURL, "://") == -1 {
		rtspURL = rtspScheme + "://" + rtspURL
	}

	// Parse RTSP URL
	u, err = url.Parse(rtspURL)
	if err != nil {
		return
	}

	// Add default RTSP port if empty
	if u.Port() == "" {
		u.Host += fmt.Sprintf(":%v", 554)
	}

	return
}
