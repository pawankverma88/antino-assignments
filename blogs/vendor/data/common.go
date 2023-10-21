package data

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/kennygrant/sanitize"
)

// CleanText returns a sanitized string (strips html tags) with spaces trimmed.
// If lowerCase is true, the function will also convert the entire string into lowercase.
// If trimDoubleSpaces is true, all double spaces will be trimmed down to one single space.
func CleanText(text string, lowerCase bool, trimDoubleSpaces bool) string {
	sanitizedText := sanitize.HTML(text)
	sanitizedText = strings.TrimSpace(sanitizedText)
	sanitizedText = strings.Replace(sanitizedText, "\t", " ", -1)
	if lowerCase {
		sanitizedText = strings.ToLower(sanitizedText)
	}
	if trimDoubleSpaces {
		for {
			if strings.Contains(sanitizedText, "  ") {
				sanitizedText = strings.Replace(sanitizedText, "  ", " ", -1)
			} else {
				break
			}
		}
	}

	return sanitizedText
}

// JSONMessageContent :
type JSONMessageContent struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

// JSONWrappedContent :
type JSONWrappedContent struct {
	StatusCode int         `json:"statusCode"`
	Content    interface{} `json:"content"`
}

// JSONMessage returns a preformatted ISG JSON with the response code and message.
func JSONMessage(code int, msg string) []byte {
	jsonString := JSONMessageContent{
		StatusCode: code,
		Message:    msg,
	}

	result, err := json.MarshalIndent(jsonString, "", "    ")
	if err != nil {
		fmt.Println(err)
	}

	return result
}

// JSONMessageWrappedObj returns an encoded JSON of the object provided.
func JSONMessageWrappedObj(code int, obj interface{}) []byte {
	jsonString := JSONWrappedContent{
		StatusCode: code,
		Content:    obj,
	}

	result, err := json.MarshalIndent(jsonString, "", "    ")
	if err != nil {
		fmt.Println(err)
	}
	return result
}

// RespondJSON :
func RespondJSON(w http.ResponseWriter, status int, payload string) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(JSONMessage(status, payload))

}

// RespondJSONObject :
func RespondJSONObject(w http.ResponseWriter, status int, obj interface{}) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(JSONMessageWrappedObj(status, obj))

}
