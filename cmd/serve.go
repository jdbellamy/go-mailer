package cmd

import (
	"github.com/spf13/cobra"
	"net/http"
	"log"
	"github.com/fatih/color"
	"fmt"
	"github.com/jdbellamy/go-mailer/rest"
)

// serveCmd represents the serve command
var port int
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		port := fmt.Sprintf(":%d", port)
		router := rest.NewRouter()
		log.Printf("Listening on port %s\n", color.BlueString(port))
		log.Fatal(http.ListenAndServe(port, router))
	},
}

func init() {
	RootCmd.AddCommand(serveCmd)
	serveCmd.PersistentFlags().IntVar(&port, "port", 8080, "Server port (default is :8080)")
}
