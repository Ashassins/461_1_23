# 1: get repo link and token

Probably would use params and/or headers

- https://docs.gofiber.io/api/ctx#params
- https://docs.gofiber.io/api/ctx#get

# 2: make requests to github

You can use net/http according to this tutorial. Will have to set some request headers according to the GitHub API.

- https://www.digitalocean.com/community/tutorials/how-to-make-http-requests-in-go

Request endpoints are:
(sub-bullets are list of keys to access relevant data)

- https://docs.github.com/en/rest/repos/repos?apiVersion=2022-11-28#get-a-repository
	- "license"
- https://docs.github.com/en/rest/issues/issues?apiVersion=2022-11-28#list-repository-issues
	- index, "created_at"
	- index, "closed_at"
- https://docs.github.com/en/rest/commits/commits?apiVersion=2022-11-28
	- index, "author", "login"

# 3: return relevant information

Define a struct data type, populate one after sending GitHub API requests, then return it suing `c.JSON()`. See example:

https://docs.gofiber.io/api/ctx#json