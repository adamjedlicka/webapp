INSERT INTO Permissions (ID, Name, IsAdmin)
	VALUES (UUID(), "admin", true);

INSERT INTO Permissions (ID, Name, IsAdmin)
	VALUES (UUID(), "Project Manager", false);

INSERT INTO Users (ID, UserName, Password, Permission_ID)
	VALUES (UUID(), "admin", "admin", (SELECT ID FROM Permissions WHERE Name = "admin"));

INSERT INTO Users (ID, UserName, Password, Permission_ID)
	VALUES (UUID(), "jedliad1", "asdfghjkl", (SELECT ID FROM Permissions WHERE Name = "Project Manager"));

INSERT INTO Firms (ID, Name)
	VALUES (UUID(), "SoftCorp s.r.o.");

INSERT INTO Firms (ID, Name)
	VALUES (UUID(), "Google a.s.");

INSERT INTO Projects (ID, Name, Code, StartDate, User_ID, Firm_ID)
	VALUES (UUID(), "Vyvoj systemu ISSZP", "ISSZP", CURDATE(),
		(SELECT ID FROM Users WHERE UserName = "jedliad1"),
		(SELECT ID FROM Firms WHERE Name = "SoftCorp s.r.o."));

INSERT INTO Tasks (ID, Name, StartDate, PlanEndDate, User_ID_Maintainer, User_ID_Worker, Project_ID)
	VALUES (UUID(), "Naplnit aplikaci zakladnimy daty", CURDATE(), "2017-04-20",
		(SELECT ID FROM Users WHERE UserName = "admin"),
		(SELECT ID FROM Users WHERE UserName = "admin"),
		(SELECT ID From Projects WHERE Code = "ISSZP"));