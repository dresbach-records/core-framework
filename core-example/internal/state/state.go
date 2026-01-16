package state

type State map[string]any

func New() State {
	return State{}
}
