package main
import "fmt"
type student struct{
	nama string 
	nim int
	gpa float64
}
const NMAX int = 100
func main(){
	var A[NMAX]student
	var name string
	var n, i, id, j, index int
	var ipk, terkecil, compalier float64
	fmt.Scan(&n)
	for i =0; i < n; i++{
		fmt.Scan(&name, &id, &ipk)
		A[i].nama = name
		A[i].nim = id
		A[i].gpa = ipk
	}
	for i = 0; i < n; i++{
		terkecil = A[i].gpa
		index = i
		for j = i+1; j < n; j++{
			if A[j].gpa < terkecil{
				terkecil = A[j].gpa
				index = j
			}
		}
		compalier = A[i].gpa
		A[i].gpa = A[index].gpa
		A[index].gpa = compalier
	}
	for i = 0; i < n; i++{
		fmt.Println(A[i].gpa)
	}
}