CREATE TABLE Routes(
    ID UUID PRIMARY KEY,
    departure VARCHAR(3) NOT NULL,
    arrival VARCHAR(3) NOT NULL,
    CONSTRAINT unique_route_pair UNIQUE (departure, arrival)
);

ALTER TABLE Flights ADD COLUMN Route UUID NOT NULL REFERENCES Routes(ID) ON DELETE CASCADE;
ALTER TABLE Flights DROP COLUMN departure;
ALTER TABLE Flights DROP COLUMN arrival;