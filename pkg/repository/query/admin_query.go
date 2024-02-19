package query

const CreateFlightQuery = `
	INSERT INTO Flights 
	(StartTime, EndTime, DeparturePoint, Destination, EconomyPrice, BusinessPrice, DeluxePrice, 
	TotalEconomyTickets, TotalBusinessTickets, TotalDeluxeTickets, AvailableEconomyTickets, 
	AvailableBusinessTickets, AvailableDeluxeTickets) 
	VALUES 
	($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
	RETURNING ID
`
const GetOneFlightQuery = `
	SELECT * FROM Flights WHERE ID = $1;
`

const GetAllFlightQuery = `
	SELECT * FROM Flights;
`

const DeleteFlightQuery = `
	DELETE FROM Flights WHERE id = $1
`

const UpdateFlightQuery = `
	UPDATE Flights SET
    StartTime = $2,
    EndTime = $3,
    DeparturePoint = $4,
    Destination = $5,
    EconomyPrice = $6,
    BusinessPrice = $7,
    DeluxePrice = $8,
    TotalEconomyTickets = $9,
    TotalBusinessTickets = $10,
    TotalDeluxeTickets = $11,
    AvailableEconomyTickets = $12,
    AvailableBusinessTickets = $13,
    AvailableDeluxeTickets = $14
	WHERE
    ID = $1
`