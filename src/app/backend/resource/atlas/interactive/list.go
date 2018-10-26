package interactive

import (
  "github.com/kubernetes/dashboard/src/app/backend/api"
)

type ExampleList struct {
  ListMeta api.ListMeta `json:"listMeta"`
  ExampleSpecs []ExampleSpec `json:"exampleSpecs"`
  Errors []error `json:"errors"`
}

func GetExampleList(spec ExampleSpec) *ExampleList {
  return &ExampleList{
    ListMeta: api.ListMeta{TotalItems: 3},
    ExampleSpecs: []ExampleSpec{spec,spec,spec},
  }
}
