package main

import (
	"log"
	"sort"
)

type IData struct {
	id int
}

func SorterExample() {
	data := []*IData{
		&IData{id: 123},
		&IData{id: 243},
		&IData{id: 564},
		&IData{id: 122},
	}
	for _, v := range data {
		log.Println(v.id)
	}

	log.Println("\nPost sort")

	sort.Slice(data, func(i, j int) bool {
		return data[i].id < data[j].id
	})

	for _, v := range data {
		log.Println(v.id)
	}

}
