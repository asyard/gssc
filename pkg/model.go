package pkg

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
