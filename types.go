package enova

type UpdateParamsData struct {
	Rows Rows `xml:"Rows>Row"`
}

type Rows []Row

type Row struct {
	XML interface{} `xml:"XML"`
}
