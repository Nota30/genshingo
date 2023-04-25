package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	genshin "github.com/Nota30/genshingo"
)

func main() {
	fmt.Print("Enter your cookie: ")
	reader := bufio.NewReader(os.Stdin)
	cookie, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return
	}

	cookie = strings.TrimSuffix(cookie, "\n")

	init := genshin.Init{
		Cookie: cookie,
		Server: "os",
	}

	fmt.Print("Enter the uid: ")
	reader2 := bufio.NewReader(os.Stdin)
	uid, err := reader2.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return
	}

	uid = strings.TrimSuffix(uid, "\n")

	user, err := init.GetUserData(uid)
	if err != nil {
		println(err)
	}

	fmt.Println("------------------ YOUR DATA ------------------")
	fmt.Println("Nickname:", user.Data.Info.Nickname)
	fmt.Println("Adventure Rank:", user.Data.Info.Level)
	fmt.Println("Region:", user.Data.Info.Region)
	fmt.Println("Current Spiral Abyss:", user.Data.Stats.SpiralAbyss)
	fmt.Println("--------------------------------")
	fmt.Println("Domains Unlocked:", user.Data.Stats.Domains)
	fmt.Println("Waypoints Unlocked:", user.Data.Stats.Waypoints)
	fmt.Println("Achievements:", user.Data.Stats.Achievements)
	fmt.Println("You have been active for", user.Data.Stats.ActiveDays, "days")
	fmt.Println("Anemoculus Collected:",user.Data.Stats.Anemoculus)
	fmt.Println("Geoculus Collected:",user.Data.Stats.Geoculus)
	fmt.Println("Electroculus Collected:",user.Data.Stats.Electroculus)
	fmt.Println("Dendroculus Collected:",user.Data.Stats.Dendroculus)
	fmt.Println("Luxurious Chests Collected:",user.Data.Stats.LuxuriousChests)
	fmt.Println("Precious Chests Collected:",user.Data.Stats.PreciousChests)
	fmt.Println("Exquisite Chests Collected:",user.Data.Stats.ExquisiteChests)
	fmt.Println("Common Chests Collected:",user.Data.Stats.CommonChests)
	fmt.Println("Magic Chests Collected:",user.Data.Stats.MagicChests)
}