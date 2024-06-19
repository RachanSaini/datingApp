# Dating App API

## Setup & Run

### Prerequisites

- MySql

### Running the Application

1. Clone the repository
```
git clone <repository_url>
cd <repository_directory>
```

2. Set the environment variables. Create a `.env` file in the root of the project with the following content:

```
DB_HOST=localhost
DB_PORT=3306
DB_USER=<mysql_user>
DB_PASSWORD=<mysql_passwors>
DB_NAME=<database_name>
SECRET_KEY=<your_secret_key>
```

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

Example response:

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

5. Login and obtain the authentication token.

```
curl --location 'http://localhost:8080/auth/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "ginny@example.com",
    "password": "ginny"
}'

```

Example response:

```
{
    "token": "eyJhbGciOi........RGQgAt3c"
}

```
Copy the token as it will be needed for authentication in subsequent requests.

6. Discover all other profiles.

```
curl --location 'http://localhost:8080/discover' \
--header 'Authorization: Bearer <token>'
```
Example response:

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

7. Swipe on desired profile with user ID and choice (YES or NO).

```
curl --location 'http://localhost:8080/discover/swipe' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer <token> '{
    "otherUserId": 22,
    "choice": "YES"
}'
```

Example response if it's not a match:

```
{
    "results": {
        "matchID": 1,
        "matched": false
    }
}
```

Example response if it's a match:

```
{
    "results": {
        "matchID": 1,
        "matched": true
    }
}
```