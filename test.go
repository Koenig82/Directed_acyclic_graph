package main

import (
	"fmt"
	"ou3/dag"
	"strconv"
	"strings")

type intWeight struct{
	weight int
}
func (w intWeight) GetWeightAsInt() int {
	return w.weight
}
func (w intWeight) ShowWeightVal() string {
	return strconv.Itoa(w.weight)
}
func (w intWeight) AddWeight(x dag.WeightUnit) dag.WeightUnit {
	return intWeight{w.weight + x.GetWeightAsInt()}
}
func (w intWeight) SubtractWeight(x dag.WeightUnit) dag.WeightUnit {
	return intWeight{w.weight - x.GetWeightAsInt()}
}
func (w intWeight) LessThan(x dag.WeightUnit) bool {
	return w.weight < x.GetWeightAsInt()
}
func (w intWeight) GreaterThan(x dag.WeightUnit) bool {
	return w.weight > x.GetWeightAsInt()
}
func (w intWeight) EqualTo(x dag.WeightUnit) bool {
	return w.weight == x.GetWeightAsInt()
}

type charWeight struct{
	weight string
}
func (w charWeight) GetWeightAsInt() int {
	var total int
	for _, r := range w.weight {
        c := int(r)
		reverseOrder := 0
		for c > 0 {
        	remainder := c % 10
        	reverseOrder *= 10
        	reverseOrder += remainder 
        	c /= 10
    	}
        total += reverseOrder
    }
	return total
}
func (w charWeight) ShowWeightVal() string {
	return w.weight
}
func (w charWeight) AddWeight(x dag.WeightUnit) dag.WeightUnit {
	return charWeight{w.ShowWeightVal() + x.ShowWeightVal()}
}
func (w charWeight) SubtractWeight(x dag.WeightUnit) dag.WeightUnit {
	if strings.HasSuffix(w.weight, x.ShowWeightVal()){
		w.weight = w.weight[:len(w.weight)-len(x.ShowWeightVal())]
	}
	return charWeight{w.ShowWeightVal()}
}
func (w charWeight) LessThan(x dag.WeightUnit) bool {
	return w.GetWeightAsInt() < x.GetWeightAsInt()
}
func (w charWeight) GreaterThan(x dag.WeightUnit) bool {
	return w.GetWeightAsInt() > x.GetWeightAsInt()
}
func (w charWeight) EqualTo(x dag.WeightUnit) bool {
	return w.GetWeightAsInt() == x.GetWeightAsInt()
}

func GetWeightUnit(wu dag.WeightUnit) dag.WeightUnit{
	return wu
}

