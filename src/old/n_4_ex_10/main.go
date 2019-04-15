package main

import "fmt"

func main() {
	m := map[string][]string{
		`Bond_James`:       []string{`Shaken, not stirred`, ` Martinis`, `Women`},
		`moneypennt_ miss`: []string{`James Bond`, `Literature`, `Computer science`},
		`no_dr`:            []string{`Being evil`, `Ice cream`, `Sunsets`},
	}
	for k, v := range m {
		fmt.Println("this is the record for", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}
	m["Nhimey the great"] = []string{`Max`, `golden retriever german sheperd mix`, `Brooklyn nine nine`}
	for k, v := range m {
		fmt.Println("this is the record for", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
		delete(m, "Nhimey the great")
		for k, v := range m {
			fmt.Println("this is the record for", k)
			for i, v2 := range v {
				fmt.Println("\t", i, v2)
			}
		}
	}
}