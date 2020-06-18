package main

import (
	"context"
	"flag"
	"fmt"   //implements formatted I/O
	"log"
	"os"    //provides a platform-independent interface to os functionality
	"time"
	"unicode"
	
	"google.golang.org/grpc"
)


type Parsed struct {
	shift int32
	text  string
	encode bool
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

func validate(text string, shift int, mode string) (*Parsed, bool) {
	var shifti32 int32
	p := &Parsed{}

	shifti32 = int32(shift)
	p.shift = shifti32
	p.text = text

	if len(p.text) < 1 {
		fmt.Println("invalid text: must be at least one character")
		return nil, false
	}

	if !isAlpha(p.text) {
		fmt.Println("invalid text: must be alphabetical characters and spaces only")
		return nil, false
	}

	if mode == "e" {
		p.encode = true
	} else if mode == "d" {
		p.encode = false
	} else {
		fmt.Printf("mode flag should be e or d, not %s\n", mode)
		return nil, false
	}

	return p, true
}

func main() {
	var text string
	var shift int
	var mode string
	var ok bool
	var p *Parsed 

	//Set up connection to server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := NewCipherServiceClient(conn) //provides the client stub

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	flag.StringVar(&text, "t", "hello world", "t (text) should be a string (enclosed by \" \") of only alphabetical characters and spaces")
	flag.IntVar(&shift, "s", 0, "s (shift) should be an integer")
	flag.StringVar(&mode, "m", "e", "m (mode) should be a character e or d: e to encode text, d to decode text")
	flag.Parse()

	p, ok = validate(text, shift, mode)
	if !ok {
		os.Exit(0)
	} 
	
	res, err := client.ProcessMessage(ctx, &Input{Text: p.text, Shift: p.shift, Encode: p.encode})
	if err != nil {
		log.Fatalf("Could not encode: %v", err)
	}
	
	fmt.Println(res.GetMsg())
}
