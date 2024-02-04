CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Создание таблицы Users
CREATE TABLE Users (
    ID UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    Phone VARCHAR(20),
    Email VARCHAR(100) NOT NULL,
    Pass VARCHAR(100) NOT NULL,
    CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UpdatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Создание таблицы Flights
-- Создание таблицы Flights с UUIDv4 ID
CREATE TABLE Flights (
    ID UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    StartTime TIMESTAMP NOT NULL,
    EndTime TIMESTAMP NOT NULL,
    DeparturePoint VARCHAR(255) NOT NULL,
    Destination VARCHAR(255) NOT NULL,
    EconomyPrice DECIMAL(10, 2) NOT NULL,
    BusinessPrice DECIMAL(10, 2) NOT NULL,
    DeluxePrice DECIMAL(10, 2) NOT NULL,
    TotalEconomyTickets INT NOT NULL,
    TotalBusinessTickets INT NOT NULL,
    TotalDeluxeTickets INT NOT NULL,
    AvailableEconomyTickets INT NOT NULL,
    AvailableBusinessTickets INT NOT NULL,
    AvailableDeluxeTickets INT NOT NULL,
    CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


-- Создание типа ENUM для ранга билета
CREATE TYPE TicketRank AS ENUM ('Economy', 'Business', 'Deluxe');

-- Создание таблицы Tickets с использованием ENUM
CREATE TABLE Tickets (
    ID UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    FlightID UUID REFERENCES Flights(ID),
    UserID UUID REFERENCES Users(ID),
    Rank TicketRank NOT NULL,
    Price DECIMAL(10, 2) NOT NULL,
    CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
