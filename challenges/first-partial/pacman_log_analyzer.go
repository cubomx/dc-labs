package main

import (
	"fmt"
	"os"
	"log"
	"io/ioutil"
	"strings"
	"strconv"
)

type data struct {
	installedDate, lastUpdate, removalDate string
	manyUpdates int
}

type generalInfo struct {
	instPackg, remPackg, upgrPackg, currInstalled int
}

func  generateFile(packages map[string]*data, info *generalInfo, keys []string) {
	file, err := os.Create("packages_report.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	systemInfo := "Pacman Packages Report\n----------------------\n"
	systemInfo += "- Installed packages : " + strconv.Itoa(info.instPackg) + "\n"
	systemInfo += "- Removed packages   : " + strconv.Itoa(info.remPackg) + "\n"
	systemInfo += "- Upgraded packages  : " + strconv.Itoa(info.upgrPackg) + "\n"
	systemInfo += "- Current installed  : " + strconv.Itoa(info.currInstalled) + "\n\n"
	systemInfo += "List of packages\n----------------\n"
	_, error := file.WriteString(systemInfo)
	if error != nil {
		fmt.Println(err)
		file.Close()
		os.Exit(0)
	}
	for _, val := range keys{
		info := packages[val]
		data := "- Package Name        : " + val + "\n"
		data += "  - Install date      : " + info.installedDate + "\n"
		data += "  - Last update date  : " + info.lastUpdate + "\n"
		data += "  - How many updates  : " + strconv.Itoa(info.manyUpdates) + "\n"
		data += "  - Removal date      : " + info.removalDate + "\n"
		_, err := file.WriteString(data)
		if err != nil {
			fmt.Println(err)
			file.Close()
			os.Exit(0)
		}
	}
	fmt.Println("Successful writing")
	err = file.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}

func readMetadata(lines [] string) (map[string]*data, generalInfo) {
	packages := make(map[string] *data)
	info := generalInfo{0, 0, 0, 0}
	for _, value := range lines{
		separated := strings.Fields(value)
		if len(separated) >= 5{
			switch separated[3] {
			case "installed":{
				date := separated[0][1:] + " " + separated[1][:len(separated[1]) - 1]
				packages[separated[4]] = &data{date, date, "-", 0}
				info.instPackg++
			}
			case "upgraded" :{
				date := separated[0][1:] + " " + separated[1][:len(separated[1]) - 1]
				packages[separated[4]].manyUpdates++
				packages[separated[4]].lastUpdate = date
				if packages[separated[4]].manyUpdates == 1 {
					info.upgrPackg++
				}
			}
			case "removed" :{
				date := separated[0][1:] + " " + separated[1][:len(separated[1]) - 1]
				packages[separated[4]].removalDate = date
				info.remPackg++
			}
			case "reinstalled" : {
				date := separated[0][1:] + " " + separated[1][:len(separated[1]) - 1]
				packages[separated[4]].removalDate = "-"
				packages[separated[4]].installedDate = date
				packages[separated[4]].lastUpdate = date
				info.remPackg--
			}
			}
		}
	}
	return packages, info
}

func main() {
	fmt.Println("Pacman Log Analyzer")
	if len(os.Args) < 2 {
		fmt.Println("You must send at least one pacman log file to analize")
		fmt.Println("usage: ./pacman_log_analizer <logfile>")
		os.Exit(1)
	}
	text, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal("Unable to read file: %v", err)
	}
	lines := strings.Split(string(text), "\n")
	packages, info := readMetadata(lines)

	info.currInstalled = info.instPackg - info.remPackg
	fmt.Println(info.currInstalled)
	keys := make([]string, 0, len(packages))

	for key := range packages {
		keys = append(keys, key)
	}
	generateFile(packages, &info, keys)
}