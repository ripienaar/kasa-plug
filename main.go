package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/alecthomas/template"
	"github.com/ripienaar/hs1xxplug"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	ipaddress  string
	jsonFormat bool
	hs1xx      *hs1xxplug.Plug
)

func main() {
	app := kingpin.New("kasa-plug", "Manages TP-Link Kasa HS1xx Smart Plugs")
	app.Version("0.0.1")
	app.Flag("plug", "IP address or hostname of the plug").Required().StringVar(&ipaddress)

	icmd := app.Command("info", "Retrieves device information")
	icmd.Flag("json", "Show JSON output").BoolVar(&jsonFormat)

	ecmd := app.Command("energy", "Retrieves energy usage information")
	ecmd.Flag("json", "Show JSON output").BoolVar(&jsonFormat)

	app.Command("on", "Turns the plug on")
	app.Command("off", "Turns the plug off")
	app.Command("reboot", "Reboots the plug")
	app.Command("status", "Retrieves power state")

	cmd := kingpin.MustParse(app.Parse(os.Args[1:]))

	hs1xx = hs1xxplug.NewPlug(ipaddress)

	switch cmd {
	case app.GetCommand("info").FullCommand():
		info()

	case app.GetCommand("on").FullCommand():
		kingpin.FatalIfError(hs1xx.PowerOn(), "Could not power on")

	case app.GetCommand("off").FullCommand():
		kingpin.FatalIfError(hs1xx.PowerOff(), "Could not power off")

	case app.GetCommand("reboot").FullCommand():
		kingpin.FatalIfError(hs1xx.Reboot(), "Could not reboot")

	case app.GetCommand("status").FullCommand():
		status()

	case app.GetCommand("energy").FullCommand():
		energy()
	}
}

func energy() {
	type combined struct {
		hs1xxplug.Info
		hs1xxplug.Energy
	}

	energynfo, err := hs1xx.Energy()
	kingpin.FatalIfError(err, "Could not retrieve plug information")

	nfo, err := hs1xx.Info()
	kingpin.FatalIfError(err, "Could not retrieve plug information")

	data := combined{
		Energy: *energynfo,
		Info:   *nfo,
	}

	report := `  Power State: {{if .On}}On{{else}}Off{{end}}
        Alias: {{.Alias}}
Power On Time: {{.OnTime}} ({{.OnTimeSeconds}} seconds)
 Current Watt: {{.PowerUseWatt}} W
          Amp: {{.Amp}} A
   Total Watt: {{.TotalWatt}} kWh
         Volt: {{.Volt}} V
`

	if jsonFormat {
		out, err := json.Marshal(data)
		kingpin.FatalIfError(err, "Could not JSON encode data")
		fmt.Printf(string(out))
	} else {
		tmpl, err := template.New("plug").Parse(report)
		kingpin.FatalIfError(err, "Could not process template")

		err = tmpl.Execute(os.Stdout, data)
		kingpin.FatalIfError(err, "Could not process template")
	}
}

func status() {
	nfo, err := hs1xx.PowerState()
	kingpin.FatalIfError(err, "Could not retrieve plug information")

	if nfo == hs1xxplug.PowerOn {
		fmt.Println("On")
	} else if nfo == hs1xxplug.PowerOff {
		fmt.Println("Off")
	} else {
		fmt.Println("Unknown")
	}
}
func info() {
	nfo, err := hs1xx.Info()
	kingpin.FatalIfError(err, "Could not retrieve plug information")

	report := `TP-Link Kasa Plug '{{.Alias}}' @ {{.Address}}

  Device:

              Alias: {{.Alias}}
              Model: {{.DeviceName}} ({{.Model}})
        Power State: {{if .On}}On{{else}}Off{{end}}
                LED: {{if eq .LEDOff 0}}Off{{else}}On{{end}}
      Power On Time: {{.OnTime}} ({{.OnTimeSeconds}} seconds)

  Versions:

           Software: {{.SoftwareVersion}}
           Hardware: {{.HardwareVersion}}

  Network:

        MAC Address: {{.MAC}}
    Network Address: {{.Address}}
    Signal Strength: {{.SignalStrength}} dBm

`

	if jsonFormat {
		out, err := json.Marshal(nfo)
		kingpin.FatalIfError(err, "Could not JSON encode data")
		fmt.Printf(string(out))
	} else {
		tmpl, err := template.New("plug").Parse(report)
		kingpin.FatalIfError(err, "Could not process template")

		err = tmpl.Execute(os.Stdout, nfo)
		kingpin.FatalIfError(err, "Could not process template")
	}
}
