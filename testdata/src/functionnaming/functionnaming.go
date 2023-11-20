package functionnaming

func helloWorld() {}

func hello_world() {} // want "Function \"hello_world\" does not follow Go's naming conventions as it contains an underscore. Instead use Camel Case, for example \"examplePrivateFunctionName\" for private functions or \"ExamplePrivateFunctionName\" for public functions."

func HELLOWORLD() {} // want "Function \"HELLOWORLD\" does not follow Go's naming conventions as it contains only uppercase letters. Instead use Camel Case, for example \"examplePrivateFunctionName\" for private functions or \"ExamplePrivateFunctionName\" for public functions."