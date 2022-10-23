package redux

type Middleware func(state *State) *State

func ComposeActions(actions... Action) Action {
    if len(actions) == 0 {
		return Action {
			Type: "NULL",
			Payload: func(state *State) *State {
				return state
			},
		}
	}
    
    if len(actions) == 1 {
		return actions[0]
	}

	return Action {
		Type: "Compose Action",
		Payload: func(state *State) *State {
            for _, action := range actions {
				state = action.Payload(state)
			}
			return state
		},
	}
}

func CombineReducers (reducers... Reducer) Reducer {
	return func(state *State, action Action) *State {
		for _, reducer := range reducers {
			if reducer != nil {
				state = reducer(state, action)
			}
		}
		return state
	}
}