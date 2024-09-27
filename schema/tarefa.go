package schema

import (
	"time"

	"gorm.io/gorm"
)

type Tarefa struct {
	gorm.Model
	Titulo string
	Descricao string
	Feita bool
	DataConclusao string
}

type TarefaDTO struct {
	ID uint `json:"id"`
	Titulo string `json:"titulo"`
	Descricao string `json:"descricao"`
	Feita bool `json:"feita"`
	DataConclusao string `json:"data_conclusao"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}