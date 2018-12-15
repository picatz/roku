package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"

	"github.com/picatz/roku"
	"github.com/spf13/cobra"
)

func main() {
	// handle CTRL+C quit
	cleanup := func() {
		os.Exit(0)
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			cleanup()
		}
	}()

	var cmdDevices = &cobra.Command{
		Use:   "devices",
		Short: "Locate all of the roku devices on the local network",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			devices, err := roku.Find(roku.DefaultWaitTime)

			if err != nil {
				panic(err)
			}

			for _, device := range devices {
				info, err := device.DeviceInfo()
				if err != nil {
					fmt.Println("error:", err)
					return
				}
				wrapped := struct {
					Device string           `json:"device,omitempty"`
					Info   *roku.DeviceInfo `json:"info,omitempty"`
				}{
					Device: device.String(),
					Info:   info,
				}

				bytes, err := json.Marshal(&wrapped)
				if err != nil {
					fmt.Println("error:", err)
					return
				}
				fmt.Println(string(bytes))
			}
		},
	}

	var cmdFindRemote = &cobra.Command{
		Use:   "find-remote [device(s)]",
		Short: "Find one (or more) roku remote by making it beep",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			devices, err := roku.Find(roku.DefaultWaitTime)

			if err != nil {
				panic(err)
			}

			for _, device := range devices {
				for _, arg := range args {
					if device.String() == arg {
						err := device.FindRemote()
						if err != nil {
							fmt.Println("error:", err)
							return
						}
					}
				}
			}
		},
	}

	var cmdActiveApp = &cobra.Command{
		Use:   "active-app [device(s)]",
		Short: "Get the currently active application running on the roku(s)",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			devices, err := roku.Find(roku.DefaultWaitTime)

			if err != nil {
				panic(err)
			}

			for _, device := range devices {
				for _, arg := range args {
					if device.String() == arg {
						app, err := device.ActiveApp()
						if err != nil {
							fmt.Println("error:", err)
							return
						}
						wrapped := struct {
							Device string    `json:"device,omitempty"`
							App    *roku.App `json:"app,omitempty"`
						}{
							Device: device.String(),
							App:    app,
						}
						bytes, err := json.Marshal(wrapped)
						if err != nil {
							fmt.Println("error:", err)
							return
						}
						fmt.Println(string(bytes))
					}
				}
			}
		},
	}

	var rootCmd = &cobra.Command{Use: "roku"}
	rootCmd.AddCommand(cmdDevices)
	rootCmd.AddCommand(cmdFindRemote)
	rootCmd.AddCommand(cmdActiveApp)
	rootCmd.Execute()
}
