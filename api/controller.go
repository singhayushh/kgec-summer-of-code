package api

import (
	"github.com/gin-gonic/gin"
)

type Repo struct {
	Name        string   `json:"name,omitempty"`
	Url         string   `json:"url,omitempty"`
	Description string   `json:"description,omitempty"`
	Tags        []string `json:"tags,omitempty"`
}

var projects = []Repo{
	{
		Name:        "Parkify",
		Url:         "​​https://github.com/DSCKGEC/parkify",
		Description: "Hassle-free way to book your parking space with easy cancellations and timeline extensions",
		Tags:        []string{},
	},
	{
		Name:        "Libraryly",
		Url:         "​​https://github.com/DSCKGEC/Libraryly",
		Description: "A software solution to handle the primary functions of a library like managing books, members and issues",
		Tags:        []string{},
	},
	{
		Name:        "Seat&Eat",
		Url:         "​​https://github.com/DSCKGEC/SeatAndEat",
		Description: "This is a beautifully designed commercial website template for a restaurant. This project is a beginner-friendly project. It is a static website and mostly based on the front end.",
		Tags:        []string{},
	},
	{
		Name:        "DevBook",
		Url:         "​​https://github.com/DSCKGEC/DevBook",
		Description: "DevBook is a social development environment for front-end designers and developers. A project editor that supports HTML, CSS, and JavaScript, where you can show off your work, build test cases to learn and debug and find inspiration",
		Tags:        []string{},
	},
	{
		Name:        "Leucos",
		Url:         "​​https://github.com/DSCKGEC/Leucos",
		Description: "Real-time chatting application within private rooms, from anywhere in the world",
		Tags:        []string{},
	},
	{
		Name:        "Taskify",
		Url:         "​​https://github.com/DSCKGEC/Taskify",
		Description: "It is a task management system for everyone. It is designed to help you manage your tasks and projects from ideation to delivery. This task manager helps to bring in only the necessary parts – without all the annoying clutter.",
		Tags:        []string{},
	},
	{
		Name:        "Codeaon",
		Url:         "​​https://github.com/DSCKGEC/Codeaon",
		Description: "A one-stop website for all developers and coders. This platform helps us to get all YouTube tutorials, blogs and articles for a certain topic in one place. This project proves to be a very handy and effective yet resourceful tool for learners.",
		Tags:        []string{},
	},
	{
		Name:        "C-Coins",
		Url:         "​​https://github.com/DSCKGEC/C-Coins",
		Description: "Defining the basic behaviour of Blockchains and Web-2.0 imitation",
		Tags:        []string{},
	},
	{
		Name:        "kitkat.virus",
		Url:         "​​https://github.com/DSCKGEC/kitkat.virus",
		Description: "A simple python based keylogger",
		Tags:        []string{},
	},
	{
		Name:        "CSGO-Professional",
		Url:         "​​https://github.com/DSCKGEC/CSGO-Professional",
		Description: "Scraping, Cleaning and Exploring a Dataset about CSGO Athletes from hltv.org",
		Tags:        []string{},
	},
	{
		Name:        "Learn-Machine-Learn",
		Url:         "​​https://github.com/DSCKGEC/Learn-Machine-Learn",
		Description: "A machine learning repository to demonstrate regression and classification problems",
		Tags:        []string{},
	},
	{
		Name:        "Text Recognition",
		Url:         "​​https://github.com/DSCKGEC/OCR-TextRecognition",
		Description: "Recognizes text from any image and prints them",
		Tags:        []string{},
	},
	{
		Name:        "Movie Recommendation",
		Url:         "​​https://github.com/DSCKGEC/MovieRecommendationSystem",
		Description: "A Movie Recommendation System along with Data Analysis and Data Visualization and Revenue Prediction Model",
		Tags:        []string{},
	},
	{
		Name:        "NASA Collision Detection",
		Url:         "​​https://github.com/DSCKGEC/NASA_nearest_earth_object_classifier",
		Description: "Classifying hazardous asteroids based on NASA Dataset.",
		Tags:        []string{},
	},
	{
		Name:        "Resumie",
		Url:         "​​https://github.com/DSCKGEC/Resumie",
		Description: "Resumie is an android CV & Portfolio app. Now carry your resume with you wherever you go!",
		Tags:        []string{},
	},
	{
		Name:        "Health Tracker App",
		Url:         "​https://github.com/DSCKGEC/Health-Tracker-App",
		Description: "Health Tracker is your one stop solution to keep your heath related information in a secured way",
		Tags:        []string{},
	},
	{
		Name:        "Samsung Gallery",
		Url:         "​​https://github.com/DSCKGEC/samsung-gallery-clone",
		Description: "Clone of Samsung's Default Gallery",
		Tags:        []string{},
	},
	{
		Name:        "Wallpaper App",
		Url:         "​​https://github.com/DSCKGEC/flutter-wallx-wallpaperApp",
		Description: "",
		Tags:        []string{},
	},
}

// RenderHome ...
func (g *GitHubAPI) RenderHome(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{
		"title":  "KGEC Summer of Code 2021",
		"isHome": true,
	})
}

// RenderDashboard ...
func (g *GitHubAPI) RenderDashboard(c *gin.Context) {
	c.HTML(200, "dashboard.html", gin.H{
		"title":  "Dashboard | KGEC Summer of Code 2021",
		"issues": g.issues,
		"pulls":  g.pulls,
	})
}

// RenderLeaderboard ...
func (g *GitHubAPI) RenderLeaderboard(c *gin.Context) {
	c.HTML(200, "leaderboard.html", gin.H{
		"title":  "Leaderboard | KGEC Summer of Code 2021",
		"issues": g.issues,
		"pulls":  g.pulls,
	})
}

func (g *GitHubAPI) GetData(c *gin.Context) {
	c.JSON(200, gin.H{
		"pulls":  g.pulls,
		"issues": g.issues,
	})
}

// Refresh ...
func (g *GitHubAPI) Refresh(c *gin.Context) {
	g.issues = g.issues[:0]
	g.pulls = g.pulls[:0]

	g.FetchIssueStats()
	g.FetchPullStats()
	c.JSON(200, gin.H{
		"message": "done",
	})
}

// RenderProject ...
func (g *GitHubAPI) RenderProject(c *gin.Context) {

	c.HTML(200, "projects.html", gin.H{
		"title":  "Projects",
		"button": "View Leaderboard",
		"repos":  projects,
	})
}

// RenderAbout ...
func (g *GitHubAPI) RenderAbout(c *gin.Context) {
	c.HTML(200, "about.html", gin.H{
		"title": "About",
	})
}
