package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"path/filepath"

	"github.com/isfonzar/filecrypt"
)

var badFiles, goodFiles = 0, 0

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		printMessage()
		result, _, _ := reader.ReadLine()
		resultStr := string(result)
		switch resultStr {
		case "1":
			encryptHandle()
		case "2":
			decryptHandle()
		case "3":
			fmt.Println("Bye!")
			os.Exit(0)
		default:
			fmt.Println("Couldn't find command!")
		}
	}
}
func printMessage() {
	fmt.Print("1. Encrypt" +
		"\n2. Decrypt" +
		"\n3. Exit" +
		"\nYour choice: ")
}
func encryptHandle() {
	var files []string
	fmt.Print("Name of the file/folder: ")
	file, _, _ := bufio.NewReader(os.Stdin).ReadLine()
	fileStr := string(file)
	if !validateFile(fileStr) {
		panic("File not found")
	}

	fi, _ := os.Stat(fileStr)
	if fi.IsDir() {
		alert := filepath.Walk(fileStr, func(path string, info os.FileInfo, err error) error {
			files = append(files, path)
			return nil
		})
		if alert != nil {
			panic(alert)
		}
	}

	password := getPassword()
	if len(files) > 1 {
		fmt.Println("\nEncrypting...")
		for i := 1; i < len(files); i++ {
			location, _ := os.Stat(files[i])
			if !location.IsDir() {
				if state := encryptPanicHandler(files[i], password); state != true {
					badFiles++
				} else {
					goodFiles++
				}
			}
		}
		fmt.Println("\nEncrypted:", goodFiles, "\nFailed:", badFiles)
		waitForKeyPress()
	} else {
		fmt.Println("\nEncrypting...")
		if state := encryptPanicHandler(fileStr, password); state != false {
			fmt.Println("File successfully encrypted")
		}

		waitForKeyPress()
	}
}

func getPassword() []byte {
	fmt.Print("Enter password: ")
	password1Bytes := getInputWinToByte()
	fmt.Print("\nConfirm password: ")
	password2Bytes := getInputWinToByte()
	if !validatePassword(password1Bytes, password2Bytes) {
		fmt.Print("\nPasswords do not match. Please try again.\n")
		return getPassword()
	}
	return password1Bytes
}

func encryptPanicHandler(file string, password []byte) (state bool) {
	defer func() {
		if rec := recover(); rec != nil {
			fmt.Println("\nEncryption operation failed at file:" + file + "\n")
			state = false
		}
	}()
	filecrypt.Encrypt(file, password)
	state = true
	return state
}

func decryptHandle() {

	var files []string
	fmt.Print("Name of the file: ")
	file, _, _ := bufio.NewReader(os.Stdin).ReadLine()
	fileStr := string(file)
	if !validateFile(fileStr) {
		panic("File not found")
	}
	fi, _ := os.Stat(fileStr)
	if fi.IsDir() {
		alert := filepath.Walk(fileStr, func(path string, info os.FileInfo, err error) error {
			files = append(files, path)
			return nil
		})
		if alert != nil {
			panic(alert)
		}
	}
	fmt.Print("Enter password: ")
	password1Bytes := getInputWinToByte()
	if len(files) > 1 {
		fmt.Println("\nDecrypting...")
		for i := 1; i < len(files); i++ {
			location, _ := os.Stat(files[i])
			if !location.IsDir() {
				if state := decryptPanicHandler(files[i], password1Bytes); state != true {
					badFiles++
				} else {
					goodFiles++
				}
			}
		}
		fmt.Println("\nDecrypted:", goodFiles, "\nFailed:", badFiles)
		waitForKeyPress()
	} else {
		fmt.Println("\nDecrypting...")
		if state := decryptPanicHandler(fileStr, password1Bytes); state != false {
			fmt.Println("File successfully decrypted")
			waitForKeyPress()
		}
	}
}

func decryptPanicHandler(file string, password []byte) (status bool) {
	defer func() {
		if rec := recover(); rec != nil {
			fmt.Println("\nDecryption operation failed at file:" + file + "\n")
			status = false
		}
	}()
	filecrypt.Decrypt(file, password)
	status = true
	return status
}

func validatePassword(password1 []byte, password2 []byte) bool {
	if !bytes.Equal(password1, password2) {
		return false
	}
	return true
}

func validateFile(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}
	return true
}

func getInputWinToByte() []byte {
	input, _, _ := bufio.NewReader(os.Stdin).ReadLine()
	inputBytes := []byte(string(input))
	return inputBytes
}

func waitForKeyPress() {
	_, _, _ = bufio.NewReader(os.Stdin).ReadLine()
}