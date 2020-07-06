// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"log"
	"sort"

	"github.com/solo-io/solo-kit/pkg/api/v1/clients/kube/crd"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/solo-kit/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func NewRouteTable(namespace, name string) *RouteTable {
	routetable := &RouteTable{}
	routetable.SetMetadata(core.Metadata{
		Name:      name,
		Namespace: namespace,
	})
	return routetable
}

func (r *RouteTable) SetMetadata(meta core.Metadata) {
	r.Metadata = meta
}

func (r *RouteTable) SetStatus(status core.Status) {
	r.Status = status
}

func (r *RouteTable) MustHash() uint64 {
	hashVal, err := r.Hash(nil)
	if err != nil {
		log.Panicf("error while hashing: (%s) this should never happen", err)
	}
	return hashVal
}

func (r *RouteTable) GroupVersionKind() schema.GroupVersionKind {
	return RouteTableGVK
}

type RouteTableList []*RouteTable

func (list RouteTableList) Find(namespace, name string) (*RouteTable, error) {
	for _, routeTable := range list {
		if routeTable.GetMetadata().Name == name && routeTable.GetMetadata().Namespace == namespace {
			return routeTable, nil
		}
	}
	return nil, errors.Errorf("list did not find routeTable %v.%v", namespace, name)
}

func (list RouteTableList) AsResources() resources.ResourceList {
	var ress resources.ResourceList
	for _, routeTable := range list {
		ress = append(ress, routeTable)
	}
	return ress
}

func (list RouteTableList) AsInputResources() resources.InputResourceList {
	var ress resources.InputResourceList
	for _, routeTable := range list {
		ress = append(ress, routeTable)
	}
	return ress
}

func (list RouteTableList) Names() []string {
	var names []string
	for _, routeTable := range list {
		names = append(names, routeTable.GetMetadata().Name)
	}
	return names
}

func (list RouteTableList) NamespacesDotNames() []string {
	var names []string
	for _, routeTable := range list {
		names = append(names, routeTable.GetMetadata().Namespace+"."+routeTable.GetMetadata().Name)
	}
	return names
}

func (list RouteTableList) Sort() RouteTableList {
	sort.SliceStable(list, func(i, j int) bool {
		return list[i].GetMetadata().Less(list[j].GetMetadata())
	})
	return list
}

func (list RouteTableList) Clone() RouteTableList {
	var routeTableList RouteTableList
	for _, routeTable := range list {
		routeTableList = append(routeTableList, resources.Clone(routeTable).(*RouteTable))
	}
	return routeTableList
}

func (list RouteTableList) Each(f func(element *RouteTable)) {
	for _, routeTable := range list {
		f(routeTable)
	}
}

func (list RouteTableList) EachResource(f func(element resources.Resource)) {
	for _, routeTable := range list {
		f(routeTable)
	}
}

func (list RouteTableList) AsInterfaces() []interface{} {
	var asInterfaces []interface{}
	list.Each(func(element *RouteTable) {
		asInterfaces = append(asInterfaces, element)
	})
	return asInterfaces
}

// Kubernetes Adapter for RouteTable

func (o *RouteTable) GetObjectKind() schema.ObjectKind {
	t := RouteTableCrd.TypeMeta()
	return &t
}

func (o *RouteTable) DeepCopyObject() runtime.Object {
	return resources.Clone(o).(*RouteTable)
}

func (o *RouteTable) DeepCopyInto(out *RouteTable) {
	clone := resources.Clone(o).(*RouteTable)
	*out = *clone
}

var (
	RouteTableCrd = crd.NewCrd(
		"routetables",
		RouteTableGVK.Group,
		RouteTableGVK.Version,
		RouteTableGVK.Kind,
		"rt",
		false,
		&RouteTable{})
)

func init() {
	if err := crd.AddCrd(RouteTableCrd); err != nil {
		log.Fatalf("could not add crd to global registry")
	}
}

var (
	RouteTableGVK = schema.GroupVersionKind{
		Version: "v1",
		Group:   "gateway.solo.io",
		Kind:    "RouteTable",
	}
)
