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
var readCmd = &cobra.Command{
	Use:   "read",
	Short: "reads a xml application configuration",
	Long: `reads a xml application configuration file by
  by application --name or -n appname.
  
  Examples:
  ./xmlctl read --name app1`,

	RunE: func(cmd *cobra.Command, args []string) error {
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			fmt.Printf("error retrieving name: %s\n", err.Error())
			return err
		}
		fmt.Println("reading", name, "...")

    dsn := configuration.NewXmldConfig()

    db, _ := sql.Open("postgres", dsn.Dsn)

		qsrv := repository.New(db)

		hand := fsys.Handler{
			Qsrv:    qsrv,
			AppName: name,
		}

		hand.ReedFile()
    return nil
	},
}

func init() {
	readCmd.Flags().StringP("name", "n", "", "Name of application")
	//uploadCmd.MarkPersistentFlagRequired("filename")
	//uploadCmd.MarkPersistentFlagRequired("name")
	rootCmd.AddCommand(readCmd)
}

