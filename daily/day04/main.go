package main

func main() {

	//Nested loop
	// for i := 1; i <= 4; i++ {
	// 	for j := 1; j <= 4; j++ {
	// 		fmt.Print(i+j, " ")
	// 	}
	// }

	//Pattern building
	// for i := 1; i <= 5; i++ {
	// 	for j := 1; j <= i; j++ {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println()
	// }

	//String are immutable in golang so if you want to modify the string first convert into []byte and do operation then back to string.
	// var name string
	// name = "test"
	// name += "ing"

	// name_new := []byte(name)
	// name_new[0] = 'a'
	// res := string(name_new)
	// fmt.Println(res)

	// name := "test"
	//Normal loop to access string
	// for i := 0; i < len(name); i++ {
	// 	fmt.Print(string(name[i]))
	// }

	//range loop to access string
	// for _, char := range name {
	// 	fmt.Print(string(char))
	// }

	

}
