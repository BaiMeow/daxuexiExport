package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

type student struct {
	name  string
	id    string
	times int
}

func main() {
	file, err := os.Open("./data.csv")
	if err != nil {
		log.Fatal(err)
	}
	reader1 := csv.NewReader(file)
	reader1.Read()
	data, err := reader1.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	file2, err := os.Open("./students.csv")
	if err != nil {
		log.Fatal(err)
	}
	reader2 := csv.NewReader(file2)
	var students []*student
	stu, err := reader2.Read()
	for err == nil {
		students = append(students, &student{stu[0], stu[1], 0})
		stu, err = reader2.Read()
	}

	for _, v := range data {
		for _, v2 := range v {
			if v2 == "" {
				continue
			}
			for _, stu := range students {
				if strings.Contains(v2, stu.name) || strings.Contains(v2, stu.id) {
					stu.times++
					continue
				}
			}
		}
	}
	for _, v := range students {
		fmt.Println(v.name, ":", v.times)
	}
}
