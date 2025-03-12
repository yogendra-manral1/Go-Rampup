package constants


type Constants struct{
	ContextKeys ContextKeys
}

type ContextKeys struct {
	EMAIL string
}

var constants Constants

func init(){
	contextKeys := ContextKeys{EMAIL: "email"}
	constants = Constants{
		ContextKeys: contextKeys,
	}
}

func GetConstants() *Constants {
	return &constants
}
