package enums

type ItemRarity int8
type ItemCategory int8
type Element int8
type IdType int8
type CalcType int8
type Stat int8
type Class int8

const (
	Wizard Class = iota
	Paladin
	Ranger
)

const (
	Common ItemRarity = iota
	Rare
	Legendary
	Epic
)

const (
	Spear ItemCategory = iota
	Wand
	Bow
)

const (
	Neutral Element = iota
	Thunder
	Water
	Fire
	Earth
	Air
	Shadow
	Light
)

const (
	Strenght     Stat = iota // Increase any damage you deal
	Dexterity                // Increase the chance to deal critical damage
	Intelligence             // Decrease the overall mana cost of spells
	Agility                  // Increase the chance to dodge damages
	Vitality                 // Increase your overall health regen raw
	Lucky                    // Increase your chance to do combos
	Endurance                // Increase your max health
	Defense                  // Decrease any damage
	Faith                    // Increase the chance of miracles
)

const (
	Raw CalcType = iota
	Percentage
)

const (
	HealthRegen      IdType = iota // Boost the health regen %
	HealthRegenSpeed               // Increase health regen speed
	ManaRegen                      // Increase mana regen
	PhysicalDamage                 // Boost physical damage
	SpellDamage                    // Boost spell damage
	PhysicalDefense                // Decrease the physical damage you get
	MagicDefense                   // Decrease the spell damage you get
	StatBonus                      // Boost stat points
	WaterDamage                    // Boost water damage
	EarthDamage                    // Boost earth damage
	ThunderDamage                  // Boost thunder damage
	AirDamage                      // Boost air damage
	FireDamage                     // Boost fire damage
	ShadowDamage                   // Boost shadow damage
	LightDamage                    // Boost light damage
	ElementalDamage                // Boost all the elemental damage
)
