package application

import (
	"os"
	"sort"

	"bygui86/konf/commands/completion"
	"bygui86/konf/commands/list"
	"bygui86/konf/commands/set"
	"bygui86/konf/commands/split"
	"bygui86/konf/commands/view"
	"bygui86/konf/logger"

	"github.com/urfave/cli"
)

type KubeConfiguratorApp struct {
	app *cli.App
}

func Create() *KubeConfiguratorApp {
	app := cli.NewApp()
	addGlobalConfig(app)
	//addGlobalFlags(app)
	//addBefore(app)
	addCommands(app)

	lastConfig(app)
	return &KubeConfiguratorApp{
		app: app,
	}
}

func addGlobalConfig(app *cli.App) {
	app.Name = "konf"
	app.Usage = "Kubernetes Configurator for an easy daily job"
	app.Version = "0.0.1"
	app.UseShortOptionHandling = true
	app.EnableBashCompletion = true
}

// TODO add logging flags?
//func addGlobalFlags(app *cli.App) {
//	app.Flags = []cli.Flag{
//		cli.StringFlag{
//			Name:  "config, c",
//			Usage: "Load configuration from `FILE`",
//			FilePath: "./config", // default value set from file (takes precedence over default values set from the environment "EnvVar")
//		},
//	}
//}

// TODO add logger init?
//func addBefore(app *cli.App) {
//	app.Before = func(ctx *cli.Context) error {
//		return nil
//	}
//}

func addCommands(app *cli.App) {
	app.Commands = []cli.Command{
		*split.BuildCommand(),
		*list.BuildCommand(),
		*view.BuildCommand(),
		*set.BuildCommand(),
		*completion.BuildCommand(),
	}
}

func lastConfig(app *cli.App) {
	// sort flags in help section
	sort.Sort(cli.FlagsByName(app.Flags))
	// sort commands in help section
	sort.Sort(cli.CommandsByName(app.Commands))
}

func (k *KubeConfiguratorApp) Start() {
	err := k.app.Run(os.Args)
	if err != nil {
		logger.SugaredLogger.Errorf("❌ Error starting application: %s", err.Error())
		os.Exit(2)
	}
}
