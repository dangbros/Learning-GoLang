package main

import (
	"fmt"
	"os"
)

func main() {
	// f, err := os.Open("file_handling/example.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// // fileInfo, err := f.Stat()
	// // if err != nil {
	// // 	panic(err)
	// // }
	// // fmt.Println("file name:", fileInfo.Name())
	// // fmt.Println("file size:", fileInfo.Size())
	// // fmt.Println("folder or not:", fileInfo.IsDir())

	// defer f.Close()

	// buf := make([]byte, 12)

	// d, err := f.Read(buf)
	// if err != nil {
	// 	panic(err)
	// }

	// for i := 0; i < len(buf); i++ {
	// 	println("data", d, string(buf[i]))
	// }
	//

	// data, err := os.ReadFile("file_handling/example.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(string(data))

	// dir, err := os.Open(".")
	// if err != nil {
	// 	panic(err)
	// }

	// defer dir.Close()

	// fileInfo, err := dir.ReadDir(-1)

	// for _, fi := range fileInfo {
	// 	fmt.Println(fi.Name(), fi.IsDir())
	// }

	// create a file

	// f, err := os.Create("file_handling/example2.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer f.Close()

	// f.WriteString("Hi, Go!")
	// f.WriteString("Bye, Go!")

	// bytes := []byte("Hello Golang!")

	// f.Write(bytes)

	//read and write to another file(streaming fashion)

	// sourceFile, err := os.Open("file_handling/example.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer sourceFile.Close()

	// destFile, err := os.Create("file_handling/example2.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer destFile.Close()

	// reader := bufio.NewReader(sourceFile)
	// writer := bufio.NewWriter(destFile)

	// for {
	// 	b, err := reader.ReadByte()
	// 	if err != nil {
	// 		if err.Error() != "EOF" {
	// 			panic(err)
	// 		}
	// 		break
	// 	}
	// 	e := writer.WriteByte(b)
	// 	if e != nil {
	// 		panic(e)
	// 	}
	// }

	// writer.Flush()
	// fmt.Println("written to new file successfully! ")

	//delte a file

	err := os.Remove("file_handling/example2.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("file deleted successfully")
}
