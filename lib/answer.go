package answer

type Answer struct {
	Type	string
	Desc  string
	Get	func(index int) string
	Rand	func() string
	DB    []string
}

var Answers map[string]*Answer

func Init(){
Answers = make(map[string]*Answer)
}

func Register(a *Answer){
	Answers[a.Type] = a
}
