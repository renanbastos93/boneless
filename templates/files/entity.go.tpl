package {{.ComponentName}}

import (
	"time"

	"github.com/ServiceWeaver/weaver"
	"{{.Module}}/internal/{{.ComponentName}}/store"
)

type ExampleOut struct {
	weaver.AutoMarshal
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
	weaver.AutoMarshal
	Message string
}

func (e ExampleIn) ToStore() (params store.CreateExampleParams) {
	t := time.Now()
	params.ID = int32(t.UTC().UnixNano()) // fake id
	params.CreatedAt = t.UnixMilli()
	params.Message = e.Message
	return params
}
