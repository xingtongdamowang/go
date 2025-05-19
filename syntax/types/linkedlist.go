package main

import "time"

type LinkedList struct {
	head *node
	tail *node

	Len int

	CreateTime time.Time
}

type node struct {
	prev *node
	next *node
}

func (l LinkedList) Add(idx int, val any) {

}

func (l *LinkedList) AddV1(idx int, val any) {

}

type Integer int

func useInt() {
	var i1 = 10
	var i2 = Integer(10)
	println(i1)
	println(i2)
}
