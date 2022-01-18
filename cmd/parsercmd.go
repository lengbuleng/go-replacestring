package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"replacefile/config"
	"replacefile/parser"
	"syscall"

	"github.com/spf13/cobra"
)

const (
	AppName = "replace doc"
)

var (
	configYml string
	StartCmd  = &cobra.Command{
		Use:          "parser",
		Short:        "Start data parser",
		Example:      AppName + "parser -c conf/settings.yml",
		SilenceUsage: true,
		PreRun: func(cmd *cobra.Command, args []string) {
			Setup()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			tip()
			return Run()
		},
	}
)

func tip() {
	usage := "欢迎使用 " + AppName + " " + Version + " 可以使用 " + "-h" + " 查看命令"
	fmt.Printf("%s\n", usage)
}
func init() {
	StartCmd.PersistentFlags().StringVarP(&configYml, "config", "c", "settings.yml", "Start with provided configuration file")
}

func Execute() {
	if err := StartCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}

func Setup() {
	//1. 读取配置
	config.Setup(configYml)
}

func Run() error {

	fmt.Println("Parser started:")
	fmt.Printf("Source files directory: %+v\n", config.ParserConfig.SrcPath)
	// fmt.Println("Enter Control + C Shutdown Server")

	parser.Parse()
	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, os.Interrupt, os.Kill, syscall.SIGTERM)
	<-osSignals
	fmt.Println("Shutdown Parser ... ")

	return nil
}
