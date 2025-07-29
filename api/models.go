package api

type APIResponse struct {
    Response    string `json:"response"`
    ResultsFor  string `json:"results-for"`
}

type APIResponseError struct {
    Response    string `json:"response"`
    Error  string `json:"error"`
}