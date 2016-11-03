package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/urfave/cli"
)

const APP_VER = "0.0.1"

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

var GenStandardC = cli.Command{
	Name:   "c",
	Usage:  "gen language c",
	Action: gen_c,
	Flags: []cli.Flag{
		cli.StringFlag{Name: "name", Value: "default", Usage: "filename"},
	},
}

func gen_c(ctx *cli.Context) {
	name := ctx.String("name")
	if _, e := os.Stat(name + ".h"); e == nil {
		os.Remove(name + ".h")
	}
	if _, e := os.Stat(name + ".c"); e == nil {
		os.Remove(name + ".c")
	}

	header := `#ifndef _@NAME@_H_
#define _@NAME@_H_
	
#include "foundation/itype.h"
#include "foundation/core/iref.h"
	
/* Set up for C function definitions, even when using C++ */
#ifdef __cplusplus
extern "C" {
#endif

/* Ends C function definitions when using C++ */
#ifdef __cplusplus
}
#endif

#endif /* _@NAME@_H_ */
`
	header = strings.Replace(header, "@NAME@", strings.ToUpper(name), -1)
	ioutil.WriteFile(name+".h", []byte(header), 0644) // -rw-r--r--

	d, _ := os.Getwd()
	ds := strings.Split(d, "isee/code/")
	if len(ds) >= 2 {
		source := `#include "@@HEADER@@"

`
		hd := filepath.Join(ds[1], name+".h")
		source = strings.Replace(source, "@@HEADER@@", hd, -1)
		ioutil.WriteFile(name+".c", []byte(source), 0644) // -rw-r--r--
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "touch-file"
	app.Usage = "gen the *.h/*.c"
	app.Version = APP_VER
	app.Commands = []cli.Command{
		GenStandardC,
	}
	app.Flags = append(app.Flags, []cli.Flag{}...)
	app.Run(os.Args)
}
