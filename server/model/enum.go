package model

import "fmt"

type Relationship struct {
	Relation                     string
	RelationMale, RelationFemale *Relationship `json:"-"`
	Sex
}
type Department int
type Sex int
type Property string

var (
	Father          Relationship = Relationship{Relation: "father", Sex: Male}
	Mother          Relationship = Relationship{Relation: "mother", Sex: Female}
	Son             Relationship = Relationship{Relation: "son", Sex: Male}
	Daughter        Relationship = Relationship{Relation: "daughter", Sex: Female}
	Brother         Relationship = Relationship{Relation: "brother", Sex: Male}
	Sister          Relationship = Relationship{Relation: "sister", Sex: Female}
	Husband         Relationship = Relationship{Relation: "husband", Sex: Male}
	Wife            Relationship = Relationship{Relation: "wife", Sex: Female}
	DivorcedWife    Relationship = Relationship{Relation: "divorced wife", Sex: Female}
	DivorcedHusband Relationship = Relationship{Relation: "divorced husband", Sex: Male}
	GGGrandfather   Relationship = Relationship{Relation: "Great-great grandfather", Sex: Male}
	GGGrandmother   Relationship = Relationship{Relation: "Great-great grandmother", Sex: Female}
	GGrandfather    Relationship = Relationship{Relation: "Great grandfather", Sex: Male}
	GGrandmother    Relationship = Relationship{Relation: "Great grandmother", Sex: Female}
	Grandfather     Relationship = Relationship{Relation: "Grandfather", Sex: Male}
	Grandmother     Relationship = Relationship{Relation: "Grandmother", Sex: Female}
	GrandSon        Relationship = Relationship{Relation: "Grandson", Sex: Male}
	Granddaughter   Relationship = Relationship{Relation: "Granddaughter", Sex: Female}
	GGrandSon       Relationship = Relationship{Relation: "Great grandson", Sex: Male}
	GGranddaughter  Relationship = Relationship{Relation: "Great granddaughter", Sex: Female}
	GGGrandSon      Relationship = Relationship{Relation: "Great-great grandson", Sex: Male}
	GGGranddaughter Relationship = Relationship{Relation: "Great-great granddaughter", Sex: Female}
	Uncle           Relationship = Relationship{Relation: "uncle", Sex: Male}
	Aunt            Relationship = Relationship{Relation: "aunt", Sex: Female}
	Cousin          Relationship = Relationship{Relation: "cousin", Sex: Female}
	BoyCousin       Relationship = Relationship{Relation: "boy-cousin", Sex: Male}
	Nephew          Relationship = Relationship{Relation: "nephew", Sex: Male}
	Niece           Relationship = Relationship{Relation: "niece", Sex: Female}
	FatherInLaw     Relationship = Relationship{Relation: "father-in-law", Sex: Male}
	MotherInLaw     Relationship = Relationship{Relation: "mother-in-law", Sex: Female}
	SonInLaw        Relationship = Relationship{Relation: "son-in-law", Sex: Male}
	DaughterInLaw   Relationship = Relationship{Relation: "daughter-in-law", Sex: Female}
	BrotherInLaw    Relationship = Relationship{Relation: "brother-in-law", Sex: Male}
	SisterInLaw     Relationship = Relationship{Relation: "sister-in-law", Sex: Female}
	StepFather      Relationship = Relationship{Relation: "stepfather", Sex: Male}
	StepMother      Relationship = Relationship{Relation: "stepmother", Sex: Female}
	StepSon         Relationship = Relationship{Relation: "stepson", Sex: Male}
	StepDaughter    Relationship = Relationship{Relation: "stepdaughter", Sex: Female}
	StepSister      Relationship = Relationship{Relation: "stepsister", Sex: Female}
	StepBrother     Relationship = Relationship{Relation: "stepbrother", Sex: Male}
	HalfBrother     Relationship = Relationship{Relation: "half-brother", Sex: Male}
	HalfSister      Relationship = Relationship{Relation: "half-sister", Sex: Female}
)

var RelationshipsArray []Relationship
var RelationshipsMap map[string]Relationship
var RelationshipsNames []string

const (
	HC  Department = 0 // home care
	CGT Department = 1 // centrum of grief therapy
)

const (
	Unknown Sex = 0
	Male    Sex = 1
	Female  Sex = 2
)

const (
	PROPERTY_CODE6         Property = "bucket_code6"
	PROPERTY_PASSWORD_HASH Property = "bucket_password_hash"
)

func GetOpoRelation(relationship string, sex Sex) Relationship {
	fmt.Println("personToDto", relationship, sex)

	if r, ok := RelationshipsMap[relationship]; ok {
		fmt.Println("FOUND", r)
		if sex == Male {
			return *r.RelationMale
		} else {
			return *r.RelationFemale
		}
	}
	return Relationship{}
}

