# Toxic REST API

A REST API that only demonstrates the concept of JWT authentication. 🦠

## How to run?

1. Firstly, install [Docker](https://www.docker.com/).
2. Then run `sh run.sh` and enjoy.

## HTTP API

```
POST /api/v1/login
```

Login supports only 'root' as both the username and password. If you attempt to use any other username or password,
it will respond: `Get out of here with the wrong credentials!!!`.
The response body contains a token that you can use to authenticate your requests."

```
GET /api/v1/whoami
```

Retrieve your username and password data.
If you do not include the bearer token as an authorization header, it will respond with:
`Get out of here without a bearer token!!!`.
