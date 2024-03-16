package request

type Createdatarequest struct {
	Name string `"validate:"required, min=0, max=25" json="name"`
}
