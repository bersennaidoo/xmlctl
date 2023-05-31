/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
  "database/sql"

  _ "github.com/lib/pq"

	"github.com/bersennaidoo/xmlctl/internal/application/fsys"
  "github.com/bersennaidoo/xmlctl/pkg/configuration"
	"github.com/bersennaidoo/xmlctl/internal/repository"
  
	"github.com/spf13/cobra"
)

// uploadCmd represents the upload command
var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "uploads a xml application configuration",
	Long: `uploads a xml application configuration file by
  passing in --filename or -f flag followed by the
  filepath of xml configuration file and --name or -name appname.
  
  Examples:
  ./xmlctl upload -f or --filename test.xml -n or --name app1`,

	RunE: func(cmd *cobra.Command, args []string) error {
		filename, err := cmd.Flags().GetString("filename")
		if err != nil {
			fmt.Printf("error retrieving filename: %s\n", err.Error())
			return err
		}

		name, err := cmd.Flags().GetString("name")
		if err != nil {
			fmt.Printf("error retrieving name: %s\n", err.Error())
			return err
		}
		fmt.Println("Uploading", filename, "...")

    dsn := configuration.NewXmldConfig()

    db, _ := sql.Open("postgres", dsn.Dsn)

		qsrv := repository.New(db)

		hand := fsys.Handler{
			Qsrv:    qsrv,
			AppName: name,
			Fname:   filename,
		}

		app := hand.UploadFile()
		fmt.Println(string(app.Config.([]byte)))
    return nil
	},
}

func init() {
	uploadCmd.Flags().StringP("filename", "f", "", "Filepath of filename to be uploaded")
	uploadCmd.Flags().StringP("name", "n", "", "Name of application")
	//uploadCmd.MarkPersistentFlagRequired("filename")
	//uploadCmd.MarkPersistentFlagRequired("name")
	rootCmd.AddCommand(uploadCmd)
}
