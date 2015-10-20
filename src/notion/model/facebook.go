// Structs which match the API responses from facebook

package model

type FbCurrentUser struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type FbProfilePic struct {
	Data struct {
		IsSet bool   `json:"is_silhouette"`
		Url   string `json:"url"`
	} `json:"data"`
}
