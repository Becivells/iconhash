package main

import "os"

func main() {

	if len(HashUrl) != 0 {
		content, err := FromUrlGetContent(HashUrl)
		if err != nil {
			panic(err)
		}
		PrintResult(Mmh3Hash32(StandBase64(content)))
		os.Exit(0)
	}

	if len(Hashfile) != 0 {
		content, err := FromfileGetContent(Hashfile)
		if err != nil {
			panic(err)
		}
		PrintResult(Mmh3Hash32(StandBase64(content)))
		os.Exit(0)
	}

	if len(ImageBase64) != 0 {
		content, err := FromfileGetContent(ImageBase64)
		if err != nil {
			panic(err)
		}
		PrintResult(Mmh3Hash32(SplitChar76(content)))
		os.Exit(0)
	}
}
