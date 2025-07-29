class HashMap:
    def __init__(self, size):
        # Initialize the hashmap with given size, all slots None initially
        self.hashmap = [None for _ in range(size)]
        self._deleted = object()  # Sentinel object to mark deleted slots

    def __repr__(self):
        # Represent only the active key-value pairs (skip None and deleted)
        buckets = []
        for v in self.hashmap:
            if v is not None and v is not self._deleted:
                buckets.append(v)
        return str(buckets)

    def __key_to_index(self, key):
        """
        Simple hash function: sum of ASCII codes of characters modulo hashmap size.
        Returns the initial bucket index for a given key.
        """
        total = 0
        for c in key:
            total += ord(c)
        return total % len(self.hashmap)

    # -------------------- Search --------------------
    # Time Complexity: O(1) average case

    def get(self, key):
        """
        Retrieve value for the given key.
        Raises Exception if key not found.
        """
        index = self.__key_to_index(key)
        original_index = index
        first_iteration = True

        # Probe linearly until we find the key or an empty slot
        while self.hashmap[index] is not None:
            # If this slot is active and keys match, return value
            if self.hashmap[index] != self._deleted and self.hashmap[index][0] == key:
                return self.hashmap[index][1]

            # If we've looped back to the original index, stop searching
            if not first_iteration and index == original_index:
                break

            index = (index + 1) % len(self.hashmap)
            first_iteration = False

        raise Exception("Key not found")

    # -------------------- Insertion --------------------
        # Time Complexity: O(1) average case
        
    def insert(self, key, value):
        """
        Insert or update the key with the given value.
        Resizes the hashmap if load factor exceeds threshold.
        """
        self.__resize_if_needed()
        self.__insert_without_resize(key, value)

    def __insert_without_resize(self, key, value):
        """
        Insert key-value pair without resizing.
        Uses linear probing to find an empty or deleted slot or the same key.
        """
        index = self.__key_to_index(key)

        while (
            self.hashmap[index] is not None
            and self.hashmap[index] != self._deleted
            and self.hashmap[index][0] != key
        ):
            index = (index + 1) % len(self.hashmap)

        self.hashmap[index] = (key, value)

    # -------------------- Deletion --------------------
    # Time Complexity: O(1) average case

    def delete(self, key):
        """
        Delete the key-value pair by marking the slot as deleted with sentinel.
        Raises Exception if key not found.
        """
        index = self.__key_to_index(key)
        original_index = index
        first_iteration = True

        # Probe to find the key
        while self.hashmap[index] is not None:
            if self.hashmap[index] != self._deleted and self.hashmap[index][0] == key:
                # Mark this slot as deleted instead of None to preserve probe chain
                self.hashmap[index] = self._deleted
                return

            if not first_iteration and index == original_index:
                break

            index = (index + 1) % len(self.hashmap)
            first_iteration = False

        raise Exception("Key not found")

    # -------------------- Utilities --------------------

    def size(self):
        """
        Returns the number of active key-value pairs stored.
        """
        count = 0
        for b in self.hashmap:
            if b is not None and b != self._deleted:
                count += 1
        return count

    def __current_load(self):
        """
        Calculate current load factor: active entries / capacity.
        """
        return self.size() / len(self.hashmap)

    def __resize_if_needed(self):
        """
        Resize hashmap by doubling size if load factor exceeds 0.7.
        Rehashes all active entries.
        """
        load = self.__current_load()
        if load < 0.7:
            return

        old_hashmap = self.hashmap
        self.hashmap = [None for _ in range(len(old_hashmap) * 2)]

        for kvp in old_hashmap:
            if kvp is not None and kvp != self._deleted:
                self.__insert_without_resize(kvp[0], kvp[1])
