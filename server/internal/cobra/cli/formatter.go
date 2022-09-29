package cli

import (
	"encoding/xml"
	"fmt"

	"server/internal/db"

	"github.com/goccy/go-json"
)

type displayer interface {
	display(tasks db.Tasks) ([]byte, error)
}

type defaultDisplayer struct{}

// ---

func getDefaultDisplayer() defaultDisplayer {
	return defaultDisplayer{}
}

func (d defaultDisplayer) display(tasks db.Tasks) ([]byte, error) {
	return []byte(fmt.Sprint(tasks)), nil
}

// ---

type jsonDisplayer struct{}

func getJSONDisplayer() jsonDisplayer {
	return jsonDisplayer{}
}

func (d jsonDisplayer) display(tasks db.Tasks) ([]byte, error) {
	return json.MarshalIndent(tasks, "", "  ")
}

// ---

type xmlDisplayer struct{}

func getXMLDisplayer() xmlDisplayer {
	return xmlDisplayer{}
}

func (d xmlDisplayer) display(tasks db.Tasks) ([]byte, error) {
	return xml.MarshalIndent(tasks, "", "  ")
}
