package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"io/ioutil"
	"net/http"
	"strconv"
)

/*
TODO: Create a post listing model as a Go struct
This model has the following properties:

Description = text
Subcomments = # of Kids
*/

type ApiClient struct {
	BaseURI string
	Version string
	Suffix  string
}

func NewApiClient() ApiClient {
	return ApiClient{
		BaseURI: "https://hacker-news.firebaseio.com/",
		Version: "v0",
		Suffix:  ".json",
	}
}

// HN User
// See https://github.com/HackerNews/API#users for definition

// TODO: See if Martini has its own conventions around json
type User struct {
	ID        string `json:"id"`
	Delay     int    `json:"delay"`
	Created   int    `json:"created"`
	Karma     int    `json:"karma"`
	About     string `json:"about"`
	Submitted []int  `json:"submitted"`
}

// HN Story
type Story struct {
	By    string `json:"by"`
	ID    int    `json:"id"`
	Kids  []int  `json:"kids"`
	Score int    `json:"score"`
	Time  int    `json:"time"`
	Title string `json:"title"`
	Type  string `json:"type"`
	URL   string `json:"url"`
}

// TODO: Figure out Item struct definition
type Item struct {
}

func (client ApiClient) GetItem(id int) (Item, error) {
	url := client.BaseURI + client.Version + "/item/" + strconv.Itoa(id) + client.Suffix

	body, err := client.MakeHTTPRequest(url)
	if err != nil {
		return {}, err
	}

	return Item{}, nil
}

func (client ApiClient) GetUser(name string) (User, error) {
	// Attempt to get the user name using the HN api
	url := client.BaseURI + client.Version + "/user/" + name + client.Suffix

	var u User

	body, err := client.MakeHTTPRequest(url)
	if err != nil {
		return u, err
	}

	err = json.Unmarshal(body, &u)
	if err != nil {
		return u, err
	}

	return u, nil
}

func (client ApiClient) GetHiringPost() (Story, error) {
	// Retrieve an list of posts from the who-is-hiring bot and filter to find this months Who is Hiring post
	// Returns: A [data structure] of posts

	// Make call to API to look for the 'whoishiring' user
	u, err := client.GetUser("whoishiring")

	// Iterate the first 5 items belonging to the user
	// Make a request to the HN Api for each item to get all item details
	// Find the item title that contains 'Ask HN: Who is hiring?' in the title and the type of the item is a story
	fmt.Println("User details are: ")
	fmt.Println(u)

	// TODO: Learn how to do for each loops

	if err != nil {
		fmt.Println(err)
		fmt.Println(u)
	}

	// TODO: Return populated struct
	return Story{}, nil
}

// TODO: See if Martini has this function built in
func (client ApiClient) MakeHTTPRequest(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	if response.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf(http.StatusText(http.StatusNotFound))
	}
	return body, nil
}

func main() {
	server := martini.Classic()
	// Use martini contrib render middleware
	server.Use(render.Renderer())

	client := NewApiClient()

	server.Get("/listings", func(r render.Render) {
		// Return a formatted JSON object of 50 job listings along with their keywords for location, job type, etc
		post, err := client.GetHiringPost()

		fmt.Println(err)
		r.JSON(200, post)
	})

	server.Run()
}
