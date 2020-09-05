package somepack

type From struct {
	City *City `json:"city"`
}

type Budget struct {
	CostCenter    string    `json:"cost_center"`
	CostItem      string    `json:"cost_item"`
	InvestProject string    `json:"invest_project"`
	InvestTask    string    `json:"invest_task"`
	Expenses      *Expenses `json:"expenses"`
	BalanceUnit   string    `json:"balance_unit"`
}

type Participant struct {
	Name       string  `json:"name"`
	Id         string  `json:"id"`
	EmployeeId float64 `json:"employee_id"`
}

type Movements struct {
	Id      string   `json:"id"`
	From    *From    `json:"from"`
	To      *To      `json:"to"`
	Payment *Payment `json:"payment"`
}

type To struct {
	City *City `json:"city"`
}

type Status struct {
	Title string `json:"title"`
	Id    string `json:"id"`
}

type PaymentCompany struct {
	RoomId  string `json:"room_id"`
	Booking string `json:"booking"`
	Id      string `json:"id"`
}

type Example struct {
	Participant    *Participant `json:"participant"`
	Movements      *Movements   `json:"movements"`
	Status         *Status      `json:"status"`
	Prepayment     bool         `json:"prepayment"`
	WorkingWeekend []string     `json:"working_weekend"`
	Id             string       `json:"id"`
	Goal           string       `json:"goal"`
	DateEnd        string       `json:"date_end"`
	Residence      *Residence   `json:"residence"`
	DateStart      string       `json:"date_start"`
	Agreement      bool         `json:"agreement"`
	Budget         *Budget      `json:"budget"`
}

type Info struct {
	Id           string `json:"id"`
	Carrier      string `json:"carrier"`
	FlightNumber string `json:"flight_number"`
	Date         string `json:"date"`
	Comment      string `json:"comment"`
}

type Payment struct {
	Type          string  `json:"type"`
	TransportType string  `json:"transport_type"`
	Ticket        *Ticket `json:"ticket"`
}

type Expenses struct {
	Id    string  `json:"id"`
	Type  *Type   `json:"type"`
	Value float64 `json:"value"`
}

type City struct {
	Id          string `json:"id"`
	EnName      string `json:"en_name"`
	RusName     string `json:"rus_name"`
	IataCode    string `json:"iata_code"`
	Travelclick string `json:"travelclick"`
}

type Ticket struct {
	Id   string `json:"id"`
	Info *Info  `json:"info"`
}

type Residence struct {
	PaymentUser    *PaymentUser    `json:"payment_user"`
	PaymentCompany *PaymentCompany `json:"payment_company"`
}

type PaymentUser struct {
	DateOut   string `json:"date_out"`
	HotelName string `json:"hotel_name"`
	Comment   string `json:"comment"`
	DateIn    string `json:"date_in"`
}

type Type struct {
	Id          string `json:"id"`
	Description string `json:"description"`
}
