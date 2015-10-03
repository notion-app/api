// Direct mapping to the user model in the database
package model

type User struct {
  Id string `db:"id,primarykey"`
  Name string `db:"name"`
  Role string `db:"role"`
  Verified bool `db:"verified"`
  AuthMethod string `db:"auth_method"`
  FbUserId string `db:"fb_user_id"`
  FbAuthToken string `db:"fb_auth_token"`
  FbExpiresIn string `db:"fb_expires_in"`
  FbProfilePic string `db:"fb_profile_pic"`
}
