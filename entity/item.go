package entity

import "crysalia/common/enums"

type Item struct {
	Id              string                         `bson:"id"`
	Name            string                         `bson:"name"`
	Level           int64                          `bson:"level"`
	Texture         int64                          `bson:"texture"`
	Rarity          enums.ItemRarity               `bson:"rarity"`
	Category        enums.ItemCategory             `bson:"category"`
	Description     []string                       `bson:"description"`
	Identifications []ItemIdentification           `bson:"identifications"`
	BaseDamages     map[enums.Element]IntegerRange `bson:"base_damages"`
}
