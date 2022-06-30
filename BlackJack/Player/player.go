package player

type Player struct {
	Name  string
	Score int
}

func CreateNewPlayer(Name string) *Player {
	var newPlayer Player
	newPlayer.Name = Name
	newPlayer.Score = 0
	return &newPlayer
}
