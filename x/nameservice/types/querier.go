package types

import "strings"

// Query endpoints supported by the nameservice querier
const (
// TODO: Describe query parameters, update <action> with your query
// Query<Action>    = "<action>"
)

// QueryResResolve Queries Result Payload for a resolve query
type QueryResResolve struct {
	Value string `json:"value"`
}

// implement fmt.Stringer
func (r QueryResResolve) String() string {
	return r.Value
}

// Below you will be able how to set your own queries:

// QueryResList Queries Result Payload for a query
type QueryResList []string

// implement fmt.Stringer
func (n QueryResList) String() string {
	return strings.Join(n[:], "\n")
}
