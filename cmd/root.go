package cmd

import (
	"errors"
	"fmt"
	"net"
	"time"

	dga "github.com/brutalgg/gobermann/internal/domaingeneratingalgorithm"
	"github.com/brutalgg/gobermann/pkg/cli"
	"github.com/brutalgg/gobermann/plugins/dga/locky"
	"github.com/brutalgg/gobermann/plugins/dga/necurs"
	"github.com/brutalgg/gobermann/plugins/dga/nymaim2"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:              "gobermann",
	Version:          "",
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
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Include verbose messages from program execution")
	rootCmd.PersistentFlags().BoolP("dryrun", "r", true, "Determines whether DNS traffic is sent over the wire")
	rootCmd.PersistentFlags().IntP("burst", "b", 15, "Number of requests in a burst of DNS traffic")
	rootCmd.PersistentFlags().IntP("delay", "d", 1, "Delay between requets in a burst in seconds")
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
	alg, _ := cmd.Flags().GetString("alg")
	dns, _ := cmd.Flags().GetString("dns")

	if _, err := selectDGA(alg); err != nil {
		cli.Fatal("%v is not one of the supported algorithms", alg)
	}
	// trying to figure out if this test could be better
	cli.Info("Testing Connection to %v", dns)
	if _, err := net.Dial("tcp", fmt.Sprintf("%v:53", dns)); err != nil {
		cli.Fatal("could not connect to server %v", dns)
	}

}

func run(cmd *cobra.Command, args []string) {
	// todo implement other flags
	alg, _ := cmd.Flags().GetString("alg")
	for {
		dga, ok := selectDGA(alg)
		if ok != nil {
			cli.Fatal("Empty DGA detected. How did you even hit this message?")
		}
		burst(25, 0, dga)
		cli.Info("completed burst")
		time.Sleep(time.Minute)
	}
}

func burst(burst int, delay int, d dga.DomainGenerator) {
	// todo implement sending of messages
	// var msg dns.Msg
	// fqdn := dns.Fqdn("google.com")
	// msg.SetQuestion(fqdn, dns.TypeA)
	// dns.Exchange(&msg, "8.8.8.8:53")
	for i := 0; i < burst; i++ {
		cli.Info(d.GenerateDomain())
		time.Sleep(time.Second * time.Duration(delay))
	}
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
