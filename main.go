package main

import (
	"log"
	"net/http"

	"github.com/PaiAkshay998/EdTech_CSF/handler"
	"github.com/PaiAkshay998/EdTech_CSF/utils"
)

func main() {
	config := utils.GetConfiguration()
	utils.Init(config)
	utils.Logger.Infof("Server running on port %s",config.ServerPort)
	http.HandleFunc("/searchStudentMaterial", handler.SearchStudentMaterial)
	http.HandleFunc("/searchTeacherMaterial", handler.SearchTeacherMaterial)
	log.Fatal(http.ListenAndServe(config.ServerPort, nil))
}
