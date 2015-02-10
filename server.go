package main

import "github.com/go-martini/martini"

func getHiringPost() {
  // Retrieve an list of posts from the who-is-hiring bot and filter to find this months Who is Hiring post
  // Returns: A [data structure] of posts
}

func main() {
  server := martini.Classic()
  server.Get("/listings", func() string {
    // Return a formatted JSON object of 50 job listings along with their keywords for location, job type, etc
    return "Index page"
  })
  server.Run()
}
