# 461_1_23
Mimi Chon
Anna Shen
Emile Baez
Ben Schwartz


# GitHub API Notes
endpoints : /license, /responsiveness, /contributors
need to add /correctness, etc.

using localhost:3000

e.g. to add headers to get terminal responses
curl -H "InputURL: https://github.com/qiangxue/go-rest-api" -H "GitHubToken: <YOURGPAT>" localhost:3000/api/license

for me this returns : {"license":{"key":"mit","name":"MIT License","url":"https://api.github.com/licenses/mit"}}