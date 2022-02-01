package handlers

import (
	"github.com/gorilla/mux"
	"github.com/todolist/data"
	"log"
	"net/http"
)

type Items struct {
	l *log.Logger
}

func NewItems(l *log.Logger) *Items {
	return &Items{l}
}

func (i *Items) GetAllItems(rw http.ResponseWriter, r *http.Request) {
	i.l.Println("Hit getAllItems handler")

	items, err := data.GetItems()
	if err != nil {
		i.l.Println(err)
		return
	}

	err = items.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (i *Items) GetSingleItems(rw http.ResponseWriter, r *http.Request) {
	i.l.Println("getSingleItems handler")

	vars := mux.Vars(r)

	item, err := data.GetItem(vars["id"])
	if err != nil {
		i.l.Println(err)
		return
	}

	err = item.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (i *Items) AddItems(rw http.ResponseWriter, r *http.Request) {
	i.l.Println("Hit addItems handler")

	//body, err := ioutil.ReadAll(r.Body)
	//if err != nil {
	//	http.Error(rw, "Unable to read body", http.StatusInternalServerError)
	//	return
	//}
	err := data.AddItem(r.Body)
	if err != nil {
		http.Error(rw, "Unable to read body", http.StatusBadRequest)
		return
	}
}

func (i *Items) UpdateItems(rw http.ResponseWriter, r *http.Request) {
	i.l.Println("Hit updateItems handler")
}

func (i *Items) DeleteItems(rw http.ResponseWriter, r *http.Request) {
	i.l.Println("Hit deleteItems handler")
}
