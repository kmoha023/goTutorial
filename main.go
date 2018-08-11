package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	StringsInBinary()
	IOpackage()
	FilesAndFolder()

}

// StringsInBinary ...
func StringsInBinary() {
	fmt.Println("\nStrings as binary")
	//Sometimes we need to work with strings as binary data. To convert a string to a slice of bytes (and vice-versa) do this:
	//[84 101 115 116] ASCII Alphabets Decimal value
	arr := []byte("Test")
	fmt.Println(arr)
	str := []byte{'T', 'e', 's', 't'}
	fmt.Println(str)
	string := string(arr)
	fmt.Println(string)
	fmt.Println("\n--------------------------")

}

// IOpackage ...
func IOpackage() {
	//-------Buffer struct of bytes to read or write
	fmt.Println("\nBuffer struct of Bytes")
	var buf bytes.Buffer
	buf.Write([]byte("Test"))
	// http.Post("http://example.com/", "text/plain", &buf)

	//if only wants to read from a string (efficient)
	str := strings.NewReader("Test from another server..")
	http.Post("http://localhost:9090/", "text/plain", str)
	b := make([]byte, 8)
	for {
		n, err := str.Read(b)
		fmt.Println("\n readed string ", n)
		if err == io.EOF {
			break
		}
	}
	//To convert buffer to []byte
	buf.Bytes()
	fmt.Println("\n--------------------------")
}

// FilesAndFolder ...
func FilesAndFolder() {
	// Files and Folders
	//-------Buffer struct of bytes to read or write
	fmt.Println("\n Method I ----- Files and Folders")
	//To Open a file use Open function from os package
	file, err := os.Open("text.txt")
	if err != nil {
		log.Fatal("Error opening file", err)
		return
	}
	defer file.Close()

	//get the file size
	stat, err := file.Stat()
	if err != nil {
		log.Fatal("Error getting stats..", err)
		return
	}

	//Read from file
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		log.Fatal("Error reading file...", err)
		return
	}

	str := string(bs)
	fmt.Println(str)
	fmt.Println("\n Method II ----- Files and Folders")

	bs2, err := ioutil.ReadFile("text.txt")
	if err != nil {
		log.Fatal("Error in reading in method 2", err)
	}
	str2 := string(bs2)
	fmt.Println(str2)

}
