/*
At composite pattern, a root instance hosts child instances in a tree form.
You can think about this pattern is like a recursive function.
But it is a recursive data type, and it has recursive behavior.
*/

package main

import (
	"fmt"
)

func main() {
	branch := &Branch{name: "Root Branch"}
	childBranch := &Branch{name: "Child Branch"}
	branch.addLeaf(Leaf{name: "Root Branch's Leaf 1"})
	branch.addLeaf(Leaf{name: "Root Branch's Leaf 2"})
	childBranch.addLeaf(Leaf{name: "Child Branch's Leaf 1"})
	childBranch.addLeaf(Leaf{name: "Child Branch's Leaf 2"})
	branch.addBranch(*childBranch)
	branch.perform()
}

type IComposite interface {
	perform()
}

type Leaf struct {
	name string
}

func (leaf *Leaf) perform() {
	fmt.Println("Leaf name: ", leaf.name)
}

type Branch struct {
	name     string
	leafs    []Leaf
	branches []Branch
}

func (branch *Branch) perform() {
	fmt.Println("Branch name: ", branch.name)
	for _, leaf := range branch.leafs {
		leaf.perform()
	}
	for _, branch := range branch.branches {
		branch.perform()
	}
}

func (branch *Branch) addLeaf(childLeaf Leaf) {
	branch.leafs = append(branch.leafs, childLeaf)
}

func (branch *Branch) addBranch(childBranch Branch) {
	branch.branches = append(branch.branches, childBranch)
}
