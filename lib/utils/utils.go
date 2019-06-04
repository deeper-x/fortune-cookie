package utils

import (
	"encoding/json"
	"math/rand"
	"time"

	"github.com/graphql-go/graphql"
)

// GenRandom todo doc
func GenRandom(totToCount int) int {
	rand.Seed(time.Now().UnixNano())

	return rand.Intn(totToCount)
}

// ReturnJSONData todo doc
func ReturnJSONData(graphqlResponse *graphql.Result) []byte {
	JSONResponse, _ := json.Marshal(graphqlResponse)

	return JSONResponse
}
