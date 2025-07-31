from collections import deque


class Graph:
    def __init__(self):
        self.graph = {}

    def add_node(self, u):
        if u not in self.graph:
            self.graph[u] = set()

    def add_edge(self, u, v):
        if u not in self.graph:
            self.graph[u] = set()
        if v not in self.graph:
            self.graph[v] = set()
        self.graph[u].add(v)
        self.graph[v].add(u)

    def edge_exists(self, u, v):
        return (
            u in self.graph
            and v in self.graph
            and v in self.graph[u]
            and u in self.graph[v]
        )

    def adjacent_nodes(self, node):
        return self.graph[node]

    def unconnected_vertices(self):
        unconnected_vertices = []
        for k in self.graph:
            if not self.graph[k]:
                unconnected_vertices.append(k)
        return unconnected_vertices

    def bfs(self, v):
        visited = set()
        queue = deque([v])
        result = []

        while queue:
            curr = queue.popleft()
            if curr in visited:
                continue
            visited.add(curr)
            result.append(curr)

            # Deterministic traversal
            for neighbor in sorted(self.graph[curr]):
                if neighbor not in visited:
                    queue.append(neighbor)
        return result

    def dfs(self, v):
        visited = set()
        stack = [v]
        result = []

        while stack:
            curr = stack.pop()
            if curr in visited:
                continue
            visited.add(curr)
            result.append(curr)

            # Deterministic traversal
            for neighbor in sorted(self.graph[curr], reverse=True):
                if neighbor not in visited:
                    stack.append(neighbor)

        return result
