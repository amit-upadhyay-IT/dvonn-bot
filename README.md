# dvonn-bot

There can be several ways of writing good bots like AlphaGo is written using the Monte Carlo tree search. Few other
bots can be written using the minimax algo, these categories of algorithms can be used along with the reinforcement learning
to produce good results.

I'm writing a different kind of bots for this game `dvonn`:
- **BeginnerBot**: decides by randomly selecting moves from the possible set of moves. This is a dumb bot without any intelligence and so it's faster in terms of the decision and lighter in terms of size.
- **AdvancedBot**: It has some intelligence based on the previously played games.
- **IntermediateBot**: It is a mixture of BeginnerBot and AdvancedBot, by mixture I mean that some move it will decide using BeginnerBot and the other moves it will decide using the AdvancedBot.
- **GrandmasterBot**: This one also takes a decision based on the previously played games, it's just that here the number of games played is huge and so, its decision making capability is nice.

## An approach to the intelligent Bot

Here, I am writing bot with a pretty naive algorithm (very different from MCTS, MINIMAX) with some optimizations totally based on the specific domain.
Well, as this bot is designed specifically for a kind of domain so this bot will be tightly coupled with the dvonn game only or similar kind of games with a small set of state spaces.

This algorithm can briefly state into:
1) **Expansion of the Game Tree**: expanding the underlying decision graph whenever any new state is encountered in the sample state space.
2) **Confidence calculation** using parameters from the above Game Tree.

Just to sound little fancier, let me name this algorithm as an **out-and-out tree search**.

**So, on what basis are we going to do expansion?**
Well, for this I have generated lots of games played between two dummy bot. Using that data, I am iterating over the games played and expanding the Game Tree and assigning a few parameters value based on the result of that particular game.

So, unlike other game bots which are generally reinforcement learning, this **out-and-out tree search** algorithm will be a kind of supervised learning algorithm as lots of training data is available here.

**Confidence calculation**:
- Uses a formula (`float64(vi) + math.Sqrt(math.Log(float64(visitCount)/float64(ni)))`) to calculate confidence value for nodes, it is a function of winning the reward, node visit count and parent visit count.


## Problem faced:
- Here one huge problem I faced was the size of the models.

Eg: I tried generating game samples for 1,00,000 games. And the data set size came around 60~70MBs. On generating the confidence model tree, the size came out around 180MBs, and on compressing the data the size got reduced around 25 MBs, which is also a huge number.


**Few points to note in the above model is**:
- Tried matching the similar game states in the Game tree, I couldn't find one single game played at least twice in the games samples. Anyway, it is close to impossible that two games will occur with the same moves as if you see the different ways for filling out the chips in the placement phase would be 49! (a huge huge possibility).
- I was only able to find common moves for around `2~3` steps from beginning and that too around `40~50 times` in the whole sample set of 1 Lac games.

So, the above model concluded nothing. Because even to have at-lease one repetition of the game I would need around 50 Lac ~ 1 crore sample sets. Whose size you can imagine now. 

### A good observation:
- The depth of the confidence tree of the game tree was only up to 91. i.e. the search time would be a constant operation i.e. O(1) time complexity. But, yes space complexity is really really huge for at least an intermediate bot. Well, of course this kinda decision making capability you can't find in algorithms like MCTS :P

## Optimisation in the previous model

- If you notice, the first phase of the game (i.e. placement phase) has some strategies involved but they don't matter much if the two opponents are smart enough to understand the basics of the game, they will easily be able to figure out the placement phase decisions. Thus, I decided to have some conditions for the placement phase and do a fair distribution of the chips.
- So, removing the possible combinations in the placement phase is was would make the model light weight. As there would be huge possible actions available.
- Now, after the placement phase is removed from the learning model, we can see a huge improvement in terms of size and of course the model depth.

