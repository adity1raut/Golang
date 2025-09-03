package looop

func main() {

	for i := 0; i < 10; i++ {
		println(i)
	}

	for j := 0; j < 10; j++ {
		if j == 5 {
			break
		}
		println(j)
	}

	for k := 0; k < 10; k++ {
		if k%2 == 0 {
			continue
		}
		println(k)
	}

	// Using a goto statement
	i := 0
LOOP:
	if i < 10 {
		println(i)
		i++
		goto LOOP
	}
}
