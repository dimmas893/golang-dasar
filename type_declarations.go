package main

import "fmt"

func main() {
	type noKTP string
	type aliasboolean bool
	var noKTPananda noKTP = "1234132312312"
	var valueBool aliasboolean = true
	fmt.Println(noKTPananda)
	fmt.Println(valueBool)
}
