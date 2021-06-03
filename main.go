package main

import (
	"github.com/urfave/cli"
	"gorm_gen/generator"
	"os"
)

func main(){
//./gorm_gen -u=username -p=password -d=database -t=ALL -dir=./model  -host=host -port=port
	app  := cli.NewApp()
	app.Usage = "generator model for jinzhu/gorm"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "port",
			Value: "port",
			Usage: "port of mysql",
		},
		cli.StringFlag{
			Name:  "host",
			Value: "host",
			Usage: "host of mysql",
		},
		cli.StringFlag{
			Name:  "username,u",
			Value: "username",
			Usage: "Username of mysql",
		},
		cli.StringFlag{
			Name:  "password, p",
			Value: "password",
			Usage: "Password of mysql",
		},
		cli.StringFlag{
			Name:  "database, d",
			Value: "database",
			Usage: "select database",
		},
		cli.StringFlag{
			Name:  "table, t",
			Usage: "table name",
			Value: "ALL",
		},
		cli.StringFlag{
			Name:  "dir",
			Usage: "path which models will be stored",
			Value: "models",
		},
	}

	app.Action = generator.Generate
	app.Run(os.Args)
}
