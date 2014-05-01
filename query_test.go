package v1go_test

import "github.com/32bitkid/v1go"
import "testing"

func TestSomething(t *testing.T) {

	status := v1go.QueryFrom("Status")
	status.Select("Name")

	owners := v1go.QueryFrom("Owners")
	owners.Select("Name", "Nickname")

	q := v1go.QueryFrom("Story")

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
