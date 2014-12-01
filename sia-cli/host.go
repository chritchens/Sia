package main

import (
	"fmt"

	"github.com/NebulousLabs/Andromeda/siacore"
	"github.com/NebulousLabs/Andromeda/siad"
)

// Pulls a bunch of information and announces the host to the network.
func becomeHostWalkthrough(e *siad.Environment) (err error) {
	// Get a volume of days to freeze the coins.
	// Burn will be equal to price.
	// Frequency will be 100.

	// Get a volume of storage to sell.
	fmt.Print("Amount of storage to sell (in MB): ")
	var storage uint64
	_, err = fmt.Scanln(&storage)
	if err != nil {
		return
	}

	// Get a price to sell it at.
	fmt.Print("Price of storage (siacoins per kilobyte): ")
	var price uint64
	_, err = fmt.Scanln(&price)
	if err != nil {
		return
	}

	// Get a volume of coins to freeze.
	fmt.Print("How many coins to freeze (more is better): ")
	var freezeCoins uint64
	_, err = fmt.Scanln(&freezeCoins)
	if err != nil {
		return
	}

	fmt.Print("How many blocks to freeze the coins (more is better): ")
	var freezeBlocks uint64
	_, err = fmt.Scanln(&freezeBlocks)
	if err != nil {
		return
	}

	// Create the host announcement structure.
	hostSettings := siad.HostAnnouncement{
		IPAddress:             e.NetAddress(),
		MinFilesize:           1024 * 1024, // 1mb
		MaxFilesize:           storage * 1024 * 1024,
		MinDuration:           2000,
		MaxDuration:           10000,
		MinChallengeFrequency: siacore.BlockHeight(250),
		MaxChallengeFrequency: siacore.BlockHeight(100),
		MinTolerance:          10,
		Price:                 siacore.Currency(price),
		Burn:                  siacore.Currency(price),
		CoinAddress:           e.CoinAddress(),
		// SpendConditions and FreezeIndex handled by HostAnnounceSelg
	}
	e.SetHostSettings(hostSettings)

	// Have the wallet make the announcement.
	_, err = e.HostAnnounceSelf(siacore.Currency(freezeCoins), siacore.BlockHeight(freezeBlocks)+e.Height(), 10)
	if err != nil {
		return
	}

	return
}
