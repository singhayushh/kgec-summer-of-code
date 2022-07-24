package api

import (
	"html"

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
		Url:         html.EscapeString("https://github.com/DSCKGEC/parkify"),
		Description: "Hassle-free way to book your parking space with easy cancellations and timeline extensions.",
		Tags:        []string{"html", "css", "javascript", "node-js", "express", "mongodb", "bootstrap"},
	},
	{
		Name:        "Libraryly",
		Url:         html.EscapeString("https://github.com/DSCKGEC/Libraryly"),
		Description: "A software solution to handle the primary functions of a library like managing books, members and issues.",
		Tags:        []string{"html", "css", "javascript", "node-js", "express", "mongodb"},
	},
	{
		Name:        "Seat&Eat",
		Url:         html.EscapeString("https://github.com/DSCKGEC/SeatAndEat"),
		Description: "SeatAndEat is a beautifully designed commercial website template for a restaurant.",
		Tags:        []string{"html", "css", "javascript"},
	},
	{
		Name:        "DevBook",
		Url:         html.EscapeString("https://github.com/DSCKGEC/DevBook"),
		Description: "DevBook is a development environment for front-end designers and developers that supports HTML, CSS, and JavaScript.",
		Tags:        []string{"react", "html", "css", "javascript"},
	},
	{
		Name:        "Leucos",
		Url:         html.EscapeString("https://github.com/DSCKGEC/Leucos"),
		Description: "Real-time chatting application within private rooms, from anywhere in the world.",
		Tags:        []string{"html", "css", "javascript", "node-js", "express", "mongodb"},
	},
	{
		Name:        "Taskify",
		Url:         html.EscapeString("https://github.com/DSCKGEC/Taskify"),
		Description: "Taskify is a task management system for everyone to help you manage your tasks and projects from ideation to delivery.",
		Tags:        []string{"html", "css", "javascript", "node-js", "express", "mongodb"},
	},
	{
		Name:        "Codeaon",
		Url:         html.EscapeString("https://github.com/DSCKGEC/Codeaon"),
		Description: "A one-stop website for all programmers that helps us to get all YouTube tutorials and articles for a certain topic in one place.",
		Tags:        []string{"html", "css", "javascript", "node-js", "express", "mongodb"},
	},
	{
		Name:        "C-Coins",
		Url:         html.EscapeString("https://github.com/DSCKGEC/C-Coin"),
		Description: "Defining the basic behaviour of Blockchains and Web-2.0 imitation.",
		Tags:        []string{"python"},
	},
	{
		Name:        "kitkat.virus",
		Url:         html.EscapeString("https://github.com/DSCKGEC/kitkat.v1rus"),
		Description: "A simple python based keylogger.",
		Tags:        []string{"python"},
	},
	{
		Name:        "CSGO-Professional",
		Url:         html.EscapeString("https://github.com/DSCKGEC/CS-GO-Professionals"),
		Description: "Scraping, Cleaning and Exploring a Dataset about CSGO Athletes from hltv.org",
		Tags:        []string{"jupyter notebook"},
	},
	{
		Name:        "Learn-Machine-Learn",
		Url:         html.EscapeString("https://github.com/DSCKGEC/Learn-Machine-Learn"),
		Description: "A machine learning repository to demonstrate regression and classification problems.",
		Tags:        []string{"jupyter notebook"},
	},
	{
		Name:        "Text Recognition",
		Url:         html.EscapeString("https://github.com/DSCKGEC/OCR-TextRecognition"),
		Description: "Recognizes text from any image and prints them.",
		Tags:        []string{"python"},
	},
	{
		Name:        "Movie Recommendation",
		Url:         html.EscapeString("https://github.com/DSCKGEC/MovieRecommendationSystem"),
		Description: "A Movie Recommendation System along with Data Analysis and Data Visualization and Revenue Prediction Model.",
		Tags:        []string{"jupyter notebook"},
	},
	{
		Name:        "NASA Collision Detection",
		Url:         html.EscapeString("https://github.com/DSCKGEC/NASA_nearest_earth_object_classifier"),
		Description: "Classifying hazardous asteroids based on NASA Dataset.",
		Tags:        []string{"jupyter notebook"},
	},
	{
		Name:        "Resumie",
		Url:         html.EscapeString("https://github.com/DSCKGEC/Resumie"),
		Description: "Resumie is an android CV & Portfolio app. Now carry your resume with you wherever you go!",
		Tags:        []string{"java"},
	},
	{
		Name:        "Health Tracker App",
		Url:         html.EscapeString("https://github.com/DSCKGEC/Health-Tracker-App"),
		Description: "Health Tracker is your one stop solution to keep your heath related information in a secured way.",
		Tags:        []string{"dart", "c++", "cmake", "html", "swift", "c"},
	},
	{
		Name:        "Galleriz",
		Url:         html.EscapeString("https://github.com/DSCKGEC/Galleriz"),
		Description: "A beautiful and functional gallery app for smartphones.",
		Tags:        []string{"javascript", "java", "objective-c", "ruby", "starlark"},
	},
	{
		Name:        "WallX",
		Url:         html.EscapeString("https://github.com/DSCKGEC/WallX"),
		Description: "A cool wallpaper app for smartphones.",
		Tags:        []string{"dart", "html", "swift"},
	},
}

// RenderHome ...
func (g *GitHubAPI) RenderHome(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{
		"title":  "KGEC Summer of Code 2022",
		"isHome": true,
	})
}

// RenderDashboard ...
func (g *GitHubAPI) RenderDashboard(c *gin.Context) {
	c.HTML(200, "dashboard.html", gin.H{
		"title":  "Dashboard | KGEC Summer of Code 2022",
		"issues": g.issues,
		"pulls":  g.pulls,
	})
}

// RenderLeaderboard ...
func (g *GitHubAPI) RenderLeaderboard(c *gin.Context) {
	c.HTML(200, "leaderboard.html", gin.H{
		"title":  "Leaderboard | KGEC Summer of Code 2022",
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
