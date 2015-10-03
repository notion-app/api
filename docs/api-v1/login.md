
# `POST /v1/login`

Request Body (debug) (dev environment only)

Note: Your name becomes your unique ID in the database.

```
{
  authMethod: "debug",
  name: "Michael Hockerman"
}
```

Request Body (facebook)

```
{
  authMethod: "facebook",
  accessToken: "abc123",
  facebookId: "12345",
  expiresIn: "..."
}
```

Response (success)

Note: Returns (201) if user is created, (202) if user already exists

```
{
  userId: "unique notion id of the user",
  token: "...",
  name: "Michael Hockerman",
  profilePic: "http://..."
}
```
