package model

type Relationship string

const (
	Father          Relationship = "father"
	Mother          Relationship = "mother"
	Son             Relationship = "son"
	Daughter        Relationship = "daughter"
	Brother         Relationship = "brother"
	Sister          Relationship = "sister"
	Husband         Relationship = "husband"
	Wife            Relationship = "wife"
	DivorcedWife    Relationship = "divorced wife"
	DivorcedHusband Relationship = "divorced husband"
	GGGrandfather   Relationship = "Great-great grandfather"
	GGGrandmother   Relationship = "Great-great grandmother"
	GGrandfather    Relationship = "Great grandfather"
	GGrandmother    Relationship = "Great grandmother"
	Grandfather     Relationship = "Grandfather"
	Grandmother     Relationship = "Grandmother"
	GrandSon        Relationship = "Grandson"
	Granddaughter   Relationship = "Granddaughter"
	GGrandSon       Relationship = "Great grandson"
	GGranddaughter  Relationship = "Great granddaughter"
	GGGrandSon      Relationship = "Great-great grandson"
	GGGranddaughter Relationship = "Great-great granddaughter"
	Uncle           Relationship = "uncle"
	Aunt            Relationship = "aunt"
	Cousin          Relationship = "cousin"
	BoyCousin       Relationship = "boy-cousin"
	Nephew          Relationship = "nephew"
	Niece           Relationship = "niece"
	FatherInLaw     Relationship = "father-in-law"
	MotherInLaw     Relationship = "mother-in-law"
	SonInLaw        Relationship = "son-in-law"
	DaughterInLaw   Relationship = "daughter-in-law"
	BrotherInLaw    Relationship = "brother-in-law"
	SisterInLaw     Relationship = "sister-in-law"
	StepFather      Relationship = "stepfather"
	StepMother      Relationship = "stepmother"
	StepSon         Relationship = "stepson"
	StepDaughter    Relationship = "stepdaughter"
	StepSister      Relationship = "stepsister"
	StepBrother     Relationship = "stepbrother"
	HalfBrother     Relationship = "half-brother"
	HalfSister      Relationship = "half-sister"
)

type Department int

const (
	HC  Department = 0 // home care
	CGT Department = 1 // centrum of grief therapy
)

type Sex int

const (
	Male   Sex = 0
	Female Sex = 1
)
