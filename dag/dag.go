package dag

import "fmt"

type DAG struct {
	idCount int
	vertices map[int]vertex
	edges []edge
}

type vertex struct {
	id int
	weight WeightUnit
}
type edge struct {
	a, b vertex
	weight WeightUnit
}


type WeightUnit interface {
	GetWeightAsInt() int
	ShowWeightVal() string
	AddWeight(x WeightUnit) WeightUnit
	SubtractWeight(x WeightUnit) WeightUnit
	LessThan(x WeightUnit) bool
	GreaterThan(x WeightUnit) bool
	EqualTo(x WeightUnit) bool
}

func (dag *DAG) Add_vertex(weight WeightUnit) int{
	newVertex := vertex{dag.idCount, weight}
	if dag.vertices == nil {
    	dag.vertices = make(map[int]vertex)
	}
	dag.vertices[dag.idCount] = newVertex
	dag.idCount++
	return newVertex.id
}

func (dag *DAG) Add_edge(a int, b int, w WeightUnit) error{
	
	var va vertex
	var vb vertex
	if val, ok := dag.vertices[a]; ok {
    	va = val
	}else{
		return fmt.Errorf("Add_edge: Node A does not exist")
	}
	if val, ok := dag.vertices[b]; ok {
    	vb = val
	} else {
		return fmt.Errorf("Add_edge: Node B does not exist")
	}	
	
	dagCopy := *dag
	dagCopy.edges = append(dagCopy.edges, edge{va, vb, w})
	if _, err := dagCopy.Topological_ordering(); err != nil {
		return fmt.Errorf("Add_edge: edge creates cycle")
	} else {
		dag.edges = append(dag.edges, edge{va, vb, w})
		return nil
	}	
}


func (dag *DAG) Topological_ordering() ([]int, error){

	//get starting nodes and a graph
	var sortedNodes []int
	var graph map[vertex][]vertex
	var startNodes map[vertex]bool
	startNodes = make(map[vertex]bool)
	graph = make(map[vertex][]vertex)
	for _, v := range dag.vertices{
		startNodes[v] = true
		for _, e := range dag.edges{
			if e.b.id ==  v.id {
				delete(startNodes, v)			
				graph[v] = append(graph[v], e.a)
			}
		}
	}

	//while startNodes is non-empty do
	for {
        if len(startNodes) == 0 { break }

		//add n to tail of sortedNodes
		for n, _ := range startNodes{
			sortedNodes = append(sortedNodes, n.id)
			//for each node m with an edge e from n to m do
			for m, v := range graph{
				//remove edge e from the graph
				var tempEdges []vertex
				for _, e := range v{
					if e.id != n.id {
						tempEdges = append(tempEdges, e)
					}
				}
				graph[m] = tempEdges
				//if m has no other incoming edges then
				if len(graph[m]) < 1 {
					//insert m into startNodes
					startNodes[m] = true
					delete(graph, m)
				}
			}
			//remove a node n from startNodes
			delete(startNodes, n);
		}
	}

	if len(graph) > 0 {
		return nil, fmt.Errorf("Topological_ordering: Graph has cycle")
	} else {
		return sortedNodes, nil
	}
}

func (dag *DAG) Show_DAG() error{

	if len(dag.vertices) < 1 {
		return fmt.Errorf("Show_dag: No vertices in graph")
	}
	var verts map[vertex]bool
	verts = make(map[vertex]bool)
	for _, v := range dag.vertices{
		verts[v] = true
		for _, e := range dag.edges{
			if e.b.id == v.id || e.a.id == v.id {
				verts[v] = false
			}
		}
	}
	for k, v := range verts{
		if v == true {
			fmt.Printf("(id:%d,w:%s)\n",k.id, k.weight.ShowWeightVal())
		}
	}
	for _, e := range dag.edges {
		fmt.Printf("(id:%d,w:%s) -%s-> (id:%d,w:%s)\n",e.a.id, e.a.weight.ShowWeightVal(),e.weight.ShowWeightVal(),e.b.id,e.b.weight.ShowWeightVal())
	}
	return nil
}

func (dag *DAG) Weight_of_longest_path(a int, b int) (string, error) {

	var distances map[int]WeightUnit
	var postAddWeightToId map[int]bool
	distances = make(map[int]WeightUnit)
	postAddWeightToId = make(map[int]bool)
	for _, v := range dag.vertices{
		distances[v.id] = v.weight.SubtractWeight(v.weight)
	}
	
	var err error
	var topOrder []int
	if topOrder, err = dag.Topological_ordering(); err != nil {
		return "", fmt.Errorf("weight_of_longest_path: failed to create topological order",err.Error())
	}

	//Do following for every vertex u in topological order.
	for _, u := range topOrder{
		//Do following for every adjacent vertex v of u
		for _, e := range dag.edges{
			//if (dist[v] < dist[u] + weight(u, v))
			if e.b.id == u && distances[u].LessThan(distances[e.a.id].AddWeight(e.a.weight.AddWeight(e.weight))) {
				//dist[v] = dist[u] + weight(u, v)
				distances[u] = (distances[e.a.id].AddWeight(e.a.weight.AddWeight(e.weight)))
				postAddWeightToId[u] = true
			}
		}
	}

	for id, _ := range postAddWeightToId{
		distances[id] = distances[id].AddWeight(dag.vertices[id].weight)
	}
	
	return distances[b].ShowWeightVal(), nil
}
