import unittest

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
        return words_acc

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

class TestTrie(unittest.TestCase):

    def test_add_and_exists(self):
        trie = Trie()
        # Initially, no words
        self.assertFalse(trie.exists("hello"))

        # Add a word and test existence
        trie.add("hello")
        self.assertTrue(trie.exists("hello"))

        # Test that prefixes without end symbol don't count
        self.assertFalse(trie.exists("hell"))
        self.assertFalse(trie.exists("helloo"))

        # Add another word
        trie.add("hell")
        self.assertTrue(trie.exists("hell"))
        self.assertTrue(trie.exists("hello"))

    def test_exists_manual_root(self):
        # Build a trie manually for the word 'test'
        trie = Trie()
        trie.root = {
            't': {
                'e': {
                    's': {
                        't': {
                            trie.end_symbol: True
                        }
                    }
                }
            }
        }
        self.assertTrue(trie.exists("test"))
        self.assertFalse(trie.exists("tes"))
        self.assertFalse(trie.exists("testing"))

    def test_words_with_prefix(self):
        trie = Trie()
        # Add some words
        for word in ["apple", "app", "apply", "apt", "bat", "batch"]:
            trie.add(word)

        # Test prefix 'app'
        self.assertEqual(
            sorted(trie.words_with_prefix("app")),
            ["app", "apple", "apply"]
        )

        # Test prefix 'bat'
        self.assertEqual(
            sorted(trie.words_with_prefix("bat")),
            ["bat", "batch"]
        )

        # Test prefix not in trie
        self.assertEqual(trie.words_with_prefix("banana"), [])

    def test_words_with_prefix_manual_root(self):
        # Manually build a trie for 'cat' and 'car'
        trie = Trie()
        trie.root = {
            'c': {
                'a': {
                    't': {trie.end_symbol: True},
                    'r': {trie.end_symbol: True}
                }
            }
        }
        self.assertEqual(
            sorted(trie.words_with_prefix("ca")),
            ["car", "cat"]
        )

    def test_longest_common_prefix(self):
        trie = Trie()
        for word in ["car", "cat", "cab"]:
            trie.add(word)
        self.assertEqual(trie.longest_common_prefix(), "ca")

        trie2 = Trie()
        for word in ["dog", "cat", "mouse"]:
            trie2.add(word)
        self.assertEqual(trie2.longest_common_prefix(), "")

        trie3 = Trie()
        trie3.add("hello")
        self.assertEqual(trie3.longest_common_prefix(), "hello")

        trie4 = Trie()
        # Empty trie should have empty prefix
        self.assertEqual(trie4.longest_common_prefix(), "")

    def test_find_matches(self):
        trie = Trie()
        for word in ["apple", "bat", "cat"]:
            trie.add(word)

        document = "I have an apple and a bat."
        matches = trie.find_matches(document)
        self.assertEqual(matches, {"apple", "bat"})

        # Test no matches
        document2 = "No matches here"
        matches2 = trie.find_matches(document2)
        self.assertEqual(matches2, set())

    def test_advanced_find_matches(self):
        trie = Trie()
        for word in ["apple", "bat", "cat"]:
            trie.add(word)

        document = "I have 4pple and b@t."
        variations = {"4": "a", "@": "a"}  # Map '4'->'a' and '@'->'a'
        matches = trie.advanced_find_matches(document, variations)
        # Matches are returned with original substrings from document
        self.assertEqual(matches, {"4pple", "b@t"})

        # Test with no variations
        document2 = "apple and bat"
        matches2 = trie.advanced_find_matches(document2, {})
        self.assertEqual(matches2, {"apple", "bat"})

if __name__ == "__main__":
    unittest.main()
