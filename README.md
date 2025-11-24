# Configs Management Service Documentation
This is a service to manage configuration

## Prerequisite
- Go 1.25
- SQLite3
- Redis
- Docker
- Docker-compose

## Setup
### Setup Project
Clone the project
```bash
  git clone https://github.com/bobyindra/configs-management-service
```

Go to the project directory
```bash
  cd configs-management-service
```

Copy environment variables
```bash
  cp env.sample .env
```

### Setup Database
The schema will be migrated using golang-migrate tools, ensure you have installed it before. If you haven't, please follow this to install the golang-migrate

For Mac OS (using Homebrew)
```bash
  brew install golang-migrate
```
For any OS with Go installed
```bash
  go install github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

Create a directory with name `db` in the root project
```bash
  mkdir -p db
```

Run migration
```bash
  make migrate
```

Inject Users (For testing purpose because the API for registering an user is not provided yet in this project). We need the JWT token to hit the endpoint
```bash
  make inject-user
```

## Run the Application
### Run the server via docker compose - (Recommended)
**To start all of the servers (App and Redis)**
```bash
  make compose-up
```
*Notes:* Please wait around 1 minutes (for download dependencies and build the project), and then it can be accessed on 127.0.0.1:8080

**To stop the servers**
```bash
  make compose-down
```

### Run the server standalone
1. Run Redis
```bash
  sudo service redis-server start
  # or
  sudo systemctl start redis-server
  # or (macOS)
  brew services start redis
```
2. Run the App
```bash
  make run
```
It can be accessed on 127.0.0.1:8080

### Run test
Run unit test
```bash
  make test
```

Run unit test with coverage report
```bash
  make coverage
```

Run integration test
```bash
  make integration-test
```

## Work with service
### API Documentation
Before work with the endpoint, see the APIs Documentation [here](https://github.com/bobyindra/configs-management-service/blob/main/openapi.yaml) - Use [editor.swagger.io](https://editor.swagger.io/) to visualize it or VCS extension if you have

### Postman Collection
See the Postman Collection [here](https://github.com/bobyindra/configs-management-service/blob/main/collections/Configs%20Management%20Service.postman_collection.json)

### Supported Configs Schema
You can see the supported configs schema [here](https://github.com/bobyindra/configs-management-service/tree/main/module/configuration/schema)

#### Add Config Shcema
Want to add a new config schema? follow this instruction:
- Create a new json file or copy your schema into this [folder](https://github.com/bobyindra/configs-management-service/tree/main/module/configuration/schema)
- Register your schema in the schema_registry [here](https://github.com/bobyindra/configs-management-service/blob/main/module/configuration/schema/schema_registry.go) by following this
```go
  const (
    PAYMENT_CONFIG    = "payment-config"
    ...
    WORDING_CONFIG    = "wording-config"
    [YOUR_NEW_CONFIG] = "[your-config]" //<- Register your config name here
  )
```
```go
  var schemaFileMap = map[string]string{
	PAYMENT_CONFIG:  "payment_config.json",
	...
	WORDING_CONFIG:  "wording_config.json",
    [YOUR_NEW_CONFIG]: "your_config.json", //<- Put your new config here
  }  
```
- Rerun the service. That's it, you're all set!

### Get Token
As mentioned earlier, we have predifined users to access the endpoint. Here are the details:
| username | password        | Description                                                                   |
| :------- | :-------------- | :---------------------------------------------------------------------------- |
| `rouser` | `readonlyuser`  | This user only have access read operations such as get config                 |
| `rwuser` | `readwriteuser` | This user have access to both read and write operations such as create config |

hit this in your terminal to get the token (you can change the username and password) or feel free to hit it via postman (**Postman Collection is provided if needed [here](#postman-collection)**)
```bash
curl --location '127.0.0.1:8080/api/v1/auth/login' \
--header 'Content-Type: application/json' \
--data '{
    "username": "rwuser",
    "password": "readwriteuser"
}'
```

### Access Endpoint
Please follow the API documentation above to access the endpoint, don't forget to put the Authorization Token. (**Postman Collection is provided if needed [here](#postman-collection)**)

## Architecture Overview
![Project Architecture](https://drive.google.com/thumbnail?id=1X5hE3gBOUMYjwBJoU7MvwfVqfW3EPshb&sz=w350)

At the moment, the service architecture contains a simple relation between service, sql database, and redis. In the future, the service will have more comprehensive architecture to support real world usage. See the future development plans [here](#future-development)

### Project Structure
```
  └── config-management-service
    ├── build
    │ └── rest   <- contains Dockerfile
    ├── cmd
    │ ├── inject   <- temporary for testing purpose
    │ ├── migrate   <- SQLite migration library
    │ └── rest
    ├── db    <- SQLite db for this project
    ├── internal    <- store all global config and util for all module
    └── module
      └── configuration  <- the application logic is here
        ├── config
        ├── db
        │ └── migrations  <- contains all migrations schema
        ├── entity
        ├── helper
        ├── internal
        │ ├── auth
        │ ├── encryption
        │ ├── handler
        │ ├── middleware
        │ ├── repository
        │ └── usecase
        ├── schema
        └── test
          └── integration  <- all integration tests for this module are stored here

