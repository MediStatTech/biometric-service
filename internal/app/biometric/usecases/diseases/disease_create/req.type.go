package disease_create

type Request struct {
	Name string
	Code string
}

type Response struct {
	DiseaseID string
}
