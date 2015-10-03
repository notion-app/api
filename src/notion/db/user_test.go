
package db

import (
  "notion/model"
  "os"
  "testing"
)

func TestMain(m *testing.M) {
  // Insert some test data
  Init()
  dbmap.Insert(&model.User{
    Id: "1",
    Name: "Michael Hockerman",
    Role: "student",
    Verified: false,
    AuthMethod: "facebook",
    FbUserId: "1",
    FbAuthToken: "1",
    FbExpiresIn: "1",
    FbProfilePic: "1",
  })
  // Run the tests
  os.Exit(m.Run())
  // Clean up the test data
}
