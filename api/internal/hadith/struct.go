package hadith

type Hadith struct {
	Number int    `json:"number"`
	Arab   string `json:"arab"`
	ID     string `json:"id"`
}

type Translate struct {
	ID string `json:"id"`
}

type Data struct {
	Arabic    string    `json:"arabic"`
	Number    int       `json:"number"`
	Translate Translate `json:"translate"`
}

type ResponseHadith200Ok struct {
	StatusCode  int    `json:"status_code"`
	Author      string `json:"author"`
	TotalHadith int    `json:"total_hadith"`
	Page        int    `json:"page"`
	PerPage     int    `json:"per_page"`
	Data        []Data `json:"data"`
}
