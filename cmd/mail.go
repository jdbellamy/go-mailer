package cmd

import (
	"github.com/spf13/cobra"
	"github.com/jdbellamy/go-mailer/mail"
	. "github.com/jdbellamy/go-mailer/middleware"
	"github.com/uber-go/zap"
	"time"
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
			Port: 25,
		}
		lf := cmd.LocalFlags()
		m := mail.Email{
			Sender: func() string { r, _ := lf.GetString("from"); return r}(),
			Recipients: func() []string { r, _ := lf.GetStringArray("to"); return r}(),
			Body: func() string { r, _ := lf.GetString("body"); return r}(),
			Subject: func() string { r, _ := lf.GetString("subject"); return r}(),
		}

		spinner := NewSpinner()
		time.Sleep(2 * time.Second)

		if err := mailer.Send(&m); err != nil {
			spinner.Stop()
			Z.Error("Unexpected error while sending message",
				zap.Error(err),
				zap.Object("email", m),
				zap.Object("mailer", mailer))
		} else {
			spinner.Stop()
			Z.Info("Email was successfully sent",
				zap.Object("message", m),
				zap.Object("mailer", mailer))
		}
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
