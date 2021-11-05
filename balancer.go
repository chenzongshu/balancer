package main

import (
"fmt"
"k8s.io/component-base/logs"
"os"
"balancer/cmd/app"
)

func main() {
	out := os.Stdout
	cmd := app.NewBalancerCommand(out)
	cmd.AddCommand(app.NewVersionCommand())

	logs.InitLogs()
	defer logs.FlushLogs()

	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}