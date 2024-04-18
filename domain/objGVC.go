package domain

type ObjetoGvc struct {
	NameUsername string `json:"NameUsername"`
	Host         string `json:"host"`
	Dyn          string `json:"dyn"`
	Forcerport   string `json:"forcerport"`
	Comedia      string `json:"comedia"`
	ACL          string `json:"acl"`
	Port         string `json:"port"`
	Status       string `json:"status"`
	Description  string `json:"description"`
}

type AllObjGvc struct {
	ObjetosGvc []ObjetoGvc
}
