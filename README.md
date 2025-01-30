# Requirements
- .env file on the root directory with the database credentials for "DBUSER" and "DBPASS".
- Go installed.
- Postgres SQL 

# How to start server
Simply run launch.sh - The server will start locally at port '8080'

# End points available
- GET ```/projects``` - to obtain all projects
- GET ```/projects/{id}``` - to get a project by its id
- POST ```/projects``` - to submit a new project
- PUT ```/projects/{id}``` - to update a project by its id
- DELETE ```/projects/{id}``` - to delete a project by its id
