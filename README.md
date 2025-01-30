# Provico
Provico is a server backend designed to submit and store project ideas. Its main goal is to provide 
different options if you are stuck with your next project or are looking for inspiration.

**Important**:
Provico is meant to be a learning project to start getting used to both backend development and the Go language.
As such I would not recommend using this server in a profesional matter, or do so knowing what you are doing.

### Requirements
- .env file on the root directory with the database credentials for "DBUSER" and "DBPASS".
- Go installed.
- Postgres SQL 

### How to start server
Simply run launch.sh - The server will start locally at port '8080'

### End points available
- GET ```/projects``` - to obtain all projects
- GET ```/projects/{id}``` - to get a project by its id
- POST ```/projects``` - to submit a new project
- PUT ```/projects/{id}``` - to update a project by its id
- DELETE ```/projects/{id}``` - to delete a project by its id
