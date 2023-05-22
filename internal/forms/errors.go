package forms

type errors map[string][]string

// add the error for given form field
func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)
}

// Get retrives the first error messgae
func (e errors) Get(field string) string {
	es := e[field]
	if len(es) == 0 {
		return ""
	}
	return es[0]
}
