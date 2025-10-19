package main
/*
Interfaces are Soooo Cooool!!!!!
*/

// this is the contract that the db must confirm to "plug in"
type dbcontract interface {
	Close()
	InsertUser(uname string)error
	SelectSingeUser(uname string)(string, error)
}

type Application struct {
	db dbcontract // db is not any specifc like pg,mongo but interface

}


func NewApp(db dbcontract) * Application{
	return &Application{db:db}
}

func main() {
	// db, err := mysql.NewConfig().DBName would return instance of mysql
	// the db must implement all the same methods defined in the interface
	// we can package the db with our application now we can call on db fucntions so fluently
}

/*

we can switch out the db easy af since we just need our db to implement the expectedf interface
*/