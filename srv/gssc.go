package srv

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

func HeadersToday() (resultData HeadersTodayResponse) {
	resp, err := http.Get(GsscApiHeadersToday)
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
		}
	}(resp.Body)

	if resp.StatusCode != 200 {
		log.Fatal(":(")
	}

	body, err := ioutil.ReadAll(resp.Body) // response body is []byte
	if err != nil {
		log.Fatal(err)
	}

	var todayHeadersResp HeadersTodayResponse
	if err := json.Unmarshal(body, &todayHeadersResp); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}

	return todayHeadersResp
}

func EntriesToday(headerId int) (resultData EntriesTodayResponse) {
	resp, err := http.Get(GsscApiEntriesOfHeaderToday + strconv.Itoa(headerId) + "&tarih=" + time.Now().Format("2006-01-02"))
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	if resp.StatusCode != 200 {
		log.Fatal(":(")
	}

	body, err := ioutil.ReadAll(resp.Body) // response body is []byte
	if err != nil {
		log.Fatal(err)
	}

	var todayResp EntriesTodayResponse
	if err := json.Unmarshal(body, &todayResp); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}

	return todayResp
}

func AgendaToday() (resultData AgendaTodayResponse) {
	resp, err := http.Get(GsscApiHeadersAgenda)
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
		}
	}(resp.Body)

	if resp.StatusCode != 200 {
		log.Fatal(":(")
	}

	body, err := ioutil.ReadAll(resp.Body) // response body is []byte
	if err != nil {
		log.Fatal(err)
	}
	var todayAgendaResp AgendaTodayResponse
	if err := json.Unmarshal(body, &todayAgendaResp); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}

	return todayAgendaResp
}

func GetEntry(entryId string) (resultData SingleEntry) {
	resp, err := http.Get(GsscApiEntryDetail + entryId)
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
		}
	}(resp.Body)

	if resp.StatusCode != 200 {
		log.Fatal(":(")
	}

	body, err := ioutil.ReadAll(resp.Body) // response body is []byte
	if err != nil {
		log.Fatal(err)
	}
	var singleEntry SingleEntry
	if err := json.Unmarshal(body, &singleEntry); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}

	return singleEntry
}

func GetAuthor(authorNick string) (resultData AuthorData) {
	resp, err := http.Get(GsscApiProfileDetail + authorNick)
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
		}
	}(resp.Body)

	if resp.StatusCode != 200 {
		log.Fatal(":(")
	}

	body, err := ioutil.ReadAll(resp.Body) // response body is []byte
	if err != nil {
		log.Fatal(err)
	}

	var authorData AuthorData
	if err := json.Unmarshal(body, &authorData); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}

	return authorData
}

func GetAuthorEntryStats(authorNick string) (resultData AuthorEntryStats) {
	resp, err := http.Get(GsscApiProfileStatsDetail + authorNick + "&liste=1")
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
		}
	}(resp.Body)

	if resp.StatusCode != 200 {
		log.Fatal(":(")
	}

	body, err := ioutil.ReadAll(resp.Body) // response body is []byte
	if err != nil {
		log.Fatal(err)
	}

	var authorEntryData AuthorEntryStats
	if err := json.Unmarshal(body, &authorEntryData); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}

	return authorEntryData
}
