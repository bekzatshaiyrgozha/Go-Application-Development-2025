package main

import (
	"html/template"
	"log"
	"net/http"
)

type Storage interface {
	Add(*Photo) error
	GetPhotos(int) ([]*Photo, error)
}

// -----------------------------

type PhotolistHandler struct {
	St   Storage
	Tmpl *template.Template
}

func (h *PhotolistHandler) List(w http.ResponseWriter, r *http.Request) {
	items, err := h.St.GetPhotos(userID)
	if err != nil {
		log.Println("cant get items", err)
		http.Error(w, "storage error 1", http.StatusInternalServerError)
		return
	}

	err = h.Tmpl.Execute(w,
		struct {
			Items []*Photo
		}{
			items,
		})
	if err != nil {
		log.Println("cant execute template", err)
		http.Error(w, "template error", http.StatusInternalServerError)
		return
	}
}

func (h *PhotolistHandler) Upload(w http.ResponseWriter, r *http.Request) {
	uploadData, _, err := r.FormFile("my_file")
	if err != nil {
		log.Println("cant parse file", err)
		http.Error(w, "request error", http.StatusInternalServerError)
		return
	}
	defer uploadData.Close()

	md5Sum, err := SaveFile(uploadData)
	if err != nil {
		log.Println("cant save file", err)
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}

	realFile := "./../images/" + md5Sum + ".jpg"
	err = MakeThumbnails(realFile, md5Sum)
	if err != nil {
		log.Println("cant resize file", err)
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}

	err = h.St.Add(&Photo{UserID: userID, Path: md5Sum})
	if err != nil {
		log.Println("cant store item", err)
		http.Error(w, "storage error 2", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/photos", 302)
}

var (
	userID = 0
)
