# Pomotodo - Product 

## 1. Project Introduction
    
   - The pomotodo project based on vue and gin, which separates the front and rear of the full stack. Pomotodo contains full workflow management. Stay focused at work and get more done with Pomotodo.

## 2. How to run

-  Required

   - PostgreSQL database
   - gin
   - vue
   - IDE recommendation: VSCode
   - docker-compose
  

    ```bash
    # clone the project
    git clone https://github.com/jingyuan-git/pomotodo

    # create an database for pomotodo in PostgreSQL
    # and than will write database-related information into the configuration
    # in `server/conf/app.ini`
    ``` 

### 2.1 server project

- conf

    ```bash
    You should modify `server/conf/app.ini`

    [server]
    ; debug or release
    RunMode = debug
    ; Host = localhost
    HttpPort = 8000
    ...

    [database]
    Type = postgres
    User = 
    Password = 
    Host = 127.0.0.1
    Port = 5432
    Name = 
    TimeZone = Australia/Melbourne
    DataPath = ../data/
    ...
    ```

- build and run

    ``` bash
    cd server

    # use go mod And install the go dependency package
    go mod tidy

    # Compile 
    go build -o server main.go (windows the compile command is go build -o server.exe main.go )

    # Run binary
    ./server (windows The run command is server.exe)
    ```

- Project information and existing API

    ```
    [GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
    - using env:   export GIN_MODE=release
    - using code:  gin.SetMode(gin.ReleaseMode)

    [GIN-debug] POST   /api/v1/pomos/list        --> server/routers/v1.GetPomos (6 handlers)
    [GIN-debug] POST   /api/v1/pomos/create      --> server/routers/v1.CreatePomo (6 handlers)
    [GIN-debug] POST   /api/v1/pomos/count       --> server/routers/v1.CountPomos (6 handlers)
    [GIN-debug] POST   /api/v1/todos/list        --> server/routers/v1.GetTodos (6 handlers)
    [GIN-debug] POST   /api/v1/todos/create      --> server/routers/v1.CreateTodo (6 handlers)
    [GIN-debug] POST   /api/v1/todos/delete      --> server/routers/v1.DeleteTodo (6 handlers)
    [GIN-debug] POST   /api/v1/todos/update      --> server/routers/v1.UpdateTodo (6 handlers)

    Listening port is 8000
    ```

### 2.2 web project

- Config
    ```
    You can config in `.env.development` or `.env.production`, for server api.

    NODE_ENV=development
    VITE_APP_BASE_URL = 'http://127.0.0.1:8000'
    ```

- Project setup
    ```
    cd web
    yarn install
    ```

- Compiles and hot-reloads for development
    ```
    yarn serve
    ```

- Compiles and minifies for production
    ```
    yarn build
    ```

- Lints and fixes files
    ```
    yarn lint
    ```

- Customize configuration
See [Configuration Reference](https://cli.vuejs.org/config/).

# 3. Project Display

