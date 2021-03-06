// Package mysql_generate is a tool to generate basic models and database scripts from the database
// By dyy
package main

import (
	"flag"
	"fmt"
	mysql "github.com/lights-T/lib-go/database/mysql_generate/model"
	"os"
	"strings"
)

var (
	addr     string
	database string
	table    string
	help     bool
	V        bool
	v        bool
)

const CurrentVersion = "1.0.3"

func main() {

	flag.BoolVar(&help, "help", false, "get help")
	flag.StringVar(&addr, "a", "", "mysql connection address,like a user:password@tcp(127.0.0.1)")
	flag.StringVar(&database, "d", "", "mysql database name,like d_user")
	flag.StringVar(&table, "t", "", "mysql table name,like t_user")

	flag.BoolVar(&v, "v", false, "get version")
	flag.BoolVar(&V, "V", false, "get version")
	flag.Parse()

	flag.Usage = Usage

	if help {
		flag.Usage()
		return
	}
	if v || V {
		Version()
		os.Exit(0)
	}

	if len(addr) == 0 || len(database) == 0 {
		a, ok := os.LookupEnv("DATABASE_URL")
		if ok {
			database = strings.TrimRight(strings.TrimLeft(a, "/"), "?")
			addr = strings.TrimRight(a, "/")
		} else {
			flag.Usage()
			return
		}
	}
	mysql.SaveConfig(addr, database, table)
	mysql.Init()
}

func Usage() {
	fmt.Fprintf(os.Stderr, `
Generation Version: %s
Usage: NewModel [-adt] [-a address] [-d database] [-t table ]

Options:
`, CurrentVersion)
	flag.PrintDefaults()
}

func Version() {
	fmt.Fprintf(os.Stderr, `
Generation Version: %s
`, CurrentVersion)
}
