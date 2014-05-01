package v1api_test

import "github.com/32bitkid/v1api"
import "testing"

func TestSomething(t *testing.T) {

	status := v1api.QueryFrom("Status")
	status.Select("Name")

	owners := v1api.QueryFrom("Owners")
	owners.Select("Name", "Nickname")

	q := v1api.QueryFrom("Story")

	q.Select("Estimate", owners)
	q.Where("Estimate", "1")
	q.Filter("Name=$Who")
	q.With("$Who", "John Doe")
	q.Sort("+Something")
	q.GroupBy(status)

	if json, err := q.ToJSON(); err == nil {
		t.Logf("%s\n\n", json)
	} else {
		t.Error("ACK", err)
	}
}
