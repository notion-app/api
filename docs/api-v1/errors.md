# Errors

All endpoints have the capability of throwing multiple different types of Errors.
This page documents errors which are general to most, if not all, endpoints.
Specific endpoints might be documented to throw other errors.

# General Format

```
{
  "status": 400,
  "message": "This is a human readable message for the developer of the API"
}
```

# Bad Request `400`

Happens when the request body is malformed in some way.

```
{
  "status": 400,
  "message": "Request body malformed. Must include a userId field."
}
```

# Unauthorized `401`

Happens when an access token provided is not valid for the user.
Can also be thrown if an external access token provided is no longer valid (like facebook).

```
{
  "status": 401,
  "message": "The notion access token provided is not valid",
  "service": "notion"
}
```

```
{
  "status": 401,
  "message": "The facebook access token provided is not valid.",
  "service": "facebook"
}
```

# ISE `500`

```
{
  "status": 500,
  "message": "Internal server error"
}
```
