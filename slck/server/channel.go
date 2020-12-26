package server

// chat rooms
type channel struct {
	name    string
	clients map[*client]bool
}
