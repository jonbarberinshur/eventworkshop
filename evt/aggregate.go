package evt

// Aggregate interface
type Aggregate interface {
	EventHandlers() map[string]EventHandler
}
