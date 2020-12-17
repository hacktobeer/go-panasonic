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
	version = "development"
	commit  = "development"
	date    = "development"
)

func readConfig() {
	viper.SetConfigName("gopanasonic")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
}

func main() {
	versionFlag := flag.Bool("version", false, "Show build version information")
	listFlag := flag.Bool("list", false, "List available devices")
	deviceFlag := flag.String("device", "", "Device to issue command to")
	onFlag := flag.Bool("on", false, "Turn device on")
	offFlag := flag.Bool("off", false, "Turn device off")
	tempFlag := flag.Float64("temp", 0, "Set the temperature (in Celsius)")
	modeFlag := flag.String("mode", "", "Set mode: auto,heat,cool,dry,fan")
	statusFlag := flag.Bool("status", false, "Display current status of device")
	historyFlag := flag.String("history", "", "Display history: day,week,month,year")

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

	var client cloudcontrol.Client
	err := client.CreateSession("", user, pass)
	if err != nil {
		log.Fatalln(err)
	}

	if *listFlag {
		log.Println("TODO Listing available devices.....")
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
