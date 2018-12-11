package CardInfo

type CardType int

const (
	Card_Minion CardType = iota + 1
	Card_Weapon
	Card_Spell
	Card_Hero
)

type CardSet int

const (
	Set_Basic Class = iota + 1
	Set_Classic
	Set_HOF // Hall of Fame
	Set_CON // Curse of Naxxramas
	Set_GVG // Goblins vs Gnomes
	Set_BRM // Blackrock Mountain
	Set_TGT // The Grand Tournament
	Set_LOE // The League of Explorers
	Set_WOG // Whispers of the Old Gods
	Set_ONK // One Night in Karazhan
	Set_MSG // Mean Streets of Gadgetzan
	Set_JUG // Journey to Un'Goro
	Set_KFT // Knights of the Frozen Throne
	Set_KAC // Kobolds & Catacombs
	Set_TWW // The Witchwood
	Set_TBP // The Boomsday Project
	Set_RKR // Rastakhan's Rumble
)

type Class int

const (
	Class_Neutral Class = iota + 1
	Class_Druid
	Class_Hunter
	Class_Mage
	Class_Paladin
	Class_Priest
	Class_Rogue
	Class_Shaman
	Class_Warlock
	Class_Warrior
)

type Rarity int

const (
	Rarity_Free Rarity = iota + 1
	Rarity_Common
	Rarity_Rare
	Rarity_Epic
	Rarity_Legendary
)

type MinionType int

const (
	Minion_Beast MinionType = iota + 1
	Minion_Demon
	Minion_Dragon
	Minion_Elemental
	Minion_Mech
	Minion_Murloc
	Minion_Pirate
	Minion_Totem
	Minion_General
	Minion_All
)

type CardInfo struct {
	CardType    CardType
	CardSet     CardSet
	Class       int
	Rarity      int
	Collectible bool
	Golden      bool
	Cost        int
}

type MinionCardInfo struct {
	CardInfo
	MinionType MinionType
	Attack     int
	HP         int
	// TODO: effect
}

type WeaponCardInfo struct {
	CardInfo
	Attack int
	HP     int
	// TODO: effect
}

type SpellCardInfo struct {
	CardInfo
	// TODO: effect
}
