package pwtcp

type serpStackStruct struct {
	Request struct {
		Success            bool    `json:"success"`
		ProcessedTimestamp int     `json:"processed_timestamp"`
		SearchURL          string  `json:"search_url"`
		TotalTimeTaken     float64 `json:"total_time_taken"`
	} `json:"request"`

	LocalResults []struct {
		Position    int    `json:"position"`
		Title       string `json:"title"`
		Coordinates struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"coordinates"`
		ImageURL   interface{} `json:"image_url"`
		Address    string      `json:"address"`
		Extensions interface{} `json:"extensions"`
		Rating     interface{} `json:"rating"`
		Reviews    interface{} `json:"reviews"`
		Type       string      `json:"type"`
		Price      interface{} `json:"price"`
		URL        string      `json:"url"`
	} `json:"local_results"`

	OrganicResults []struct {
		Position     int    `json:"position"`
		Title        string `json:"title"`
		URL          string `json:"url"`
		DisplayedURL string `json:"displayed_url"`
		Snippet      string `json:"snippet"`
		Sitelinks    struct {
			Inline []struct {
				Title string `json:"title"`
				URL   string `json:"url"`
			} `json:"inline"`
		} `json:"sitelinks,omitempty"`
	} `json:"organic_results"`
}
