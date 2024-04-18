// Code generated by Cuelang code gen. DO NOT EDIT!
// Cuelang source: github.com/specterops/bloodhound/-/tree/main/packages/cue/schemas/

package graphschema

import graph "github.com/specterops/bloodhound/dawgs/graph"

type KindDescriptor struct {
	Kind graph.Kind
	Name string
}

func (s KindDescriptor) GetName() string {
	if s.Name == "" {
		return s.Kind.String()
	}
	return s.Name
}

type Path struct {
	Outbound      KindDescriptor
	Inbound       KindDescriptor
	Relationships []KindDescriptor
}
