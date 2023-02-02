package main

import (
	"encoding/json"
    "fmt"
    "io"
    "net/http"
    "net/url"
    "os"
    "strings"
    "time"

    "github.com/gofiber/fiber/v2"
    
    "github.com/19chonm/461_1_23/cli/commands"
    // "github.com/gofiber/fiber/v2/middleware/cors"
    
    // "github.com/19chonm/461_1_23/database"
    // "github.com/19chonm/461_1_23/url"
)
// const serverPort = 4000


type Response interface {
    Validate() bool
}

type LicenseResponse struct {
    License struct {
        Key *string  `json:"key"`
        Name *string `json:"name"`
        Url *string  `json:"url"`
    } `json:"license"`
}

type IssueResponse struct {
    CreatedAt *string `json:"created_at"`
    ClosedAt *string  `json:"closed_at"`
    PullRequest *struct {} `json:"pull_request"`
}

type ContributorStatsResponse[] struct {
    Author struct {
        Login *string `json:"login"`
    } `json:"author"`
    Weeks []struct {
        Week *int64 `json:"w"`
        Commits *int `json:"c"`
    } `json:"weeks"`
}

func (self LicenseResponse) Validate() bool {
    return self.License.Key != nil && self.License.Name != nil && self.License.Url != nil;
}

func (self IssueResponse) Validate() bool {
    return self.CreatedAt != nil && self.ClosedAt != nil
}

func (self ContributorStatsResponse) Validate() bool {
    for _, contributor := range self {
        if contributor.Author.Login == nil {
            return false
        }
        for _, week := range contributor.Weeks {
            if week.Week == nil || week.Commits == nil {
                return false
            }
        }
    }
    return true
}

/* API RESPONSE TYPES */

type Responsiveness struct {
    AvgLifespan float64 `json:"avg_lifespan"`
    NumSampled int `json:"num_sampled"`
}

type Contributor struct {
    Name string `json:"name"`
    RecentCommits int `json:"recent_commits"`
}

// Parse required headers and validate them
func processHeaders(c *fiber.Ctx) (user string, repo string, token string, err error) {
    // get headers
    inputUrl := c.Get("InputURL")
    token = c.Get("GitHubToken")
    // example command 
    /*
        curl -H "InputURL: https://github.com/qiangxue/go-rest-api" -H "GitHubToken: <YOURTOKEN>" localhost:3000/api/license

        // Should RETURN:{"license":{"key":"mit","name":"MIT License","url":"https://api.github.com/licenses/mit"}}%                             
    */
    // validate URL
    if inputUrl == "" {
        err = fmt.Errorf("InputURL not provided")
        return
    }
    urlObject, err := url.Parse(inputUrl)
    if err != nil {
        return
    }
    if urlObject.Host != "github.com" {
        err = fmt.Errorf("InputURL %s is not a GitHub URL", urlObject)
        return
    }
    path := strings.Split(urlObject.EscapedPath(), "/")[1:]
    if len(path) != 2 {
        err = fmt.Errorf("InputURL %s does not point to a GitHub repository", urlObject)
        return
    }
    user, repo = path[0], path[1]

    // validate token
    if token == "" {
        err = fmt.Errorf("GitHubToken not provided")
        return
    }
    return
}

// Build and a request to the given endpoint; return HTTP response
func sendGithubRequestHelper(endpoint string, token string) (res *http.Response, err error, statusCode int) {
    // build GitHub API request
    req, _ := http.NewRequest(http.MethodGet, endpoint, nil)
    req.Header.Add("Accept", "application/vnd.github+json")
    req.Header.Add("Authorization", "Bearer " + token)
    req.Header.Add("X-GitHub-Api-Version", "2022-11-28")

    res, err = http.DefaultClient.Do(req)

    if err != nil {
        err = fmt.Errorf("Failed to send HTTP request")
        statusCode = 500 // Internal server error
    } else if res.StatusCode != 200 {
        statusCode = res.StatusCode // forward API error code
        buf := make([]byte, 1000)
        n, _ := res.Body.Read(buf)
        fmt.Printf("%s\n", buf[:n])
        err = fmt.Errorf("GitHub request responded with error code %d", statusCode)
    }
    return
}

// Decode HTTP response using JSON decoder
func decodeResponse[T any](res *http.Response) (jsonRes T, err error) {
    decoder := json.NewDecoder(res.Body)
    for {
        err = decoder.Decode(&jsonRes)
        if err == io.EOF {
            err = nil
            return
        } else if err != nil {
            return
        }
    }
}

// Set a query parameter on an HTTP request
func setQueryParameter(endpoint *string, parameter string, value string) (err error) {
    var urlObject *url.URL 
    urlObject, err = url.Parse(*endpoint)
    if err != nil {
        return
    }
    query := urlObject.Query()
    query.Set(parameter, value)
    urlObject.RawQuery = query.Encode()
    *endpoint = urlObject.String()
    return
}

// Send GitHub API request and return response of type T
func sendGithubRequest[T Response](endpoint string, token string) (jsonRes T, err error, statusCode int) {
    res, err, statusCode := sendGithubRequestHelper(endpoint, token)
    if err != nil {
        return
    }

    jsonRes, err = decodeResponse[T](res) 

    if !jsonRes.Validate() {
        err = fmt.Errorf("Failed to parse GitHub response")
        statusCode = 500 // Internal server error
        return
    }

    // assert that there is no pagination
    linkHeader := res.Header.Get("link")
    if (linkHeader != "") {
        err = fmt.Errorf("Did not expect pagination")
        statusCode = 500 // Internal server error
        return
    }

    return // success
}

