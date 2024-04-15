package domain

type Ramal struct {
	Sip int `json:"sip"`
	// Ip    string `json:"ip"`
	// InUse bool   `json:"inuse"`
	// Empresa string `json:"empresa"`
}

type RamalSolo struct {
	Ramais []Ramal `json:"ramais"`
}

type RamaisRegistrados struct {
	RamaisRegistrados []Ramal `json:"ramais_registrados"`
}
