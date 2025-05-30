package entity

type Bus struct {
	BusID          string   `json:"bus_id"`          // unique id
	OperatorName   string   `json:"operator_name"`   // ex: KPN
	BusNumber      string   `json:"bus_number"`      // TN-01-1234
	BusType        string   `json:"bus_type"`        // ex: AC Sleeper
	TotalSeats     int      `json:"total_seats"`     // 40
	AvailableSeats int      `json:"available_seats"` // 30
	Facilities     []string `json:"facilities"`      // ["AC", "Blanket", "Charging"]
	From           string   `json:"from"`            // Chennai
	To             string   `json:"to"`              // Coimbatore
	DepartureTime  string   `json:"departure_time"`  // "2025-06-01T22:00:00Z"
	ArrivalTime    string   `json:"arrival_time"`    // "2025-06-02T06:00:00Z"
	Fare           float64  `json:"fare"`            // 750.00
	Status         string   `json:"status"`          // ACTIVE / INACTIVE / MAINTENANCE
}
