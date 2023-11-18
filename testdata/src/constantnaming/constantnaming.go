package constantnaming

var notConst string

const myConst = "test"

const my_const_name = "test" // want "Constant \"my_const_name\" does not follow Go's naming conventions as it contains an underscore. Instead use CamelCase, for example: exampleConstantName."

const MYCONSTNAME = "test" // want "Constant \"MYCONSTNAME\" does not follow Go's naming conventions as it contains only uppercase letters. Instead use CamelCase, for example: exampleConstantName."
