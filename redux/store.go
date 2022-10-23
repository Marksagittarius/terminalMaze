package redux

import (
	"fmt"
	"sync"
)

type Listener func(state *State) error

type Action struct {
	Type    string
	Payload func(*State) *State
}

type Reducer func(state *State, action Action) *State

type State struct {
	HashMap map[string]any
}

type Store struct {
	sync.RWMutex
	currentState     *State
	currentReducer   Reducer
	currentListeners []Listener
	nextListeners    []Listener
	middlewares      []Middleware
	isDispatching    bool
	isSubscribed     bool
}

func CreateStore(reducer Reducer, preloadState *State) *Store {
	return &Store{
		currentReducer:   reducer,
		currentState:     preloadState,
		currentListeners: make([]Listener, 0),
		nextListeners:    make([]Listener, 0),
		middlewares:      make([]Middleware, 0),
		isDispatching:    false,
		isSubscribed:     false,
	}
}

func (store *Store) GetState() (*State, error) {
	if store.isDispatching {
		return nil, fmt.Errorf("You may not call store.getState() while the reducer is executing." +
			"The reducer has already received the state as an argument. " +
			"Pass it down from the top reducer instead of reading it from the store.")
	}

	return store.currentState, nil
}

func (store *Store) Subscribe(listener ...Listener) (func() error, error) {
	if listener == nil {
		return nil, fmt.Errorf("Expected listener to be a function.")
	}

	if store.isDispatching {
		return nil, fmt.Errorf("You may not call store.subscribe() while the reducer is executing. " +
			"If you would like to be notified after the store has been updated, subscribe from a " +
			"component and invoke store.getState() in the callback to access the latest state. " +
			"See http://redux.js.org/docs/api/Store.html#subscribe for more details.")

	}

	store.Lock()
	preLength := len(store.nextListeners)
	store.nextListeners = append(store.nextListeners, listener...)
	store.Unlock()

	return func() error {
		if !store.isSubscribed {
			return nil
		}

		if store.isDispatching {
			return fmt.Errorf("You may not unsubscribe from a store listener while the reducer is executing. " +
				"See http://redux.js.org/docs/api/Store.html#subscribe for more details.")
		}

		store.Lock()
		store.isSubscribed = false
		store.nextListeners = store.nextListeners[0:preLength]
		store.Unlock()

		return nil
	}, nil
}

func (store *Store) Dispatch(actions ...Action) ([]Action, error) {
	if actions == nil {
		return nil, fmt.Errorf("Actions may not have an undefined type property. ")
	}

	if store.isDispatching {
		return nil, fmt.Errorf("Reducers may not dispatch actions.")
	}

	store.Lock()
	store.isDispatching = true

	for _, middleware := range store.middlewares {
		store.currentState = store.currentReducer(store.currentState, Action{
			Type:    "Middleware",
			Payload: middleware,
		})
	}

	for _, action := range actions {
		store.currentState = store.currentReducer(store.currentState, action)
	}
	store.isDispatching = false
	store.Unlock()

	store.Lock()
	store.currentListeners = store.nextListeners
	for _, listener := range store.currentListeners {
		err := listener(store.currentState)
		if err != nil {
			return nil, err
		}
	}
	store.Unlock()

	return actions, nil
}

func (store *Store) ReplaceReducer(nextReducer Reducer) error {
	if nextReducer == nil {
		return fmt.Errorf("Expected the nextReducer to be a function.")
	}

	store.Lock()
	store.currentReducer = nextReducer
	store.Unlock()
	store.Dispatch(Action{
		Type: "ActionTypes.REPLACE",
		Payload: func(state *State) *State {
			return state
		},
	})

	return nil
}

func (store *Store) ApplyMiddlewares(middlewares ...Middleware) (func() error, error) {
	if middlewares == nil {
		return nil, fmt.Errorf("Expected middleware to be a function.")
	}

	if store.isDispatching {
		return nil, fmt.Errorf("You may not call store.applyMiddleware() while the reducer is executing. " +
			"If you would like to be notified after the store has been updated, subscribe from a " +
			"component and invoke store.getState() in the callback to access the latest state. " +
			"See http://redux.js.org/docs/api/Store.html#subscribe for more details.")

	}

	store.Lock()
	preLength := len(store.middlewares)
	store.middlewares = append(store.middlewares, middlewares...)
	store.Unlock()

	return func() error {

		if store.isDispatching {
			return fmt.Errorf("You may not remove a store middleware while the reducer is executing. " +
				"See http://redux.js.org/docs/api/Store.html#subscribe for more details.")
		}

		store.Lock()
		store.middlewares = store.middlewares[0:preLength]
		store.Unlock()

		return nil
	}, nil
}
