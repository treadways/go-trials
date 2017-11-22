package main
import "fmt"

func main() {
	var myname string
	str1 := "What is your name?"
	fmt.Println(str1)
	fmt.Scanf("%s", &myname)

	var age int
	fmt.Println(" ")
	fmt.Print("What is your age ", myname,"?")
	fmt.Println(" ")
	fmt.Scanf("%d", &age)
	fmt.Println(" ")
	fmt.Print("If your age is ", age)

	if age > 17 {
		fmt.Print(", it is alright if you any movie ", myname, ".")
		fmt.Println(" ")
	}
	if age < 17 && age > 12 {
		fmt.Print(", you should not watch R-rated movies. But you could watch anything up to PG-13 ", myname, ".")
		fmt.Println(" ")
	}
	if age < 13 {
		fmt.Print(", PG and G movies are most suited for you ", myname, ".")
		fmt.Println(" ")
	}
	fmt.Println(" ");
}