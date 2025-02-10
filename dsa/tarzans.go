package dsa

func tarzans(g *Graph) [][]string {
	id := 0
	ids, low, vis := make(map[string]int), make(map[string]int), make(map[string]bool)
	var startVertex string

	g.ForEachVertex(func(vtx *Vertex) {
		if startVertex == "" {
			startVertex = vtx.GetKey()
		}
	})

	if startVertex == "" {
		return nil
	}

	var bridges [][]string

	var dfs func(v, parent string)
	dfs = func(v, parent string) {
		vis[v] = true
		ids[v] = id
		low[v] = id
		id++

		g.GetVertex(v).ForEachEdge(func(n string, w int) {
			if n == parent {
				return
			}
			if !vis[n] {
				dfs(n, v)
				low[v] = min(low[v], low[n])
				if ids[v] < low[n] {
					bridges = append(bridges, []string{v, n})
				}
			} else {
				low[v] = min(low[v], ids[n])
			}
		})
	}

	dfs(startVertex, "")

	return bridges
}
