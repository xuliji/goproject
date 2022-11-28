package main

import "goproject/src/go_code/chapter11/interface_vs_extend/classes"

func main() {
	monkey := classes.YoungMonkey{classes.OldMonkey{Name: "悟空"}}
	classes.ShowSkills(&monkey)
}
