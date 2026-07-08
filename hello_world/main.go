package main

import "fmt"

const (
	english  = "English"
	spanish  = "Spanish"
	french   = "French"
	japanese = "Japanese"
)

const (
	englishHelloPrefix  = "Hello, "
	spanishHelloPrefix  = "Hola, "
	frenchHelloPrefix   = "Bonjour, "
	japaneseHelloPrefix = "こんにちは、"
)

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case japanese:
		prefix = japaneseHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func worldPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = "Mundo"
	case french:
		prefix = "Monde"
	case japanese:
		prefix = "Sekai"
	default:
		prefix = "World"
	}
	return
}

func Hello(name string, language string) string {
	prefix := greetingPrefix(language)
	world := worldPrefix(language)
	
	if name == "" {
		name = world
	}
	return fmt.Sprintf("%s%s!", prefix, name)

}

func main() {
	fmt.Println(Hello("Brandon", english))
	fmt.Println(Hello("Brandon", spanish))
	fmt.Println(Hello("Brandon", french))
	fmt.Println(Hello("Brandon", japanese))
	fmt.Println(Hello("", ""))
	fmt.Println(Hello("", english))
	fmt.Println(Hello("", spanish))
	fmt.Println(Hello("", french))
	fmt.Println(Hello("", japanese))
}
