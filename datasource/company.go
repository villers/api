package datasource

import "github.com/villers/api/datamodel"

// Movies is our imaginary data source.
var Companies = map[int64]datamodel.Company{
	1: {
		ID:   1,
		Name: "Google",
	},
	2: {
		ID:   2,
		Name: "Facebook",
	},
}
