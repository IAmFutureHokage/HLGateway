package dto

type Post struct {
	ID    string `json:"id"`
	Code  string `json:"code"`
	Name  string `json:"name"`
	River string `json:"river"`
}

type PostRequest struct {
	Post Post `json:"post"`
}

type DeletePostRequest struct {
	ID string `json:"id"`
}

type PostResponse struct {
	Post Post `json:"post"`
}

type GetPostsPageResponse struct {
	Posts         []Post `json:"posts"`
	PageNumber    uint32 `json:"page_number"`
	MaxPageNumber uint32 `json:"max_page_number"`
}

type PostsResponse struct {
	Posts []Post `json:"posts"`
}
