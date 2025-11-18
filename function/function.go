package main

// func add(a, b int) int {
// return a + b
// }

// func getLanguages() (string, string, string) {
// return "golang", "javascript", "python"
// }

func processit(fn func(int) int) {
	fn(1)
}
func main() {
	// res := add(3, 5)
	// fmt.Println(res)
	// lang1, lang2, lang3 := getLanguages()
	//
	// fmt.Println(lang1)
	// fmt.Println(lang2)
	// fmt.Println(lang3)

	fn := func(a int) int {
		return 2
	}
	processit(fn)
}
