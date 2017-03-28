INSERT INTO Users (Username, Password, IsEmployee)
	VALUES ("admin", "admin", true);

INSERT INTO Projects (Name, Code, StartDate, User_ID)
	VALUES ("Projekt 1", "Proj1", CURDATE(), 1);

INSERT INTO Tasks (Name, Code, StartDate, Project_ID)
	VALUES ("Úkol 1", "Ukol1", CURDATE(), 1);