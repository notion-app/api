# `GET /v1/login`

Creates and logs in user accounts.

The access token returned must be provided as the url parameter `token` in all other requests.

Request (facebook)

```
{
  "auth_method": "facebook",
  "access_token": "12345"
}
```

Response (201) (User Created)

```
{
  "user_id": "..."
  "name": "..."
  "token": "..."
  "profile_pic": "..."
}
```

Response (202) (User Logged In)

```
same as 201
```
