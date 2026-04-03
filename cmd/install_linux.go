package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/InazumaV/V2bX/common/exec"
	"github.com/spf13/cobra"
)

var targetVersion string

var (
	updateCommand = cobra.Command{
		Use:   "update",
		Short: "Update xnxx version",
		Run: func(_ *cobra.Command, _ []string) {
			exec.RunCommandStd("bash",
				"<(curl -Ls https://raw.githubusercontents.com/InazumaV/V2bX-script/master/install.sh)",
				targetVersion)
		},
		Args: cobra.NoArgs,
	}
	uninstallCommand = cobra.Command{
		Use:   "uninstall",
		Short: "Uninstall xnxx",
		Run:   uninstallHandle,
	}
)

func init() {
	updateCommand.PersistentFlags().StringVar(&targetVersion, "version", "", "update target version")
	command.AddCommand(&updateCommand)
	command.AddCommand(&uninstallCommand)
}

func uninstallHandle(_ *cobra.Command, _ []string) {
	var yes string
	fmt.Println(Warn("Are you sure you want to uninstall xnxx? (Y/N)"))
	fmt.Scan(&yes)
	if strings.ToLower(yes) != "y" {
		fmt.Println("Uninstall canceled")
	}
	_, err := exec.RunCommandByShell("systemctl stop xnxx&&systemctl disable xnxx")
	if err != nil {
		fmt.Println(Err("exec cmd error: ", err))
		fmt.Println(Err("Uninstall failed"))
		return
	}
	_ = os.RemoveAll("/etc/systemd/system/xnxx.service")
	_ = os.RemoveAll("/etc/xnxx/")
	_ = os.RemoveAll("/usr/local/xnxx/")
	_ = os.RemoveAll("/bin/xnxx")
	_, err = exec.RunCommandByShell("systemctl daemon-reload&&systemctl reset-failed")
	if err != nil {
		fmt.Println(Err("exec cmd error: ", err))
		fmt.Println(Err("Uninstall failed"))
		return
	}
	fmt.Println(Ok("Uninstalled successfully"))
}
