/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	apperrors "github.com/wakuwaku3/example-golang-cobra/lib/app_errors"
	"github.com/wakuwaku3/example-golang-cobra/usecase/echo"
)

// echoCmd represents the echo command
var echoCmd = &cobra.Command{
	Use:   "echo",
	Short: "echo コマンドをラップしています",
	Long:  `echo コマンドは引数に指定した文字列をそのまま出力します。`,
	Run: func(cmd *cobra.Command, args []string) {
		apperrors.Handle(echo.NewUsecase(args).Execute())
	},
}

func init() {
	rootCmd.AddCommand(echoCmd)
	echoCmd.Flags().SetInterspersed(false)
}
