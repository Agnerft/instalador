package domain

type Cliente struct {
	Cliente          string `json:"cliente"`
	Documento        string `json:"documento"`
	Quantidade_ramal int    `json:"quantidade_ramal"`
	Link             string `json:"link"`
	Id               int    `json:"id"`
	Link_sip         string `json:"link_sip"`
	// RamaisRegistrados []Ramal
}

// type ClientesRegistrados struct {
// 	Clientes []Cliente
// }
