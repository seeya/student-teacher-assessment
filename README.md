# Introduction

# Hosted API

# Setup Instructions

## MySQL and phpMyAdmin

Create an `.env` file with the following keys field saving it in the root directory.

```
MYSQL_ROOT_PASSWORD=CX98JU!M9#39f0tdaoAZjKwBrjnZKQRs
MYSQL_DATABASE=school
MYSQL_USER=admin
MYSQL_PASSWORD=6dE66&teSrc0thl&fBm*r!D*jqzxEQWZ
MYSQL_HOST=mysql
PMA_PORT=8081
```

Run the following command to spin up the containers.

```
# Spin up containers for MySQL and phpMyAdmin
docker-compose up -d
```

Access phpMyAdmin at `http://localhost:<PMA_PORT>`.

# New Branch

git push --set-upstream origin database
git checkout -b database
git push origin master:database
