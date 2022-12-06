package main

func main() {
	var i interface{} = 1
	n := i.(int)
	print(n)
}
