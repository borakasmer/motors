/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
//https://github.com/spf13/cobra-cli/blob/main/README.md
// 1-) go install github.com/spf13/cobra-cli@latest (****Cobra Global yuklu degil ise..)
// 2-) go env" With the command "GOPATH" and "GOROOT" folders of GO can be seen. On Mac after "go install ..." the "go/bin/fuel" file under GOPATH => should be copied under "go/bin/" folder under GOROOT. (***Cobra Global yuklu degil ise..)
// 3-) CD Users/borakasmer/GolandProjects/Cbr650R
// 4-) cobra-cli init
*/
package main

import "github.com/borakasmer/motors/cmd"

func main() {
	cmd.Execute()
}
