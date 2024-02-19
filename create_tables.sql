-- Drop the database if it exists
DROP DATABASE IF EXISTS go;

-- Create the database
CREATE DATABASE go;

-- Drop the table if it exists
DROP TABLE IF EXISTS Userdata;

-- Create the table
CREATE TABLE Userdata(
                         UserID SERIAL PRIMARY KEY,
                         Name VARCHAR(100),
                         Surname VARCHAR(100),
                         Username VARCHAR(100) UNIQUE,
                         Description VARCHAR(100)
);

-- Insert data into the table
INSERT INTO Userdata (Name, Surname, Username, Description)
VALUES
    ('John', 'Doe', 'john_doe', 'Lorem ipsum dolor sit amet'),
    ('Jane', 'Smith', 'jane_smith', 'Consectetur adipiscing elit'),
    ('Michael', 'Johnson', 'michael_johnson', 'Sed do eiusmod tempor incididunt'),
    ('Emily', 'Williams', 'emily_williams', 'Ut labore et dolore magna aliqua'),
    ('David', 'Brown', 'david_brown', 'Ut enim ad minim veniam'),
    ('Sarah', 'Jones', 'sarah_jones', 'Quis nostrud exercitation ullamco'),
    ('James', 'Miller', 'james_miller', 'Laboris nisi ut aliquip ex ea commodo consequat'),
    ('Jessica', 'Davis', 'jessica_davis', 'Duis aute irure dolor in reprehenderit'),
    ('Matthew', 'Garcia', 'matthew_garcia', 'Voluptate velit esse cillum dolore eu fugiat'),
    ('Laura', 'Martinez', 'laura_martinez', 'Nulla pariatur. Excepteur sint occaecat'),
    ('Daniel', 'Hernandez', 'daniel_hernandez', 'Cupidatat non proident, sunt in culpa qui officia'),
    ('Ashley', 'Lopez', 'ashley_lopez', 'Deserunt mollit anim id est laborum'),
    ('Christopher', 'Gonzalez', 'christopher_gonzalez', 'Lorem ipsum dolor sit amet'),
    ('Amanda', 'Rodriguez', 'amanda_rodriguez', 'Consectetur adipiscing elit'),
    ('Brian', 'Perez', 'brian_perez', 'Sed do eiusmod tempor incididunt'),
    ('Jennifer', 'Wilson', 'jennifer_wilson', 'Ut labore et dolore magna aliqua'),
    ('Joshua', 'Torres', 'joshua_torres', 'Ut enim ad minim veniam'),
    ('Megan', 'Gutierrez', 'megan_gutierrez', 'Quis nostrud exercitation ullamco'),
    ('Ryan', 'Flores', 'ryan_flores', 'Laboris nisi ut aliquip ex ea commodo consequat'),
    ('Nicole', 'Reyes', 'nicole_reyes', 'Duis aute irure dolor in reprehenderit');
