/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	apperrors "github.com/wakuwaku3/example-golang-cobra/lib/app_errors"
)

// sampleCmd represents the sample command
var sampleCmd = &cobra.Command{
	Use:   "sample",
	Short: "sample コマンドです",
	Long:  `Sample コマンドはサブコマンドのサンプルです。`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("sample called")
	},
}

var (
	name string
	age  int
	icon string
)

func init() {
	rootCmd.AddCommand(sampleCmd)

	sampleCmd.Flags().StringVarP(&name, "name", "n", "", "名前を指定します")
	sampleCmd.Flags().IntVarP(&age, "age", "a", 0, "年齢を指定します")
	sampleCmd.Flags().StringVarP(&icon, "icon", "s", "", "icon ファイルを指定します")

	apperrors.Handle(apperrors.Wrap(sampleCmd.MarkFlagRequired("name")))
	apperrors.Handle(apperrors.Wrap(sampleCmd.MarkFlagRequired("age")))
	apperrors.Handle(apperrors.Wrap(sampleCmd.MarkFlagFilename("icon", "png", "jpg", "jpeg")))
}
