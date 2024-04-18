-- Create Users table
CREATE TABLE Users (
UserID SERIAL PRIMARY KEY,
Email VARCHAR(255) UNIQUE NOT NULL,
Password VARCHAR(255) NOT NULL,
DepositAmount DECIMAL DEFAULT 0
);

-- Create EquipmentTypes table
CREATE TABLE EquipmentTypes (
TypeID SERIAL PRIMARY KEY,
Name VARCHAR(255) NOT NULL,
Availability BOOLEAN NOT NULL,
RentalCosts DECIMAL NOT NULL,
Category VARCHAR(255) NOT NULL
);

-- Create RentalHistory table
CREATE TABLE RentalHistory (
RentalID SERIAL PRIMARY KEY,
UserID INT REFERENCES Users(UserID),
EquipmentID INT REFERENCES EquipmentTypes(TypeID),
RentalDate DATE NOT NULL,
ReturnDate DATE,
CONSTRAINT fk_user_rental FOREIGN KEY (UserID) REFERENCES Users(UserID),
CONSTRAINT fk_equipment_rental FOREIGN KEY (EquipmentID) REFERENCES EquipmentTypes(TypeID)
);

-- Create Transactions table
CREATE TABLE Transactions (
TransactionID SERIAL PRIMARY KEY,
RentalID INT REFERENCES RentalHistory(RentalID),
TransactionDate DATE NOT NULL,
Amount DECIMAL(10, 2) NOT NULL,
PaymentMethod VARCHAR(100),
CONSTRAINT fk_rental_transaction FOREIGN KEY (RentalID) REFERENCES RentalHistory(RentalID)
);

-- Create EquipmentRentals table
CREATE TABLE EquipmentRentals (
RentalID INT REFERENCES RentalHistory(RentalID),
TransactionID INT REFERENCES Transactions(TransactionID),
RentalDate DATE NOT NULL,
ReturnDate DATE,
PRIMARY KEY (RentalID, TransactionID),
CONSTRAINT fk_rental_current FOREIGN KEY (RentalID) REFERENCES RentalHistory(RentalID),
CONSTRAINT fk_transaction_current FOREIGN KEY (TransactionID) REFERENCES Transactions(TransactionID)
);

-- Create Payments table
CREATE TABLE Payments (
PaymentID SERIAL PRIMARY KEY,
UserID INT REFERENCES Users(UserID),
TransactionID INT REFERENCES Transactions(TransactionID),
PaymentDate DATE NOT NULL,
Amount DECIMAL(10, 2) NOT NULL,
CONSTRAINT fk_user_payment FOREIGN KEY (UserID) REFERENCES Users(UserID),
CONSTRAINT fk_transaction_payment FOREIGN KEY (TransactionID) REFERENCES Transactions(TransactionID)
);
