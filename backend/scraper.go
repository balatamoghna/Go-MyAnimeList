package backend

import (
	"strconv"
	"strings"

	"github.com/gocolly/colly"

	m "trademarkia.com/animeAPI/model"
)

//GetByID function used to get Anime model by given ID
func GetByID(id int) m.Anime {
	details := m.Anime{
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
	c := colly.NewCollector(
		colly.AllowedDomains("myanimelist.net"),

		//Caching done along with database ensuring no data is re-scraped unnecessarily
		colly.CacheDir("./animelist_cache"),
	)

	c.OnHTML("h1", func(e *colly.HTMLElement) {
		details.Title = e.Text
	})

	c.OnHTML("#content", func(e *colly.HTMLElement) {
		details.Image = e.ChildAttr("img", "data-src")
		e.ForEach("a", func(_ int, v *colly.HTMLElement) {
			if v.Attr("class") == "iframe js-fancybox-video video-unit promotion" {
				details.Video = v.Attr("href")
			}
		})
		e.ForEach("p", func(i int, h *colly.HTMLElement) {
			if h.Attr("itemprop") == "description" {
				details.Synopsis = h.Text
			}
		})
		details.Score, _ = strconv.ParseFloat(e.ChildText("div.score-label"), 64)

		e.ForEach("div.spaceit", func(i int, h *colly.HTMLElement) {
			if h.ChildText("span") == "Episodes:" {
				details.Episode, _ = strconv.Atoi(strings.Split(h.Text, "\n")[2])
			}
			if h.ChildText("span") == "Type:" {
				details.Type = strings.Split(h.Text, "\n")[2]
			}
			if h.ChildText("span") == "Status:" {
				details.Status = strings.Split(h.Text, "\n")[2]
			}
			if h.ChildText("span") == "Source:" {
				details.Source = strings.Split(h.Text, "\n")[2]
			}
			if h.ChildText("span") == "Duration:" {
				details.Duration = strings.Split(h.Text, "\n")[2]
			}
			if h.ChildText("span") == "Rating:" {
				details.Rating = strings.Split(h.Text, "\n")[2]
			}
			if h.ChildText("span") == "Ranked:" {
				s := strings.Split(h.Text, "\n")[2]
				s = s[3 : len(s)-1]
				details.Rank, _ = strconv.Atoi(s)
			}
			if h.ChildText("span") == "Broadcast:" {
				details.Broadcast = strings.Split(h.Text, "\n")[2]
			}

		})
	})

	c.Visit("https://myanimelist.net/anime/" + strconv.Itoa(id))
	return details
}
