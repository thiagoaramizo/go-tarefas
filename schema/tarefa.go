package schema

import "gorm.io/gorm"

type Tarefa struct {
	gorm.Model
	Titulo string
	Descricao string
	Feita bool
	DataConclusao string
}