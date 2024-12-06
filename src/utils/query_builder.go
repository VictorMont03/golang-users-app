package utils

import (
	"net/url"

	"go.mongodb.org/mongo-driver/bson"
)

type QueryOptions struct {
	Where map[string]interface{} `json:"where"`
	Limit int                    `json:"limit"`
	Page  int                    `json:"page"`
	Skip  int                    `json:"skip"`
}

type QueryBuilder struct {
	query bson.D
}

func NewQueryBuilder() *QueryBuilder {
	return &QueryBuilder{
		query: bson.D{},
	}
}

func (qb *QueryBuilder) ApplyWhere(where url.Values) *QueryBuilder {
	for key, value := range where {
		qb.query = append(qb.query, bson.E{Key: key, Value: value[0]})
	}
	return qb
}

func (qb *QueryBuilder) Build() bson.D {
	return qb.query
}