func main() {
	fmt.Println("Creating DAG")
	newDag := dag.DAG{}

	fmt.Println("Trying to show empty dag:")
	if err := newDag.Show_DAG(); err != nil {
    	fmt.Println(err)
	}

	fmt.Println("Adding vertices with weight 2 and the following ids:")
	vert0 := newDag.Add_vertex(intWeight{2})
	vert1 := newDag.Add_vertex(intWeight{2})
	vert2 := newDag.Add_vertex(intWeight{2})
	vert3 := newDag.Add_vertex(intWeight{2})
	vert4 := newDag.Add_vertex(intWeight{2})
	vert5 := newDag.Add_vertex(intWeight{2})
	fmt.Println(vert0, vert1, vert2, vert3, vert4, vert5)

	fmt.Println("Showing dag with only unconnected vertices:")
	if err := newDag.Show_DAG(); err != nil {
    	fmt.Println(err)
	}

	fmt.Println("Adding the following edges(all with weight 4):")
	fmt.Println(vert0,"->",vert1)
	fmt.Println(vert1,"->",vert2)
	fmt.Println(vert3,"->",vert2)
	fmt.Println(vert0,"->",vert2)
	fmt.Println(vert2,"->",vert4)
	if err := newDag.Add_edge(vert0,vert1,intWeight{4}); err != nil {
    	fmt.Println(err)
	}
	if err := newDag.Add_edge(vert1,vert2,intWeight{4}); err != nil {
    	fmt.Println(err)
	}
	if err := newDag.Add_edge(vert3,vert2,intWeight{4}); err != nil {
    	fmt.Println(err)
	}
	if err := newDag.Add_edge(vert0,vert2,intWeight{4}); err != nil {
    	fmt.Println(err)
	}
	if err := newDag.Add_edge(vert2,vert4,intWeight{4}); err != nil {
    	fmt.Println(err)
	}
	
	fmt.Println("Showing dag with some connected vertices:")
	if err := newDag.Show_DAG(); err != nil {
    	fmt.Println(err)
	}

	fmt.Println("trying to add edge to nonexisting vertice A")
	if err := newDag.Add_edge(500,vert5,intWeight{5}); err != nil {
    	fmt.Println(err)
	}

	fmt.Println("trying to add edge to nonexisting vertice B")
	if err := newDag.Add_edge(vert5,500,intWeight{5}); err != nil {
    	fmt.Println(err)
	}

	fmt.Println("trying to add edge from",vert4,"to",vert1,"creating a cycle")
	if err := newDag.Add_edge(vert4,vert1,intWeight{5}); err != nil {
    	fmt.Println(err)
	}

	fmt.Println("A topological ordering of the dag:")
	if ordering, err := newDag.Topological_ordering(); err != nil {
		fmt.Println(err)
	}else{
		fmt.Println(ordering)
	}
	
	if longest, err := newDag.Weight_of_longest_path(vert0, vert4, GetWeightUnit, GetWeightUnit); err != nil {
    	fmt.Println(err)
	}else{
		fmt.Printf("Weight of the longest path between vertex %d and %d = %s\n",vert0, vert4, longest)
	}
	
	fmt.Println("\nCreating a new DAG with char weights ordered by reversed alphabetical order")
	newDag2 := dag.DAG{}

	fmt.Println("Adding vertices with weight \"c\" and the following ids:")
	a := newDag2.Add_vertex(charWeight{"c"})
	b := newDag2.Add_vertex(charWeight{"c"})
	c := newDag2.Add_vertex(charWeight{"c"})
	d := newDag2.Add_vertex(charWeight{"c"})
	e := newDag2.Add_vertex(charWeight{"c"})
	f := newDag2.Add_vertex(charWeight{"c"})
	fmt.Println(a, b, c, d, e, f)

	fmt.Println("Adding the following edges with different char weights:")
	fmt.Println(a,"->",b)
	fmt.Println(a,"->",c)
	fmt.Println(b,"->",d)
	fmt.Println(c,"->",d)
	fmt.Println(c,"->",e)
	fmt.Println(d,"->",f)
	fmt.Println(e,"->",f)
	if err := newDag2.Add_edge(a,b,charWeight{"c"}); err != nil {
    	fmt.Println(err)
	}
	if err := newDag2.Add_edge(a,c,charWeight{"c"}); err != nil {
    	fmt.Println(err)
	}
	if err := newDag2.Add_edge(b,d,charWeight{"a"}); err != nil {
    	fmt.Println(err)
	}
	if err := newDag2.Add_edge(c,d,charWeight{"b"}); err != nil {
    	fmt.Println(err)
	}
	if err := newDag2.Add_edge(c,e,charWeight{"a"}); err != nil {
    	fmt.Println(err)
	}
	if err := newDag2.Add_edge(d,f,charWeight{"d"}); err != nil {
    	fmt.Println(err)
	}
	if err := newDag2.Add_edge(e,f,charWeight{"a"}); err != nil {
    	fmt.Println(err)
	}
	
	fmt.Println("Showing dag:")
	if err := newDag2.Show_DAG(); err != nil {
    	fmt.Println(err)
	}

	fmt.Println("A topological ordering of the dag:")
	if ordering, err := newDag2.Topological_ordering(); err != nil {
		fmt.Println(err)
	}else{
		fmt.Println(ordering)
	}

	if longest, err := newDag2.Weight_of_longest_path(a, f, GetWeightUnit, GetWeightUnit); err != nil {
    	fmt.Println(err)
	}else{
		fmt.Printf("Weight of the longest path between vertex %d and %d = %s\n",a, f, longest)
	}
}


