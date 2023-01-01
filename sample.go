// You can edit this code!
// Click here and start typing.
package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
	"sync"
	"unicode/utf8"
)

// func main() {
	var once sync.Once
	onceBody := func() {
		fmt.Println("Only once")
	}

	// this is a channel where receive signals, data type of the signal is boolean
	done := make(chan bool)

	for i := 0; i < 10; i++ {
		fmt.Println("upper loop", i)
		go func() {
			// following line runs on random time which we cant predict
			once.Do(onceBody)
			done <- true
		}()
	}
	for i := 0; i < 10; i++ {
		fmt.Println("down loop", i)
		<-done
	}
	/*
		sample := []int{1, 2, 3, 4}
		fmt.Println(sample[len(sample)-1])

		/*
			if time.Sunday.String() == "Sunday" {
				fmt.Println("equal")
			}

			//loc, _ := time.LoadLocation("Australia/Adelaide")
			timeNow, _ := time.Parse("2006-01-02", time.Now().Format("2006-01-02"))
			fmt.Println(timeNow)
	*/
	//t := time.Date(2022, 9, 5, 23, 0, 0, 0, time.UTC)
	//fmt.Println("time1 - ", t)
	//fmt.Println(time.Now().UTC())
	//now, _ := time.Parse("2006-01-02", time.Now().Add(24*time.Hour).UTC().Format("2006-01-02"))
	//
	//fmt.Println("time2 - ", now)
	//today, _ := time.Parse("2006-01-02", "2022-09-05")
	//fmt.Println("time3 - ", today)
	//today.After(now)

	// this is how we get the daw of the tomorrow or end of today

	/*
		specDate := "2022-09-04"
		timeZone := "UTC"
		loc, _ := time.LoadLocation(timeZone)

		specTime, _ := time.ParseInLocation("2006-01-02", specDate, loc)
		fmt.Println(specTime)

		now2 := time.Now().In(loc).Truncate(24 * time.Hour)
		fmt.Println(now2)

		fmt.Println(specTime.Before(now2))
	*/
	/*
		// preparing the header
		headers := make(http.Header)
		headers.Set("account-id", "d8339377-8b31-48cb-a188-ab4092b9b12d")
		headers.Set("trace-id", "d8339377-8b31-48cb-a188-ab4092b9b12d")
		headers.Set("user-id", "d8339377-8b31-48cb-a188-ab4092b9b12d")
		headers.Set("Content-Type", "application/json")
		fmt.Println("header is- ", headers)

		// preparing the request body
		requestBody := struct {
			StartDate string `json:"startDate"`
			EndDate   string `json:"endDate"`
			Currency  string `json:"currency" default:"AUD"`
		}{
			StartDate: "2022-09-03",
			EndDate:   "2023-09-03",
			Currency:  "AUD",
		}
		fmt.Println("request body is", requestBody)
		jsonStr, _ := json.Marshal(requestBody)
		fmt.Println("request body is in byte arr", jsonStr)

		// preparing the URL
		baseURL := "http://localhost:9000" + fmt.Sprintf("/timeline/%s?budgetType=BUDGET&&timelineType=STREAM&&streamId=%s&&debug=true", "8e8c0b68-49d9-416c-a419-f96ea380947e", "f8806aa0-93af-40bb-a5d5-2568816c630c")
		fmt.Println("URL is", baseURL)

		// preparing the request
		request, _ := http.NewRequestWithContext(context.Background(), http.MethodPost, baseURL, bytes.NewReader(jsonStr))
		request.Header = headers
		fmt.Println("request is ", request)

		// hit the request
		client := http.Client{
			Timeout: time.Second * 5,
		}
		response, _ := client.Do(request)
		defer response.Body.Close()
		fmt.Println(response)

		respBody, _ := io.ReadAll(response.Body)
		//fmt.Println("response body is - ", respBody)
		s := string(respBody)
		fmt.Println("string response is - ", s)
	*/

	/*
		fmt.Println("program begins")
		start := time.Now()
		num := "20000000"
		//fmt.Println(num[0], num[1])
		if num[0] == '2' {
			fmt.Println("this is 2")
		}

		elapsed := time.Since(start)
		log.Printf("took %s", elapsed)

		///////////////////////////////////////////////
		start2 := time.Now()

		num2, _ := strconv.Atoi(num)
		//fmt.Println(num2[0], num2[1])
		if 20000000 <= num2 && num2 <= 29999999 {
			fmt.Println("this is 2")
		}

		elapsed2 := time.Since(start2)
		log.Printf("took %s", elapsed2)

		// when the condition becomes recursive its complex in memory wise and time wise also
	*/
	//for i := 0; i <= 10; i++ {
	//	x, _ := generateClientNumber()
	//	fmt.Println(x)
	//}

}

// generateClientNumber to generate random, non-sequential 8 digit number
func generateClientNumber() (string, error) {
	randVal, err := rand.Int(rand.Reader, big.NewInt(99999999-0))

	if randVal.Int64() >= 20000000 && randVal.Int64() <= 30000000 {
		fmt.Println("between the range ", randVal)
		return generateClientNumber()
	}

	if err != nil {
		return "", err
	}
	randString := randVal.String()

	clientNum := strings.Repeat("0", 8-utf8.RuneCountInString(randString)) + randString

	return clientNum, nil
}
