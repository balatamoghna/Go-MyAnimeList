# Go-MyAnimeListAPI
A GoLang scraper implementation for MyAnimeList API.

Created as a project submission task for the company Trademarkia. Completely written in GoLang and uses external libraries.

Libraries used:
---------------
* GORM - The fantastic ORM library for Golang
* Fiber - Express inspired web framework written in Go
* Colly - Elegant Scraper and Crawler Framework for Golang

Features:
---------------
Get details of the anime by inputting the MAL ID.
Details include:
* 	ID        
* 	Title    
* 	Image     
* 	Video    
* 	Synopsis  
* 	Score     
* 	Rank      
* 	Type    
* 	Episode  
* 	Status   
* 	Broadcast
* 	Source   
* 	Duration 
* 	Rating

_More can be added relatively easily_

Run project
---------------
Download the project and run this command in your terminal
```go run .\anime_api.go ```

Documentation
---------------
godocs can be used to generate the documentation shown below (albiet with lesser descriptiveness)
```godoc -http=:6060
http://localhost:6060/pkg/trademarkia.com/animeAPI/ 
```

### The project has two subfolders namely backend and model
* The backend consists of the scraper, object model relationship (along with driver to local MySQL database) and an indexer for ease of scalability.
* The model contains the Anime model used to represent the structure of the information. Also kept seperately to aid in ease of readablitly and scalability.

### Functions used:
* AboutMe - return details about myself (contact number, college, register number).
* Homepage - return the task's details (Also is the default root page defined in the router)
* FindAnime - get Anime model from database or save to database then retrive, uses fiber context in order to get id parameter.
* GetAnime - return Anime model from database.
* GetByID - return Anime model by given ID.
* InitialMigration - used to establish connection to the local MySQL database. (DNS value set to be const, if you are using your own server make sure to change the DNS).
* SaveAnime - Saves Anime to database. 

### Constants used:
* DNS alone is a constant.
```const DNS = "root:pass@tcp(127.0.0.1:3306)/godb?charset=utf8mb4&parseTime=True&loc=Local"```

### Anime Model
```
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
```

## Screenshots

![Running the project](https://user-images.githubusercontent.com/480968/131229224-bdb01065-495b-4a6a-9108-6502787886ea.png)

![The root page (invokes Homepage function)](https://user-images.githubusercontent.com/480968/131229375-4a891aec-bd17-408a-b94b-601ea2ac8c70.png)

![About me page (invokes AboutMe function)](https://user-images.githubusercontent.com/480968/131229394-e80fc19c-d3f9-4def-a302-4e7b8c03a31c.png)

![Anime details page (invokes FindAnime function)](https://user-images.githubusercontent.com/480968/131229404-5ef1498e-e774-4fb9-8d06-c860f150f7d7.png)

![Example ID given in question](https://user-images.githubusercontent.com/480968/131229439-29430d6d-20a4-47f0-9ade-4f8a4026c0be.png)


## Video
[![Watch the video here](https://img.youtube.com/vi/ZnL5h8xfhL4/0.jpg)](https://youtu.be/ZnL5h8xfhL4)

## Unit Tests
Generic tests have been included in the project in accordance to given bonus point.
To try the unit tests, run:
```go test -v```
### Some of the tests given are:
A test case to test the scrapper. (Includes 3 Test cases within, all 3 pass)
* Initial test with 1 test case alone took ~1.27s
* Second test with 3 test cases took ~3.02s
* Third test onwards took ~0.02s
(All tests were taken consecutively, the results were due to colly's caching. Thus focusing on performance, apart from caching there is also a local database keeping readable track of the data.)

A test case to test the request handler (fiber)
* Test case always give 0.00s. (It simply checks if test page is served)
