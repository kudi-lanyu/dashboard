package interactive

import (
	"github.com/kubernetes/dashboard/src/app/backend/api"
	"k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
)

type DeploySpec struct {
	// Name of the application.
	Name string `json:"name"`

	// Arguments for the specified container command or container entrypoint (if command is not
	// specified here).
	Ptr *string `json:"ptr"`

	// Number of replicas of the image to maintain.
	Nums int32 `json:"nums"`

	List []NestedStruct `json:"list"`

	Map map[string]NestObject `json:"map"`

	Bool bool `json:"bool"`
}

type ExampleSpec struct {
	ObjectMeta api.ObjectMeta     `json:"objectMeta"`
	TypeMeta   api.TypeMeta       `json:"typeMeta"`
	Ready      v1.ConditionStatus `json:"ready"`
	// Name of the application.
	Name string `json:"name"`

	// Number of replicas of the image to maintain.
	Nums int32 `json:"nums"`

	List []NestedStruct `json:"list"`

	Map map[string]NestObject `json:"map"`

	Bool bool `json:"bool"`
}

type NestObject struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type NestedStruct struct {
	Field1 string `json:"field_1"`
	Field2 int    `json:"field_2"`
}

func GetExampleSpec(client kubernetes.Interface) *ExampleSpec {
	nestObject := NestObject{Name: "nestObjectName", Value: "nestObjectValue"}
	return &ExampleSpec{
		ObjectMeta: api.ObjectMeta{},
		TypeMeta:   api.TypeMeta{"define by myself"},
		Ready:      v1.ConditionTrue,
		Name:       "specname",
		Nums:       32,
		List:       []NestedStruct{{Field1: "field1", Field2: 123}, {Field1: "field11", Field2: 123}},
		Map:        map[string]NestObject{"k": nestObject, "l": nestObject},
		Bool:       true,
	}
}
