package main

import (
	"fmt"
	"log"

	"github.com/deeper-x/fortune-cookie/lib/schema"
	"github.com/deeper-x/fortune-cookie/lib/utils"
	"github.com/graphql-go/graphql"
)

func main() {
	// Preparing data schema
	schema := schema.Create()

	// Preparing query
	query := `{tellMe}`
	params := graphql.Params{Schema: schema, RequestString: query}

	// Running query on data schema
	runIt := graphql.Do(params)

	if len(runIt.Errors) > 0 {
		log.Fatalf("No data, error: %+v", runIt.Errors)
	}

	// Fetching data
	result := utils.ReturnJSONData(runIt)
	fmt.Printf("%s \n", result)
}
