# Genshingo

A wrapper to fetch your genshin user data!

## Prerequisites
- [Go](https://go.dev/)
- HoyoLab Cookies - You can get you hoyolab cookies by visiting [hoyolab](https://www.hoyolab.com/) and logging in. After you login `Ctrl + Shift + I` then click on the `Console` Tab and then type `document.cookie` and thats the cookie!

## Quick Start
Import the `genshingo package first.
```go
package main

import (
	genshin "github.com/Nota30/genshingo"
)

func main() {
	init := genshin.Init{
		Cookie: "cookie goes here",
		Server: "os",
	}

	user, err := init.GetUserData("uid")
	if err != nil {
		println(err)
	}

	println(user.Data.Info.Nickname)
}
```

## Issues
- If you don't get a users data that means they haven't set their profile to `public` in hoyolab. You can make it public by going to hoyolab settings.
- Currently only works with the `os` servers so the `cn` servers won't work. Will be added in a later release maybe?
