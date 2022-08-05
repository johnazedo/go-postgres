# Gopostgres
A simple way to use migrations and postgres with go.

## Usage

Create this fields in your .env file:
```shell
POSTGRES_USER=your_user
POSTGRES_PASSWORD=your_password
POSTGRES_DB=your_database_name
POSTGRES_HOST=your_host
POSTGRES_PORT=your_port
```

After that call the `GetDatabase` function in your main function:
```go
import "github.com/johnazedo/gopostgres"

func main() {
    db, err := gopostgres.GetDatabase()
    if err != nil {
        panic(err)
    }
}
```