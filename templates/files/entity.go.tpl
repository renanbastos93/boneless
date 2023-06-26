package {{.ComponentName}}

import (
	"time"

	"{{.Module}}/{{.ComponentName}}/store"
)

type ExampleOut struct {
	ID        int       `json:"id,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	Message   string    `json:"message,omitempty"`
}

type AllExamplesOut []ExampleOut

func (e ExampleOut) FromStore(in store.Example) ExampleOut {
	e.ID = int(in.ID)
	e.CreatedAt = time.UnixMilli(in.CreatedAt)
	e.Message = in.Message
	return e
}

func (e AllExamplesOut) FromStore(in []store.Example) AllExamplesOut {
	e = make(AllExamplesOut, 0, len(in))
	for _, v := range in {
		e = append(e, ExampleOut{}.FromStore(v))
	}
	return e
}

type ExampleIn struct {
	Message string
}

func (e ExampleIn) ToStore() (params store.CreateExampleParams) {
	t := time.Now()
	params.ID = int32(t.UTC().UnixNano()) // fake id
	params.CreatedAt = t.UnixMilli()
	params.Message = e.Message
	return params
}
