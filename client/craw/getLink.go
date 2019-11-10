package craw

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os/exec"
	"regexp"
	"strings"

	"github.com/trungtvq/go-utils/fileutils"
)

func GetLinkFromFile(path string) {
	files, err := fileutils.FileToSlice(path)
	if err != nil {
		return
	}
	for _, f := range files {
		if len(f) > 3 {
			GetLink(f)
		}
	}
}

//GetLink https://www.fshare.vn/file/3WZ2XAERFLDUVM7?token=1573397899
func GetLink(link string) string {
	fmt.Println("link", link)
	client := &http.Client{}
	re := regexp.MustCompile(`\?token=([0-9a-zA-Z]*)`)
	link = re.ReplaceAllLiteralString(link, "")
	re = regexp.MustCompile(`(.*)file\/`)
	link = re.ReplaceAllLiteralString(link, "")

	addr := "https://www.fshare.vn/download/get"
	b := `_csrf-app=qEHEoutVtM57_V81vLr_rzFlnCxiHKuM_GZD1DSGvFjjK77x3mGEvTmXNWHq38eCHCOtAQpV6O2FEyidV_b2Lw%3D%3D&linkcode=` + link + `&withFcode5=0&fcode=`
	_, err := http.NewRequest("POST", addr+"?"+b, nil)
	if err != nil {
		fmt.Println(err)
	}

	apiUrl := "https://www.fshare.vn"
	resource := "/download/get"

	u, _ := url.ParseRequestURI(apiUrl)
	u.Path = resource
	urlStr := u.String()                                          // "https://api.com/user/"
	r, _ := http.NewRequest("POST", urlStr, strings.NewReader(b)) // URL-encoded payload

	r.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	r.Header.Add("Cookie", "_uidcms=157304601194335132; _ga=GA1.2.282030455.1573046012; __yoid__=5b7816b58084471185bb9353509ea09f; _gid=GA1.2.1034586615.1573394962; fshare-app=g2ban7tncmnm6tcrn8876oe1qr; _identity-app=5ce037c846a63c7dc49c2687468c282c628f5afbfb77500ae58420e60a6c913fa%3A2%3A%7Bi%3A0%3Bs%3A13%3A%22_identity-app%22%3Bi%3A1%3Bs%3A56%3A%22%5B12179678%2C%22K2tGaIuQPlWw3r84H7LOAS_kPBVeLi48%22%2C1573481388%5D%22%3B%7D; _gat_gtag_UA_97071061_1=1")
	//r.Header.Add("origin", "https://www.fshare.vn")

	resp, err := client.Do(r)
	if err != nil {
		fmt.Println(err)
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	type data struct {
		Url  string
		Name string
	}
	var d data
	json.Unmarshal(bodyBytes, &d)
	go func() {
		cmd := exec.Command(`C:\Program Files (x86)\Internet Download Manager\IDMan.exe`, `/d`, d.Url)
		log.Printf("Running command and waiting for it to finish...")
		cmd.Run()
	}()

	return d.Url
}

//Login ... not working yet
func Login(username string, pass string) string {
	client := &http.Client{}
	username = url.QueryEscape(username)
	pass = url.QueryEscape(pass)
	b := `_csrf-app=EFptTiBktktuAPibtLHBv8Y4MJ-TrCXn43ANqbcm0cgpaCwZYVX5DDxRkKLt8vHKtWoC19jrfK2vFXfo5kKGoA%3D%3D&LoginForm%5Bemail%5D=` + username + `&LoginForm%5Bpassword%5D=` + pass + `&LoginForm%5BrememberMe%5D=0`

	apiUrl := "https://www.fshare.vn"
	resource := "/site/login"

	u, _ := url.ParseRequestURI(apiUrl)
	u.Path = resource
	urlStr := u.String()                                          // "https://api.com/user/"
	r, _ := http.NewRequest("POST", urlStr, strings.NewReader(b)) // URL-encoded payload

	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Cookie", "apluuid=8266fd21-5357-4cfe-1eab-c09788e889bb; _a3rd1537438164=0-7; _a3rd1537437937=0-1; _uidcms=1573398512193215652; _ga=GA1.2.1619653431.1573398513; _gid=GA1.2.532598729.1573398513; __yoid__=a7fde9a28d984d218d5532ca67507c69; fshare-app=h9kmuf5m9frruufvi3sjnn27va; _gat_gtag_UA_97071061_1=1")
	//r.Header.Add("origin", "https://www.fshare.vn")

	resp, err := client.Do(r)
	if err != nil {
		fmt.Println(err)
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)

	return string(bodyBytes)
}
