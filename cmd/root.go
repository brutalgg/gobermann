package cmd

import (
	"github.com/brutalgg/cli"
	"github.com/brutalgg/gobermann/pkg/dnsspam"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:              "gobermann",
	Version:          "1.0.1alpha",
	PersistentPreRun: preChecks,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: run,
}

func init() {
	// Add additional commands to our CLI interface
	//rootCmd.AddCommand(Cmd)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringP("loglevel", "l", "info", "Include verbose messages from program execution [error, warn, info, debug]")
	rootCmd.PersistentFlags().BoolP("dryrun", "r", true, "When enabled dns traffic will not be sent over the wire")
	rootCmd.PersistentFlags().IntP("burst", "b", 15, "Number of requests in a burst of DNS traffic")
	rootCmd.PersistentFlags().IntP("delay", "d", 500, "Delay between requests in a burst in milliseconds")
	rootCmd.PersistentFlags().IntP("interval", "i", 720, "Delay between bursts in minutes")
	rootCmd.PersistentFlags().StringP("dns", "s", "1.1.1.1", "Target DNS Server")
	rootCmd.PersistentFlags().StringP("alg", "a", "locky", "The domain generating algorithm to use.[locky, nymaim2, necurs, monero]")
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		cli.Fatalln(err)
	}
}

func preChecks(cmd *cobra.Command, args []string) {
	l, _ := cmd.Flags().GetString("loglevel")
	switch l {
	case "error":
		cli.SetPrintLevel(3)
	case "warn":
		cli.SetPrintLevel(2)
	case "debug":
		cli.SetPrintLevel(0)
	default:
		cli.SetPrintLevel(1)
	}

	alg, _ := cmd.Flags().GetString("alg")
	if err := dnsspam.CheckAlgorithm(alg); err != nil {
		cli.Fatal("%v is not one of the supported algorithms", alg)
	}

	if cmd.Flags().Changed("dryrun") {
		dns, _ := cmd.Flags().GetString("dns")
		cli.Debug("Testing Connection to %v", dns)
		if err := dnsspam.DNSQuery("google.com", dns); err != nil {
			cli.Fatalln(err)
		}
	} else {
		cli.Debugln("dryrun detected: no dns requets will be sent")
	}
}

func run(cmd *cobra.Command, args []string) {
	//l, _ := cmd.Flags().GetString("loglevel")
	a, _ := cmd.Flags().GetString("alg")
	i, _ := cmd.Flags().GetInt("interval")
	b, _ := cmd.Flags().GetInt("burst")
	d, _ := cmd.Flags().GetInt("delay")
	s, _ := cmd.Flags().GetString("dns")

	spam := dnsspam.Spammer{
		//LogLevel:  l,
		Algorithm: a,
		Interval:  i,
		Burst:     b,
		Delay:     d,
		DNSServer: s,
		DryRun:    !cmd.Flags().Changed("dryrun"),
	}
	spam.Run()
}
