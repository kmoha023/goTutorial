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
	fmt.Println("\nFiles and Folders")
	//To Open a file use Open function from os package
	file, err := os.Open("text.txt")
	if err != nil {
		log.Fatal("Error opening file", err)
		return
	}
	defer file.Close()

	//get the file size
	stat, error1 := getFileInfo(file)
	if error1 != nil {
		errorHandler(error1)
	} else {
		fmt.Printf("File %s size is %d ", stat.Name(), stat.Size())
	}

	//read file method 1
	msg1, err1 := ReadFromFileMethod1(file, stat)
	if err1 != nil {
		errorHandler(err1)
	} else {
		PrintToConsole(msg1)
	}

	//read file method 2
	msg2, err2 := ReadFromFileMethod2("text.txt")
	if err2 != nil {
		errorHandler(err2)
	} else {
		PrintToConsole(msg2)
	}

	createdFile, err3 := CreateFile("text2.txt", "Hello world! from created file")
	if err3 != nil {
		errorHandler(err3)
	} else {
		createdFilestat, error1 := getFileInfo(createdFile)
		if error1 != nil {
			errorHandler(error1)
		} else {
			fmt.Printf("Created file %s size is %d ", createdFilestat.Name(), createdFilestat.Size())
		}
	}
	msg4, err4 := ReadFromFileMethod2("text2.txt")
	if err4 != nil {
		errorHandler(err4)
	} else {
		PrintToConsole(msg4)
	}

}

func errorHandler(err error) {
	log.Fatal(err)
}

//PrintToConsole ...
func PrintToConsole(msg string) {
	fmt.Println(msg)
}

func getFileInfo(file *os.File) (os.FileInfo, error) {
	fmt.Println("Getting fileInfo....")
	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}
	return stat, nil
}

//ReadFromFileMethod1 ...
func ReadFromFileMethod1(file *os.File, stat os.FileInfo) (string, error) {
	//Read from file
	fmt.Println("\nReading Method I....")
	bs := make([]byte, stat.Size())
	_, err := file.Read(bs)
	if err != nil {
		return "", err
	}
	return string(bs), nil
}

//ReadFromFileMethod2 ...
func ReadFromFileMethod2(file string) (string, error) {
	fmt.Println("Reading Method II....")
	bs2, err := ioutil.ReadFile(file)
	if err != nil {
		return "", err
	}
	return string(bs2), nil
}

//CreateFile ...
func CreateFile(name string, text string) (*os.File, error) {
	//createFile
	fmt.Println("Creating File... ", name)
	file, err := os.Create(name)
	if err != nil {
		return file, err
	}
	file.WriteString(text)
	return file, nil
}
