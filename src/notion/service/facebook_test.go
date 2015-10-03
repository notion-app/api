
package service

import (
  "os"
  "testing"
)

func TestFacebookGetCurrentUserValid(t *testing.T) {
  valid, user, err := Facebook{}.GetCurrentUser(os.Getenv("FACEBOOK_AUTH"))
  if err != nil {
    t.Fatal(err)
  }
  if !valid {
    t.Fatal("Access token provided under $FACEBOOK_AUTH is not valid; maybe you need to update it?")
  }
  if len(user.Id) == 0 || len(user.Name) == 0 {
    t.Fatal("No id or name returned by facebook")
  }
}

func TestFacebookGEtCurrentUserInvalid(t *testing.T) {
  valid, _, err := Facebook{}.GetCurrentUser("here's an invalid code")
  if err != nil {
    t.Fatal(err)
  }
  if valid {
    t.Fatal("Method returned valid when it should be returning invalid")
  }
}
