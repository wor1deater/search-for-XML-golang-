package main

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
)


type Users struct {
	Row []struct {
		Id        int `xml:"id"`
		FirstName string `xml:"first_name"`
		LastName string `xml:"last_name"`
		Balance   string `xml:"balance"`
		Age       int `xml:"age"`
		Gender    string `xml:"gender"`
		About     string `xml:"about"`
	} `xml:"row"`
}

type SearchTask struct {
	Query      string
	OrderField string
	OrderBy    int
	Limit      int
	Offset     int
}

func handler(w http.ResponseWriter, r *http.Request) {
	/* 
		Собираем GET-params из запроса и передаем их в функцию
	*/
	query := r.FormValue("query")
	orderField := r.FormValue("order_field")
	orderBy := r.FormValue("order_by")
	limit := r.FormValue("limit")
	offset := r.FormValue("offset")
	o, err := strconv.Atoi(orderBy)
	l, err := strconv.Atoi(limit)
	off, err := strconv.Atoi(offset)
	xmlData, err := os.ReadFile("C:/Users/yaros/go/src/GO learn/4/99_hw/dataset.xml")
	if err != nil {
		fmt.Errorf("error: %v", err)
		panic(err)
	}
	data := readFromXml(xmlData)
	searchRequest := SearchTask{Query : query, OrderField: orderField, OrderBy: o, Limit: l, Offset: off}
	SearchServer(searchRequest, data, w)
}

func main() {
	/* 
		Ставим handler на главную страницу
	*/
	http.HandleFunc("/", handler)
	fmt.Println("Starting server at :8080")
	http.ListenAndServe(":8080", nil)
}


func SearchServer(task SearchTask, data *Users, w http.ResponseWriter) {
	/* 
		Собираем нужные параметры и сортируем данные, согласно запросу
	*/

	type AgeStruct struct {
		Key int
		Value int
	}

	id := []int{}
	names := []string{}
	ages := make(map[int]int)
	for i:=0; i < 35; i++ {
		if strings.Contains(data.Row[i].About, task.Query) || (strings.Contains(data.Row[i].FirstName + " " + data.Row[i].LastName, task.Query)) {
			ages[data.Row[i].Id] = data.Row[i].Age
			id = append(id, data.Row[i].Id)
			names = append(names, data.Row[i].FirstName + " " + data.Row[i].LastName)
		}
	}

	switch task.OrderField {
	case "Id":
		switch task.OrderBy {
		case 0:
			//
		case 1:
			sort.Ints(id)
		case -1:
			sort.Sort(sort.Reverse(sort.IntSlice(id)))
		}
		offset := id[task.Offset:]
		result := []int{}

		for i := 0; i < task.Limit; i++{
			result = append(result, offset[i])
		}
		fmt.Fprintln(w, "Sort by ID: ")
		fmt.Fprintln(w, "")
		for i:=0; i < len(result); i++ {
			for x:=0; x < 35; x++ {
				if result[i] == data.Row[x].Id {
					fmt.Fprintln(w, data.Row[x].Id, data.Row[x].FirstName + " " + data.Row[x].LastName + ", Age:", data.Row[x].Age, ", Balance: " + data.Row[x].Balance)
				}
			}
		}
	case "", "Name":
		switch task.OrderBy {
		case 0:
			//
		case 1:
			sort.Strings(names)
		case -1:
			sort.Sort(sort.Reverse(sort.StringSlice(names)))
		}
		offset := names[task.Offset:]
		result := []string{}

		for i := 0; i < task.Limit; i++{
			result = append(result, offset[i])
		}
		fmt.Fprintln(w, "Sort by names: ")
		fmt.Fprintln(w, "")
		for i:=0; i < len(result); i++ {
			for x:=0; x < 35; x++ {
				if result[i] == (data.Row[x].FirstName + " " + data.Row[x].LastName) {
					fmt.Fprintln(w, data.Row[x].Id, data.Row[x].FirstName + " " + data.Row[x].LastName + ", Age:", data.Row[x].Age, ", Balance: " + data.Row[x].Balance)
				}
			}
		}
	case "Age":
		var agekeys []AgeStruct
		for key, value := range ages {
			agekeys = append(agekeys , AgeStruct {key, value})
		}
		switch task.OrderBy {
		case 0:
			//
		case 1:
			sort.Slice(agekeys , func(i, j int) bool {
				return agekeys[i].Value < agekeys[j].Value
			})
		case -1:
			sort.Slice(agekeys , func(i, j int) bool {
				return agekeys[i].Value > agekeys[j].Value
			})
		}


		fmt.Fprintln(w, "Sort by ages: ")
		fmt.Fprintln(w, "")
		c := task.Limit
		for i := task.Offset; i < len(agekeys); i++ {
			for x:=0; x < 35; x++ {
				if (agekeys[i].Key == data.Row[x].Id) {
					if c != 0 {
						fmt.Fprintln(w, data.Row[x].Id, data.Row[x].FirstName + " " + data.Row[x].LastName + ", Age:", data.Row[x].Age, ", Balance: " + data.Row[x].Balance)
						c--
					}
				}
			}
		}

	default:
		panic("Wrong orderField")
	}

}



func readFromXml(xmlData []byte) *Users{
	/* 
		Читаем XML
	*/
	v := new(Users)
	err := xml.Unmarshal(xmlData, v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return nil
	}
	return v
}
