## Design details

Using the command design pattern here, where I will treat the different types of bots as the receiver instances.

```go
type DvonnBot interface {

    playNextMove() bool

}
```

Now, there can be multiple types of Bots like `BeginnerBot`, `IntermediateBot`, `AdvancedBot`, `GrandmasterBot`.
Each of them will have to implement `playNextMove()` method.

eg:
```java
type BeginnerBot struct implements DvonnBot {
    // Implement playNextMove()
}
```
etc...

Now, I will be using a Command type which will be coupled with the receiver object.

```go
type Command struct {
    Execute() void
}
```
Now, we can make some very specific commands. Note that a Command object is like an operation which needs to be performed by the receiver, so we will have as many Command class as we have the number of operations to perform.
For eg: in this case we just have one operation to perform i.e. play next move.

```go
type PlayRandomMove struct implement Command {
    bot DvonnBot
    PlayRandomMove(bot DvonnBot) {
        this.bot = bot
    }
    func Execute() {
        bot.execute()
    }
}
``` 

We can also have an invoker object which should have be coupled with the Command, i.e. when method of this object is invoked then the command should get executed.

```go
type MoveInvoker struct {
    command Command
    MoveInvoker(command Command) {
        this.command = command
    }
    func move() {
        command.Execute()
    }
}
```


