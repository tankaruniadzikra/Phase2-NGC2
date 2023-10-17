CREATE TABLE Heroes (
    ID INT AUTO_INCREMENT PRIMARY KEY,
    Name VARCHAR(255) NOT NULL,
    Universe VARCHAR(255) NOT NULL,
    Skill VARCHAR(255) NOT NULL,
    ImageURL VARCHAR(255) NOT NULL
);

CREATE TABLE Villains (
    ID INT AUTO_INCREMENT PRIMARY KEY,
    Name VARCHAR(255) NOT NULL,
    Universe VARCHAR(255) NOT NULL,
    ImageURL VARCHAR(255) NOT NULL
);

INSERT INTO Heroes (Name, Universe, Skill, ImageURL)
VALUES
    ('Iron Man', 'Marvel', 'Genius inventor and engineer', 'ironman.jpg'),
    ('Thor', 'Marvel', 'God of Thunder', 'thor.jpg'),
    ('Black Widow', 'Marvel', 'Master spy and hand-to-hand combatant', 'blackwidow.jpg');

INSERT INTO Villains (Name, Universe, ImageURL)
VALUES
    ('Loki', 'Marvel', 'loki.jpg'),
    ('Red Skull', 'Marvel', 'redskull.jpg'),
    ('Thanos', 'Marvel', 'thanos.jpg');

-- "root:@tcp(127.0.0.1:3306)/phase2-ngc2"