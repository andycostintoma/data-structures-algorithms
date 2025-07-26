class Graph:
    def __init__(self):
        self.graph = {}

    def add_edge(self, u, v):
        if u not in self.graph:
            self.graph[u] = set()
        if v not in self.graph:
            self.graph[v] = set()
        self.graph[u].add(v)
        self.graph[v].add(u)

    def edge_exists(self, u, v):
        if u in self.graph and v in self.graph:
            return (v in self.graph[u]) and (u in self.graph[v])
        return False

    def adjacent_nodes(self, node):
        return self.graph[node]
    
    def unconnected_vertices(self):
        unconnected_vertices = []
        for k in self.graph:
            if not self.graph[k]:
                unconnected_vertices.append(k)
        return unconnected_vertices

    def breadth_first_search(self, v):
        visited = []
        to_visit = []
        to_visit.append(v)
        while to_visit:
            s = to_visit.pop(0)
            visited.append(s)
            sorted_neighbors = sorted(self.graph[s])
            for neighbor in sorted_neighbors:
                if neighbor not in visited and neighbor not in to_visit:
                    to_visit.append(neighbor)
        return visited