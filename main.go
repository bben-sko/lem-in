package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile(os.Args[1])
	if err != nil {
		return
	}
	sss := make(map[string][]string)
	var ants int
	var table []string
	star := ""
	end := ""
	file_spl := strings.Split(string(file), "\r\n")
	for i := 0; i < len(file_spl); i++ {
		if i == 0 {
			ants, err = strconv.Atoi(file_spl[i])
			if err != nil {
				fmt.Println(err)
			}
			continue
		}
		strnodes := strings.Split(file_spl[i], " ")
		if len(strnodes) == 3 {
			table = append(table, strnodes[0])

		}
		s := strings.Split(file_spl[i], "-")
		if len(s) == 2 {
			sss[s[0]] = append(sss[s[0]], s[1])
			sss[s[1]] = append(sss[s[1]], s[0])
		} else if strnodes[0] == "##start" {
			s := strings.Split(file_spl[i+1], " ")
			star = s[0]

			if err != nil {
				fmt.Println(err)
				return
			}
		} else if strnodes[0] == "##end" {
			s := strings.Split(file_spl[i+1], " ")
			end = s[0]
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
	fmt.Println(ants, table, sss, star, end)
	Handpath(ants, table, sss, star, end)

}

func Handpath(ants int, table []string, relation map[string][]string, star string, end string) {
	_ = ants
	_ = table
	var result []string
	var path [][]string
	j := 0
	done := false
	s := star

	for !done {
		result = append(result, star)
		for star != end {

			indix := relation[star][j]
			if Is(indix, result) || Check(indix, path,end) {
				j++
				if len(relation[star]) <= j {
					done = true
					break
				}
			} else if Find_end(relation[star], end) {
				//fmt.Println(indix)
				result = append(result, end)
				break
			} else {
				result = append(result, indix)
				star = indix
				j = 0
			}
		}
		if result != nil && result[len(result)-1] == end {
			path = append(path, result)
			result = nil
			j = 0
			star = s
		}
	}
	fmt.Println(path)

}

func Is(i string, find []string) bool {
	if len(find) == 0 {
		return false
	}
	for j := 0; j < len(find); j++ {
		if find[j] == i {
			return true
		}
	}
	return false
}
func Check(i string, path [][]string, end string) bool {
	if len(path) == 0 {
		return false
	}
	
	for k := 0; k < len(path); k++ {
		for l := 1; l < len(path[k])-1; l++ {
			if path[k][l] == i  {
				return true
			}
		}
	}
	if i == end {
		return false
	}

	return false
}
func Find_end(tbl []string, end string) bool {
	for i := 0; i < len(tbl); i++ {
		if tbl[i] == end {
			return true
		}
	}
	return false
}
