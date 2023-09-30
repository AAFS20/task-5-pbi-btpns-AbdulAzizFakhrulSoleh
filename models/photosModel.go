package models

type Photos struct {
	Idp       *int    `json:"id_photo"`
	Title     *string `json:"title"`
	Photo_url string  `json:"photoUrl"`
	Caption   string  `json:"caption"`
	User_id   string  `json:"userid"`
}
