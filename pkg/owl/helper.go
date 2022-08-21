/*
This file is auto-generated by OWL2Go (https://git.rwth-aachen.de/acs/public/ontology/owl/owl2go).

Copyright 2020 Institute for Automation of Complex Power Systems,
E.ON Energy Research Center, RWTH Aachen University

This project is licensed under either of
- Apache License, Version 2.0
- MIT License
at your option.

Apache License, Version 2.0:

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

MIT License:

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package owl

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"git.rwth-aachen.de/acs/public/ontology/owl/owl2go/pkg/rdf"
	"github.com/piprate/json-gold/ld"
)

// AddObjectToGraph adds the specified object to the graph
func AddObjectToGraph(g *rdf.Graph, typeIRI string, res Thing) (node *rdf.Node) {
	var ok bool
	if node, ok = g.Nodes[res.IRI()]; !ok {
		if isIRI(res.IRI()) {
			node = &rdf.Node{Term: rdf.NewIRI(res.IRI())}
		} else {
			node = &rdf.Node{Term: rdf.NewBlankNode(res.IRI())}
		}
		g.Nodes[res.IRI()] = node
	}
	var typ *rdf.Node
	if typ, ok = g.Nodes[typeIRI]; !ok {
		typ = &rdf.Node{Term: rdf.NewIRI(typeIRI)}
		g.Nodes[typeIRI] = typ
	}
	pred := &rdf.Edge{
		Pred:    rdf.NewIRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#type"),
		Subject: node,
		Object:  typ,
	}
	g.Edges = append(g.Edges, pred)
	node.Edge = append(node.Edge, pred)
	typ.InverseEdge = append(typ.InverseEdge, pred)
	var nInd *rdf.Node
	if nInd, ok = g.Nodes["http://www.w3.org/2002/07/owl#NamedIndividual"]; !ok {
		nInd = &rdf.Node{Term: rdf.NewIRI("http://www.w3.org/2002/07/owl#NamedIndividual")}
		g.Nodes["http://www.w3.org/2002/07/owl#NamedIndividual"] = nInd
	}
	pred = &rdf.Edge{
		Pred:    rdf.NewIRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#type"),
		Subject: node,
		Object:  nInd,
	}
	g.Edges = append(g.Edges, pred)
	node.Edge = append(node.Edge, pred)
	nInd.InverseEdge = append(nInd.InverseEdge, pred)
	return
}

// AddClassPropertyToGraph adds the specified property to the graph
func AddClassPropertyToGraph(g *rdf.Graph, propIRI string, subjNode *rdf.Node, obj Thing) {
	if obj == nil {
		return
	}
	objNode, ok := g.Nodes[obj.IRI()]
	if !ok {
		if isIRI(obj.IRI()) {
			objNode = &rdf.Node{Term: rdf.NewIRI(obj.IRI())}
		} else {
			objNode = &rdf.Node{Term: rdf.NewBlankNode(obj.IRI())}
		}
		g.Nodes[obj.IRI()] = objNode
	}
	pred := &rdf.Edge{
		Pred:    rdf.NewIRI(propIRI),
		Object:  objNode,
		Subject: subjNode,
	}
	subjNode.Edge = append(subjNode.Edge, pred)
	objNode.InverseEdge = append(objNode.InverseEdge, pred)
	g.Edges = append(g.Edges, pred)
	return
}

// AddIntPropertyToGraph adds the specified property to the graph
func AddIntPropertyToGraph(g *rdf.Graph, propIRI string, subjNode *rdf.Node, obj int) {
	lit, _ := rdf.NewLiteral(obj, "")
	objNode, ok := g.Nodes[fmt.Sprintf("%v", obj)]
	if !ok {
		objNode = &rdf.Node{Term: lit}
		g.Nodes[fmt.Sprintf("%v", obj)] = objNode
	}
	pred := &rdf.Edge{
		Pred:    rdf.NewIRI(propIRI),
		Object:  objNode,
		Subject: subjNode,
	}
	subjNode.Edge = append(subjNode.Edge, pred)
	objNode.InverseEdge = append(objNode.InverseEdge, pred)
	g.Edges = append(g.Edges, pred)
	return
}

// AddFloatPropertyToGraph adds the specified property to the graph
func AddFloatPropertyToGraph(g *rdf.Graph, propIRI string, subjNode *rdf.Node, obj float64) {
	lit, _ := rdf.NewLiteral(obj, "")
	objNode, ok := g.Nodes[fmt.Sprintf("%v", obj)]
	if !ok {
		objNode = &rdf.Node{Term: lit}
		g.Nodes[fmt.Sprintf("%v", obj)] = objNode
	}
	pred := &rdf.Edge{
		Pred:    rdf.NewIRI(propIRI),
		Object:  objNode,
		Subject: subjNode,
	}
	subjNode.Edge = append(subjNode.Edge, pred)
	objNode.InverseEdge = append(objNode.InverseEdge, pred)
	g.Edges = append(g.Edges, pred)
	return
}

// AddStringPropertyToGraph adds the specified property to the graph
func AddStringPropertyToGraph(g *rdf.Graph, propIRI string, subjNode *rdf.Node, obj string) {
	if obj == "" {
		return
	}
	lit, _ := rdf.NewLiteral(obj, "")
	objNode, ok := g.Nodes[obj]
	if !ok {
		objNode = &rdf.Node{Term: lit}
		g.Nodes[obj] = objNode
	}
	pred := &rdf.Edge{
		Pred:    rdf.NewIRI(propIRI),
		Object:  objNode,
		Subject: subjNode,
	}
	subjNode.Edge = append(subjNode.Edge, pred)
	objNode.InverseEdge = append(objNode.InverseEdge, pred)
	g.Edges = append(g.Edges, pred)
	return
}

// AddBoolPropertyToGraph adds the specified property to the graph
func AddBoolPropertyToGraph(g *rdf.Graph, propIRI string, subjNode *rdf.Node, obj bool) {
	lit, _ := rdf.NewLiteral(obj, "")
	objNode, ok := g.Nodes[fmt.Sprintf("%v", obj)]
	if !ok {
		objNode = &rdf.Node{Term: lit}
		g.Nodes[fmt.Sprintf("%v", obj)] = objNode
	}
	pred := &rdf.Edge{
		Pred:    rdf.NewIRI(propIRI),
		Object:  objNode,
		Subject: subjNode,
	}
	subjNode.Edge = append(subjNode.Edge, pred)
	objNode.InverseEdge = append(objNode.InverseEdge, pred)
	g.Edges = append(g.Edges, pred)
}

// AddInterfacePropertyToGraph adds the specified property to the graph
func AddInterfacePropertyToGraph(g *rdf.Graph, propIRI string, subjNode *rdf.Node, obj interface{}) {
	lit, _ := rdf.NewLiteral(obj, "")
	objNode, ok := g.Nodes[fmt.Sprintf("%v", obj)]
	if !ok {
		objNode = &rdf.Node{Term: lit}
		g.Nodes[fmt.Sprintf("%v", obj)] = objNode
	}
	pred := &rdf.Edge{
		Pred:    rdf.NewIRI(propIRI),
		Object:  objNode,
		Subject: subjNode,
	}
	subjNode.Edge = append(subjNode.Edge, pred)
	objNode.InverseEdge = append(objNode.InverseEdge, pred)
	g.Edges = append(g.Edges, pred)
}

// AddDurationPropertyToGraph adds the specified property to the graph
func AddDurationPropertyToGraph(g *rdf.Graph, propIRI string, subjNode *rdf.Node, obj time.Duration) {
	var temp time.Duration
	if obj == temp {
		return
	}
	lit, _ := rdf.NewLiteral(obj, "")
	var objNode *rdf.Node
	if objNode, ok := g.Nodes[lit.String()]; !ok {
		objNode = &rdf.Node{Term: lit}
		g.Nodes[lit.String()] = objNode
	}
	pred := &rdf.Edge{
		Pred:    rdf.NewIRI(propIRI),
		Object:  objNode,
		Subject: subjNode,
	}
	subjNode.Edge = append(subjNode.Edge, pred)
	objNode.InverseEdge = append(objNode.InverseEdge, pred)
	g.Edges = append(g.Edges, pred)
}

// AddTimePropertyToGraph adds the specified property to the graph
func AddTimePropertyToGraph(g *rdf.Graph, propIRI string, subjNode *rdf.Node, obj time.Time) {
	var temp time.Time
	if obj == temp {
		return
	}
	lit, _ := rdf.NewLiteral(obj, rdf.XsdTime)
	objNode, ok := g.Nodes[lit.String()]
	if !ok {
		objNode = &rdf.Node{Term: lit}
		g.Nodes[lit.String()] = objNode
	}
	pred := &rdf.Edge{
		Pred:    rdf.NewIRI(propIRI),
		Object:  objNode,
		Subject: subjNode,
	}
	subjNode.Edge = append(subjNode.Edge, pred)
	objNode.InverseEdge = append(objNode.InverseEdge, pred)
	g.Edges = append(g.Edges, pred)
}

// AddDateTimePropertyToGraph adds the specified property to the graph
func AddDateTimePropertyToGraph(g *rdf.Graph, propIRI string, subjNode *rdf.Node, obj time.Time) {
	var temp time.Time
	if obj == temp {
		return
	}
	lit, _ := rdf.NewLiteral(obj, rdf.XsdDateTime)
	objNode, ok := g.Nodes[lit.String()]
	if !ok {
		objNode = &rdf.Node{Term: lit}
		g.Nodes[lit.String()] = objNode
	}
	pred := &rdf.Edge{
		Pred:    rdf.NewIRI(propIRI),
		Object:  objNode,
		Subject: subjNode,
	}
	subjNode.Edge = append(subjNode.Edge, pred)
	objNode.InverseEdge = append(objNode.InverseEdge, pred)
	g.Edges = append(g.Edges, pred)
}

// AddDatePropertyToGraph adds the specified property to the graph
func AddDatePropertyToGraph(g *rdf.Graph, propIRI string, subjNode *rdf.Node, obj time.Time) {
	var temp time.Time
	if obj == temp {
		return
	}
	lit, _ := rdf.NewLiteral(obj, rdf.XsdDate)
	objNode, ok := g.Nodes[lit.String()]
	if !ok {
		objNode = &rdf.Node{Term: lit}
		g.Nodes[lit.String()] = objNode
	}
	pred := &rdf.Edge{
		Pred:    rdf.NewIRI(propIRI),
		Object:  objNode,
		Subject: subjNode,
	}
	subjNode.Edge = append(subjNode.Edge, pred)
	objNode.InverseEdge = append(objNode.InverseEdge, pred)
	g.Edges = append(g.Edges, pred)
}

// AddDateTimeStampPropertyToGraph adds the specified property to the graph
func AddDateTimeStampPropertyToGraph(g *rdf.Graph, propIRI string, subjNode *rdf.Node, obj time.Time) {
	var temp time.Time
	if obj == temp {
		return
	}
	lit, _ := rdf.NewLiteral(obj, rdf.XsdDateTimeStamp)
	objNode, ok := g.Nodes[lit.String()]
	if !ok {
		objNode = &rdf.Node{Term: lit}
		g.Nodes[lit.String()] = objNode
	}
	pred := &rdf.Edge{
		Pred:    rdf.NewIRI(propIRI),
		Object:  objNode,
		Subject: subjNode,
	}
	subjNode.Edge = append(subjNode.Edge, pred)
	objNode.InverseEdge = append(objNode.InverseEdge, pred)
	g.Edges = append(g.Edges, pred)
}

// AddGYearPropertyToGraph adds the specified property to the graph
func AddGYearPropertyToGraph(g *rdf.Graph, propIRI string, subjNode *rdf.Node, obj time.Time) {
	var temp time.Time
	if obj == temp {
		return
	}
	lit, _ := rdf.NewLiteral(obj, rdf.XsdYear)
	objNode, ok := g.Nodes[lit.String()]
	if !ok {
		objNode = &rdf.Node{Term: lit}
		g.Nodes[lit.String()] = objNode
	}
	pred := &rdf.Edge{
		Pred:    rdf.NewIRI(propIRI),
		Object:  objNode,
		Subject: subjNode,
	}
	subjNode.Edge = append(subjNode.Edge, pred)
	objNode.InverseEdge = append(objNode.InverseEdge, pred)
	g.Edges = append(g.Edges, pred)
}

// AddGDayPropertyToGraph adds the specified property to the graph
func AddGDayPropertyToGraph(g *rdf.Graph, propIRI string, subjNode *rdf.Node, obj time.Time) {
	var temp time.Time
	if obj == temp {
		return
	}
	lit, _ := rdf.NewLiteral(obj, rdf.XsdDay)
	objNode, ok := g.Nodes[lit.String()]
	if !ok {
		objNode = &rdf.Node{Term: lit}
		g.Nodes[lit.String()] = objNode
	}
	pred := &rdf.Edge{
		Pred:    rdf.NewIRI(propIRI),
		Object:  objNode,
		Subject: subjNode,
	}
	subjNode.Edge = append(subjNode.Edge, pred)
	objNode.InverseEdge = append(objNode.InverseEdge, pred)
	g.Edges = append(g.Edges, pred)
}

// AddGYearMonthPropertyToGraph adds the specified property to the graph
func AddGYearMonthPropertyToGraph(g *rdf.Graph, propIRI string, subjNode *rdf.Node, obj time.Time) {
	var temp time.Time
	if obj == temp {
		return
	}
	lit, _ := rdf.NewLiteral(obj, rdf.XsdYearMonth)
	objNode, ok := g.Nodes[lit.String()]
	if !ok {
		objNode = &rdf.Node{Term: lit}
		g.Nodes[lit.String()] = objNode
	}
	pred := &rdf.Edge{
		Pred:    rdf.NewIRI(propIRI),
		Object:  objNode,
		Subject: subjNode,
	}
	subjNode.Edge = append(subjNode.Edge, pred)
	objNode.InverseEdge = append(objNode.InverseEdge, pred)
	g.Edges = append(g.Edges, pred)
}

// AddGMonthPropertyToGraph adds the specified property to the graph
func AddGMonthPropertyToGraph(g *rdf.Graph, propIRI string, subjNode *rdf.Node, obj time.Time) {
	var temp time.Time
	if obj == temp {
		return
	}
	lit, _ := rdf.NewLiteral(obj, rdf.XsdMonth)
	objNode, ok := g.Nodes[lit.String()]
	if !ok {
		objNode = &rdf.Node{Term: lit}
		g.Nodes[lit.String()] = objNode
	}
	pred := &rdf.Edge{
		Pred:    rdf.NewIRI(propIRI),
		Object:  objNode,
		Subject: subjNode,
	}
	subjNode.Edge = append(subjNode.Edge, pred)
	objNode.InverseEdge = append(objNode.InverseEdge, pred)
	g.Edges = append(g.Edges, pred)
}

// ParseXsdDuration parses xsdDuration
func ParseXsdDuration(in string) (out time.Duration, err error) {
	p := strings.TrimPrefix(in, "P")
	str := ""
	h := strings.Split(p, "H")
	if len(h) > 1 {
		if n, err := strconv.Atoi(h[0]); err == nil && n > 0 {
			str += h[0] + "h"
		}
	}
	m := strings.Split(h[len(h)-1], "M")
	if len(m) > 1 {
		if n, err := strconv.Atoi(m[0]); err == nil && n > 0 {
			str += m[0] + "m"
		}
	}
	s := strings.Split(m[len(m)-1], "S")
	if len(s) > 1 {
		if n, err := strconv.ParseFloat(s[0], 32); err == nil && n > 0 {
			str += s[0] + "s"
		}
	}
	out, err = time.ParseDuration(str)
	return
}

// isIRI checks if string is valid iri
func isIRI(iri string) (ok bool) {
	ok = false
	if ld.IsURL(iri) {
		ok = true
	}
	return
}
