package request

type Updatedatarequest struct {
	Uuid int    `validate:"required"`
	Name string `validate:"required, max=250, min=0" json:"name"`
}
