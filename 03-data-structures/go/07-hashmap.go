package main

type entry struct {
	key     string
	value   any
	deleted bool
}

type HashMap struct {
	entries []entry
}
