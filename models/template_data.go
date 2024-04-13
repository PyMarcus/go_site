package models

// TemplateData holds data sent from handlers template
type TemplateData struct {
	StringMap    map[string]string
	IntegerMap   map[string]int
	FloatMap     map[string]float64
	Data         map[string]interface{}
	CSRFToken    string
	FlashMessage string
	Warning      string
	Error        string
}
