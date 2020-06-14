package main

import (
	"bufio" //implements buffered I/O by wrapping an io reader or writer object
	"context"
	"fmt"   //implements formatted I/O
	"log"
	"os"    //provides a platform-independent interface to os functionality
	"strconv"
	"strings"
	"time"
	"unicode"
	
	"google.golang.org/grpc"
)


type Parsed struct {
	shift int32
	text  string
}

func isAlpha(s string) bool {
	isValid := true
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsSpace(r) {
			continue
		} else {
			isValid = false
		}
	}
	return isValid
}

func validate(input string) (*Parsed, bool) {
	var lasti32 int32
	parts := strings.Split(input, " ")
	p := &Parsed{}

	last := parts[len(parts)-1]
	lasti, err := strconv.Atoi(last)

	if err != nil {
		fmt.Println("Expected input: text shift, invalid shift: must be integer")
		fmt.Printf("Please try again: ")
		return nil, false
	}

	lasti32 = int32(lasti)
	p.shift = lasti32
	p.text = strings.Join(parts[:len(parts)-1], " ")

	if !isAlpha(p.text) {
		fmt.Println("Expected input: text shift, invalid text: must be alphabetical characters only")
		fmt.Printf("Please try again: ")
		return nil, false
	}

	return p, true
}

func main() {
	var parsed *Parsed
	var line string
	var ok bool

	//Set up connection to server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := NewCipherServiceClient(conn) //this provides the client stub

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	scanner := bufio.NewScanner(os.Stdin) //returns a new scanner to read from os.standard input
	fmt.Println("Enter the text you'd like to encode (using only alphabetical characters) followed by a number. Type exit to quit program.")
	for scanner.Scan() {
		line = scanner.Text()
		if line == "exit" {
			os.Exit(0)
		}

		parsed, ok = validate(line)
		if !ok {
			//keep scanning, input not valid
			continue
		} else {
			break
		}
	}
	if err := scanner.Err(); err != nil { //returns first non-EOF error encountered by scanner
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	
	res, err := client.EncodeMessage(ctx, &Input{Text: parsed.text, Shift: parsed.shift, Encode: false}) //req protobuffer object
	if err != nil {
		log.Fatalf("Could not encode: %v", err)
	}
	
	fmt.Println(res.GetMsg())
}
