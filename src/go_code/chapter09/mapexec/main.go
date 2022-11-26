package main

import "fmt"

func ModifyUsers(users map[string]map[string]string, name string) {
	if users[name] != nil {
		users[name]["pwd"] = "888888"
	} else {
		users[name] = map[string]string{
			"nickname": "hape",
			"pwd":      "666666",
		}
	}
}

func main() {

	users := map[string]map[string]string{
		"marry": {
			"nickname": "请叫我BTree",
			"pwd":      "1111111",
		},
	}
	ModifyUsers(users, "jack")
	fmt.Println(users)
}
