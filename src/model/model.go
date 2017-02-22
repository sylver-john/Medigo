package model

type Drug struct {
	Cis                        string
	Nom                        string
	FormePharmaceutique        string
	VoiesAdmnistration         []string
	StatusAutorisation         string
	TypeAutorisation           string
	EtatCommercialisation      string
	DateMiseSurLeMarche        string
	StatutBdm                  string
	NumeroAutorisationEuropeen string
	Titulaire                  []string
	SurveillanceRenforcee      bool
	Presentation               []Presentation
	Composition                []Composition
	Generique                  []Generique
	ConditionsPrescription     []ConditionsPrescription
	Score                      string
}

type Presentation struct {
	Cis                     string
	Cip                     float32
	Libelle                 string
	StatutAdministratif     string
	EtatCommercialisation   string
	DateDeclaration         string
	CIP13                   int
	AgrementCollectivite    bool
	TauxRemboursement       []string
	Prix                    float32
	IndicationRemboursement string
}

type Composition struct {
	Cis                                          string
	ElementPharmaceutique                        string
	CodeSubstance                                float32
	DenominationSubstance                        string
	EtatCommercialisation                        string
	DosageSubstance                              string
	ReferenceDosage                              string
	LienSubtanceActivesEtFractionsTherapeutiques string
}

type Generique struct {
	IdentifiantGenerique string
	LibelleGenerique     string
	Cis                  string
	Type                 string
}

type ConditionsPrescription struct {
	Cis string
}

type Titulaire struct {
	Id string
	Titulaire []string
}

type Drugs [][]Drug
