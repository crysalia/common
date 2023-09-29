package entity

import "crysalia/common/enums"

type ItemIdentification struct {
	Id   string         `bson:"id"`
	Type enums.CalcType `bson:"calc_type"`
}

type RangeableItemIdentification struct {
	Range IntegerRange `bson:"range"`
	ItemIdentification
}

type BonusStatItemIdentification struct {
	Stat   enums.Stat `bson:"stat"`
	Amount int64      `bson:"amount"`
	ItemIdentification
}
