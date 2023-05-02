package datahandle

import (
    // melakukan encoding dikirimkan sebagai response dalam data json
    "encoding/json"
    // mengatur komunikasi http client and server
    "net/http"
)

type ItemData struct {
	Name        string
	StudentId   string
	Address     string
}

type Response struct {
    Status      string
    StatucCode  int
    Data        []ItemData
}

type Link struct {
    Link        string
    Method      string
    Description string
}

var Datas[]ItemData

var Links = [] Link {  
    Link{Link:"/",Method:"GET",Description:"See all API link"},
    Link{Link:"/student",Method:"GET",Description:"Get all student data"},
    Link{Link:"/student",Method:"POST",Description:"Store a student data"},
}

// handle request url "/"
func HomeUrl(w http.ResponseWriter, r *http.Request) {
    // switch dari request user method
    switch r.Method {
        case http.MethodGet:
            HomePage(w, r)
        default:
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

func HomePage(w http.ResponseWriter, r *http.Request) {
    // set header content type menjadi json
    w.Header().Set("Content-Type", "application/json")

    // mengambil data dari varibel untuk dijadikan sebagai response stream
    res := json.NewEncoder(w).Encode(Links)

    // jika encode tidak berhasil
    if res != nil {
        http.Error(w, res.Error(), http.StatusInternalServerError)
        return
    }
    
    // status response pada header
    w.WriteHeader(http.StatusOK)
}

// handle request url "/student"
func HandleUrl(w http.ResponseWriter, r *http.Request) {
    // switch dari request user method
    switch r.Method {
        case http.MethodGet:
            ShowAllData(w, r)
        case http.MethodPost:
            AddData(w, r)
        default:
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

func ShowAllData(w http.ResponseWriter, r *http.Request) {
    // set header content type menjadi json
    w.Header().Set("Content-Type", "application/json")

    // mengambil data dari varibel untuk dijadikan sebagai response stream
    res := json.NewEncoder(w).Encode(&Response{Status:"OK",StatucCode:200,Data:Datas})

    // jika encode tidak berhasil
    if res != nil {
        http.Error(w, res.Error(), http.StatusInternalServerError)
        return
    }
    
    // status response pada header
    w.WriteHeader(http.StatusOK)
}

func AddData(w http.ResponseWriter, r *http.Request) {
    // data yang dimasukkan user
    var data ItemData

    // set header content type menjadi json
    w.Header().Set("Content-Type", "application/json")

    // get data yang dimasukkan user dari body yang dikirim saat melakukan request
    // lalu decode
    res := json.NewDecoder(r.Body).Decode(&data)

    // jika decode tidak berhasil
    if res != nil {
        http.Error(w, res.Error(), http.StatusBadRequest)
        return
    }

    Datas = append(Datas,data)

    w.WriteHeader(http.StatusCreated)

    json.NewEncoder(w).Encode(&Response{"Created", 201, Datas})
}