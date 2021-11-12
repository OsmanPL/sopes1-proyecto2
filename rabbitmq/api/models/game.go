package models

type Game struct {
	ID       int    `json:"id"`
	GameName string `json:"gameName"`
	Winner   string `json:"winner"`
	Players  int    `json:"players"`
	Worker   string `json:"worker"`
}
