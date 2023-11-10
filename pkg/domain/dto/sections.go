package dto

type SectionsToEdit struct {
	Header interface{} `json:"header"`
	Body   Codigo      `json:"body"`
}

type Codigo struct {
	MAPISECTOEDIT Sections `json:"MAPISECTOEDIT"`
}
type Sections struct {
	Sections []Section `json:"sections"`
}

type Section struct {
	Section_code string `json:"section_code"`
	Modification string `json:"modification"`
}

func (s *Sections) ToStringGuay() string {
	var sb = "["
	for i := 0; i < len(s.Sections); i++ {
		sb = sb + "{\"section_code\":\"" + s.Sections[i].Section_code + "\",\"modification\":\"" + s.Sections[i].Modification + "\"}"
		if i < (len(s.Sections) - 1) {
			sb = sb + ","
		}
	}
	sb = sb + "]"
	return sb
}
