package gssc

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
	resp, err := http.Get("https://rerererarara.net/api/v1/basliklar/bugun")
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

	var todayHeadersResp HeadersTodayResponse
	if err := json.Unmarshal(body, &todayHeadersResp); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}

	return todayHeadersResp
}

type HeadersTodayResponse struct {
	Header     string `json:"liste"`
	Date       string `json:"tarih"`
	TotalCount int    `json:"baslik_sayisi"`
	Headers    []struct {
		HeaderText string `json:"baslik"`
		HeaderID   int    `json:"baslik_id"`
		Count      int    `json:"sayi"`
		Link       string `json:"link"`
	} `json:"basliklar"`
}

func EntriesToday(headerId int) (resultData EntriesTodayResponse) {
	resp, err := http.Get("https://rerererarara.net/api/v1/entryler/birgun?baslik_id=" + strconv.Itoa(headerId) + "&tarih=" + time.Now().Format("2006-01-02"))
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

	var todayResp EntriesTodayResponse
	if err := json.Unmarshal(body, &todayResp); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}

	return todayResp
}

type EntriesTodayResponse struct {
	BaslikID    int    `json:"baslik_id"`
	Baslik      string `json:"baslik"`
	Tarih       string `json:"tarih"`
	EntrySayisi int    `json:"entry_sayisi"`
	Entryler    []struct {
		EntryID       int    `json:"entry_id"`
		SiraNo        int    `json:"sira_no"`
		Mesaj         string `json:"mesaj"`
		Yazar         string `json:"yazar"`
		Tarih         string `json:"tarih"`
		TarihDuzeltme string `json:"tarih_duzeltme"`
	} `json:"entryler"`
	ToplamSayfa int `json:"toplam_sayfa"`
	Sayfa       int `json:"sayfa"`
}

func AgendaToday() (resultData AgendaTodayResponse) {
	resp, err := http.Get("https://rerererarara.net/api/v1/basliklar/gundem/basliklar")
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

	var todayAgendaResp AgendaTodayResponse
	if err := json.Unmarshal(body, &todayAgendaResp); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}

	return todayAgendaResp
}

type AgendaTodayResponse struct {
	List        string `json:"liste"`
	Date        string `json:"tarih"`
	Day         string `json:"gun"`
	HeaderCount int    `json:"baslik_sayisi"`
	Headers     []struct {
		HeaderText string `json:"baslik"`
		HeaderID   int    `json:"baslik_id"`
		Count      int    `json:"sayi"`
		Link       string `json:"link"`
	} `json:"basliklar"`
}

func GetEntry(entryId string) (resultData SingleEntry) {
	resp, err := http.Get("https://rerererarara.net/api/v1/entry/detay?entry_id=" + entryId)
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

	var singleEntry SingleEntry
	if err := json.Unmarshal(body, &singleEntry); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}

	return singleEntry
}

type SingleEntry struct {
	HeaderID        int         `json:"baslik_id"`
	HeaderText      string      `json:"baslik"`
	TotalEntryCount int         `json:"entry_sayisi"`
	EntryID         int         `json:"entry_id"`
	OrderId         int         `json:"sira_no"`
	EntryText       string      `json:"mesaj"`
	Author          string      `json:"yazar"`
	Date            string      `json:"tarih"`
	DateModified    interface{} `json:"tarih_duzeltme"`
}

func GetAuthor(authorNick string) (resultData AuthorData) {
	resp, err := http.Get("https://rerererarara.net/api/v1/kullanici/profil?nick=" + authorNick)
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

	var authorData AuthorData
	if err := json.Unmarshal(body, &authorData); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}

	return authorData
}

type AuthorData struct {
	Author struct {
		Nick      string `json:"nick"`
		Genre     string `json:"nesil"`
		Status    string `json:"durum"`
		Karma     string `json:"karma"`
		Online    bool   `json:"online"`
		LastEntry string `json:"son_entry"`
	} `json:"kullanici"`
	EntryStats struct {
		Total int `json:"toplam"`
		Week  int `json:"hafta"`
		Month int `json:"ay"`
		Today int `json:"bugun"`
	} `json:"entry_sayilari"`
}

func GetAuthorEntryStats(authorNick string) (resultData AuthorEntryStats) {
	resp, err := http.Get("https://rerererarara.net/api/v1/kullanici/liste?nick=" + authorNick + "&liste=1")
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

	var authorEntryData AuthorEntryStats
	if err := json.Unmarshal(body, &authorEntryData); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}

	return authorEntryData
}

type AuthorEntryStats struct {
	AuthorityName string `json:"kullanici"`
	Lists         struct {
		LastEntriesText            string `json:"1"`
		PopularThisWeekEntriesText string `json:"2"`
		MostLikedEntriesText       string `json:"3"`
		MostLikedLongEntriesText   string `json:"4"`
		MostDislikedEntriesText    string `json:"5"`
		MostFavedEntriesText       string `json:"6"`
		FavedEntriesText           string `json:"7"`
	} `json:"listeler"`
	List    string `json:"liste"`
	Entries []struct {
		EntryID           int         `json:"entry_id"`
		HeaderText        string      `json:"baslik"`
		EntryText         string      `json:"mesaj"`
		EntryDate         string      `json:"tarih"`
		EntryModifiedDate interface{} `json:"tarih_duzeltme"`
	} `json:"entryler"`
	PageId int `json:"sayfa"`
}
