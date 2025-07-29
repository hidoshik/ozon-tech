package integration

type PowerStats struct {
    Intelligence string `json:"intelligence"`
    Strength     string `json:"strength"`
    Speed        string `json:"speed"`
    Durability   string `json:"durability"`
    Power        string `json:"power"`
    Combat       string `json:"combat"`
}

type Biography struct {
    FullName     string   `json:"full-name"`
    AlterEgos    string   `json:"alter-egos"`
    Aliases      []string `json:"aliases"`
    PlaceOfBirth string   `json:"place-of-birth"`
    FirstAppearance string `json:"first-appearance"`
    Publisher    string   `json:"publisher"`
    Alignment    string   `json:"alignment"`
}

type Appearance struct {
    Gender    string   `json:"gender"`
    Race      string   `json:"race"`
    Height    []string `json:"height"`
    Weight    []string `json:"weight"`
    EyeColor  string   `json:"eye-color"`
    HairColor string   `json:"hair-color"`
}

type Work struct {
    Occupation string `json:"occupation"`
    Base       string `json:"base"`
}

type Connections struct {
    GroupAffiliation string `json:"group-affiliation"`
    Relatives        string `json:"relatives"`
}

type Image struct {
    URL string `json:"url"`
}

type Hero struct {
    ID          string      `json:"id"`
    Name        string      `json:"name"`
    PowerStats  PowerStats  `json:"powerstats"`
    Biography   Biography   `json:"biography"`
    Appearance  Appearance  `json:"appearance"`
    Work        Work        `json:"work"`
    Connections Connections `json:"connections"`
    Image       Image       `json:"image"`
}

type APIResponse struct {
    Response    string `json:"response"`
    ResultsFor  string `json:"results-for"`
	Results		[]Hero `json:"results"`
}
