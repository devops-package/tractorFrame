package cli

import (
	"github.com/spf13/cobra"
)

type Commander interface {
	Cmd() *cobra.Command
}

func Root() *cobra.Command {
	//var version bool
	cmd := &cobra.Command{
		Use:     "firmeve",
		Short:   "Firmeve Framework",
		Version: "0",
	}
	//cmd.PersistentFlags().StringP("config", "C", "", "Config directory path(required)")
	//err := cmd.MarkFlagRequired("config")
	//if err != nil {
	//	firmeve.F(`logger`).(logging.Loggable).Fatal(err.Error())
	//}
	cmd.SetVersionTemplate("{{with .Short}}{{printf \"%s \" .}}{{end}}{{printf \"Version %s\" .Version}}\n")

	return cmd
}
