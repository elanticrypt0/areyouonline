package cmd

import (
	"fmt"
	"os"

	"github.com/elantycrypt0/go4it/src"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "listen",
	Short: "Escuchar un puerto",
	Long:  `Esta es una descripción con mayor detalle para este comando de ejemplo`,
	Run: func(cmd *cobra.Command, args []string) {
		src.CheckServer()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
