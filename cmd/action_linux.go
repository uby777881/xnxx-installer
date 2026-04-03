package cmd

import (
	"fmt"
	"time"

	"github.com/InazumaV/V2bX/common/exec"
	"github.com/spf13/cobra"
)

var (
	startCommand = cobra.Command{
		Use:   "start",
		Short: "Start xnxx service",
		Run:   startHandle,
	}
	stopCommand = cobra.Command{
		Use:   "stop",
		Short: "Stop xnxx service",
		Run:   stopHandle,
	}
	restartCommand = cobra.Command{
		Use:   "restart",
		Short: "Restart xnxx service",
		Run:   restartHandle,
	}
	logCommand = cobra.Command{
		Use:   "log",
		Short: "Output xnxx log",
		Run: func(_ *cobra.Command, _ []string) {
			exec.RunCommandStd("journalctl", "-u", "xnxx.service", "-e", "--no-pager", "-f")
		},
	}
)

func init() {
	command.AddCommand(&startCommand)
	command.AddCommand(&stopCommand)
	command.AddCommand(&restartCommand)
	command.AddCommand(&logCommand)
}

func startHandle(_ *cobra.Command, _ []string) {
	r, err := checkRunning()
	if err != nil {
		fmt.Println(Err("check status error: ", err))
		fmt.Println(Err("xnxx failed to launch."))
		return
	}
	if r {
		fmt.Println(Ok("xnxx is already running and does not need to be started again. Please select Restart if you need to restart."))
	}
	_, err = exec.RunCommandByShell("systemctl start xnxx.service")
	if err != nil {
		fmt.Println(Err("exec start cmd error: ", err))
		fmt.Println(Err("xnxx failed to start"))
		return
	}
	time.Sleep(time.Second * 3)
	r, err = checkRunning()
	if err != nil {
		fmt.Println(Err("check status error: ", err))
		fmt.Println(Err("xnxx failed to start"))
	}
	if !r {
		fmt.Println(Err("xnxx may fail to start. Please check the log information using xnxx log later."))
		return
	}
	fmt.Println(Ok("xnxx started successfully. Please use xnxx log to view the runtime log."))
}

func stopHandle(_ *cobra.Command, _ []string) {
	_, err := exec.RunCommandByShell("systemctl stop xnxx.service")
	if err != nil {
		fmt.Println(Err("exec stop cmd error: ", err))
		fmt.Println(Err("xnxx failed to stop"))
		return
	}
	time.Sleep(2 * time.Second)
	r, err := checkRunning()
	if err != nil {
		fmt.Println(Err("check status error:", err))
		fmt.Println(Err("xnxx failed to stop"))
		return
	}
	if r {
		fmt.Println(Err("xThe nxx stop failed, possibly because the stop time exceeded two seconds. Please check the log information later."))
		return
	}
	fmt.Println(Ok("xnxx stopped successfully"))
}

func restartHandle(_ *cobra.Command, _ []string) {
	_, err := exec.RunCommandByShell("systemctl restart xnxx.service")
	if err != nil {
		fmt.Println(Err("exec restart cmd error: ", err))
		fmt.Println(Err("xnxx failed to restart"))
		return
	}
	r, err := checkRunning()
	if err != nil {
		fmt.Println(Err("check status error: ", err))
		fmt.Println(Err("xnxx failed to restart"))
		return
	}
	if !r {
		fmt.Println(Err("xnxx may fail to start. Please check the log information using xnxx log later."))
		return
	}
	fmt.Println(Ok("xnxx restarted successfully"))
}
