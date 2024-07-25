package syntaxBase

import (
	"fmt"
	"io"
	"log"
	"os"
)

func ReadFile() {

	file, err := os.Open("./option_test.go")
	if err != nil {
		log.Fatalf("open file error: %v", err)
		return
	}
	defer file.Close()
	var content []byte
	for {
		buf := make([]byte, 1024)
		n, err := file.Read(buf)
		if err == io.EOF {
			log.Printf("read file error: %v", err)
			break
		}
		content = append(content, buf[:n]...)
	}

	fmt.Printf("read file content:%s", string(content))

}
