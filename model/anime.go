package model

//Anime details model
type Anime struct {
	ID        int     `json:"id"`
	Title     string  `json:"title"`
	Image     string  `json:"image"`
	Video     string  `json:"video"`
	Synopsis  string  `json:"synopsis"`
	Score     float64 `json:"score"`
	Rank      int     `json:"rank"`
	Type      string  `json:"type"`
	Episode   int     `json:"episode"`
	Status    string  `json:"status"`
	Broadcast string  `json:"broadcast"`
	Source    string  `json:"source"`
	Duration  string  `json:"duration"`
	Rating    string  `json:"rating"`
}
