package main

import (
	"fmt"
	"os"
)

func main() {

	// File Info:

	// f, err := os.Open("example.txt")

	// if err != nil {
	// 	// if error found
	// 	panic(err)
	// }

	// fileInfo, err := f.Stat()

	// if err != nil {
	// 	// if error found
	// 	panic(err)
	// }

	// fmt.Println("file name:", fileInfo.Name())
	// fmt.Println("file or folder:", fileInfo.IsDir())
	// fmt.Println("file size:", fileInfo.Size())
	// fmt.Println("file permission:", fileInfo.Mode())
	// fmt.Println("file modified at:", fileInfo.ModTime())

	// --

	// File Read:

	// Normally we read the data of the file and store that in buffer(temporary space)
	// buffer: array of space

	// f, err := os.Open("example.txt")

	// if err != nil {
	// 	panic(err)
	// }

	// defer f.Close() // run defer as soon as possible, cz, it declared in last, then might possible that, it doesn't reach it's end

	// buf := make([]byte, 12)

	// d, err := f.Read(buf)

	// if err != nil {
	// 	panic(err)
	// }

	// for i := 0; i < len(buf); i++ {
	// 	println("data", d, string(buf[i]))
	// }

	// --

	// Easy way: but can not used for every time, cz Readfile loads all the data at once, if any huge data like:videos, movies are loaded, then gbs of data has to load to the memory, then might possible that resources of that application is not enough to handle remaning data

	// So, if size is small, then can easily use them, but when file size is huge, use based on the requirements

	// data, err := os.ReadFile("example.txt")

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(string(data))

	// --

	// read folder::

	// dir, err := os.Open(".") // . refers current directory

	// if err != nil {
	// 	panic(err)
	// }

	// defer dir.Close() //there's details in that file, so at first we have to close that file

	// fileInfo, err := dir.ReadDir(-1) // can mention, which number of file n>0 like:1,2... that we require, if need all files, then use n<=0

	// // as ReadDir returns slice, so get that data, we have to loop
	// for _, fi := range fileInfo {
	// 	fmt.Println(fi.Name(), fi.IsDir(), fi.Type().IsRegular())
	// }

	// --

	// create a file:

	// f, err := os.Create("example2.txt")

	// if err != nil {
	// 	panic(err)
	// }

	// defer f.Close()

	// // add string:
	// f.WriteString("Golang is a good language, added to example2")

	// // add bytes:
	// bytes := []byte("Hello Golang, added to example2 ")

	// f.Write(bytes)

	// --

	// read and write to another file(streaming fashion)

	// sourceFile, err := os.Open("example.txt")

	// if err != nil {
	// 	panic(err)
	// }

	// defer sourceFile.Close()

	// desFile, err := os.Create("example3.txt")

	// if err != nil {
	// 	panic(err)
	// }

	// defer desFile.Close()

	// reader := bufio.NewReader(sourceFile)
	// writer := bufio.NewWriter(desFile)

	// for {
	// s, err:=reader.ReadString("\n") // read line by line
	// 	b, err := reader.ReadByte()	//read byte by byte

	// 	if err != nil {
	// 		if err.Error() != "EOF" {
	// 			panic(err)
	// 		}

	// 		break
	// 	}

	//  e := writer.WriteString(s)	// write each line
	// 	e := writer.WriteByte(b)	// write each byte

	// 	if e != nil {
	// 		panic(e)
	// 	}

	// }

	// writer.Flush()		// Flush the buffer to ensure all data is written to disk

	// fmt.Println("written to new file successfully")

	// If we want just copy the data, then use like this:

	// copiedData, err := io.Copy(desFile, sourceFile)

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("Copied data successfully\n", copiedData)

	// --

	// Delete a file:

	err := os.Remove("example2.txt")

	if err != nil {
		panic(err)
	}

	fmt.Println("File deleted successfully")

}
