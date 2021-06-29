package mtss

import (
	"encoding/json"

	dto "github.com/piqba/mtss-cli/mtss/dto"
)

// MTSS models mtss
type MTSS struct {
	ID            int     `bson:"_id"`
	Company       string  `bson:"organismo"`
	Position      string  `bson:"cargo"`
	Taken         int     `bson:"cantidad"`
	Entity        string  `bson:"entidad"`
	Province      string  `bson:"provincia"`
	Municipality  string  `bson:"municipio"`
	Availability  int     `bson:"ocupadas"`
	Activity      string  `bson:"actividad"`
	Pay           float32 `bson:"salario"`
	SchoolLevel   string  `bson:"nivelEscolar"`
	Details       string  `bson:"observaciones"`
	EntityMail    string  `bson:"correo_entidad"`
	EntityAddress string  `bson:"direccion_entidad"`
	EntityPhone   string  `bson:"telefono_entidad"`
	RegisterDate  string  `bson:"fecha_registro"`
	UniqueStamp   string  `bson:"unique_stamp"`
	Enabled       bool    `bson:"habilitada"`
	Source        string  `bson:"source"`
	TypeWork      string  `bson:"type_work"`
}

type MtssRepository interface {
	FetchAllFromAPI(string) ([]MTSS, error)
	CreateOne(string, MTSS) error
}

func (mt *MTSS) MarshalBinary() ([]byte, error) {
	return json.Marshal(mt)
}

func (mt *MTSS) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, &mt); err != nil {
		return err
	}

	return nil
}

func (mt *MTSS) ToMAP() (toHashMap map[string]interface{}, err error) {

	fromStruct, _ := json.Marshal(mt)
	if err := json.Unmarshal(fromStruct, &toHashMap); err != nil {
		return toHashMap, err
	}
	// 	for field, val := range inInterface {
	// 		fmt.Println("KV Pair: ", field, val)
	// }
	return toHashMap, nil
}

func (mt *MTSS) ToDto() dto.MtssResponse {
	return dto.MtssResponse{
		ID:            mt.ID,
		Company:       mt.Company,
		Position:      mt.Position,
		Taken:         mt.Taken,
		Entity:        mt.Entity,
		Province:      mt.Province,
		Municipality:  mt.Municipality,
		Availability:  mt.Availability,
		Activity:      mt.Activity,
		Pay:           mt.Pay,
		SchoolLevel:   mt.SchoolLevel,
		Details:       mt.Details,
		EntityMail:    mt.EntityMail,
		EntityAddress: mt.EntityAddress,
		EntityPhone:   mt.EntityPhone,
		RegisterDate:  mt.RegisterDate,
		UniqueStamp:   mt.UniqueStamp,
		Enabled:       mt.Enabled,
		Source:        mt.Source,
		TypeWork:      mt.TypeWork,
	}
}
