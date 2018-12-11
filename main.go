package main

import (
	CardInfo "hearthlethal/card"
)

func main() {
	if !CardInfo.CheckMetaFileExist() {
		CardInfo.GetAllCards()
	}
}
