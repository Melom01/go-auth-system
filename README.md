# AUTH SYSTEM (codename: sentinel)

## OVERVIEW

This project implements an authentication system using Golang, with the following technologies:

- Router: Mux
- Database: PostgreSQL
- Authentication Provider: Keycloak
- Tokens: JWT (JSON Web Tokens)
- Protocol: OAuth2

## FEATURES

1. CRUD operations on the users database
2. OTP functionality via email


## PROJECT STRUCTURE

```
.
├── apperror                        #            
├── config                          #
├── controller                      #
├── database                        #
├── emailer                         #
│   └── templates                   #
├── logger                          #    
├── middleware                      #
├── model                           #
├── routing                         #
├── service                         #
├── tests                           #
│   ├── config_test                 #
│   ├── logger_test                 #
│   └── utils_test                  #    
│       └── mock_data               #
├── utils                           #
├── go.mod                          #
├── go.sum                          #
├── main.go                         #
└── README.md                       #  
```

## INSTALLATION

### Prerequisites:

- [Go](https://golang.org/dl/) (version 1.23+)
- [PostgreSQL](https://www.postgresql.org/)
- [Keycloak](https://www.keycloak.org/)

### Clone the repository

```bash
git clone https://github.com/Melom01/go-auth-system.git
cd go-auth-system
```

### Set up dependencies
Install Go modules:

```bash
go mod tidy
```

---

## CONFIGURATION

Update the `config.json` file with your configuration:
```JSON

```
- **DB_HOST**: Database host. For example, **localhost** or **127.0.0.1**.
- **DB_PORT**: Database port. For example, **5432**.
- **DB_NAME**: Database name.
- **DB_USER**: Database username.
- **DB_PASSWORD**: Your password to access the database.
- **FIREBASE_CREDENTIALS**: The path to your Firebase Admin SDK JSON file.


## USAGE

### Run the application

```shell
go run main.go
```

### Sort imports
#### Making the `sort_imports.sh` script executable

Before running the script for the first time, you need to make it executable. You can do this by running the following command in your terminal:
```shell
chmod +x sort_imports.sh
```
This step is only required once.

#### Running the script
After making the script executable, you can sort all imports throughout the application by running:
```shell
./sort_imports.sh
```
The script will automatically process all Go files, excluding any specified patterns (e.g., files with `_gen.go` in the di folder).
Once completed, you will see a confirmation message.

## LICENSE

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
