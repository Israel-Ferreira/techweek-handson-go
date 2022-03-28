package exceptions

import "errors"

var ErrorNotFoundProduct = errors.New("error: produto não encontrado")
var ErrorBodyIsNotValid = errors.New("error: requisição inválida")
