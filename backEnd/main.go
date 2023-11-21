package main

import (
	"github.com/milimansilla44/proyecto-arqiSoftware/tree/master/backEnd/app"
	"github.com/milimansilla44/proyecto-arqiSoftware/tree/master/backEnd/db"
)

func main() {
	db.StartDbEngine()
	app.StartRoute()
}
