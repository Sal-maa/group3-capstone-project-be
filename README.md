<!-- ABOUT THE PROJECT -->
### 💻 &nbsp;About The Project

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

This application provides REST API for Company Asset Management System. Among its features are
- **User**: Login, get user detail by user id, get all users, update user, get history of past requests of user
- **Asset**: Create asset, get asset detail by asset short-name, get all assets and filter them by category, update maintenance status of asset, get asset statistics, get asset usage history
- **Request**: Request to borrow/procure an asset, get borrow/procure request detail by request id, get all borrow/procure requests and filter them by status or request date or asset category, update status of borrow/procure request.

### 🕮 &nbsp;Documentation

See open API documentation [here](https://app.swaggerhub.com/apis-docs/bagusbpg6/group3-capstone-API/1.0.0).

### 🖳 &nbsp;Installation
To run the application, clone this repository first,
```bash
git clone https://github.com/bagusbpg/group3-capstone-project-be.git
```
Then, set up configuration file in JSON.
```json
{
    "app_port": ,
    "jwt_secret": ,
    "database": {
        "db_driver": ,
        "db_host": ,
        "db_port": ,
        "db_username": ,
        "db_password": ,
        "db_name":
    },
    "aws": {
        "aws_accesskeyid": ,
        "aws_secretkey": ,
        "aws_region": ,
        "aws_bucket":
    }
}
```
or using environment variables,
```bash
PORT=
JWT_SECRET=
DB_DRIVER=
DB_HOST=
DB_PORT=
DB_USERNAME=
DB_PASSWORD=
DB_NAME=
AWS_ACCESSKEYID=
AWS_SECRETKEY=
AWS_REGION=
AWS_BUCKET=
```
Download the required libraries or packages
```bash
cd group3-capstone-project-be
```
```go
go mod tidy
go get
```
Create the database along with all necessary tables as defined in ERD. Then, run the application
```go
source <env-file>
go run ./app/main.go
```
To run the application using docker, build the docker image first,
```bash
docker build -t <image-name>[:tag] .
```
Then run docker image as container,
```bash
docker run -d --env-file <env file> -p <host-port>:<container-port> --name <container-name> <image-name>
```
### 🛠 &nbsp;Tech Stacks

![Golang](https://img.shields.io/badge/-Golang-05122A?style=flat&logo=go&logoColor=4479A1)&nbsp;
![MySQL](https://img.shields.io/badge/-MySQL-05122A?style=flat&logo=mysql&logoColor=4479A1)&nbsp;
![AWS](https://img.shields.io/badge/-AWS-05122A?style=flat&logo=amazon)&nbsp;
![Postman](https://img.shields.io/badge/-Postman-05122A?style=flat&logo=postman)&nbsp;
![GitHub](https://img.shields.io/badge/-GitHub-05122A?style=flat&logo=github)&nbsp;
![Visual Studio Code](https://img.shields.io/badge/-Visual%20Studio%20Code-05122A?style=flat&logo=visual-studio-code&logoColor=007ACC)&nbsp;
![Docker](https://img.shields.io/badge/-Docker-05122A?style=flat&logo=docker)&nbsp;
![Ubuntu](https://img.shields.io/badge/-Ubuntu-05122A?style=flat&logo=ubuntu)&nbsp;

## How to Run

Clone the project

```bash
  git clone https://github.com/bagusbpg/group3-capstone-project-be.git
```

Go to the project directory

```bash
  cd group3-capstone-project-be
```

Install dependencie

```bash
  go mod tidy
```

Required .env file

```
PORT=
JWT_SECRET=
DB_DRIVER=mysql
DB_HOST=
DB_PORT=
DB_USERNAME=
DB_PASSWORD=
DB_NAME=
AWS_ACCESSKEYID=
AWS_SECRETKEY=
AWS_REGION=
AWS_BUCKET=
```

Start the server

```bash
  go run app/main.go
```

<!-- CONTACT -->
### 📭 &nbsp;Contact

[![GitHub Bagus](https://img.shields.io/badge/-Bagus-white?style=flat&logo=github&logoColor=black)](https://github.com/bagusbpg)
[![GitHub Salmaa](https://img.shields.io/badge/-Salmaa-white?style=flat&logo=github&logoColor=black)](https://github.com/Sal-maa)
[![GitHub Yahya](https://img.shields.io/badge/-Yahya-white?style=flat&logo=github&logoColor=black)](https://github.com/zakariyahya)

<p align="center">:copyright: 2022</p>
</h3>
