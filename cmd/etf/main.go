// Copyright 2018 The go-ethereum Authors
// This file is part of go-ethereum-etf
//

//etf cmd

package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/cmd/utils"
	"gopkg.in/urfave/cli.v1"
	"os"
)

var gitCommit = "" // Git SHA1 commit hash of the release (set via linker flags)

var (
	app = utils.NewApp(gitCommit, "the evm command line interface")

	CMD = cli.StringFlag{
		Name: "cmd",
		Usage: "value likes:cAddr/" +
			"\n\r\rcAddr:get contract address  (before 4700000,etf block fork)",
	}
	DATADIR = cli.StringFlag{
		Name:  "statedb",
		Usage: "if cmd is cAddr ,datadir is local stateDB position",
	}
	DBHEAD = cli.StringFlag{
		Name:  "stateroot",
		Usage: "if cmd is cAddr ,head is the get block headHash",
	}
)

func init() {
	//app := cli.NewApp()
	app.Name = "etf"
	app.Usage = "make an etf cmd tool "
	app.Action = func(c *cli.Context) error {

		if c.String("cmd") == "cAddr" {
			var dir string
			dir = c.String("statedb")
			if dir == "" {
				fmt.Println("err! cmd likes :etf --cmd cAddr --statedb stateDBPath --stateroot 0x800cb7****edf6b2")
				return nil
			}
			stateroot := c.String("stateroot")
			if stateroot == ""{
				fmt.Println("err! cmd likes :etf --cmd cAddr --statedb StateDBPath --stateroot 0x800cb7****edf6b2")
				return nil
			}
			//fmt.Println("local blockchain addr[", c.String("datadir"),"]","[",dir,"]")
			fmt.Println(dir)
			err := GetContractAddr(dir,stateroot)
			if err != nil {
				fmt.Println("cmd run error ")
			}

		} else {
			fmt.Println("cmd error ,the help cmd is :etf --help")
		}
		return nil
	}

	app.Flags = []cli.Flag{
		CMD,
		DATADIR,
		DBHEAD,
	}

}

func main() {

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
