package main

import "fmt"

func main() {
	x := make([]string, 50, 50)
	x = []string{`Alabama`, `Alaska`, `Arizona`, `Arkansas`, `California`, `Colorado`, `Connecticut`, `Delaware`, `Florida`, `Georgia`, `Hawaii`, `Idaho`, `Illinois`, `Indiana`, `Lowa`, `Kansas`, `Kentucky`, `Louisiana`, `Maine`, `Maryland`, `Massachuesetts`, `Michigan`, `Minnesota`, `Missipi`, `Missouri`, `Montana`, `Nebraska`, `Nevada`, `New Hampshire`, `New Jersey`, `New mexico`, `New york`, `North Carolina`, `North Dakota`, `Ohio,`, `Oklahoma`, `Oregon`, `Pensylvannia`, `Rhode Island`, ` South Carolina`, `South Dakota`, `Tennessee`, `Texas`, `Utah`, `Vermont`, `Virginia`, `Washington`, `West Virginia`, `Wisconsin`, `Wyoming`}
	for banana := 0; banana < len(x); banana++ {

		fmt.Println(banana, x[banana])
	}
}
