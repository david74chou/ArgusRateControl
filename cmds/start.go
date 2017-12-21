package cmds

import (
	"bytes"
	"fmt"
	"os/exec"
	"path"
	"strconv"

	"github.com/david74chou/ArgusRateControl/amtk"
	"github.com/david74chou/ArgusRateControl/log"
	"github.com/spf13/cobra"
)

var cmdStartRateControl = &cobra.Command{
	Use:           "start",
	Short:         "Start RateControl testing",
	Run:           startRateControlRun,
	SilenceErrors: true,
}

type rcProfile struct {
	mode             string
	resolution       string
	targetBitrate    int
	maxTargetBitrate int
	compression      int
}

const MP4_RECORDING_TIME = 60

var rcProfiles = []rcProfile{
	{"cbr", "320x180", 100, 300, 80},
	{"vbr", "640x360", 100, 300, 65},
	{"vbr", "640x360", 100, 300, 60},
}

func init() {
	cmdStartRateControl.Flags().SortFlags = false

	RootCmd.AddCommand(cmdStartRateControl)
}

func startRateControlRun(cmd *cobra.Command, args []string) {
	url, err := normalizeURL(rtspURL)
	if err != nil {
		log.WARN(cmd, "invalid url: %s", rtspURL)
		return
	}

	api := amtk.New(&amtk.AMTKAPIParams{
		APIHostURL:  url.Hostname(),
		APIUser:     "admin",
		APIPassword: "",
	})

	for i, p := range rcProfiles {

		log.INFO(cmd, "#%v: %v", i, p)

		if err = api.SetResolution(p.resolution); err != nil {
			log.WARN(cmd, err.Error())
			return
		}
		if err = api.SetCompression(p.compression); err != nil {
			log.WARN(cmd, err.Error())
			return
		}
		if err = api.SetRateControl(p.mode, p.targetBitrate, p.maxTargetBitrate); err != nil {
			log.WARN(cmd, err.Error())
			return
		}

		fileName := path.Join(
			targetDir,
			fmt.Sprintf("argus-%s-%s-%dkbps-%dkbps-compress%d.mp4", p.resolution, p.mode, p.targetBitrate, p.maxTargetBitrate, p.compression),
		)

		if err = ffmpegRTSPRecording(cmd, url.String(), fileName); err != nil {
			log.WARN(cmd, err.Error())
			return
		}
	}
}

func ffmpegRTSPRecording(c *cobra.Command, url, fileName string) error {

	// Start RTSP server
	cmd := exec.Command(
		"ffmpeg",
		"-y",
		"-t", strconv.Itoa(MP4_RECORDING_TIME),
		"-rtsp_transport", "tcp",
		"-i", url,
		"-vcodec", "copy",
		"-an",
		fileName)

	var b bytes.Buffer
	cmd.Stdout = &b
	cmd.Stderr = &b

	cmd.Start()

	// Wait a while
	//time.Sleep(5 * time.Second)

	//cmd.Process.Signal(syscall.SIGTERM)

	err := cmd.Wait()
	if err != nil {
		log.WARN(c, "%s", b.Bytes())
		return err
	}

	return nil
}
