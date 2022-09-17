package service

type GroupingRequest struct {
	GroupNum     int           `json:"group_num"`
	UsersChoices []UsersChoice `json:"users_choices"`
}
type UsersChoice struct {
	UserID  int   `json:"user_id"`
	Choices []int `json:"choices"`
}

func GroupingService() {

}
