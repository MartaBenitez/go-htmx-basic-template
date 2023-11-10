package dto

type SectionsToEdit struct {
	NAPISALSECTOEDIT SectionsEdit `json:"NAPISALSECTOEDIT"`
}

type SectionsEdit struct {
	Sections_to_edit Sections `json:"sections_to_edit"`
	Id_query         string   `json:"id_query"`
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
