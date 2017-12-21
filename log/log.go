package log

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	ARGUSRC = "argusrc"
)

func INFO(cmd *cobra.Command, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	cmd.OutOrStdout().Write([]byte(fmt.Sprintf("%s: %s\n", ARGUSRC, msg)))
}

func WARN(cmd *cobra.Command, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	cmd.OutOrStdout().Write([]byte(fmt.Sprintf("%s: <WARN> %s\n", ARGUSRC, msg)))
}
