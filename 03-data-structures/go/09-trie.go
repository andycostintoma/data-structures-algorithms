package main

import (
	"fmt"
	"sort"
	"strings"
)

type node struct {
	children map[rune]*node
	isEnd    bool
}

func newNode() *node {
	return &node{children: make(map[rune]*node)}
}

type Trie struct {
	root *node
}

func NewTrie() *Trie {
	return &Trie{root: newNode()}
}

// Add inserts a word into the trie.
func (t *Trie) Add(word string) {
	curr := t.root
	for _, ch := range word {
		if _, ok := curr.children[ch]; !ok {
			curr.children[ch] = newNode()
		}
		curr = curr.children[ch]
	}
	curr.isEnd = true
}

// Exists checks whether a complete word is in the trie.
func (t *Trie) Exists(word string) bool {
	curr := t.root
	for _, ch := range word {
		next, ok := curr.children[ch]
		if !ok {
			return false
		}
		curr = next
	}
	return curr.isEnd
}

// WordsWithPrefix returns all stored words that start with the given prefix.
func (t *Trie) WordsWithPrefix(prefix string) []string {
	curr := t.root
	for _, ch := range prefix {
		next, ok := curr.children[ch]
		if !ok {
			return nil
		}
		curr = next
	}
	var results []string
	t.collectWords(curr, prefix, &results)
	return results
}

// collectWords recursively collects all complete words from the current node.
func (t *Trie) collectWords(curr *node, prefix string, results *[]string) {
	if curr.isEnd {
		*results = append(*results, prefix)
	}
	letters := make([]rune, 0, len(curr.children))
	for ch := range curr.children {
		letters = append(letters, ch)
	}
	sort.Slice(letters, func(i, j int) bool { return letters[i] < letters[j] })
	for _, ch := range letters {
		t.collectWords(curr.children[ch], prefix+string(ch), results)
	}
}

// LongestCommonPrefix returns the longest prefix common to all words in the trie.
func (t *Trie) LongestCommonPrefix() string {
	var builder strings.Builder
	curr := t.root
	for len(curr.children) == 1 && !curr.isEnd {
		for ch, next := range curr.children {
			builder.WriteRune(ch)
			curr = next
		}
	}
	return builder.String()
}

// FindMatches finds all trie words that appear as substrings in the input.
func (t *Trie) FindMatches(text string) map[string]bool {
	matches := make(map[string]bool)
	runes := []rune(text)
	for i := range runes {
		curr := t.root
		for j := i; j < len(runes); j++ {
			ch := runes[j]
			next, ok := curr.children[ch]
			if !ok {
				break
			}
			curr = next
			if curr.isEnd {
				matches[string(runes[i:j+1])] = true
			}
		}
	}
	return matches
}

// AdvancedFindMatches allows substitutions via the provided mapping.
func (t *Trie) AdvancedFindMatches(text string, variations map[rune]rune) map[string]bool {
	matches := make(map[string]bool)
	runes := []rune(text)
	for i := range runes {
		curr := t.root
		for j := i; j < len(runes); j++ {
			ch := runes[j]
			if alt, ok := variations[ch]; ok {
				ch = alt
			}
			next, ok := curr.children[ch]
			if !ok {
				break
			}
			curr = next
			if curr.isEnd {
				matches[string(runes[i:j+1])] = true
			}
		}
	}
	return matches
}

// PrintTree prints the structure of the trie.
func (t *Trie) PrintTree() {
	var printNode func(*node, string, int)
	printNode = func(n *node, prefix string, level int) {
		indent := strings.Repeat("  ", level)
		if n.isEnd {
			fmt.Printf("%s%s (end)\n", indent, prefix)
		} else if prefix != "" {
			fmt.Printf("%s%s\n", indent, prefix)
		}
		letters := make([]rune, 0, len(n.children))
		for ch := range n.children {
			letters = append(letters, ch)
		}
		sort.Slice(letters, func(i, j int) bool { return letters[i] < letters[j] })
		for _, ch := range letters {
			printNode(n.children[ch], string(ch), level+1)
		}
	}
	printNode(t.root, "", 0)
}
