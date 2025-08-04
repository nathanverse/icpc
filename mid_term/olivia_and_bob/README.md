# Oliver and the Game

Oliver and Bob are best friends. They spent their childhood in the beautiful city of Byteland. The people of Byteland live happily with the King.

The city has a uniform architecture with a total of N houses. The King's Mansion is a very large and beautiful wooden house with address = 1. The remaining houses in Byteland have a unique address, are connected by roads, and there is always a unique road between any two houses in the city. Note that the King's Mansion is also included in these houses.

Oliver and Bob have decided to play Hide and Seek, taking the entire city as their playing area. In one scenario of the game, it is Oliver's turn to hide and Bob has to find him.

Oliver can hide in any house in the city including the King's Mansion. Since Bob is very lazy, to find Oliver he either walks towards the King's Mansion (he stops when he gets there), or he walks away from the Mansion on any path until he reaches the last house on that path.

Oliver hides in a house (call it X) and Bob starts the game from his house (call it Y). If Bob goes to house X, then he is sure to find Oliver.

Given Q queries, you need to tell Bob whether he can find Oliver or not.

Queries can be of two types:

* `0 X Y`: Bob moves towards the King's Mansion.
* `1 X Y`: Bob moves away from the King's Mansion.

## Input

The first line of Input contains a single integer N, the total number of houses in the city. The next N-1 lines contain two space-separated integers A and B representing a road between two houses at addresses A and B.

The next line contains an integer Q representing the number of queries.

The next Q lines contain three space-separated integers representing each query as explained above.