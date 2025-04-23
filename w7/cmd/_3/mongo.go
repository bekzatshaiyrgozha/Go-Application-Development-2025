package main

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/gorilla/mux"
)

type Item struct {
	Id          primitive.ObjectID `json:"id" bson:"_id"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
	Updated     string             `json:"updated" bson:"updated"`
}

type Handler struct {
	Sess  *mongo.Session
	Items *mongo.Collection
	Tmpl  *template.Template
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	items := []*Item{}

	// bson.M{} - это типа условия для поиска
	cursor, err := h.Items.Find(context.Background(), bson.M{})
	__err_panic(err)

	err = cursor.All(context.Background(), &items)
	__err_panic(err)

	err = h.Tmpl.ExecuteTemplate(w, "index.html", struct {
		Items []*Item
	}{
		Items: items,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) AddForm(w http.ResponseWriter, r *http.Request) {
	err := h.Tmpl.ExecuteTemplate(w, "create.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) Add(w http.ResponseWriter, r *http.Request) {

	newItem := bson.M{
		"_id":         primitive.NewObjectID(),
		"title":       r.FormValue("title"),
		"description": r.FormValue("description"),
		"some_field":  123,
	}

	res, err := h.Items.InsertOne(context.Background(), newItem)
	__err_panic(err)

	fmt.Println("Inserted document with _id:", res.InsertedID)

	http.Redirect(w, r, "/", http.StatusFound)
}

func (h *Handler) Edit(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(vars["id"])
	if err != nil {
		http.Error(w, "bad id", http.StatusBadRequest)
		return
	}

	var post Item
	ctx := context.TODO()
	err = h.Items.FindOne(ctx, bson.M{"_id": id}).Decode(&post)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.Tmpl.ExecuteTemplate(w, "edit.html", post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	ctx := context.TODO()
	id, err := primitive.ObjectIDFromHex(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "bad id", http.StatusInternalServerError)
		return
	}

	update := bson.M{
		"$set": bson.M{
			"title":       r.FormValue("title"),
			"description": r.FormValue("description"),
			"updated":     "kbtu_user",
			"newField":    123,
		},
	}

	res, err := h.Items.UpdateOne(ctx, bson.M{"_id": id}, update)
	__err_panic(err)

	fmt.Println("Update - Matched:", res.MatchedCount, "Modified:", res.ModifiedCount)
	http.Redirect(w, r, "/", http.StatusFound)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := context.TODO()
	id, err := primitive.ObjectIDFromHex(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "bad id", http.StatusInternalServerError)
		return
	}

	res, err := h.Items.DeleteOne(ctx, bson.M{"_id": id})
	__err_panic(err)
	affected := res.DeletedCount

	w.Header().Set("Content-Type", "application/json")
	resp := `{"affected": ` + strconv.FormatInt(affected, 10) + `}`
	w.Write([]byte(resp))
}

func main() {
	ctx := context.TODO()
	sess, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost"))
	defer sess.Disconnect(ctx)
	__err_panic(err)

	// если коллекции не будет, то она создасться автоматически
	collection := sess.Database("kbtu").Collection("items")

	// для монги нет такого красивого дампа SQL, так что я вставляю демо-запись если коллекция пуста
	if n, _ := collection.CountDocuments(ctx, bson.M{}); n == 0 {
		collection.InsertOne(ctx, Item{
			Id:          primitive.NewObjectID(),
			Title:       "mongodb",
			Description: "Рассказать про монгу",
			Updated:     "",
		})
		collection.InsertOne(ctx, Item{
			Id:          primitive.NewObjectID(),
			Title:       "redis",
			Description: "Рассказать про redis",
			Updated:     "kbtu_user",
		})
	}

	handlers := &Handler{
		Items: collection,
		Tmpl:  template.Must(template.ParseGlob("../mongo_templates/*")),
	}

	// в целях упрощения примера пропущена авторизация и csrf
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.List).Methods("GET")
	r.HandleFunc("/items", handlers.List).Methods("GET")
	r.HandleFunc("/items/new", handlers.AddForm).Methods("GET")
	r.HandleFunc("/items/new", handlers.Add).Methods("POST")
	r.HandleFunc("/items/{id}", handlers.Edit).Methods("GET")
	r.HandleFunc("/items/{id}", handlers.Update).Methods("POST")
	r.HandleFunc("/items/{id}", handlers.Delete).Methods("DELETE")

	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", r)
}

// не используйте такой код в продакшене
// ошибка должна всегда явно обрабатываться
func __err_panic(err error) {
	if err != nil {
		panic(err)
	}
}
