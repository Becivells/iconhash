package main

import (
	"fmt"
	"os"
)

func PrintErr(err error) {
	if Debug {
		panic(err)
	}
	fmt.Printf("%s\n", err)
	os.Exit(1)
}
func main() {

	if len(HashUrl) != 0 {
		content, err := FromUrlGetContent(HashUrl)

		if err != nil {
			PrintErr(err)
		}

		PrintResult(Mmh3Hash32(StandBase64(content)))
		os.Exit(0)
	}

	if len(Hashfile) != 0 {
		content, err := FromfileGetContent(Hashfile)

		if err != nil {
			PrintErr(err)
		}

		PrintResult(Mmh3Hash32(StandBase64(content)))
		os.Exit(0)
	}

	if len(ImageBase64) != 0 {
		content, err := FromfileGetContent(ImageBase64)

		if err != nil {
			PrintErr(err)
		}

		PrintResult(Mmh3Hash32(SplitChar76(content)))
		os.Exit(0)
	}
}
