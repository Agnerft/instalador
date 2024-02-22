package domain

type Ramal struct {
	Sip   string `json:"sip"`
	Ip    string `json:"ip"`
	InUse bool   `json:"inuse"`
	// Empresa string `json:"empresa"`
}

type RamaisRegistrados struct {
	RamaisRegistrados []Ramal `json:"ramais_registrados"`
}
