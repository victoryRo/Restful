-- Database postgres

CREATE TABLE clients (
    User_ID        BIGSERIAL PRIMARY KEY,
    User_Name      VARCHAR(55) NOT NULL,
    Name           VARCHAR(55) NOT NULL,
    Created_At     TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    Is_Enabled     BOOLEAN DEFAULT TRUE NOT NULL
);

INSERT INTO clients (
    User_Name, Name
) VALUES 
    ( 'Victoria', 'Maria Victoria' ),
    ( 'Angela', 'Angela perea' ),
    ( 'Mariana', 'Mariana soza' ),
    ( 'Lucia', 'Lucia Marino' ),
    ( 'Patricia', 'Patricia goan' ),
    ( 'Susana', 'Susana susa' ),
    ( 'Veronica', 'Veronica vera' ),
    ( 'Esmeranda', 'Esmeranda marina' );

SELECT * FROM clients;
