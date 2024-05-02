# A little TodolistServer

latest version: v0.1.1

## Setup
- create your go project (go mod init example.com/....)
- run "go get github.com/DenisMathan/GuteFrageChallenge@v0.1.1"
- create main.go with:
    server := api.NewServer(cfg) <br />
      server.Start()
- cfg := configurations.Config{<br />
    ... <br />
  }
- run "go run main.go"



## Allowed Requests
### GET
- /todos
- parameters: <br />
    pagination: int 	//Limits the query according to the number <br />
    nth: int		//Sets the Offset for the return values <br />
    onlyDone: bool	//optional -> if true it returns only the done elements <br />
  


Works currently only with a mysqlDB which has a todos table with these columns:
- id
- description
- done
