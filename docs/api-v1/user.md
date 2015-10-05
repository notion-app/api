# `GET /v1/me`

Returns information about the requesting user.

Response (200)

```
{
  id: "...",
  name: "...",
  verified: false,
  auth_method: "facebook",
  fb_user_id: "...",
  fb_profile_pic: "..."
}
```

# `GET /v1/user/:user_id`

Returns information about a given user

Right now, this will only return a valid response if the user id being requested is the user Id of the requesting user.

Response (200)

```
Same as GET /v1/me
```

# `DELETE /v1/user/:user_id`

Purges a user's account from the notion service. Deletes all associated notebooks and notes.
