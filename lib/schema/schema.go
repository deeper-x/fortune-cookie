package schema

import (
	"fmt"
	"log"

	"github.com/deeper-x/fortune-cookie/lib/quotereader"
	qr "github.com/deeper-x/fortune-cookie/lib/quotereader"
	"github.com/deeper-x/fortune-cookie/lib/utils"
	"github.com/graphql-go/graphql"
)

//Create a graphql Schema
func Create() graphql.Schema {
	// Schema
	fields := graphql.Fields{
		"tellMe": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				//load data from file
				quotereader.LoadData()

				//configure fetching params
				totIn := len(qr.Data.Records)
				index := utils.GenRandom(totIn)

				//output result
				quote := qr.Data.Records[index].QuoteText
				author := qr.Data.Records[index].QuoteAuthor
				res := fmt.Sprintf("%s %s", quote, author)

				return res, nil
			},
		},
	}
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	return schema
}
