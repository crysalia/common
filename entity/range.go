package entity

type IntegerRange struct {
	From int64 `bson:"from"`
	To   int64 `bson:"to"`
}
