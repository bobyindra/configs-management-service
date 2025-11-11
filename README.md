# Configs Management Service Documentation
This is a service to manage configs
## Setup
### Prerequisite
- Go 1.25
- SQLite

### Installation
#### Setup Project
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

#### Setup Database
Run migration
```bash
  make migrate
```

Inject Users (For testing purpose because the API for registering an user is not provided yet in this project). We need the JWT token to hit the endpoint
```bash
  make inject-user
```

Run the server

```bash
  make run
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
  func GetSchemaByConfigName(cfgName string) ([]byte, error) {
    switch cfgName {
    case PAYMENT_CONFIG:
      return os.ReadFile("./module/configuration/schema/payment_config.json")
    ...
    case WORDING_CONFIG:
      return os.ReadFile("./module/configuration/schema/wording_config.json")
    case [YOUR_NEW_CONFIG]: //<- Put your new config var here
      return os.ReadFile("./module/configuration/schema/your_config.json") //<- Don't forget to update the path
    default:
      return nil, entity.ErrConfigNotFound
    }
  }
```
- Rerun the service. That's it, you're all set!

### Get Token
As mentioned earlier, we have predifined users to access the endpoint. Here are the details:
| username | password        | Description                                                                   |
| :------- | :-------------- | :---------------------------------------------------------------------------- |
| `rouser` | `readonlyuser`  | This user only have access read operations such as get config                 |
| `rwuser` | `readwriteuser` | This user have access to both read and write operations such as create config |

hit this in your terminal to get the token (you can change the username and password) or feel free to hit it via postman (**Postman Collection is provided if needed**)
```bash
curl --location '127.0.0.1:8080/api/v1/auth/login' \
--header 'Content-Type: application/json' \
--data '{
    "username": "rwuser",
    "password": "readwriteuser"
}'
```

### Access Endpoint
Please follow the API documentation above to access the endpoint, don't forget to put the Authorization Token. (**Postman Collection is provided if needed**)

## Architecture Overview

### Project Structure

### Database Schema

### Notes & Trade-offs

### Future Development

## API Documentation

## Question?