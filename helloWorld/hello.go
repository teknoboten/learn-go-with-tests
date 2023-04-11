package main

func Greeting(language string) string {
	greeting := "Hello, "
	
	switch language {
	case "french":
		greeting = "Bonjour, "
	case "spanish":
		greeting = "Hola, "
	case "german":
		greeting = "Hallo, "
	}
	return greeting
}


func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return Greeting(language) + name
}




