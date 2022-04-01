package exceptions

import "errors"

var ErrorNotFoundItem = errors.New("error: item não encontrado")
var ErrorInvalidTitle = errors.New("error: titulo inválido")
var ErrorInvalidSKU = errors.New("error: SKU inválido")

var ErrorInvalidParam = errors.New("error: parametro inválido")
var ErrorInvalidQty = errors.New("error; invalid quantity")
