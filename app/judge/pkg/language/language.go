package language

import "fmt"

type Language int64

const (
	Language_CPP     Language = 0
	Language_C       Language = 1
	Language_Python3 Language = 2
	Language_Java    Language = 3
	Language_Go      Language = 4
)

func (p Language) String() string {
	switch p {
	case Language_CPP:
		return "CPP"
	case Language_C:
		return "C"
	case Language_Python3:
		return "Python3"
	case Language_Java:
		return "Java"
	case Language_Go:
		return "Go"
	}
	return "<UNSET>"
}

func (p Language) FileName() string {
	switch p {
	case Language_CPP:
		return "main.cpp"
	case Language_C:
		return "main.c"
	case Language_Python3:
		return "main.py"
	case Language_Java:
		return "Main.java"
	case Language_Go:
		return "main.go"
	}
	return "<UNSET>"
}

func FromString(s string) (Language, error) {
	switch s {
	case "CPP":
		return Language_CPP, nil
	case "C":
		return Language_C, nil
	case "Python3":
		return Language_Python3, nil
	case "Java":
		return Language_Java, nil
	case "Go":
		return Language_Go, nil
	}
	return Language(0), fmt.Errorf("not a valid Language string")
}
