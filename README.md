### Purpose
This service was implemented to provide a RESTful API for the game Minesweeper
The framework used was Go Kit. This particular structure provides us with several features
that ensure an optimal modularization to achieve a scalable solution.
The service core contains the main functionality of the game, and several layers of middlewares can be applied to it.
At the moment, the implemented middleware serves the function of logging transactions, but this structure allows us to add other
features in the future, such as authentication or metrics.
Endpoints wraps the services with their middlewares and gives us a notion of single Endpoints each related with the transport layer(http)
Encoding and Decoding is also treated apart, so in the future this service could interact with diferent transports, such as a Publish/Subscriber architecture
(NATS - Kafka - GRPC)
Storage is in the db package. It uses a map to store the games, but this could be easily upgraded to a proper DB thanks to the design of the interface and package itself

The whole structure of this project is aimed towards a scalable and adaptable solution.

Currently deployed into an aws ec2 instance.

### Endpoints

Create Game

This endpoint creates a new game and stores it locally for future uses.
Name is the identifier of each game, rows and columns specifies the size of the board, discovered is the number of current clicked cells
Mines let us set any quantity of them.  

POST ec2-18-216-153-136.us-east-2.compute.amazonaws.com:48080/minesweeper/games


body:
{
    "name": "minetest",
    "rows": 6,
    "columns": 6,
    "discovered": 0,
    "mines": 4,
    "status": "start"
}

This request should receive a 201 response (Created).

 ----------------------------------------------------------------------------------------------------------------------------------------

Load Game

This endpoint search the desired game in the local storage and return it in Json format.
The response is an object containing properties of the game as well as the state of it.

GET ec2-18-216-153-136.us-east-2.compute.amazonaws.com:48080/minesweeper/games/minetest

{
    "clicked": false,
    "mine": false,
    "flag": false,
    "number": 0
},
Above is the basic structure of each Cell. LoadGame returns an object containing many of these.
This information would be used for visual purposes, such as painting unvisited/visited cells, flagging, etc.

-----------------------------------------------------------------------------------------------------------------------------------------

Click

This endpoint allows to interact with the current state of the game. Each request represent a click in one of the cells of the grid.
When a cell without nearby mines is clicked, all the surrounded cells are clicked, like the original minesweeper.
After each click the state of the game is checked. In the response (A JSON object) the state of the game can be seen, whether is still continuing, 
game over or victory.

PUT ec2-18-216-153-136.us-east-2.compute.amazonaws.com:48080/minesweeper/games

{
    "name": "minetest",
    "row": 4, //from 0 to maximum rows -1
    "column": 3 //from 0 to maximum columns -1
}
