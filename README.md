## Kasa Home Plug CLI tool

A small utility to manage TP Link Kasa Homeplugs.

## Usage

Full details are available using `--help`, here are some example commands:

```
$ kasa-plug --help
usage: kasa-plug --plug=PLUG [<flags>] <command> [<args> ...]

Manages TP-Link Kasa HS1xx Smart Plugs

Flags:
  --help       Show context-sensitive help (also try --help-long and --help-man).
  --plug=PLUG  IP address or hostname of the plug

Commands:
  help [<command>...]
    Show help.

  info [<flags>]
    Retrieves device information

  energy [<flags>]
    Retrieves energy usage information

  on
    Turns the plug on

  off
    Turns the plug off

  reboot
    Reboots the plug

  status
    Retrieves power state
```

```
$ kasa-plug --plug dehumidifier-plug info
TP-Link Kasa Plug 'Dehumidifier' @ dehumidifier-plug

  Device:

              Alias: Dehumidifier
              Model: Smart Wi-Fi Plug With Energy Monitoring (HS110(UK))
        Power State: On
                LED: Off
      Power On Time: 16 minutes 52 seconds (1012 seconds)

  Versions:

           Software: 1.5.7 Build 180806 Rel.135437
           Hardware: 2.1

  Network:

        MAC Address: D8:0D:17:D8:54:6D
    Network Address: dehumidifier-plug
    Signal Strength: -75 dBm
```

```
$ kasa-plug --plug dehumidifier-plug energy
  Power State: On
        Alias: Dehumidifier
Power On Time: 17 minutes 34 seconds (1054 seconds)
 Current Watt: 731.538 W
          Amp: 3.094 A
   Total Watt: 11.263 kWh
         Volt: 238.192 V
```

```
$ kasa-plug --plug dehumidifier-plug energy --json
{"sw_ver":"1.5.7 Build 180806 Rel.135437","hw_ver":"2.1","type":"IOT.SMARTPLUGSWITCH","model":"HS110(UK)","mac":"D8:0D:17:D8:54:6D","dev_name":"Smart Wi-Fi Plug With Energy Monitoring","alias":"Dehumidifier","relay_state":1,"on_time":1159,"on_time_string":"19 minutes 19 seconds","active_mode":"none","feature":"TIM:ENE","updating":0,"rssi":-75,"led_off":0,"longitude_i":144395,"latitude_i":358544,"hwId":"0750E2C15BB77902833ABF45366B8E9A","fwId":"00000000000000000000000000000000","deviceId":"8006B2F62F673091ED1A642C0B5AA7541B26675C","oemId":"AB8C79FE7869756511CDC455BDFE41EA","ntc_state":0,"power_on":true,"power_off":false,"address":"dehumidifier-plug","voltage_mv":238932,"volt":238.932,"current_ma":3123,"current_amp":3.123,"power_mw":740772,"power_w":740.772,"total_wh":11285,"total_watt":11.285}
```

## Status

It is usable and downloadable on the release page.  It uses a Go library [hs1xxplug](https://github.com/ripienaar/hs1xxplug) so you can build your own commands or extend that library to support additional features.
