package main

import (
	"fmt"
	"log"
	"net/http"

	gqltools "github.com/gunsluo/graphql-go-tools"
	"github.com/gunsluo/graphql-go-tools/example/starwars"
	"github.com/neelance/graphql-go"
	"github.com/neelance/graphql-go/relay"
)

var schema *graphql.Schema
var schemaPath = "starwars/starwars.graphql"

func init() {
	//var schemaDir = "starwars"
	//dsl, err := gqltools.ImportSchemaDslFromDir(schemaDir)
	//dsl, err := gqltools.ImportSchemaDslFromPath(schemaPath)
	//if err != nil {
	//panic(err)
	//}
	//schema = graphql.MustParseSchema(dsl, &starwars.Resolver{})
	var err error

	schema, err = gqltools.MakeExecutableSchema(schemaPath, &starwars.Resolver{})
	if err != nil {
		panic(err)
	}
}

func main() {
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(page)
	}))

	http.Handle("/query", &relay.Handler{Schema: schema})

	fmt.Println("Now server is running on port 8080.")
	fmt.Println("Test with Get      : http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

var page = []byte(`
<!DOCTYPE html>
<html>
	<head>
		<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/graphiql/0.10.2/graphiql.css" />
		<script src="https://cdnjs.cloudflare.com/ajax/libs/fetch/1.1.0/fetch.min.js"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/react/15.5.4/react.min.js"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/react/15.5.4/react-dom.min.js"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/graphiql/0.10.2/graphiql.js"></script>
	</head>
	<body style="width: 100%; height: 100%; margin: 0; overflow: hidden;">
		<div id="graphiql" style="height: 100vh;">Loading...</div>
		<script>
			function graphQLFetcher(graphQLParams) {
				return fetch("/query", {
					method: "post",
					body: JSON.stringify(graphQLParams),
					credentials: "include",
				}).then(function (response) {
					return response.text();
				}).then(function (responseBody) {
					try {
						return JSON.parse(responseBody);
					} catch (error) {
						return responseBody;
					}
				});
			}

			ReactDOM.render(
				React.createElement(GraphiQL, {fetcher: graphQLFetcher}),
				document.getElementById("graphiql")
			);
		</script>
	</body>
</html>
`)
