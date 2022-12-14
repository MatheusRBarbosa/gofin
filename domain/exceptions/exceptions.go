package exceptions

type SquiloError struct {
	Code         int
	InternalCode string
	Message      string
	Description  string
}

var (
	DEFAULT = newError(500, "0001", "Erro interno.", "Um erro interno do servidor aconteceu. Muito triste =(")

	// 1xxx - Generic validations
	UNAUTHORIZED       = newError(401, "1001", "Credenciais Inválidas!", "As credenciais que você digitou são inválidas. Veja se digitou tudo corretamente e tente novamente.")
	URL_PARAMS_MISSING = newError(500, "1002", "Parametros inválidos!", "Um ou mais parametros de URL estão faltando.")

	// 2xxx - Email
	EMAIL_ALREADY_EXIST = newError(400, "2001", "Email já está em uso.", "Caso já possua uma conta, tente recuperar a senha.")

	// 3xxx - Vault
	VAULT_NOT_FOUND        = newError(404, "3001", "Não encontrado.", "O cofre procurado não foi encontrado.")
	VAULT_NOT_USER_RELATED = newError(400, "3002", "Cofre não pertence a esse usuário.", "Cofre não pertence a esse usuário.")

	// 4xxx - Transaction
	TRANSACTION_NOT_FOUND            = newError(404, "4001", "Não encontrado.", "A transação procurada não foi encontrada.")
	TRANSACTION_WRONG_TYPE           = newError(400, "4002", "Transação não permitida.", "Verifique se o tipo de cofre permite essa transação.")
	TRANSACTION_NOT_BELONGS_TO_VAULT = newError(400, "4003", "Transação não pertence a esse cofre.", "Verifique se essa transação realmente pertence a esse cofre.")

	// 5xxx - User
	USER_NOT_FOUND = newError(404, "5001", "Não encontrado.", "O usuário procurado não foi encontrado.")
)

func (e *SquiloError) Error() string {
	return e.Message
}

func SlimError(code int, message string) *SquiloError {
	return &SquiloError{
		Message: message,
		Code:    code,
	}
}

func newError(code int, internalCode string, msg string, description string) *SquiloError {
	return &SquiloError{
		Message:      msg,
		Code:         code,
		InternalCode: internalCode,
		Description:  description,
	}
}

func GetError(e *SquiloError) *SquiloError {
	return &SquiloError{
		Message:     e.Message,
		Code:        e.Code,
		Description: e.Description,
	}
}
