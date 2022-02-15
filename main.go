package main

import (
	"github.com/huiyangz/psptk5/pkg/service"
	"github.com/spf13/cobra"
)

func main() {
	var cmdExtract = &cobra.Command{
		Use:   "extract [path of DATA.BIN] [path of EBOOT.BIN] [dir to extract files]",
		Short: "extract files from DATA.BIN",
		Long: `extract files from DATA.BIN.
	The EBOOT.BIN file contains segmentation informations,
	It needs to be decrypted to ELF first!`,
		Args: cobra.MinimumNArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			return service.Extract(args[0], args[1], args[2])
		},
	}

	var rootCmd = &cobra.Command{Use: "psptk5", CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true}}
	rootCmd.AddCommand(cmdExtract)
	rootCmd.Execute()
}
