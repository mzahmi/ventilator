package main
import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/fatih/structs"
	"github.com/go-redis/redis"
	"github.com/mzahmi/ventilator/control/modeselect"
)

// Some default values: to be determined
var DefaultParams = UserInput{
	Mode:                "NA",
	BreathType:          "NA",
	PatientTriggerType:  "NA",
	TidalVolume:         0,
	Rate:                0,
	Ti:                  0,
	TiMax:               0,
	Te:                  0,
	IR:                  0,
	ER:                  0,
	PeakFlow:            0,
	PEEP:                0,
	FiO2:                0,
	PressureTrigSense:   0,
	FlowTrigSense:       0,
	FlowCyclePercent:    0,
	PressureSupport:     0,
	InspiratoryPressure: 0,
	UpperLimitVT:        0,
	LowerLimitVt:        0,
	RiseTime:            0,
	UpperLimitPIP:       0,
	LowerLimitPIP:       0,
	MinuteVolume:        0,
	UpperLimitMV:        0,
	LowerLimitMV:        0,
	UpperLimitRR:        0,
	LowerLimitRR:        0,
}

func initParams()
{
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	
	if fileExists("params.json") 
	{
		//Load json file and put in redis		
		jsonFile, err := os.Open("params.json")
		if err != nil {
			fmt.Println(err)
		}		

		err = client.Set("PARAMS", jsonFile, 0).Err()
		if err != nil {
			fmt.Println(err)
		}
		defer jsonFile.Close()
 	} else {
		json, err := json.Marshal(DefaultParams)

		err = client.Set("PARAMS", json, 0).Err()
		if err != nil {
			fmt.Println(err)
		}
		_ = ioutil.WriteFile("params.json", json, 0644)
	}

	// Confirm params
	val, err := client.Get("PARAMS").Result()
	if err != nil {
		fmt.Println(err)
	}			
	fmt.Println(val)

}