# dvonn-bot

There can be several ways of writing a bot, the approach I am using is reading out state spaces of games played and taking further decisions.
Basically, this shall be considered as a supervised learning approach as I have a set of training data.
It's not clear how efficient/inefficient this approach is in comparision with other non-supervised learning approach, but you can create a benchmark
against Monte Carlo Tree Search.

Anyways, let me name this supervised approach as out-and-out tree search algorithm.

Basically, out-and-out tree search algorithm will majorly involve 2 steps:
1) Expansion
2) Back propagation

- Expansion is expanding the underlying decision graph whenever any new state is encountered in the sample state space.
 
- Back propagation is basically rolling out to the top and updating the states when you reach to the end of the sample state space graph.

