package main

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"flag"
	"fmt"
	"hash"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/twmb/murmur3"
)

//nolint
var (
	h                  bool
	v                  bool
	version            string
	commit             string
	date               string
	Branch             string
	GitDirty           string
	HashURL            string
	Hashfile           string
	ImageBase64        string
	UserAgent          string
	IsUint32           bool
	FofaFormat         bool
	ShodanFormat       bool
	InsecureSkipVerify bool
	ReqTimeOut         int = 10
	Debug              bool
	DefaultUA          string = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_0) AppleWebKit/535.11 (KHTML, like Gecko) Chrome/17.0.963.56 Safari/535.11"
)

// PrintVersion 打印版本信息
func PrintVersion() {
	fmt.Printf("Version: %s\n", version)
	fmt.Printf("Compile: %s\n", date)
	fmt.Printf("Commit: %s\n", commit)
	fmt.Printf("Branch: %s\n", Branch)
	fmt.Printf("GitDirty: %s\n", GitDirty)
}

//nolint:gochecknoinits // this is init
func init() {
	flag.BoolVar(&h, "h", false, "look help \n iconhash  favicon.ico \n iconhash  https://www.baidu.com/favicon.ico")
	flag.BoolVar(&v, "v", false, "version")
	flag.BoolVar(&Debug, "debug", false, "debug mode")
	flag.BoolVar(&FofaFormat, "fofa", true, "fofa search format")
	flag.BoolVar(&ShodanFormat, "shodan", false, "shodan search format \n iconhash   -file test/favicon.ico -shodan -fofa=false")
	flag.BoolVar(&IsUint32, "uint32", false, "uint32")
	flag.BoolVar(&InsecureSkipVerify, "skip-verify", true, "https InsecureSkipVerify")
	flag.StringVar(&Hashfile, "file", "", "mmh3 hash from file \n iconhash -file favicon.ico")
	flag.StringVar(&HashURL, "url", "", "mmh3 hash from url \n iconhash -url  https://www.baidu.com/favicon.ico")
	flag.StringVar(&UserAgent, "user-agent", DefaultUA, "mmh3 hash from url")
	flag.StringVar(&ImageBase64, "b64", "", "mmh3 hash image base64 from file \n iconhash   -file test/favicon.ico ")
	IconHashArgs := map[string]int{
		"-h":           1,
		"-v":           1,
		"-fofa":        1,
		"-shodan":      1,
		"-uint32":      1,
		"-file":        1,
		"-url":         1,
		"-user-agent":  1,
		"-b64":         1,
		"-debug":       1,
		"-skip-verify": 1,
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
	// nolint
	if len(os.Args) == 2 {
		arg := os.Args[1]
		if _, ok := IconHashArgs[arg]; !ok {
			FofaFormat = false
			if strings.HasPrefix(arg, "http://") || strings.HasPrefix(arg, "https://") {
				HashURL = arg
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

	if Debug {
		fmt.Print("---------------------------     var    value     --------------------------------\n")
		fmt.Printf("h                  :%t\n", h)
		fmt.Printf("v                  :%t\n", v)
		fmt.Printf("Version            :%s\n", version)
		fmt.Printf("Commit        :%s\n", commit)
		fmt.Printf("Compile            :%s\n", date)
		fmt.Printf("HashURL            :%s\n", HashURL)
		fmt.Printf("Hashfile           :%s\n", Hashfile)
		fmt.Printf("ImageBase64        :%s\n", ImageBase64)
		fmt.Printf("UserAgent          :%s\n", UserAgent)
		fmt.Printf("IsUint32           :%t\n", IsUint32)
		fmt.Printf("FofaFormat        :%t\n", FofaFormat)
		fmt.Printf("ShodanFormat       :%t\n", ShodanFormat)
		fmt.Printf("InsecureSkipVerify :%t\n", InsecureSkipVerify)
		fmt.Printf("Debug              :%t\n", Debug)
		fmt.Printf("DefaultUA          :%s\n", DefaultUA)
		defer fmt.Print("---------------------------     var    value     --------------------------------\n")
	}
}

// PrintResult 打印结果
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

// FromURLGetContent 从 URL 中获取图片内容
func FromURLGetContent(requrl string) (content []byte, err error) {
	if Debug {
		fmt.Print("---------------------------  start url  content  --------------------------------\n")
		fmt.Printf("====> url: %s\n", HashURL)
		defer fmt.Print("---------------------------  end url  content  --------------------------------\n")
	}

	client := &http.Client{
		Timeout: time.Second * time.Duration(ReqTimeOut),
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: InsecureSkipVerify}, // param
		},
	}

	req, err := http.NewRequest("GET", requrl, nil) //nolint
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", UserAgent)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close() //nolint

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if Debug {
		fmt.Printf("===> status code: %d\n", resp.StatusCode)
		fmt.Printf("====> content: \n%s\n", body)
	}

	return body, nil
}

// FromfileGetContent 从文件中获取图片内容
func FromfileGetContent(path string) (content []byte, err error) {
	if Debug {
		fmt.Print("---------------------------start From file get content--------------------------------\n")
		defer fmt.Print("---------------------------end  From file get content--------------------------------\n")
	}

	fi, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer fi.Close() // nolint
	content, err = ioutil.ReadAll(fi)
	if Debug {
		fmt.Printf("====> fileContent:\n %s\n", content)
	}

	if err != nil {
		return nil, err
	}
	return content, nil
}

/*
Mmh3Hash32 计算 mmh3 hash
*/
func Mmh3Hash32(raw []byte) string {
	var h32 hash.Hash32 = murmur3.New32()
	h32.Write(raw)
	if IsUint32 {
		return fmt.Sprintf("%d", h32.Sum32())
	}
	return fmt.Sprintf("%d", int32(h32.Sum32()))
}

// StandBase64 计算 base64 的值
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
	if Debug {
		fmt.Print("---------------------------start base64 content--------------------------------\n")
		fmt.Printf("====> base64:\n%s\n", buffer.String())
		defer fmt.Print("---------------------------end base64 content--------------------------------\n")
	}
	return buffer.Bytes()
}

// SplitChar76 按照 76 字符切分
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

	if Debug {
		fmt.Print("---------------------------start base64 content--------------------------------\n")
		fmt.Printf("====> base64 split 76:\n %s\n", buffer.String())
		defer fmt.Print("---------------------------end base64 content--------------------------------\n")
	}

	return buffer.Bytes()
}
