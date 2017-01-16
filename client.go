package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func main() {
	//Small dataset for testing all reports. Also assumming all devices move in a straight line(north)
	//data points :			1-10
	//Device-1 speed	 	0
	//Device-2 speed		0
	//Device-3 speed		70
	//Device-4 speed		70
	//Device-5 speed		30
	//The geo position to speed mapping might not make sense

	WriteDatePoint := func(device, speed string) {
		for i := 1; i <= 10; i++ {
			res, err := http.PostForm("http://localhost:3000/addData",
				url.Values{
					"device_id": {device},
					"latitude":  {"5.44991" + fmt.Sprintf("%d", (i*5))},
					"longitude": {"73.826066"},
					"time":      {"0510" + fmt.Sprintf("%d", (i*5))},
					"date":      {"20160822"},
					"status":    {"0x0A"},
					"speed":     {speed}})

			if err != nil {
				log.Fatal(err)
			}
			output, err := ioutil.ReadAll(res.Body)
			res.Body.Close()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%s", output)
		}

	}

	WriteDatePoint("device-1", "0")
	WriteDatePoint("device-2", "0")
	WriteDatePoint("device-3", "70")
	WriteDatePoint("device-4", "70")
	WriteDatePoint("device-5", "30")

}
