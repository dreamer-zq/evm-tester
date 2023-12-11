package simple

// Manager is a struct that manages a map of samplers.
type Manager struct {
	ms map[string]Sampler
}

// NewManager creates a new Manager.
//
// It returns a pointer to a Manager struct.
func NewManager() *Manager {
	ms := make(map[string]Sampler)
	ms["eTicket"] = &ETicketSampler{}
	ms["ticket"] = &TicketGameSampler{}
	return &Manager{
		ms: ms,
	}
}

// ListSamplers lists all the samplers in the Manager.
//
// It returns a slice of strings, containing the names of the samplers.
func (m *Manager) ListSamplers() (names []string) {
	for name := range m.ms {
		names = append(names, name)
	}
	return
}

// GetSampler returns the sampler with the given name.
//
// Parameters:
// - name: the name of the sampler to retrieve.
//
// Return type:
// - Sampler: the requested sampler.
func (m *Manager) GetSampler(name string) Sampler {
	sampler, ok := m.ms[name]
	if !ok {
		panic("invalid sampler name")
	}
	return sampler
}
