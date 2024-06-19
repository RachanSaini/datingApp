# Dating App API

## Setup & Run

### Prerequisites

- Docker
- Docker Compose
- MySql

### Environment Variables

Create a `.env` file in the root of the project with the following content:
```
DB_HOST=localhost
DB_PORT=3306
DB_USER=<mysql_user>
DB_PASSWORD=<mysql_passwors>
DB_NAME=<database_name>
SECRET_KEY=<your_secret_key>
```

### Running the Application

1. Clone the repository

2. Set the environment variables

3. Run the go application

```
go run main.go

```

4. Register new users.

```
curl --location 'http://localhost:8080/users/register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "ginny@example.com",
    "password": "ginny",
    "name": "Ginny Weasley",
    "gender": "Female",
    "dob": "1995-01-01T00:00:00Z",
    "location": "Hogwarts"
}'

```

You will see response as

```
{
    "user": {
        "ID": 23,
        "Email": "ginny@example.com",
        "Password": "",
        "Name": "Ginny Weasley",
        "Gender": "Female",
        "DOB": "1995-01-01T00:00:00Z",
        "Location": "71",
        "CreatedAt": "2024-06-19T11:55:49.64875+01:00",
        "UpdatedAt": "2024-06-19T11:55:49.64875+01:00"
    }
}
```

5. Login and copy the generated token as this is needed for authentication.

```
curl --location 'http://localhost:8080/auth/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "ginny@example.com",
    "password": "ginny"
}'

```

You will get token as response, don't forget to copy this token.

```
{
    "token": "eyJhbGciOi........RGQgAt3c"
}

```

6. Discover all other profiles.

```
curl --location 'http://localhost:8080/discover' \
--header 'Authorization: Bearer <token>'
```
Response would be other registered users:

```
{
    "results": [
        {
            "age": 34,
            "gender": "male",
            "id": 1,
            "name": "Iron Man"
        },
        {
            "age": 29,
            "gender": "Female",
            "id": 21,
            "name": "Redar D"
        },
        {
            "age": 29,
            "gender": "Male",
            "id": 22,
            "name": "Harry Potter"
        }
    ]
}

```

7. Swipe on desired profile with user id and choice. Choice can be YES or NO
```
curl --location 'http://localhost:8080/discover/swipe' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer <token> '{
    "otherUserId": 22,
    "choice": "YES"
}'
```

if it is not a match you will get matched "false", if its a match you will get:

```
{
    "results": {
        "matchID": 1,
        "matched": true
    }
}
```



