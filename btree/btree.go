package main

import (
	"fmt"
	"time"

	"github.com/google/btree"
)

type user struct {
	Name string
	Age uint32
	Birthday time.Time
}

func (u *user) Less(u1 btree.Item) bool {
	return u.Birthday.After(u1.(*user).Birthday)
}

func main() {
	tree := btree.New(10)

	now := time.Now().UTC()
	u1 := newUser(now, "XiaoMing", 12)
	u2 := newUser(now, "WangMaZi", 24)
	u3 := newUser(now, "Bob", 36)
	u6 := newUser(now, "John", 28)
	u8 := newUser(now, "John1", 56)

	tree.ReplaceOrInsert(u8)
	tree.ReplaceOrInsert(u1)
	tree.ReplaceOrInsert(u2)
	tree.ReplaceOrInsert(u3)
	tree.ReplaceOrInsert(u6)

	u1InTree := tree.Get(u1)
	fmt.Printf("get u1: %+v\n", u1InTree)

	u2InTree := tree.Get(u2)
	fmt.Printf("get u1: %+v\n", u2InTree)

	u3InTree := tree.Get(u3)
	fmt.Printf("get u1: %+v\n", u3InTree)

	theMin := tree.Min()
	fmt.Printf("the min item: %+v\n", theMin)

	theMax := tree.Max()
	fmt.Printf("the max item: %+v\n", theMax)

	hasU1 := tree.Has(u1)
	fmt.Printf("the max item: %+v\n", hasU1)

	delItem := tree.Delete(u6)
	fmt.Printf("deleted item: %+v\n", delItem)

	var users []*user
	tree.Descend(func(i btree.Item) bool {
		users = append(users, i.(*user))
		return true
	})
	fmt.Println("users in reverse order")
	for _, u := range users {
		fmt.Printf("user: %+v\n", *u)
	}
}

func newUser(now time.Time, name string, age uint32) *user {
	years := time.Duration(age)
	return &user{
		Name:     name,
		Age:      age,
		Birthday: now.Add((-years*365*24) * time.Hour),
	}
}
