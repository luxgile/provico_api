package projectdb

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
)

type ProjectId uint64

type Project struct {
	Id          ProjectId `json:"id"`
	Summary     string    `json:"summary"`
	Description string    `json:"description"`
	Tags        string    `json:"tags"`
}

var db *sql.DB

func Init() error {
  user, valid := os.LookupEnv("DBUSER")
  if !valid {
    return fmt.Errorf("no 'DBUSER' enviroment variable found")
  }
  password, valid := os.LookupEnv("DBPASS")
  if !valid {
    return fmt.Errorf("no 'DBPASS' enviroment variable found")
  }

	connStr := fmt.Sprintf("postgres://%s:%s@localhost/provico_testing?sslmode=disable", user, password)
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	if err = db.Ping(); err != nil {
		return err
	}

	return nil
}

func GetAllProjects() ([]Project, error) {
	rows, err := db.Query("SELECT * FROM project")
	if err != nil {
		return nil, err
	}

	var projects []Project
	for rows.Next() {
		var project Project
		if err := rows.Scan(&project.Id, &project.Summary, &project.Description, &project.Tags); err != nil {
			return nil, err
		}
		projects = append(projects, project)
	}
	return projects, nil
}

func GetProjectById(id ProjectId) (*Project, error) {
	rows, err := db.Query("SELECT * FROM project WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	var projects []Project
	for rows.Next() {
		var project Project
		if err := rows.Scan(&project.Id, &project.Summary, &project.Description, &project.Tags); err != nil {
			return nil, err
		}
		projects = append(projects, project)
	}

	if len(projects) == 0 {
		return nil, errors.New("no project found")
	}
	if len(projects) > 1 {
		return nil, errors.New("more than one project found")
	}

	return &projects[0], nil
}

func AddProject(summary, description, tags string) (ProjectId, error) {
  lastInsertId := 0
	err := db.QueryRow("INSERT INTO project (summary, description, tags) VALUES ($1, $2, $3) RETURNING id", summary, description, tags).Scan(&lastInsertId)
	if err != nil {
		return 0, err
	}

	return ProjectId(lastInsertId), nil
}

func UpdateProject(project Project) error {
	_, err := db.Exec("UPDATE project SET summary = $1, description = $2, tags = $3 WHERE id = $4", project.Summary, project.Description, project.Tags, project.Id)
	if err != nil {
		return err
	}

	return nil
}

func DeleteProject(id ProjectId) error {
	_, err := db.Exec("DELETE FROM project WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
