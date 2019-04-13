package main

import (
	"fmt"
<<<<<<< HEAD
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
=======
	"math/rand"
	"time"
)

func main() {
	a := 0
	b := 0
	c := 0
	d := 0
	e := 0
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for x := 0; x < 100; x++ {
		num := r.Intn(100)
		switch {
		case num <= 20:
			a++
		case num > 20 && num <= 40:
			b++
		case 40 < num && num <= 60:
			c++
		case 60 < num && num <= 80:
			d++
		case 80 < num && num <= 100:
			e++
		}
	}
	fmt.Printf("there are %v number form 0 to 20 \n", a)
	fmt.Printf("there are %v number form 20 to 40 \n", b)
	fmt.Printf("there are %v number form 40 to 60 \n", c)
	fmt.Printf("there are %v number form 60 to 80 \n", d)
	fmt.Printf("there are %v number form 80 to 90 \n", e)
>>>>>>> 2df98c42e1a3958c55c8bcf99c6b1f1fe1aaee92
}
