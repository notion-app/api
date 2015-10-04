// Direct mapping to the user model in the database
package model

type User struct {
  Id string `db:"id,primarykey" json:"id"`
  Name string `db:"name" json:"name"`
  Role string `db:"role" json:"role"`
  Verified bool `db:"verified" json:"verified"`
  AuthMethod string `db:"auth_method" json:"auth_method"`
  FbUserId string `db:"fb_user_id" json:"fb_user_id"`
  FbAuthToken string `db:"fb_auth_token" json:"fb_auth_token"`
  FbExpiresIn string `db:"fb_expires_in" json:"fb_expires_in"`
  FbProfilePic string `db:"fb_profile_pic" json:"fb_profile_pic"`
}
