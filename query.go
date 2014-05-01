package v1api

import "encoding/json"

func QueryFrom(assetType string) Query {
	return &query{
		From:   assetType,
		Wheres: make(map[string]string),
		Withs:  make(map[string]string),
	}
}

type Query interface {
	ToJSON() ([]byte, error)
	Select(fields ...interface{})
	Where(key, value string)
	Filter(tokens ...string)
	With(key, value string)
	Sort(token string)
	SortAscending(token string)
	SortDescending(token string)
	GroupBy(groups ...Query)
}

type query struct {
	From    string            `json:"from"`
	Selects []interface{}     `json:"select,omitempty"`
	Wheres  map[string]string `json:"where,omitempty"`
	Filters []string          `json:"filter,omitempty"`
	Withs   map[string]string `json:"with,omitempty"`
	Sorts   []string          `json:"sort,omitempty"`
	Groups  []Query           `json:"group,omitempty"`
}

func (q *query) ToJSON() ([]byte, error) {
	return json.Marshal(q)
}

func (q *query) Select(fields ...interface{}) {
	for _, field := range fields {
		q.Selects = append(q.Selects, field)
	}
}

func (q *query) Where(key, value string) {
	q.Wheres[key] = value
}

func (q *query) Filter(tokens ...string) {
	for _, token := range tokens {
		q.Filters = append(q.Filters, token)
	}
}

func (q *query) With(key, value string) {
	q.Withs[key] = value
}

func (q *query) Sort(token string) {
	q.Sorts = append(q.Sorts, token)
}

func (q *query) SortAscending(token string) {
	q.Sorts = append(q.Sorts, "+"+token)
}

func (q *query) SortDescending(token string) {
	q.Sorts = append(q.Sorts, "-"+token)
}

func (q *query) GroupBy(groups ...Query) {
	q.Groups = append(q.Groups, groups...)
}
