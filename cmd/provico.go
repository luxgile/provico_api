package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/luxgile/provico/internal/projectdb"
)

func GetAllProjectsHandler(w http.ResponseWriter, r *http.Request) {
	projects, err := projectdb.GetAllProjects()
	if err != nil {
    http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
    return
	}

  w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(projects)
	if err != nil {
    http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
    return
	}
}

func GetSingleProjectHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
    http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
    return 
	}

	project, err := projectdb.GetProjectById(projectdb.ProjectId(id))
	if err != nil {
    http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
    return
	}

	err = json.NewEncoder(w).Encode(project)
	if err != nil {
    http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
    return
	}
}

func PostProjectHandler(w http.ResponseWriter, r *http.Request) {
  var project projectdb.Project
  decoder := json.NewDecoder(r.Body)
  err := decoder.Decode(&project)
  if err != nil {
    http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
    return
  }

  id, err := projectdb.AddProject(project.Summary, project.Description, project.Tags)
	if err != nil {
    http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
    return
	}

  returnedProject, err := projectdb.GetProjectById(id)
	if err != nil {
    http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
    return
	}

	err = json.NewEncoder(w).Encode(returnedProject)
	if err != nil {
    http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
    return
	}
}

func PutProjectHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
    http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
    return
	}

  var project projectdb.Project
  decoder := json.NewDecoder(r.Body)
  err = decoder.Decode(&project)
  if err != nil {
    http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
    return
  }
    
  project.Id = projectdb.ProjectId(id)

	err = projectdb.UpdateProject(project)
	if err != nil {
    http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
    return
	}
  
  returnedProject, err := projectdb.GetProjectById(project.Id)
	if err != nil {
    http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
    return
	}

	err = json.NewEncoder(w).Encode(returnedProject)
	if err != nil {
    http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
    return
	}
}

func DeleteProjectHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
    http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
    return
	}
    
	err = projectdb.DeleteProject(projectdb.ProjectId(id))
	if err != nil {
    http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
    return
	}
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Executing middleware", r.Method)

		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	  w.Header().Set("Content-Type", "application/json")

		next.ServeHTTP(w, r)
		log.Println("Executing middleware again")
	})
}

func main() {
  err := projectdb.Init()
  if err != nil {
    panic(err)
  }

	r := mux.NewRouter()
	r.HandleFunc("/projects", GetAllProjectsHandler).Methods("GET")
	r.HandleFunc("/projects", PostProjectHandler).Methods("POST")
	r.HandleFunc("/projects/{id}", GetSingleProjectHandler).Methods("GET")
	r.HandleFunc("/projects/{id}", PutProjectHandler).Methods("PUT")
	r.HandleFunc("/projects/{id}", DeleteProjectHandler).Methods("DELETE")
	log.Fatalln(http.ListenAndServe(":8080", corsMiddleware(r)))
}
