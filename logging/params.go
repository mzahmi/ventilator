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
	"github.com/mzahmi/ventilator/control/sensors"
)

func initParams()
{
	if fileExists("params.json") 
	{
        //Load json file and put in redis
 	} else {
		//load default params
    }
}