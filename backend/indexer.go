package backend

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	m "trademarkia.com/animeAPI/model"
)

//FindAnime to get Anime model from database or save to database then retrive, uses fiber context in order to get id parameter
func FindAnime(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	anime := m.Anime{
		ID:        id,
		Title:     "",
		Image:     "",
		Video:     "",
		Synopsis:  "",
		Score:     0,
		Rank:      0,
		Type:      "",
		Episode:   0,
		Status:    "",
		Broadcast: "",
		Source:    "",
		Duration:  "",
		Rating:    "",
	}
	if len(GetAnime(id).Title) != 0 {
		anime = GetAnime(id)
	} else {
		SaveAnime(id)
		anime = GetByID(id)
	}
	DetailsJSON, _ := json.MarshalIndent(anime, "", "\t")
	return c.SendString(strings.ReplaceAll(string(DetailsJSON), "\\u0026", "&"))
}

//AboutMe return details about myself
func AboutMe(c *fiber.Ctx) error {
	s := `Hello Trademarkia!
	My name is Balatamoghna and this is my submission for the job posting GoLang Developer. Please look at my files as well as my documentation to get a clear understand of the uses as well as project process I have taken to complete this project.
	Thank you!
	
	About me:
	College: VIT,
	Location: Chennai,
	Registration Number: 18BLC1142,
	Branch: Electronics and Computer Management ( A Part CSE Branch ),
	Contact Number: +91-9551076904
	`

	return c.SendString(s)
}

//Homepage returns the task's details
func Homepage(c *fiber.Ctx) error {
	s := `Project for Trademarkia.
	Details of the project:
	Imagine the website MyAnimeList.net does not have an API. Create an API that can take in an anime ID and return details about that anime. Eg: for "https://myanimelist.net/anime/14719", 14719 is the anime ID.

	a.	Hint: scraping is probably the right way to go about it

	b.	Put an emphasis on performance. Scraping is slow (and probably unavoidable); can you compensate in other ways?

	c.	You may use external libraries.

	d.	Bonus for unit tests.

	Please look at the /about page for personal details.
	Please look at the /anime/:id page to test the scraper. (Where :id is the variable, eg. 1,5114,14719,etc)
	`
	return c.SendString(s)
}
