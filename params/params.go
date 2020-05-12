package params

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/go-redis/redis"
)

type UserInput struct {
	Mode                string
	BreathType          string
	PatientTriggerType  string
	TidalVolume         float32 // ml
	Rate                float32 // BPM
	Ti                  float32 // inhalation time
	TiMax               float32 // for PSV mode backup time control
	Te                  float32 // exhalation time
	IR                  float32 // inhalation ratio part IR:IE
	ER                  float32 // exhalation ratio part IR:IE
	PeakFlow            float32
	PEEP                float32 // 5-20 cmH2O
	FiO2                float32 // 21% - 100%
	PressureTrigSense   float32 // -0.5 to -2.0 cmH2O
	FlowTrigSense       float32 // 0.5 to 5 Lpm
	FlowCyclePercent    float32 // for flow cycling ranges from 0 to 100%
	PressureSupport     float32 // needs to be defined
	InspiratoryPressure float32 // Insp Pressure setting over PEEP
	UpperLimitVT        float32 // upper limit of tidal volume
	LowerLimitVT        float32 // lower limit of tidal volume
	RiseTime            float32 // needs to be defined
	UpperLimitPIP       float32 // upper limit of airway peak prssure
	LowerLimitPIP       float32 // lower limit of airway peak pressure
	MinuteVolume        float32 // minute volume in Lpm is the amount of gas expired per minute.
	UpperLimitMV        float32 // upper limit of minute volume
	LowerLimitMV        float32 // lower limit of minute volume
	UpperLimitRR        float32 // upper limit of monitored BPM
	LowerLimitRR        float32 // lower limit of monitored BPM
}

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
	IR:                  1,
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
	LowerLimitVT:        0,
	RiseTime:            0,
	UpperLimitPIP:       0,
	LowerLimitPIP:       0,
	MinuteVolume:        0,
	UpperLimitMV:        0,
	LowerLimitMV:        0,
	UpperLimitRR:        0,
	LowerLimitRR:        0,
}

// fileExists checks in current directory for a file with the filename specified and returns a bool
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// ReadParams gets the parameters of the UserInput struct from a redis client
func ReadParams(client *redis.Client) (DefaultParams UserInput) {
	val, err := client.Get("PARAMS").Result()
	check(err)
	//fmt.Println(val)
	json.Unmarshal([]byte(val), &DefaultParams)
	return DefaultParams
}

// WriteParams sets the parameters of the UserInput struct to a redis client
func WriteParams(client *redis.Client, DefaultParams UserInput) error {
	json, err := json.Marshal(DefaultParams)
	check(err)
	err = client.Set("PARAMS", json, 0).Err()
	check(err)
	return err
}

//InitParams initiates the paramters defined by the user input populates
// the redis client as well as saves the output into .json file
func InitParams(client *redis.Client) {

	if fileExists("params.json") {
		//Load json file and put in redis
		jsonFile, err := os.Open("params.json")
		check(err)
		byteValue, _ := ioutil.ReadAll(jsonFile)
		err = client.Set("PARAMS", byteValue, 0).Err()
		check(err)
		defer jsonFile.Close()
	} else {
		json, err := json.Marshal(DefaultParams)
		check(err)
		err = client.Set("PARAMS", json, 0).Err()
		check(err)
		_ = ioutil.WriteFile("params.json", json, 0644)
	}
	// // Confirm params
	// val, err := client.Get("PARAMS").Result()
	// check(err)
	// //fmt.Println(val)
	// err = json.Unmarshal([]byte(val), &DefaultParams)
	// check(err)
	DefaultParams = ReadParams(client)
}

// prints out the checked error err
func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
