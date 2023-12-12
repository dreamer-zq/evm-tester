package simple

// Manager is a struct that manages a map of samplers.
type Manager struct {
	ms map[string]Contract
}

// NewManager creates a new Manager.
//
// It returns a pointer to a Manager struct.
func NewManager() *Manager {
	ms := make(map[string]Contract)
	ms["eTicket"] = &ETicketSampler{}
	ms["ticket"] = &TicketGameSampler{}
	return &Manager{
		ms: ms,
	}
}

// ListContracts lists all the samplers in the Manager.
//
// It returns a slice of strings, containing the names of the samplers.
func (m *Manager) ListContracts() (names []string) {
	for name := range m.ms {
		names = append(names, name)
	}
	return
}


// GetContract returns the Contract with the given name.
//
// It takes a string parameter called name which is the name of the Contract to retrieve.
// It returns a Contract.
func (m *Manager) GetContract(name string) Contract {
	sampler, ok := m.ms[name]
	if !ok {
		panic("invalid sampler name")
	}
	return sampler
}
