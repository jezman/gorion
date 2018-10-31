package helpers

const (
	QueryEvents = `SELECT p.Name, p.FirstName, p.MidName, c.Name, TimeVal, e.Contents, a.Name
		FROM pLogData l
		JOIN pList p ON (p.ID = l.HozOrgan)
		JOIN pCompany c ON (c.ID = p.Company)
		JOIN Events e ON (e.Event = l.Event)
		JOIN AcessPoint a ON (a.GIndex = l.DoorIndex)
		WHERE TimeVal BETWEEN ? AND ?
		AND e.Event BETWEEN 26 AND 29
		ORDER BY TimeVal`
	QueryEventsByEmployeeAndDoor = `SELECT p.Name, p.FirstName, p.MidName, c.Name, TimeVal, e.Contents, a.Name
		FROM pLogData l
		JOIN pList p ON (p.ID = l.HozOrgan)
		JOIN pCompany c ON (c.ID = p.Company)
		JOIN Events e ON (e.Event = l.Event)
		JOIN AcessPoint a ON (a.GIndex = l.DoorIndex)
		WHERE TimeVal BETWEEN ? AND ?
		AND e.Event BETWEEN 26 AND 29
		AND p.Name = ?
		AND DoorIndex = ?
		ORDER BY TimeVal`
	QueryEventsByEmployee = `SELECT p.Name, p.FirstName, p.MidName, c.Name, TimeVal, e.Contents, a.Name
		FROM pLogData l
		JOIN pList p ON (p.ID = l.HozOrgan)
		JOIN pCompany c ON (c.ID = p.Company)
		JOIN Events e ON (e.Event = l.Event)
		JOIN AcessPoint a ON (a.GIndex = l.DoorIndex)
		WHERE TimeVal BETWEEN ? AND ?
		AND e.Event BETWEEN 26 AND 29
		AND p.Name = ?
		ORDER BY TimeVal`
	QueryEventsByDoor = `SELECT p.Name, p.FirstName, p.MidName, c.Name, TimeVal, e.Contents, a.Name
		FROM pLogData l
		JOIN pList p ON (p.ID = l.HozOrgan)
		JOIN pCompany c ON (c.ID = p.Company)
		JOIN Events e ON (e.Event = l.Event)
		JOIN AcessPoint a ON (a.GIndex = l.DoorIndex)
		WHERE TimeVal BETWEEN ? AND ?
		AND e.Event BETWEEN 26 AND 29
		AND DoorIndex = ?
		ORDER BY TimeVal`
	QueryEventsDenied = `SELECT p.Name, p.FirstName, p.MidName, c.Name, TimeVal, e.Contents, a.Name
		FROM pLogData l
		JOIN pList p ON (p.ID = l.HozOrgan)
		JOIN pCompany c ON (c.ID = p.Company)
		JOIN Events e ON (e.Event = l.Event)
		JOIN AcessPoint a ON (a.GIndex = l.DoorIndex)
		WHERE TimeVal BETWEEN ? AND ?
		AND e.Event IN (26, 29)
		ORDER BY TimeVal`
	QueryWorkedTimeByEmployee = `SELECT p.Name, p.FirstName, p.MidName, c.Name, min(TimeVal), max(TimeVal)
		FROM pLogData l
		JOIN pList p ON (p.ID = l.HozOrgan)
		JOIN pCompany c ON (c.ID = p.Company)
		WHERE TimeVal BETWEEN ? AND ?
		AND p.Name = ?
		GROUP BY p.Name, p.FirstName, p.MidName, c.Name, CONVERT(varchar(20), TimeVal, 104)`
	QueryWorkedTimeByCompany = `SELECT p.Name, p.FirstName, p.MidName, c.Name, min(TimeVal), max(TimeVal)
		FROM pLogData l
		JOIN pList p ON (p.ID = l.HozOrgan)
		JOIN pCompany c ON (c.ID = p.Company)
		WHERE TimeVal BETWEEN ? AND ?
		AND c.Name = ?
		GROUP BY p.Name, p.FirstName, p.MidName, c.Name, CONVERT(varchar(20), TimeVal, 104)`
	QueryWorkedTime = `SELECT p.Name, p.FirstName, p.MidName, c.Name, min(TimeVal), max(TimeVal)
		FROM pLogData l
		JOIN pList p ON (p.ID = l.HozOrgan)
		JOIN pCompany c ON (c.ID = p.Company)
		WHERE TimeVal BETWEEN ? AND ?
		GROUP BY p.Name, p.FirstName, p.MidName, c.Name, CONVERT(varchar(20), TimeVal, 104)`
	QueryEmployees = `SELECT p.Name, p.FirstName, p.MidName, c.Name FROM pList p
		JOIN pCompany c ON (c.ID = p.Company)
		ORDER BY c.Name`
	QueryEmployeesByCompany = `SELECT plist.Name, pList.FirstName, pList.MidName, c.Name FROM pList
		JOIN pCompany c ON (c.ID = Company)
		WHERE c.Name = ?
		ORDER BY pList.Name`
	QueryCompanies = `SELECT c.Name, Count(pList.Name) FROM pList
		JOIN pCompany c ON (c.ID = Company)
		GROUP BY c.Name`
	QueryDoors = "SELECT GIndex, Name FROM AcessPoint ORDER BY GIndex"
	QueryEventsValues = "SELECT Event, Contents, Comment from Events"
)
