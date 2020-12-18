package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/hacktobeer/gopanasonic/cloudcontrol"
	pt "github.com/hacktobeer/gopanasonic/types"

	"github.com/spf13/viper"
)

var (
	commit  = "development"
	date    = "development"
	version = "development"

	configFlag  = flag.String("config", "gopanasonic.yaml", "Path of YAML configuration file")
	deviceFlag  = flag.String("device", "", "Device to issue command to")
	historyFlag = flag.String("history", "", "Display history: day,week,month,year")
	listFlag    = flag.Bool("list", false, "List available devices")
	modeFlag    = flag.String("mode", "", "Set mode: auto,heat,cool,dry,fan")
	offFlag     = flag.Bool("off", false, "Turn device off")
	onFlag      = flag.Bool("on", false, "Turn device on")
	statusFlag  = flag.Bool("status", false, "Display current status of device")
	tempFlag    = flag.Float64("temp", 0, "Set the temperature (in Celsius)")
	versionFlag = flag.Bool("version", false, "Show build version information")
)

func readConfig() {
	viper.SetConfigFile(*configFlag)
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
}

func main() {
	if len(os.Args) < 2 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	flag.Parse()

	if *versionFlag {
		fmt.Printf("version: %s\n", version)
		fmt.Printf("commit: %s\n", commit)
		fmt.Printf("date: %s\n", date)
		os.Exit(0)
	}

	readConfig()
	user := viper.GetString("username")
	pass := viper.GetString("password")
	server := viper.GetString("server")

	var client cloudcontrol.Client
	err := client.CreateSession("", user, pass, server)
	if err != nil {
		log.Fatalln(err)
	}

	if *listFlag {
		log.Println("Listing available devices.....")
		devices, err := client.ListDevices()
		if err != nil {
			log.Fatalln(err)
		}

		if len(devices) != 0 {
			log.Printf("%d device(s) found:\n", len(devices))
			for _, device := range devices {
				log.Println(device)
			}
		} else {
			log.Println("error: No devices for configured account")
		}
		os.Exit(0)
	}

	// Read device from flag
	if *deviceFlag != "" {
		client.SetDevice(*deviceFlag)
	}
	// Read device from configuration file
	configDevice := viper.GetString("device")
	if configDevice != "" {
		client.SetDevice(configDevice)
	}
	// Exit if no devices are configured
	if client.DeviceGUID == "" {
		log.Fatalln("error: No device configured, please use flag or configuration file")
	}

	if *statusFlag {
		log.Println("Fetching status.....")
		status, err := client.GetDevice()
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Printf("GUID: %s\n", status.DeviceGUID)
		fmt.Println("Capabilities:")
		fmt.Printf("Auto mode: %t\n", status.AutoMode)
		fmt.Printf("Heat mode: %t\n", status.HeatMode)
		fmt.Printf("Dry mode: %t\n", status.DryMode)
		fmt.Printf("Cool mode: %t\n", status.CoolMode)
		fmt.Printf("Fan mode: %t\n", status.FanMode)
		fmt.Printf("Fan Speed mode: %d\n", status.FanSpeedMode)
		fmt.Printf("Quiet mode: %t\n", status.QuietMode)
		fmt.Printf("Eco function: %t\n", status.EcoFunction)
		fmt.Printf("EcoNavi function: %t\n", status.EcoNavi)
		fmt.Printf("iAutoX: %t\n", status.IautoX)
		fmt.Printf("NanoeX: %t\n", status.Nanoe)
		fmt.Println("Current status:")
		fmt.Printf("Status: %s\n", pt.Operate[status.Parameters.Poperate])
		fmt.Printf("Online: %t\n", status.Parameters.Ponline)
		fmt.Printf("Temperature: %0.1f\n", status.Parameters.PtemperatureSet)
		fmt.Printf("Mode: %s\n", pt.ModesReverse[status.Parameters.PoperationMode])
	}

	if *historyFlag != "" {
		log.Printf("Fetching historical data for this %s.....\n", *historyFlag)
		history, err := client.GetDeviceHistory(pt.HistoryDataMode[*historyFlag])
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("#,AverageSettingTemp,AverageInsideTemp,AverageOutsideTemp")
		for _, v := range history.HistoryEntries {
			fmt.Printf("%v,%v,%v,%v\n", v.DataNumber+1, v.AverageSettingTemp, v.AverageInsideTemp, v.AverageOutsideTemp)
		}
	}

	if *onFlag {
		log.Println("Turning device on.....")
		err = client.TurnOn()
		if err != nil {
			log.Fatalln(err)
		}
	}

	if *offFlag {
		log.Println("Turning device off....")
		err = client.TurnOff()
		if err != nil {
			log.Fatalln(err)
		}
	}

	if *tempFlag != 0 {
		log.Printf("Setting temperature to %v degrees Celsius", *tempFlag)
		err = client.SetTemperature(*tempFlag)
		if err != nil {
			log.Fatalln(err)
		}
	}

	if *modeFlag != "" {
		err = client.SetMode(pt.Modes[*modeFlag])
		if err != nil {
			log.Fatalln(err)
		}
	}
}