func init() {
	//
	Father.RelationMale = &Son
	Father.RelationFemale = &Daughter
	Mother.RelationMale = &Son
	Mother.RelationFemale = &Daughter
	Son.RelationMale = &Father
	Son.RelationFemale = &Mother
	Daughter.RelationMale = &Father
	Daughter.RelationFemale = &Mother
	Brother.RelationMale = &Brother
	Brother.RelationFemale = &Sister
	Sister.RelationMale = &Brother
	Sister.RelationFemale = &Sister
	Husband.RelationMale = &Husband
	Husband.RelationFemale = &Wife
	Wife.RelationMale = &Husband
	Wife.RelationFemale = &Wife
	DivorcedWife.RelationMale = &DivorcedHusband
	DivorcedWife.RelationFemale = &DivorcedWife
	DivorcedHusband.RelationMale = &DivorcedHusband
	DivorcedHusband.RelationFemale = &DivorcedWife
	GGGrandfather.RelationMale = &GGrandSon
	GGGrandfather.RelationFemale = &GGranddaughter
	GGGrandmother.RelationMale = &GGrandSon
	GGGrandmother.RelationFemale = &GGranddaughter
	GGrandfather.RelationMale = &GGrandSon
	GGrandfather.RelationFemale = &GGranddaughter
	GGrandmother.RelationMale = &GGrandSon
	GGrandmother.RelationFemale = &GGranddaughter
	Grandfather.RelationMale = &GrandSon
	Grandfather.RelationFemale = &Granddaughter
	Grandmother.RelationMale = &GrandSon
	Grandmother.RelationFemale = &Granddaughter
	GrandSon.RelationMale = &Grandfather
	GrandSon.RelationFemale = &Grandmother
	Granddaughter.RelationMale = &Grandfather
	Granddaughter.RelationFemale = &Grandmother
	GGrandSon.RelationMale = &GGrandfather
	GGrandSon.RelationFemale = &GGrandmother
	GGranddaughter.RelationMale = &GGrandfather
	GGranddaughter.RelationFemale = &GGrandmother
	GGGrandSon.RelationMale = &GGrandfather
	GGGrandSon.RelationFemale = &GGrandmother
	GGGranddaughter.RelationMale = &GGGrandfather
	GGGranddaughter.RelationFemale = &GGGranddaughter
	Uncle.RelationMale = &Nephew
	Uncle.RelationFemale = &Niece
	Aunt.RelationMale = &Nephew
	Aunt.RelationFemale = &Niece
	Cousin.RelationMale = &BoyCousin
	Cousin.RelationFemale = &Cousin
	BoyCousin.RelationMale = &BoyCousin
	BoyCousin.RelationFemale = &Cousin
	Nephew.RelationMale = &Uncle
	Nephew.RelationFemale = &Aunt
	Niece.RelationMale = &Uncle
	Niece.RelationFemale = &Aunt
	FatherInLaw.RelationMale = &SonInLaw
	FatherInLaw.RelationFemale = &DaughterInLaw
	MotherInLaw.RelationMale = &SonInLaw
	MotherInLaw.RelationFemale = &DaughterInLaw
	SonInLaw.RelationMale = &FatherInLaw
	SonInLaw.RelationFemale = &MotherInLaw
	DaughterInLaw.RelationMale = &FatherInLaw
	DaughterInLaw.RelationFemale = &MotherInLaw
	BrotherInLaw.RelationMale = &BrotherInLaw
	BrotherInLaw.RelationFemale = &SisterInLaw
	SisterInLaw.RelationMale = &BrotherInLaw
	SisterInLaw.RelationFemale = &SisterInLaw
	StepFather.RelationMale = &StepSon
	StepFather.RelationFemale = &StepDaughter
	StepMother.RelationMale = &StepSon
	StepMother.RelationFemale = &StepDaughter
	StepSon.RelationMale = &StepFather
	StepSon.RelationFemale = &StepMother
	StepDaughter.RelationMale = &StepFather
	StepDaughter.RelationFemale = &StepDaughter
	StepSister.RelationMale = &StepBrother
	StepSister.RelationFemale = &StepSister
	StepBrother.RelationMale = &StepBrother
	StepBrother.RelationFemale = &StepSister
	HalfBrother.RelationMale = &HalfBrother
	HalfBrother.RelationFemale = &HalfSister
	HalfSister.RelationMale = &HalfBrother
	HalfSister.RelationFemale = &HalfSister

	RelationshipsArray = []Relationship{Father, Mother, Son, Daughter, Brother, Sister, Husband, Wife, DivorcedWife,
		DivorcedHusband, GGGrandfather, GGGrandmother, GGrandfather, GGrandmother, Grandfather, Grandmother, GrandSon,
		Granddaughter, GGrandSon, GGranddaughter, GGGrandSon, GGGranddaughter, Uncle, Aunt, Cousin, BoyCousin, Nephew, Niece,
		FatherInLaw, MotherInLaw, SonInLaw, DaughterInLaw, BrotherInLaw, SisterInLaw, StepFather, StepMother, StepSon,
		StepDaughter, StepSister, StepBrother, HalfBrother, HalfSister}
	RelationshipsMap = map[string]Relationship{}
	RelationshipsNames = []string{}
	for _, r := range RelationshipsArray {
		RelationshipsMap[r.Relation] = r
		RelationshipsNames = append(RelationshipsNames, r.Relation)
	}
}
