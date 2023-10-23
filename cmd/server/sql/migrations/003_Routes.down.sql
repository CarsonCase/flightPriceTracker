-- Drop the 'Route' column from the 'Flights' table.
ALTER TABLE Flights DROP COLUMN Route;

-- Recreate the 'departure' and 'arrival' columns in the 'Flights' table.
ALTER TABLE Flights ADD COLUMN departure VARCHAR(3) NOT NULL;
ALTER TABLE Flights ADD COLUMN arrival VARCHAR(3) NOT NULL;
