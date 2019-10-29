package models

//Cell defines the different states of a single Cell
type Cell struct {
	Clicked bool `json:"clicked"` //Clicked specifies wether the Cell has been clicked or "discovered"
	Mine    bool `json:"mine"`    //Mine specifies if there's a mine in the cell or not
	Flag    bool `json:"flag"`    //This field indicates if the user has marked this Cell with a Flag (for future implementations)
	Number  int  `json:"number"`  //Number is the quantity of nearby mines this cell has.
}

//CellRow is a slice of cells. One or more of these form a board
type CellRow []Cell

//Game has the information necessary to create a new game
type Game struct {
	Name       string    `json:"name"`            //Name acts as an identifier of the Game
	Rows       int       `json:"rows"`            //How many rows the board has
	Columns    int       `json:"columns"`         //How many columns the board has
	Board      []CellRow `json:"board,omitempty"` //This is the structure itself of the board, many rows of cells
	Discovered int       `json:"discovered"`      //This is the amount of cells already discovered. Used to check if the status is victory or not
	Mines      int       `json:"mines"`           //How many mines the board has
	Status     string    `json:"status"`          //Status of the current game. In progress, Game Over, Victory.
}
