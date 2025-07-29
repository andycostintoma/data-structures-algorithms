class Set:
    def __init__(self):
        # Use built-in dict to store keys with a dummy value
        self._dict = {}
        self._DUMMY = object()  # Shared dummy value to save memory

    def __repr__(self):
        # Represent only keys in curly braces
        keys = ", ".join(repr(k) for k in self._dict.keys())
        return "{" + keys + "}"

    # -------------------- Search --------------------
    # Time Complexity: O(1) average case
    def contains(self, key):
        """
        Check if key is present in the set.
        """
        return key in self._dict

    def __contains__(self, key):
        return self.contains(key)

    # -------------------- Insertion --------------------
    # Time Complexity: O(1) average case
    def add(self, key):
        """
        Add a key to the set.
        """
        self._dict[key] = self._DUMMY

    # -------------------- Deletion --------------------
    # Time Complexity: O(1) average case
    def remove(self, key):
        """
        Remove a key from the set.
        Raises Exception if key not found.
        """
        if key in self._dict:
            del self._dict[key]
        else:
            raise Exception("Key not found")

    # -------------------- Utilities --------------------
    def size(self):
        """
        Return the number of keys in the set.
        """
        return len(self._dict)
