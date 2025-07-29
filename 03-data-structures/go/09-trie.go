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

// Exists checks if a word exists in the trie.
func (t *Trie) Exists(word string) bool {
	curr := t.root
	for _, ch := range word {
		if _, ok := curr.children[ch]; !ok {
			return false
		}
		curr = curr.children[ch]
	}
	return curr.isEnd
}

// WordsWithPrefix returns all words starting with a given prefix.
func (t *Trie) WordsWithPrefix(prefix string) []string {
	curr := t.root
	for _, ch := range prefix {
		if _, ok := curr.children[ch]; !ok {
			return nil
		}
		curr = curr.children[ch]
	}
	var results []string
	t.dfs(curr, prefix, &results)
	return results
}

// dfs is a recursive helper to collect words from a trie node.
func (t *Trie) dfs(curr *node, prefix string, results *[]string) {
	if curr.isEnd {
		*results = append(*results, prefix)
	}
	letters := make([]rune, 0, len(curr.children))
	for ch := range curr.children {
		letters = append(letters, ch)
	}
	sort.Slice(letters, func(i, j int) bool { return letters[i] < letters[j] })
	for _, ch := range letters {
		t.dfs(curr.children[ch], prefix+string(ch), results)
	}
}

// LongestCommonPrefix finds the longest shared prefix of all words.
func (t *Trie) LongestCommonPrefix() string {
	curr := t.root
	var prefix strings.Builder
	for {
		if curr.isEnd || len(curr.children) != 1 {
			break
		}
		for ch, child := range curr.children {
			prefix.WriteRune(ch)
			curr = child
			break
		}
	}
	return prefix.String()
}

// FindMatches finds all words from the trie that appear as substrings in the document.
func (t *Trie) FindMatches(document string) map[string]bool {
	matches := make(map[string]bool)
	runes := []rune(document)
	for i := 0; i < len(runes); i++ {
		curr := t.root
		for j := i; j < len(runes); j++ {
			ch := runes[j]
			if next, ok := curr.children[ch]; ok {
				curr = next
				if curr.isEnd {
					matches[string(runes[i:j+1])] = true
				}
			} else {
				break
			}
		}
	}
	return matches
}

// AdvancedFindMatches does the same as FindMatches but allows for character substitutions.
func (t *Trie) AdvancedFindMatches(document string, variations map[rune]rune) map[string]bool {
	matches := make(map[string]bool)
	runes := []rune(document)
	for i := 0; i < len(runes); i++ {
		curr := t.root
		for j := i; j < len(runes); j++ {
			ch := runes[j]
			if alt, ok := variations[ch]; ok {
				ch = alt
			}
			if next, ok := curr.children[ch]; ok {
				curr = next
				if curr.isEnd {
					matches[string(runes[i:j+1])] = true
				}
			} else {
				break
			}
		}
	}
	return matches
}

func (t *Trie) PrintTree() {
	var printNode func(n *node, prefix string, level int)
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

// Example usage
func main() {
	trie := NewTrie()
	words := []string{"apple", "app", "apply", "apt", "bat", "batch"}
	for _, w := range words {
		trie.Add(w)
	}

	fmt.Println("Exists 'app':", trie.Exists("app"))
	fmt.Println("Exists 'bat':", trie.Exists("bat"))
	fmt.Println("Exists 'bats':", trie.Exists("bats"))

	fmt.Println("Words with prefix 'ap':", trie.WordsWithPrefix("ap"))
	fmt.Println("Words with prefix 'ba':", trie.WordsWithPrefix("ba"))
	fmt.Println("Words with prefix 'z':", trie.WordsWithPrefix("z"))

	fmt.Println("Longest Common Prefix:", trie.LongestCommonPrefix())

	doc := "I have an apple and a bat."
	fmt.Println("Matches in doc:", trie.FindMatches(doc))

	variations := map[rune]rune{'4': 'a', '@': 'a'}
	doc2 := "4pple and b@t"
	fmt.Println("Advanced matches:", trie.AdvancedFindMatches(doc2, variations))

	fmt.Println("All words in trie:")
	trie.PrintTree()
}