// Send GitHub API request and return response of type T
// Follows pages, up to maxPages
func sendGithubRequestList[T Response](endpoint string, token string, maxPages int) (jsonRes[] T, err error, statusCode int) {
    err = setQueryParameter(&endpoint, "per_page", "100")
    if err != nil {
        statusCode = 500 // Internal server error
        return
    }
    jsonRes = make([] T, 0, maxPages * 100)
    for {
        var res *http.Response
        res, err, statusCode = sendGithubRequestHelper(endpoint, token)
        if err != nil {
            return
        }

        var partialJsonRes[] T = make([] T, 0, 100)
        partialJsonRes, err = decodeResponse[[] T](res)


        for _, t := range partialJsonRes {
            if !t.Validate() {
                err = fmt.Errorf("Failed to parse GitHub response")
                statusCode = 500 // Internal server error
                return
            }
        }

        jsonRes = append(jsonRes, partialJsonRes...)
        fmt.Printf("%d %d\n", len(jsonRes), cap(jsonRes))

        maxPages -= 1
        if maxPages == 0 {
            return
        }

        // handle pagination
        // https://docs.github.com/en/rest/guides/using-pagination-in-the-rest-api
        linkHeader := strings.Split(res.Header.Get("link"), ", ")
        nextFound := false
        for _, link := range linkHeader {
            if strings.HasSuffix(link, "rel=\"next\"") {
                // next URL is between <>
                endpoint = link[strings.Index(link, "<")+1:strings.Index(link, ">")]
                nextFound = true
                break
            }
        }
        if !nextFound {
            return
        }
    }
}

func main() {
	  commands.Execute() // MERGE Issue: not sure if this is in the right place
    app := fiber.New()
    api := app.Group("/api")

    // Returns information about the repository's license
    api.Get("/license", func(c *fiber.Ctx) error {
        user, repo, token, err := processHeaders(c)
        if err != nil {
            fmt.Fprintf(os.Stderr, "processHeaders(): %s\n", err.Error())
            return c.SendStatus(400) // Bad request
        }

        res, err, statusCode := sendGithubRequest[LicenseResponse](fmt.Sprintf("https://api.github.com/repos/%s/%s/license", user, repo), token)
        if err != nil {
            fmt.Fprintf(os.Stderr, "sendGithubRequest(): %s\n", err.Error())
            return c.SendStatus(statusCode)
        }

        return c.JSON(res)
    })

    // Returns the average lifespan of issues (open -> close) and the number of issues sampled
    api.Get("/responsiveness", func(c *fiber.Ctx) error {
        user, repo, token, err := processHeaders(c)
        if err != nil {
            fmt.Fprintf(os.Stderr, "processHeaders(): %s\n", err.Error())
            return c.SendStatus(400) // Bad request
        }

        res, err, statusCode := sendGithubRequestList[IssueResponse](fmt.Sprintf("https://api.github.com/repos/%s/%s/issues?state=closed", user, repo), token, 5)
        if err != nil {
            fmt.Fprintf(os.Stderr, "sendGithubRequest(): %s\n", err.Error())
            return c.SendStatus(statusCode)
        }

        totalTime := 0.0
        numIssues := 0
        for _, issue := range res {
            if issue.PullRequest != nil {
                continue // this is a pull request, not an issue
            }
            ts, err := time.Parse(time.RFC3339, *issue.CreatedAt)
            if err != nil {
                fmt.Fprintf(os.Stderr, "time.Parse(): %s\n", err.Error())
                return c.SendStatus(500) // Internal server error
            }
            te, err := time.Parse(time.RFC3339, *issue.ClosedAt)
            if err != nil {
                fmt.Fprintf(os.Stderr, "time.Parse(): %s\n", err.Error())
                return c.SendStatus(500) // Internal server error
            }
            totalTime += te.Sub(ts).Seconds()
            numIssues += 1
        }
        var responsiveness Responsiveness
        if numIssues > 0 {
            responsiveness = Responsiveness{AvgLifespan: totalTime / float64(numIssues), NumSampled: numIssues}
        } else {
            responsiveness = Responsiveness{AvgLifespan: 0, NumSampled: 0}
        }
        return c.JSON(responsiveness)
    })

    // Returns a list of contributors with recent (< 1 year old) commits and their number of recent commits
    api.Get("/contributors", func(c *fiber.Ctx) error {
        user, repo, token, err := processHeaders(c)
        if err != nil {
            fmt.Fprintf(os.Stderr, "processHeaders(): %s\n", err.Error())
            return c.SendStatus(400) // Bad request
        }

        res, err, statusCode := sendGithubRequest[ContributorStatsResponse](fmt.Sprintf("https://api.github.com/repos/%s/%s/stats/contributors", user, repo), token)
        if err != nil {
            fmt.Fprintf(os.Stderr, "sendGithubRequest(): %s\n", err.Error())
            return c.SendStatus(statusCode)
        }

        var contributors[] Contributor
        now := time.Now().Unix()
        const oneYear = 60 * 60 * 24 * 356 // approximation of seconds in a year
        for _, stats := range res {
            recentCommits := 0
            for _, week := range stats.Weeks {
                if now - *week.Week <= oneYear {
                    recentCommits += *week.Commits
                }
            }
            if recentCommits > 0 {
                contributor := Contributor{Name: *stats.Author.Login, RecentCommits: recentCommits}
                contributors = append(contributors, contributor)
            }
        }

        return c.JSON(contributors)
    })

    app.Listen(":3000")
}
