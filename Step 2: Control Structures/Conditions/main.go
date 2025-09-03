

package conditions

func main(){
	age := 20
	if age < 18 {
		println("Minor")
	} else if age >= 18 && age < 65 {
		println("Adult")
	} else {
		println("Senior")
	}

	// Using switch statement
	switch age {
	case 0, 1, 2:
		println("Infant")
	case 3, 4, 5:
		println("Toddler")
	case 6, 7, 8, 9, 10:
		println("Child")
	default:
		println("Teenager or older")
	}
}