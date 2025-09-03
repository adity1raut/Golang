

package main

func main(){

	// Using a goto statement
	i := 0
  LOOP:
	if i < 10 {
		println(i)
		i++
		goto LOOP
	}
}
