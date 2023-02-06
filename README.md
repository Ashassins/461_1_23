Team Members:
Mimi Chon, Anna Shen, Emile Baez, Ben Schwartz

# Building program
In command line, first run 
`go build -o "run"`
in order to create executable with 'run' as name.

# CLI
<`build`, `install`, `test`, `'URL_FILE'`> commands are recognized where 'URL_FILE' must be an absolute path to a file in the system.

# Architecture
![Architecture](resources/arch.jpg)   
This is the current architecture of our program. Each block represents a collection of functions towards a single functionality. Each color represents a package in Go.

# Helpful Commands
`go run .`   
`go build .`

# GitHub API Notes
endpoints : /license, /responsiveness, /contributors
need to add /correctness, etc.

using localhost:3000

e.g. to add headers to get terminal responses
curl -H "InputURL: https://github.com/qiangxue/go-rest-api" -H "GitHubToken: <YOURGPAT>" localhost:3000/api/license

for me this returns : {"license":{"key":"mit","name":"MIT License","url":"https://api.github.com/licenses/mit"}}
