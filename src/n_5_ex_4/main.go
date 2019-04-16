package main

func main(){
	s := struct{
		first string
		friends map[string]int
		towtrucks []string
	}{
		first: "James",
		friends: map[string]int{
			"Moneypenny": 555,
			"Q": 777,
			"M": 888,
		},
		towtruck []string{
			"monstertruck",
			"suvvan",
		},
	}
	fmt.Println(s)
	fmt.Println(s.friends)
	fmt.Println(s.towtruck)

	for k, v := range s.friends{
		fmt.Pritnln(k, v)
	}
	for i, val := range s.towtruck{
		fmt.Println(i, val)
	}
}