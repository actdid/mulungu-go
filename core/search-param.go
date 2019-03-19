package core

//SearchParam search param type
type SearchParam map[string]string

//NewSearchParam returns new SearchParam type
func NewSearchParam() *SearchParam {
	return &SearchParam{}
}

//Add to search param
func (sp *SearchParam) Add(key, value string) *SearchParam {
	if value != "" && key != "" {
		(*sp)[key] = value
	}
	return sp
}

//AsMap returns SearchParam as map[string]string
func (sp *SearchParam) AsMap() map[string]string {
	mapped := make(map[string]string)
	for k, v := range *sp {
		mapped[k] = v
	}
	return mapped
}
