package languages

type Language string

type Languages struct {
	English Language
	Turkish Language
}

func (l Language) String() string {
	return string(l)
}

func New() Languages {
	return Languages{
		English: "en",
		Turkish: "tr",
	}
}
