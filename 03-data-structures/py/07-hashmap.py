class HashMap:
    def insert(self, key, value):
        self.resize()
        index = self.key_to_index(key)
        self.hashmap[index] = (key, value)

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
                index = self.key_to_index(kvp[0])
                self.hashmap[index] = (kvp[0], kvp[1])

    def current_load(self):
        if len(self.hashmap) == 0:
            return 1
        filled_slots = 0
        for slot in self.hashmap:
            if slot is not None:
                filled_slots += 1
        return filled_slots / len(self.hashmap)

    # don't touch below this line

    def __init__(self, size):
        self.hashmap = [None for i in range(size)]

    def key_to_index(self, key):
        sum = 0
        for c in key:
            sum += ord(c)
        return sum % len(self.hashmap)

    def __repr__(self):
        final = ""
        for i, v in enumerate(self.hashmap):
            if v != None:
                final += f" - {str(v)}\n"
        return final
