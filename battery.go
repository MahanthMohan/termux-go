package main

import (
	"fmt"
	"os/exec"
	"encoding/json"
)

func main() {
	out, err := exec.Command("termux-battery-status").Output()
	if err != nil {
		panic(err)
	}
	var batteryStatus map[string]interface{}
	json.Unmarshal([]byte(out), &batteryStatus)

	fmt.Printf("The Battery health is %s, percentage is %.2f, %s, and at an temperature of %.2f\n",
	batteryStatus["health"].(string), batteryStatus["percentage"].(float64), batteryStatus["status"].(string), batteryStatus["temperature"].(float64))
}
