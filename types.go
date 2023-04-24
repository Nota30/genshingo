package genshin

type User struct {
	Retcode int `json:"retcode"`
	Message string `json:"message"`
	Data data `json:"data"`
}

type data struct {
	Info info `json:"role"`
	Characters []character `json:"avatars"`
	Stats stats `json:"stats"`
	Exploration []exploration `json:"world_explorations"`
}

type exploration struct {
	Name string `json:"name"`
	Reputation int `json:"level"`
	Explored int `json:"exploration_percentage"`
	Icon string `json:"icon"`
	Offerings []offering `json:"offerings"`
}

type offering struct {
	Name string `json:"name"`
	Level int `json:"level"`
	Icon string `json:"icon"`
}

type stats struct {
	ActiveDays int `json:"active_day_number"`
	Achievements int `json:"achievement_number"`
	Anemoculus int `json:"anemoculus_number"`
	Geoculus int `json:"geoculus_number"`
	Electroculus int `json:"electroculus_number"`
	Dendroculus int `json:"dendroculus_number"`
	Characters int `json:"avatar_number"`
	Waypoints int `json:"way_point_number"`
	Domains int `json:"domain_number"`
	SpiralAbyss string `json:"spiral_abyss"`
	PreciousChests int `json:"precious_chest_number"`
	LuxuriousChests int `json:"luxurious_chest_number"`
	ExquisiteChests int `json:"exquisite_chest_number"`
	CommonChests int `json:"common_chest_number"`
	MagicChests int `json:"magic_chest_number"`
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