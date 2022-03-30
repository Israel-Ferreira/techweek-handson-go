package data

var VALID_EVENTS = []string{"DELETE_PRODUCT", "NEW_PRODUCT", "UPDATE_PRODUCT"}

type EventProductMsg struct {
	Sku   string `json:"sku"`
	Title string `json:"title"`
	Event string `json:"event"`
}

func (p *EventProductMsg) IsValidEvent() error {
	var eventIsFound bool

	for _, event := range VALID_EVENTS {
		if p.Event == event {
			eventIsFound = true
		}
	}


	if !eventIsFound {
		return
	}
}
