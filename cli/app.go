package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const (
	appName      = "SDS Load Balancer (github.com/Seascape-Foundation/sds-load-balancer)"
	appUsage     = "sdslb"
	versionMajor = "0"
	versionMinor = "1"
	versionBuild = "0"
)

func getFilename(cmd *cobra.Command) (string, error) {
	fflags := cmd.Flags()

	if fflags.Changed("filename") {
		return fflags.GetString("filename")
	}

	return "", nil
}

func CreateAPP() {
	var rootCmd = &cobra.Command{
		Use: "sdslb",
		Run: func(cmd *cobra.Command, args []string) {
			fflags := cmd.Flags()
			verbose := fflags.Changed("verbose") == true

			filename, err := getFilename(cmd)

			if err != nil {
				os.Exit(1)
				return
			}

			RunServer(verbose, filename)
		},
	}

	rootCmd.Flags().BoolP("verbose", "v", false, "Help message for flag intone")
	rootCmd.Flags().StringP("filename", "f", "", "Set the filename as the configuration")

	statusCommand := &cobra.Command{
		Use: "status",
		Run: func(cmd *cobra.Command, args []string) {
			filename, err := getFilename(cmd)

			if err != nil {
				os.Exit(1)
				return
			}

			InternalStatus(filename)
		},
	}

	statusCommand.Flags().StringP("filename", "f", "", "Set the filename as the configuration")

	rootCmd.AddCommand(statusCommand)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
