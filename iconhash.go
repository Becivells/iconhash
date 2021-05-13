package main

import (
	"fmt"
	"os"
)

// PrintErr 打印错误并输出标准错误
func PrintErr(err error) {
	if Debug {
		panic(err)
	}
	fmt.Printf("%s\n", err)
	os.Exit(1)
}

func main() {
	if HashURL != "" {
		content, err := FromURLGetContent(HashURL)

		if err != nil {
			PrintErr(err)
		}

		PrintResult(Mmh3Hash32(StandBase64(content)))
		os.Exit(0)
	}

	if Hashfile != "" {
		content, err := FromfileGetContent(Hashfile)

		if err != nil {
			PrintErr(err)
		}

		PrintResult(Mmh3Hash32(StandBase64(content)))
		os.Exit(0)
	}

	if ImageBase64 != "" {
		content, err := FromfileGetContent(ImageBase64)

		if err != nil {
			PrintErr(err)
		}

		PrintResult(Mmh3Hash32(SplitChar76(content)))
		os.Exit(0)
	}
}
