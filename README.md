# Introduction

https://gist.github.com/d3hiring/4d1415d445033d316c36a56f0953f4ef

## 1. Folder structure

```bash
- app
    |_ controllers
    |_ models
    |_ queries
- init
- platform
    |_ database
    |_ migrations
- docker-compose.yml
- Dockerfile
- main.go
- main_test.go
- Makefile
```

1. `app` - contains code which handles most of the business logic
2. `init` - used for intializing `docker-compose.yml` services
3. `platform` - includes databases and migration logic
4. `docker-compose.yml` - includes mysql and phpmyadmin services
5. `Dockerfile` - builds the application into a container
6. `main.go` - main entry point
7. `main_test.go` - main test entry point
8. `Makefile` - simplify common commands

# Setup Instructions for Development

## 0. Clone Project

Clone the project and enter into the root directory

```
# Clone
git clone git@github.com:seeya/student-teacher-assessment.git

# Enter directory
cd student-teacher-assessment
```

Create an `.env` file with the following keys field saving it in the root directory.

```yaml
echo "MYSQL_ROOT_PASSWORD=CX98JU!M939f0tdaoAZjKwBrjnZKQRs
MYSQL_DATABASE=school
MYSQL_USER=admin
MYSQL_PASSWORD=6dE66teSrc0thlfBmr
MYSQL_HOST=127.0.0.1
MYSQL_PORT=3306
PMA_PORT=8081
API_PORT=3001" > .env
```

## 1. MySQL and phpMyAdmin

Run the following command to spin up the containers.

```bash
# Keep the container running in the background
docker-compose up -d
```

Access phpMyAdmin at `http://localhost:<PMA_PORT>` to check if the database `school` and `test` are created.

## 2. Migrate & Seed the database

Migration is stored in `./platform/migrations/*.sql`
Once the default tables have been created, we will seed the database

```bash
# Setup default tables
make migrate.down migrate.up

# Start the application
go run main.go

# Execute the seed endpoint
make seed
```

## 3. Test Endpoints

A `Makefile` is included to simplify running repetitive commands.

```bash
# Clean the test cache, use 1 processor and run all tests
make test
```
