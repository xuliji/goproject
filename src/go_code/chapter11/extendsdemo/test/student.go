//go:generate  sloth -out=User -fun=set,get
package test

type User struct {
	Name string
	Age  int
}
