package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"vizsla/utils"
)

func main() {

	// if len(os.Args) < 2 {
	// 	return
	// }

	// for _, arg := range os.Args[1:] {
	// 	if t, err := utils.Get(arg); err == nil {
	// 		fmt.Printf("%s:\t%s\n", arg, t)
	// 	} else {
	// 		fmt.Printf("%s: Unknown\n", arg)
	// 	}
	// }

	filepath.Walk("tests", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatalf(err.Error())
		}
		//file, _ := os.OpenFile(path, os.O_RDONLY, 0666)
		//defer file.Close()
		//var contentByte = make([]byte, 2048)
		//numByte, _ := file.Read(contentByte)
		//contentByte = contentByte[:numByte]

		//contentByte, err := ioutil.ReadFile(path)

		//if strings.HasSuffix(path, ".vdi") || strings.HasSuffix(path, ".doc") || strings.HasSuffix(path, ".ppt") {
		if t, err := utils.Get(path); err == nil {

			fmt.Printf("File Name: %s\t\t%s\n", info.Name(), t)

		}
		//}
		// 	// check := []byte{
		// 	// 	0x4d,
		// 	// 	0x45,
		// 	// 	0x54,
		// 	// 	0x41,
		// 	// 	0x2d,
		// 	// 	0x49,
		// 	// 	0x4e,
		// 	// 	0x46,
		// 	// 	0x2f}
		// 	// if compare(contentByte[30:39], check) {
		// 	// 	fmt.Println(path)
		// 	// }

		return nil
	})
}

// func compare(a []byte, b []byte) bool {
// 	if len(a) < len(b) {
// 		return false
// 	}
// 	for i, x := range b {
// 		if a[i] != x {
// 			return false
// 		}
// 	}
// 	return true
// }
