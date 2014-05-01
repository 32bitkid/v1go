package v1go_test

import "github.com/32bitkid/v1go"
import "testing"

type Workitem struct {
	Name string
}

func TestAssetLoader(t *testing.T) {
	var workItems []Workitem
	
	t.Log(workItems)
	if err := v1go.QueryFor(&workItems); err != nil {
	  t.Fatal(err)
	}
	t.Log(workItems)
	
	
}
