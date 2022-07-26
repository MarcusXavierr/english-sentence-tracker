package main

func main() {
	err := WriteFile("Testando\n", "/tmp/test")
	check(err)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
