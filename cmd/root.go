package cmd

import (
	"errors"
	"fmt"
	"time"

	dga "github.com/brutalgg/gobermann/internal/domaingeneratingalgorithm"
	"github.com/brutalgg/gobermann/pkg/cli"
	"github.com/brutalgg/gobermann/plugins/dga/locky"
	"github.com/brutalgg/gobermann/plugins/dga/necurs"
	"github.com/brutalgg/gobermann/plugins/dga/nymaim2"
	"github.com/miekg/dns"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:              "gobermann",
	Version:          "1.0alpha",
	PersistentPreRun: setup,
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
	rootCmd.PersistentFlags().IntP("delay", "d", 500, "Delay between requets in a burst in milliseconds")
	rootCmd.PersistentFlags().IntP("interval", "i", 720, "Delay between bursts in minutes")
	rootCmd.PersistentFlags().StringP("dns", "s", "1.1.1.1", "Target DNS Server")
	rootCmd.PersistentFlags().StringP("alg", "a", "locky", "The domain generating algorithm to use.[locky, nymaim2, necurs]")
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		cli.Fatalln(err)
	}
}

func setup(cmd *cobra.Command, args []string) {
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
	if _, err := selectDGA(alg); err != nil {
		cli.Fatal("%v is not one of the supported algorithms", alg)
	}

	if cmd.Flags().Changed("dryrun") {
		dns, _ := cmd.Flags().GetString("dns")
		cli.Debug("Testing Connection to %v", dns)
		if err := dnsQuery("no-reply-kt.com", dns); err != nil {
			cli.Fatalln(err)
		}
	} else {
		cli.Debugln("dryrun detected: no dns requets will be sent")
	}
}

func run(cmd *cobra.Command, args []string) {
	alg, _ := cmd.Flags().GetString("alg")
	interval, _ := cmd.Flags().GetInt("interval")

	for {
		dga, err := selectDGA(alg)
		if err != nil {
			cli.Fatalln("Empty DGA detected. How did you even hit this message?")
		}
		burst(cmd, dga)
		cli.Infoln("Waiting for interval...")
		time.Sleep(time.Minute * time.Duration(interval))
	}
}

func burst(cmd *cobra.Command, d dga.DomainGenerator) {
	burst, _ := cmd.Flags().GetInt("burst")
	delay, _ := cmd.Flags().GetInt("delay")
	dns, _ := cmd.Flags().GetString("dns")

	cli.Infoln("Starting Burst")
	for i := 0; i < burst; i++ {
		domain := d.GenerateDomain()
		cli.Debugln(domain)
		if cmd.Flags().Changed("dryrun") {
			dnsQuery(domain, dns)
		}
		time.Sleep(time.Millisecond * time.Duration(delay))
	}
	cli.Infoln("Burst Completed Successfully")
}

func selectDGA(x string) (dga.DomainGenerator, error) {
	switch x {
	case "locky":
		return locky.SeedRNG(1, 1, time.Now()), nil
	case "nymaim2":
		return nymaim2.SeedRNG(time.Now()), nil
	case "necurs":
		return necurs.SeedRNG(0, 9, time.Now()), nil
	}

	return dga.DefaultGenerator{}, errors.New("using empty generator")
}

func dnsQuery(f string, server string) error {
	var msg dns.Msg
	fqdn := dns.Fqdn(f)
	msg.SetQuestion(fqdn, dns.TypeA)
	if _, err := dns.Exchange(&msg, fmt.Sprintf("%v:53", server)); err != nil {
		return err
	}
	return nil
}
