package cmd

import (
	"fmt"
	"github.com/BurntSushi/toml"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	. "jira-bot/internal"
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"
)

var (
	confPath string
	jiraConf TomlConfig
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "jira-bot",
	Short: "Get filtered result from the JIRA server",
	Long:  ``,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		log.Error(err)
		os.Exit(1)
	}
}

func init() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			s := strings.Split(f.Function, ".")
			funcname := s[len(s)-1]
			_, filename := path.Split(f.File)
			line := strconv.Itoa(f.Line)
			return "[" + funcname + "]", "[" + filename + ":" + line + "]"
		},
		ForceColors: true,
	})
	log.SetReportCaller(true)

	rootCmd.PersistentFlags().StringVarP(&confPath, "conf", "f", "jira.toml", "Specify the configuration file path")
	rootCmd.AddCommand(wechatCmd)
}

func readConfig() {
	if _, err := toml.DecodeFile(confPath, &jiraConf); err != nil {
		log.Fatalln("Read Config file failed: " + err.Error())
	}
}
