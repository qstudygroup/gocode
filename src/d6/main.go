package main

import (
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"
)

type student struct {
	ID  int
	Gpa float64
}

func round(x, unit float64) float64 {
	return math.Round(x/unit) * unit
}

type byGpa []student

func (bg byGpa) Len() int {
	return len(bg)
}

func (bg byGpa) Swap(i, j int) {
	bg[i], bg[j] = bg[j], bg[i]
}

func (bg byGpa) Less(i, j int) bool {
	return bg[i].Gpa < bg[j].Gpa
}

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var students []student
	for i := 0; i <= 99; i++ {
		gpa := round((r.Float64()*2)+2, 0.05)
		s := student{
			ID:  i,
			Gpa: gpa,
		}
		students = append(students, s)
	}

	sort.Sort(byGpa(students))
	for _, s := range students {
		fmt.Printf("Student ID: %d, GPA: %.2f \n ", s.ID, s.Gpa)
	}

	bs, err := json.Marshal(students)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(bs))
}
