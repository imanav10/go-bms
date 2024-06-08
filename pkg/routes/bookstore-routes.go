package routes

// All Postman Routing
import(
	"github.com/gorilla/mux"
	"github.com/imanav10/go-bms/pkg/controllers"
)


var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}