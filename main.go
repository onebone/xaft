package main


import (
	"fmt"
	"regexp"
	"flag"
	"io/ioutil"
	"strings"
	"strconv"
)

func main(){
	file := flag.String("f", "notes.txt", "file of notes")
	second := flag.Float64("s", 2.5, "seconds to increase")
	parse := flag.Bool("p", false, "parses stop watch dump")

	flag.Parse()
	
	b, err := ioutil.ReadFile(*file)
	if err != nil{
		panic(err)
	}

	str := strings.Replace(string(b), "\r", "", -1)
	if *parse {
		str = parseStopWatch(str)
	}else {
		str = strings.Trim(str, " ")
		str = strings.Replace(str, "\n", "", -1)
	}

	isSecond := true
	sec := ""
	result := ""
	for i := 0; i < len(str); i++ {
		if str[i] == ',' {
			if isSecond {
				fmt.Printf("[%s]\n", sec)
				origin, err := strconv.ParseFloat(sec, 64)
				if err != nil {
					panic("Invalid format")
				}
				result += fmt.Sprintf("%.2f", origin - *second)

				sec = ""
			}

			result += ","

			isSecond = !isSecond
			continue
		}

		if isSecond {
			sec += string(str[i])
		}else{
			result += string(str[i])
		}
	}

	fmt.Println(result + ",")
}

func parseStopWatch(str string) string{
	lines := strings.Split(str, "\n")
	reg, err  := regexp.Compile("[\\d]{2,}\\. ([\\d]{1}:[\\d]{2}:[\\d]{2}\\.[\\d])?.*")
	if err != nil {
		panic(err)
	}

	ret := ""
	for _, line := range lines {
		matching := reg.FindAllStringSubmatchIndex(line, -1)
		if len(matching) > 0 {
			ret += timeToSeconds(line[matching[0][2]:matching[0][3]]) + ",45,"
			continue
		}
	}

	return ret
}

func timeToSeconds(str string) string {
	reg, err := regexp.Compile(":|\\.")
	if err != nil{
		panic(err)
	}

	strs := reg.Split(str, -1)
	hour, _ := strconv.Atoi(strs[0])
	min, _ := strconv.Atoi(strs[1])
	sec, _ := strconv.Atoi(strs[2])
	ms, _ := strconv.Atoi(strs[3])

	return fmt.Sprintf("%d.%d", hour * 3600 + min * 60 + sec, ms)
}