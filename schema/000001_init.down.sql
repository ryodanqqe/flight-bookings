BEGIN;

DROP TABLE IF EXISTS Tickets;

DROP TYPE IF EXISTS TicketRank;

DROP TABLE IF EXISTS Flights;

DROP TABLE IF EXISTS Users;

DROP EXTENSION IF EXISTS "uuid-ossp";

COMMIT;