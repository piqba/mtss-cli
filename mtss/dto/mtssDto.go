package mtss

type MtssResponse struct {
	ID            int     `json:"id"`
	Company       string  `json:"organismo"`
	Position      string  `json:"cargo"`
	Taken         int     `json:"cantidad"`
	Entity        string  `json:"entidad"`
	Province      string  `json:"provincia"`
	Municipality  string  `json:"municipio"`
	Availability  int     `json:"ocupadas"`
	Activity      string  `json:"actividad"`
	Pay           float32 `json:"salario"`
	SchoolLevel   string  `json:"nivelEscolar"`
	Details       string  `json:"observaciones"`
	EntityMail    string  `json:"correo_entidad"`
	EntityAddress string  `json:"direccion_entidad"`
	EntityPhone   string  `json:"telefono_entidad"`
	RegisterDate  string  `json:"fecha_registro"`
	UniqueStamp   string  `json:"unique_stamp"`
	Enabled       bool    `json:"habilitada"`
	Source        string  `json:"source"`
	TypeWork      string  `json:"type_work"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}
