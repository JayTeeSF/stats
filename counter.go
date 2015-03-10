package stats

type counter struct {
	context       string
	key           string
	value         int
	collected_at  int
	collected_for int
}

func (c counter) Context() string {
	return c.context
}

// perhaps this should return a person.Person object !?
func (c counter) CollectedFor() int {
	return c.collected_for
}

func (c counter) CollectedAt() int {
	return c.collected_at
}

func (c counter) Value() int {
	return c.value
}

func (c counter) Key() string {
	return c.key
}
