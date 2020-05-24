package main

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"flag"
	"fmt"
	"github.com/twmb/murmur3"
	"hash"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

var (
	h            bool
	v            bool
	Version      string
	VERSION_TAG  string
	Compile      string
	Branch       string
	GitDirty     string
	HashUrl      string
	Hashfile     string
	ImageBase64  string
	UserAgent    string
	IsUint32     bool
	FofaFormat   bool
	ShodanFormat bool
	DefaultUA    string = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_0) AppleWebKit/535.11 (KHTML, like Gecko) Chrome/17.0.963.56 Safari/535.11"
)

func PrintVersion() {
	fmt.Printf("Tag: %s\n", VERSION_TAG)
	fmt.Printf("Version: %s\n", Version)
	fmt.Printf("Compile: %s\n", Compile)
	fmt.Printf("Branch: %s\n", Branch)
	fmt.Printf("GitDirty: %s\n", GitDirty)
}

func init() {
	flag.BoolVar(&h, "h", false, "look help \n iconhash  favicon.ico \n iconhash  https://www.baidu.com/favicon.ico")
	flag.BoolVar(&v, "v", false, "version")
	flag.BoolVar(&FofaFormat, "fofa", true, "fofa search format")
	flag.BoolVar(&ShodanFormat, "shodan", false, "shodan search format \n iconhash   -file test/favicon.ico -shodan -fofa=false")
	flag.BoolVar(&IsUint32, "uint32", false, "uint32")
	flag.StringVar(&Hashfile, "file", "", "mmh3 hash from file \n iconhash -file favicon.ico")
	flag.StringVar(&HashUrl, "url", "", "mmh3 hash from url \n iconhash -url  https://www.baidu.com/favicon.ico")
	flag.StringVar(&UserAgent, "user-agent", DefaultUA, "mmh3 hash from url")
	flag.StringVar(&ImageBase64, "b64", "", "mmh3 hash image base64 from file \n iconhash   -file test/favicon.ico ")
	IconHashArgs := map[string]int{
		"-h":          1,
		"-v":          1,
		"-fofa":       1,
		"-shodan":     1,
		"-uint32":     1,
		"-file":       1,
		"-url":        1,
		"-user-agent": 1,
		"-b64":        1,
	}
	flag.Parse()

	if v {
		PrintVersion()
		return
	}

	if h || len(os.Args) == 1 {
		flag.Usage()
		return
	}

	if len(os.Args) == 2 {
		arg := os.Args[1]
		if _, ok := IconHashArgs[arg]; ok == false {
			FofaFormat = false
			if strings.HasPrefix(arg, "http://") || strings.HasPrefix(arg, "https://") {
				HashUrl = arg
			} else if _, err := os.Stat(arg); err == nil {
				Hashfile = arg

			} else {
				fmt.Print("not file or url please check\n")
				os.Exit(1)
			}

		} else {
			flag.Usage()
			os.Exit(0)
		}

	}
	if len(os.Args) > 2 && !strings.HasPrefix(os.Args[1], "-") {
		flag.Usage()
		os.Exit(1)
	}

}
func PrintResult(result string) {
	if !ShodanFormat && !FofaFormat {
		fmt.Printf("%s\n", result)
	}
	if FofaFormat {
		fmt.Printf("icon_hash=\"%s\"\n", result)
	}

	if ShodanFormat {
		fmt.Printf("http.favicon.hash:%s\n", result)
	}

}

func FromUrlGetContent(requrl string) (content []byte, err error) {
	client := &http.Client{
		Timeout: time.Second * time.Duration(10),
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, //param
		},
	}
	req, err := http.NewRequest("GET", requrl, nil)
	req.Header.Set("User-Agent", UserAgent)
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func FromfileGetContent(path string) (content []byte, err error) {
	fi, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer fi.Close()
	content, err = ioutil.ReadAll(fi)
	// fmt.Println(string(fd))
	if err != nil {
		return nil, err
	}
	return content, nil
}

func Mmh3Hash32(raw []byte) string {
	var h32 hash.Hash32 = murmur3.New32()
	h32.Write([]byte(raw))
	if IsUint32 {
		return fmt.Sprintf("%d", h32.Sum32())
	}
	return fmt.Sprintf("%d", int32(h32.Sum32()))
}

func StandBase64(braw []byte) []byte {
	bckd := base64.StdEncoding.EncodeToString(braw)
	var buffer bytes.Buffer
	for i := 0; i < len(bckd); i++ {
		ch := bckd[i]
		buffer.WriteByte(ch)
		if (i+1)%76 == 0 {
			buffer.WriteByte('\n')
		}
	}
	buffer.WriteByte('\n')
	return buffer.Bytes()

}

func SplitChar76(braw []byte) []byte {
	// 去掉 data:image/vnd.microsoft.icon;base64
	if strings.HasPrefix(string(braw), "data:image/vnd.microsoft.icon;base64,") {
		braw = braw[37:]
	}

	var buffer bytes.Buffer
	for i := 0; i < len(braw); i++ {
		ch := braw[i]
		buffer.WriteByte(ch)
		if (i+1)%76 == 0 {
			buffer.WriteByte('\n')
		}
	}
	buffer.WriteByte('\n')
	return buffer.Bytes()

}
