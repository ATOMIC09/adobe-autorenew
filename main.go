package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"crypto/tls"
	"time"
	"math/rand"
)

func run() {
	username := os.Getenv("MY_USERNAME")
    password := os.Getenv("MY_PASSWORD")
	mainUrl := "https://software.kmutnb.ac.th/"
	loginUrl := "https://software.kmutnb.ac.th/login/"
	claimUrl := "https://software.kmutnb.ac.th/adobe-reserve/add2.php"
	// client := &http.Client{
	// CheckRedirect: func(req *http.Request, via []*http.Request) error {
	//     fmt.Printf("Redirect to %s\n", req.URL.String())
	//     return nil
	// },
	// }
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Transport: tr,
	}
	mainRes, err := client.Get(mainUrl)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	mainRes.Body.Close()
	mainCookies := mainRes.Cookies()
	// for _, cookie := range mainCookies {
	// 	fmt.Printf("%s: %s\n", cookie.Name, cookie.Value)
	// }
	loginPayload := map[string]string{
		"myusername": username,
		"mypassword": password,
		"Submit":     "",
	}
	loginReq, err := http.NewRequest("POST", loginUrl, bytes.NewBuffer(jsonToFormData(loginPayload)))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	loginReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	for _, cookie := range mainCookies {
		loginReq.AddCookie(cookie)
	}
	loginRes, err := client.Do(loginReq)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	writeToFile("login_response.html", loginRes.Body)
	loginRes.Body.Close()
	claimPayload := map[string]string{
		"userId":        "",
		"date_expire":   "2027-02-08",
		"status_number": "0",
		"Submit_get":    "",
	}
	claimReq, err := http.NewRequest("POST", claimUrl, bytes.NewBuffer(jsonToFormData(claimPayload)))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	claimReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	for _, cookie := range mainCookies {
		claimReq.AddCookie(cookie)
	}
	claimRes, err := client.Do(claimReq)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	writeToFile("claim_response.html", claimRes.Body)
	claimRes.Body.Close()
}

func main() {
    for {
        run()
        fmt.Println("Claimed om Monday at", time.Now().Format("2006-01-02 15:04:05"))
        // Calculate the duration until next Monday 00:00
        currentTime := time.Now()
        daysUntilMonday := int(time.Monday - currentTime.Weekday())
        if daysUntilMonday < 0 {
            daysUntilMonday += 7
        } else {
			// Fixed bug 1 where it schedule today if it's Monday, hence infinite loop for entire of that day
			daysUntilMonday = 7
		}
        targetTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day()+daysUntilMonday, 6 + rand.Intn(2), rand.Intn(60), 0, 0, time.Local)
        duration := targetTime.Sub(currentTime)
        timer := time.NewTimer(duration)
        <-timer.C
    }
}
func writeToFile(filename string, content io.ReadCloser) {
	var bodyBytes bytes.Buffer
	_, err := io.Copy(&bodyBytes, content)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()
	file.Write(bodyBytes.Bytes())
	// fmt.Printf("Response written to %s\n", filename)
}

func jsonToFormData(data map[string]string) []byte {
	var postData bytes.Buffer
	for key, value := range data {
		postData.WriteString(key)
		postData.WriteString("=")
		postData.WriteString(value)
		postData.WriteString("&")
	}
	return postData.Bytes()
}