package genshin

type User struct {
	Retcode int `json:"retcode"`
	Message string `json:"message"`
	Data struct {
		Info info `json:"role"`
		Characters []character `json:"avatars"`
	} `json:"data"`
}

type info struct {
	AvatarUrl string `json:"AvatarUrl"`
	Nickname string `json:"nickname"`
	Region string `json:"region"`
	Level int `json:"level"`
}

type character struct {
	Id int `json:"id"`
	Image string `json:"image"`
	Name string `json:"name"`
	Element string `json:"element"`
	FriendLevel int `json:"fetter"`
	Level int `json:"level"`
	Rarity int `json:"rarity"`
	Constellation int `json:"actived_constellation_num"`
	CardImage string `json:"card_image"`
	InTeam bool `json:"is_chosen"`
}