# Genshingo

An API wrapper to fetch your genshin user data!
- You can also fetch all your characters data (name, level, constellation, etc)

## Prerequisites
- [Go](https://go.dev/)
- HoyoLab Cookies - You can get you hoyolab cookies by visiting [hoyolab](https://www.hoyolab.com/) and logging in. After you login `Ctrl + Shift + I` then click on the `Console` Tab and then type `document.cookie` and thats the cookie!

## Install
You can run the test file to run this in the command file. You can compile it and run it your self `test/test.go` or install one of the [releases](https://github.com/Nota30/genshingo/releases). **Note:** You need to go install before you run.

**The default test file looks like this:**
![image](https://user-images.githubusercontent.com/60175653/234162511-998ccaa0-187f-46ba-bfa2-1683814f2a87.png)

 
## Quick Start
Import the `genshingo` package first.
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
