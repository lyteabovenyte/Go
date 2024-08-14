package store

type Researcher struct {
	*Product
	startup  bool
	projects int
	Wanted   float64 //optional
}

type ResearcherOptions func(r *Researcher)

func NewReasearcher(p *Product, startup bool, projects int, opts ...ResearcherOptions) *Researcher {
	r := &Researcher{
		Product:  p,
		startup:  startup,
		projects: projects,
	}

	for _, option := range opts {
		option(r)
	}
	return r
}

func WithWantedSalary(wanted float64) ResearcherOptions {
	return func(r *Researcher) {
		r.Wanted = wanted
	}
}
