import json

class Trie:

    def __init__(self):
        self.root = {}
        self.end_symbol = "*"

    def exists(self, word):
        current = self.root
        for letter in word:
            if letter not in current:
                return False
            current = current[letter]
        if self.end_symbol in current:
            return True
        return False

    def add(self, word):
        current = self.root
        for letter in word:
            if letter not in current:
                current[letter] = {}
            current = current[letter]
        current[self.end_symbol] = True

    def search_level(self, current_level, current_prefix, words_acc):
        if self.end_symbol in current_level:
            words_acc.append(current_prefix)
        for letter in sorted(current_level):
            if letter == self.end_symbol:
                continue
            prefix = current_prefix + letter
            self.search_level(current_level[letter], prefix, words_acc)
        return words

    def words_with_prefix(self, prefix):
        words = []
        current_level = self.root
        for letter in prefix:
            if letter not in current_level:
                return []
            current_level = current_level[letter]
        return self.search_level(current_level, prefix, words)

    def longest_common_prefix(self):
        current = self.root
        prefix = ""
        while True:
            keys = list(current.keys())
            if self.end_symbol in keys:
                break
            if len(keys) == 1:
                char = keys[0]
                prefix += char
                current = current[char]
            else:
                break
        return prefix
    
    def find_matches(self, document):
        matches = set()
        for i in range(len(document)):
            current = self.root
            for j, c in enumerate(document[i:]):
                if c not in current:
                    break
                current = current[c]
                if self.end_symbol in current:
                    matches.add(document[i:i + j + 1])
        return matches

    def advanced_find_matches(self, document, variations):
        matches = set()
        for i in range(len(document)):
            level = self.root
            for j in range(i, len(document)):
                ch = document[j]
                if ch in variations:
                    ch = variations[ch]
                if ch not in level:
                    break
                level = level[ch]
                if self.end_symbol in level:
                    matches.add(document[i : j + 1])
        return matches