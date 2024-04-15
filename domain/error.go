package domain

// HTTPError é uma estrutura para representar erros HTTP personalizados
type HTTPError struct {
	StatusCode int    // Código de status HTTP
	Message    string // Mensagem de erro
}

func (e HTTPError) Error() string {
	return e.Message
}