```
*Notes:*
- Module: different domain should be separated from a module. The purpose is if we need to take out the module outside of this project, it can be taked out easily. No dependency between module! To communicate between module, threat it as service to service communication.

### Database Schema
![Database Schema Diagram](https://drive.google.com/thumbnail?id=1u1tUgl4KVZK9uByMoLkVE8KSbIshjD8H&sz=w500)

**Configs Table**
| Column Name   | Data Type     | Property                                          | Description                                           |
| :------------ | :------------ | :------------------------------------------------ | :---------------------------------------------------- |
| id            | integer       | Primary Key, Auto Increament, Not Null            | id of config                                          |
| name          | varchar(100)  | Not Null, Composite Index with `version`          | Column to store config name                           |
| config_values | text          | Not Null                                          | Column to store config_values                         |
| version       | smallint      | Not Null, Composite Index with `name`             | Column to store config version                        |
| created_at    | timestamp     | -                                                 | Column to store the time when the config was created  |
| actor_id      | integer       | Not Null, Foreign Key of `id` from table `users`  | Column to store actor_id for audit purpose            |


**Index on configs Table**
| Index Name              | Column        | Description                                                                                                           |
| :---------------------- | :------------ | :-------------------------------------------------------------------------------------------------------------------- |
| idx_config_name_version | name, version | This unique index will store name and version of a config because we have query to get config based name and version  |
| idx_config_name         | name          | This index is to optimize the query of get config by its name                                                         |


**Users Table**
| Column Name      | Data Type    | Property                                | Description                                         |
| :--------------- | :----------- | :-------------------------------------- | :-------------------------------------------------- |
| id               | integer      | Primary Key, Auto Increament, Not Null  | id of users                                         |
| username         | varchar(50)  | Not Null, Unique                        | Column to store username                            |
| crypted_password | text         | Not Null                                | Column to store encrypted password                  |
| role             | varchar(50)  | Not Null                                | Column to store user's role                         |
| created_at       | timestamp    | -                                       | Column to store the time when the user was created  |
| updated_at       | timestamp    | -                                       | Column to store the time when the user was updated  |


**Index on users Table**
| Index Name         | Column    | Description                                                   |
| :----------------- | :-------- | :------------------------------------------------------------ |
| idx_users_username | username  | This index is to optimize the query of get users by username  |

### Notes & Trade-offs
- This service is focus for config management (Create, Update, Rollback, Get All Versions of a Config, Get the latest Config version or specific version, predefined schema validations, and several validations as can be found on the test doc)
- This service is already addressed proper indexes in database for the current usecases.
- At the moment, this service is only supported predefined config schema. I will add support for both predefined config schema and undefined config schema, as it will support the majority of the use cases.
- This service covered the basic role permission for accessing the config management. For the sake of simplicity, currently, I put the role inside the user table since this is not the main focus on this yet. For the real case or future development, we should implement a proper RBAC mechanism following with permissions, roles, roles permissions, and users’ roles tables.
- For the authentication, this service is already supported login endpoint to get the user's JWT access token. For the testing purpose, I put the user login inside the configs management module. It's should be separated from this module since users management has a different purpose from configs management. For real world scenario, this service can be accessed from other services through internal call (service to service communication) and gRPC to increase the performance because configs service is the heavy read load on the real scenario.
- At the moment, this service is only supported SQLite (just for the simplicity). Next, I will update this service to support PosgreSQL or MySQL or MongoDB database for better performance.

### Future Development
- Support both Predefined and Undefined schema
- Implement PostgreSQL or MySQL or MongoDB
- Implement proper RBAC mechanism
- Separate this service from User Management service
- Support client call/internal endpoint (service to service communication)
- Support both gRPC and REST API

## Author
- [@bobyindra](https://github.com/bobyindra)