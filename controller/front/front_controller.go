package front

import (
	"html/template"
	"path"
	"net/http"

	"github.com/ibamaulana/golang-kreditplus/config"
	"github.com/ibamaulana/golang-kreditplus/services"
	"github.com/ibamaulana/golang-kreditplus/model"
)

func GetController(w http.ResponseWriter, r *http.Request) {
	var filepath = path.Join("view", "front.html")
	var tmpl, err = template.ParseFiles(filepath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var data = map[string]interface{}{
		"title": "Learning Golang Web",
		"name":  "Batman",
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Submit(w http.ResponseWriter) {
	var player *model.Player
	//save to database
	cfg := config.NewConfig()
	db, err := config.MysqlConnection(cfg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	//get service
	playerContract := services.NewPlayerServiceContract(db)
	tx := db.Begin()
	defer tx.Rollback()

	//run create on db
	player, err = playerContract.Create(player, tx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	tx.Commit()
}
