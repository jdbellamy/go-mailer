package cmd

import (
	"github.com/spf13/cobra"
	"github.com/jdbellamy/go-mailer/mail"
	"github.com/fatih/color"
	"fmt"
)

// mailCmd represents the mail command
var mailCmd = &cobra.Command{
	Use:   "mail",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var mailer = mail.SmtpClient{
			Server: "localhost",
			Port: 26,
		}
		m := mail.Email{}
		m.Sender, _ = cmd.LocalFlags().GetString("from")
		m.Recipients, _ = cmd.LocalFlags().GetStringArray("to")
		m.Body, _ = cmd.LocalFlags().GetString("body")
		m.Subject, _ = cmd.LocalFlags().GetString("subject")
		spinner := Spinner()
		if err := mailer.Send(&m); err != nil {
			fmt.Printf("%s: %s\n",
				color.RedString("ERROR"),
				color.WhiteString(err.Error()))
		} else {
			fmt.Printf("%s: Message sent {%s->%s}\n",
				color.GreenString("SUCCESS"),
				color.WhiteString(m.Sender),
				color.WhiteString(fmt.Sprintf("%v", m.Recipients)))
		}
		spinner.Stop()
	},
}

func init() {
	RootCmd.AddCommand(mailCmd)
	mailCmd.Flags().Bool("sendmail", false, "Use Sendmail rather than SMTP")
	mailCmd.Flags().StringP("from", "f", "a@b.com", "*From* address")
	mailCmd.Flags().StringP("body", "b", "", "Message body")
	mailCmd.Flags().StringP("subject", "s", "", "Subject header")
	mailCmd.Flags().StringArrayP("to", "t", []string{""}, "*To* address")
}
