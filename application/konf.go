package application

import (
	"os"
	"sort"

	"github.com/bygui86/konf/commands/clean"
	"github.com/bygui86/konf/commands/completion"
	"github.com/bygui86/konf/commands/list"
	"github.com/bygui86/konf/commands/rename"
	"github.com/bygui86/konf/commands/set"
	"github.com/bygui86/konf/commands/split"
	"github.com/bygui86/konf/commands/view"
	"github.com/bygui86/konf/logger"

	"github.com/urfave/cli"
)

const (
	appName    = "konf"
	appUsage   = "Kubernetes Configurator for an easy daily job"
	appVersion = "0.0.3"
)

type KubeConfiguratorApp struct {
	app *cli.App
}

func Create() *KubeConfiguratorApp {
	logger.Logger.Debug("🐛 Creating application")
	app := cli.NewApp()
	setGlobalConfig(app)
	addCommands(app)
	setLastConfig(app)
	return &KubeConfiguratorApp{
		app: app,
	}
}

func setGlobalConfig(app *cli.App) {
	logger.Logger.Debug("🐛 Setting global configurations")
	app.Name = appName
	app.Usage = appUsage
	app.Version = appVersion
	app.UseShortOptionHandling = true
	app.EnableBashCompletion = true
}

func addCommands(app *cli.App) {
	logger.Logger.Debug("🐛 Adding commands")
	app.Commands = []cli.Command{
		*split.BuildCommand(),
		*list.BuildCommand(),
		*view.BuildCommand(),
		*set.BuildCommand(),
		*clean.BuildCommand(),
		*rename.BuildCommand(),
		*completion.BuildCommand(),
	}
}

func setLastConfig(app *cli.App) {
	logger.Logger.Debug("🐛 Setting last configurations")
	// sort flags in help section
	sort.Sort(cli.FlagsByName(app.Flags))
	// sort commands in help section
	sort.Sort(cli.CommandsByName(app.Commands))
}

func (k *KubeConfiguratorApp) Start() {
	logger.Logger.Debug("🐛 Starting application")
	err := k.app.Run(os.Args)
	if err != nil {
		logger.SugaredLogger.Errorf("❌ Error starting application: %s", err.Error())
		os.Exit(2)
	}
}
