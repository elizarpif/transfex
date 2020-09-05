package somepack

type Participant struct {
	Id         string  `json:"id"`
	EmployeeId float64 `json:"employee_id"`
	Name       string  `json:"name"`
}

type Status struct {
	Id    string `json:"id"`
	Title string `json:"title"`
}

type City struct {
	Id      string `json:"id"`
	EnName  string `json:"en_name"`
	RusName string `json:"rus_name"`
}

type Payment struct {
	Type          string  `json:"type"`
	TransportType string  `json:"transport_type"`
	Ticket        *Ticket `json:"ticket"`
}

type PaymentUser struct {
	Comment   string `json:"comment"`
	DateIn    string `json:"date_in"`
	DateOut   string `json:"date_out"`
	HotelName string `json:"hotel_name"`
}

type Budget struct {
	InvestTask    string    `json:"invest_task"`
	Expenses      *Expenses `json:"expenses"`
	BalanceUnit   string    `json:"balance_unit"`
	CostCenter    string    `json:"cost_center"`
	CostItem      string    `json:"cost_item"`
	InvestProject string    `json:"invest_project"`
}

type From struct {
	City *City `json:"city"`
}

type Info struct {
	Id           string `json:"id"`
	Carrier      string `json:"carrier"`
	FlightNumber string `json:"flight_number"`
	Date         string `json:"date"`
	Comment      string `json:"comment"`
}

type PaymentCompany struct {
	RoomId  string `json:"room_id"`
	Booking string `json:"booking"`
	Id      string `json:"id"`
}

type To struct {
	City *City `json:"city"`
}

type Ticket struct {
	Id   string `json:"id"`
	Info *Info  `json:"info"`
}

type Type struct {
	Id          string `json:"id"`
	Description string `json:"description"`
}

type Expenses struct {
	Id    string  `json:"id"`
	Type  *Type   `json:"type"`
	Value float64 `json:"value"`
}

type Movements struct {
	Id      string   `json:"id"`
	From    *From    `json:"from"`
	To      *To      `json:"to"`
	Payment *Payment `json:"payment"`
}

type Example struct {
	Residence      *Residence   `json:"residence"`
	Id             string       `json:"id"`
	Goal           string       `json:"goal"`
	DateStart      string       `json:"date_start"`
	Participant    *Participant `json:"participant"`
	Agreement      bool         `json:"agreement"`
	Status         *Status      `json:"status"`
	Budget         *Budget      `json:"budget"`
	WorkingWeekend []string     `json:"working_weekend"`
	DateEnd        string       `json:"date_end"`
	Movements      *Movements   `json:"movements"`
	Prepayment     bool         `json:"prepayment"`
}

type Residence struct {
	PaymentUser    *PaymentUser    `json:"payment_user"`
	PaymentCompany *PaymentCompany `json:"payment_company"`
}
