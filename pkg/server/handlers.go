package server

//import (
//	"github.com/gorilla/mux"
//	"github.com/sirupsen/logrus"
//	"github.com/todoApp/pkg/api/todo"
//	"net/http"
//)
//
//
//type TodoServiceHandlers struct {
//	todoApi *todo.TodoService
//}
//
//func NewHTTPRouter(todoService *todo.TodoService) *mux.Router {
//	r := mux.NewRouter()
//
//	todoServiceHandlers := TodoServiceHandlers{todoApi: todoService}
//
//	getRouters := r.Methods(http.MethodGet).Subrouter()
//	getRouters.HandleFunc("/", todoServiceHandlers.GetAllItems)
//	getRouters.HandleFunc("/{id}", todoServiceHandlers.GetSingleItem)
//
//	postRouters := r.Methods(http.MethodPost).Subrouter()
//	postRouters.HandleFunc("/", todoServiceHandlers.AddItem)
//
//	putRouters := r.Methods(http.MethodPut).Subrouter()
//	putRouters.HandleFunc("/{id}", todoServiceHandlers.UpdateItems)
//
//	deleteRouters := r.Methods(http.MethodDelete).Subrouter()
//	deleteRouters.HandleFunc("/{id}", todoServiceHandlers.DeleteItem)
//
//	return r
//}
//
//func (handler *TodoServiceHandlers) GetAllItems(rw http.ResponseWriter, r *http.Request) {
//	log.Info("Hit getAllItems handler")
//
//	handler.todoApi.GetItems()
//	//if err != nil {
//	//	i.l.Println(err)
//	//	return
//	//}
//	//
//	//err = items.ToJSON(rw)
//	//if err != nil {
//	//	http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
//	//}
//}
//
//func (handler *TodoServiceHandlers) GetSingleItem(rw http.ResponseWriter, r *http.Request) {
//	//i.l.Println("getSingleItems handler")
//	//
//	//vars := mux.Vars(r)
//	//
//	//item, err := data.GetItem(vars["id"])
//	//if err != nil {
//	//	i.l.Println(err)
//	//	return
//	//}
//	//
//	//err = item.ToJSON(rw)
//	//if err != nil {
//	//	http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
//	//}
//}
//
//func (handler *TodoServiceHandlers) AddItem(rw http.ResponseWriter, r *http.Request) {
//	//	i.l.Println("Hit addItems handler")
//	//
//	//	//body, err := ioutil.ReadAll(r.Body)
//	//	//if err != nil {
//	//	//	http.Error(rw, "Unable to read body", http.StatusInternalServerError)
//	//	//	return
//	//	//}
//	//	err := data.AddItem(r.Body)
//	//	if err != nil {
//	//		http.Error(rw, "Unable to read body", http.StatusBadRequest)
//	//		return
//	//	}
//}
//
//func (handler *TodoServiceHandlers) UpdateItems(rw http.ResponseWriter, r *http.Request) {
//	//	i.l.Println("Hit updateItems handler")
//}
//
//func (handler *TodoServiceHandlers) DeleteItem(rw http.ResponseWriter, r *http.Request) {
//	//	i.l.Println("Hit deleteItems handler")
//}
