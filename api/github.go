package api

import (
	"context"
	"log"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/andanhm/go-prettytime"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

var repos [23]string = [23]string{"parkify", "Libraryly", "SeatAndEat", "codepen-clone", "Leucos", "Taskify", "Codeaon", "C-Coin", "kitkat.v1rus", "CS-GO-Professionals", "learn-machine-learn", "OCR-TextRecognition", "MovieRecommendationSystem", "NASA_nearest_earth_object_classifier", "Resumie", "Health-Tracker-App", "Galleriz", "WallX", "kgec-summer-of-code", "DevBook", "cleanurge-mcu", "cleanurge-app", "cleanurge-backend"}

var startOfTimeForIssues time.Time = time.Date(2022, time.Month(7), 10, 0, 0, 0, 0, time.UTC)
var startOfTimeForPulls time.Time = time.Date(2022, time.Month(7), 26, 0, 0, 0, 0, time.UTC)

const (
	layout = "2006-01-02T15:04:05Z"
)

// GitHubAPI .. stores github api and context
type GitHubAPI struct {
	client *github.Client
	ctx    context.Context
	issues []Issue
	pulls  []User
}

// Issue struct
type Issue struct {
	Title      string `json:"title,omitempty"`
	URL        string `json:"url,omitempty"`
	Comments   int    `json:"comments,omitempty"`
	Repository string `json:"repository,omitempty"`
	UpdatedAt  string `json:"updated_at,omitempty"`
}

// Pull struct
type Pull struct {
	Title      string `json:"title,omitempty"`
	URL        string `json:"url,omitempty"`
	Repository string `json:"repository,omitempty"`
}

// User struct
type User struct {
	Login     string `json:"username,omitempty"`
	Name      string `json:"name,omitempty"`
	AvatarURL string `json:"avatar_url,omitempty"`
	HTMLURL   string `json:"url,omitempty"`
	Pulls     []Pull `json:"pulls,omitempty"`
	Points    uint64 `json:"points"`
}

// PointsSorter sorts users by points
type PointsSorter []User

func (a PointsSorter) Len() int           { return len(a) }
func (a PointsSorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a PointsSorter) Less(i, j int) bool { return a[i].Points > a[j].Points }

// Init ...
func (g *GitHubAPI) Init() {
	g.ctx = context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("AUTH_TOKEN")},
	)
	tc := oauth2.NewClient(g.ctx, ts)
	g.client = github.NewClient(tc)
	g.FetchIssueStats()
	g.FetchPullStats()
}

// FetchIssueStats ...
func (g *GitHubAPI) FetchIssueStats() {
	count := 0
	for index, repo := range repos {
		index = index + 1
		issues, _, err := g.client.Issues.ListByRepo(g.ctx, "dsckgec", repo, &github.IssueListByRepoOptions{State: "open", Since: startOfTimeForIssues})
		if err != nil {
			continue
		}
		for _, issue := range issues {
			if issue.PullRequestLinks != nil {
				continue
			}
			count++
			t, err := time.Parse(layout, (*issue.UpdatedAt).Format(layout))
			if err != nil {
				log.Fatal(err)
			}
			newIssue := Issue{*issue.Title, *issue.HTMLURL, *issue.Comments, repo, prettytime.Format(t)}
			g.issues = append(g.issues, newIssue)
			if count > index*2 {
				break
			}
		}
	}
}

// FetchPullStats ...
func (g *GitHubAPI) FetchPullStats() {
	for _, repo := range repos {
		pulls, _, err := g.client.PullRequests.List(g.ctx, "dsckgec", repo, &github.PullRequestListOptions{State: "all"})
		if err != nil {
			continue
		}

		for _, pull := range pulls {
			if (*pull.CreatedAt).Before(startOfTimeForPulls) || pull.MergedAt == nil {
				continue
			}
			newPull := Pull{*pull.Title, *pull.HTMLURL, repo}
			flag := 0
			for i := 0; i < len(g.pulls); i++ {
				user := &g.pulls[i]
				if user.Login == *pull.User.Login {
					flag = 1
					if len(pull.Labels) > 0 {
						for j := 0; j < len(pull.Labels); j++ {
							if strings.EqualFold(*pull.Labels[j].Name, "Easy") {
								user.Points += 100
								user.addPull(newPull)
								break
							} else if strings.EqualFold(*pull.Labels[j].Name, "Medium") {
								user.Points += 200
								user.addPull(newPull)
								break
							} else if strings.EqualFold(*pull.Labels[j].Name, "Hard") {
								user.Points += 300
								user.addPull(newPull)
								break
							}
						}
					}
				}
			}
			if flag == 0 {
				newUser := User{}
				newUser.Login = *pull.User.Login
				if pull.User.GetName() == "" {
					newUser.Name = newUser.Login
				} else {
					newUser.Name = pull.User.GetName()
				}
				newUser.AvatarURL = *pull.User.AvatarURL
				newUser.HTMLURL = *pull.User.HTMLURL
				newUser.Pulls = newUser.addPull(newPull)
				newUser.Points = 0
				if len(pull.Labels) > 0 {
					for j := 0; j < len(pull.Labels); j++ {

						if strings.EqualFold(*pull.Labels[j].Name, "Easy") {
							newUser.Points += 100
							break
						} else if strings.EqualFold(*pull.Labels[j].Name, "Medium") {
							newUser.Points += 200
							break
						} else if strings.EqualFold(*pull.Labels[j].Name, "Hard") {
							newUser.Points += 300
							break
						}
					}
				}
				if newUser.Points != 0 {
					g.pulls = append(g.pulls, newUser)
				}
			}

		}
	}
	sort.Sort(PointsSorter(g.pulls))
}

func (user *User) addPull(pull Pull) []Pull {
	if len(user.Pulls) > 2 {
		return user.Pulls
	}
	user.Pulls = append(user.Pulls, pull)
	return user.Pulls
}
