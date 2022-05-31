package main

import (
	"fmt"
	"strings"
)

func main() {
	robotData := []string{"robot_arm", "robot_wheel", "robot_screw", "robot_ai", "bot_arm", "nod_arm", "nod_wheel",
		"nod_screw", "nod_ai", "robot_tech", "tag_ai", "mod_arm"}
	parts := "arm,wheel,screw,ai"
	buildRobot(robotData, parts)
}

func buildRobot(robotData []string, parts string) {
	partsList := strings.Split(parts, ",")
	robotMap := make(map[string]int)
	for i := 0; i < len(partsList); i++ {
		for j := 0; j < len(robotData); j++ {
			if strings.Contains(robotData[j], partsList[i]) {
				if robotMap[strings.Split(robotData[j], "_")[0]] != 0 {
					value := robotMap[strings.Split(robotData[j], "_")[0]]
					robotMap[strings.Split(robotData[j], "_")[0]] = value + 1
				} else {
					robotMap[strings.Split(robotData[j], "_")[0]] = 1
				}
			}
		}
	}

	for key, value := range robotMap {
		if value == len(partsList) {
			fmt.Println(key)
		}
	}
}
