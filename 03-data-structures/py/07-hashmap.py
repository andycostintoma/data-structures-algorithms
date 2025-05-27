class HashMap:
    def __init__(self, size):
        self.hashmap = [None for _ in range(size)]

    def key_to_index(self, key):
        total = 0
        for c in key:
            total += ord(c)
        return total % len(self.hashmap)

    def current_load(self):
        if len(self.hashmap) == 0:
            return 1
        filled_slots = sum(1 for slot in self.hashmap if slot is not None)
        return filled_slots / len(self.hashmap)

    def resize(self):
        if len(self.hashmap) == 0:
            self.hashmap = [None]
            return
        if self.current_load() < 0.05:
            return

        old_hashmap = self.hashmap
        self.hashmap = [None] * (len(old_hashmap) * 10)

        for kvp in old_hashmap:
            if kvp is not None:
                key, value = kvp
                self.insert(key, value)

    def insert(self, key, value):
        self.resize()
        index = self.key_to_index(key)
        original_index = index
        first_iteration = True

        while (self.hashmap[index] is not None) and (self.hashmap[index][0] != key):
            if not first_iteration and index == original_index:
                raise Exception("hashmap is full")
            index = (index + 1) % len(self.hashmap)
            first_iteration = False

        self.hashmap[index] = (key, value)
        

    def get(self, key):
        index = self.key_to_index(key)
        original_index = index
        first_iteration = True

        while self.hashmap[index] is not None:
            if self.hashmap[index][0] == key:
                return self.hashmap[index][1]
            if not first_iteration and index == original_index:
                raise Exception("sorry, key not found")
            index = (index + 1) % len(self.hashmap)
            first_iteration = False

        raise Exception("sorry, key not found")

    def __repr__(self):
        final = ""
        for i, v in enumerate(self.hashmap):
            if v is not None:
                final += f" - {str(v)}\n"
        return final
