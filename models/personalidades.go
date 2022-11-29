package models

// como eu quero criar meu json e passar os nomes
type Personalidade struct {
	Id       int    `json:"id"`
	Nome     string `json:"nome"`
	Historia string `json:"historia"`
}

// várias personalidade dentro de personalidades
var Personalidades []Personalidade
