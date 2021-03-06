# Facade Pattern


## Description
Facade provides a unified interface to a set of interfaces in a subsystem. 
Facade defines a higher-level interface that makes the subsystem easier to use.

## Motivation
When a package has multiple interfaces, it can get difficult for clients to use it.
Structuring a system into subsystems helps reduce complexity. A common design goal is to minimize the communication and dependencies between subsystems.
One way to achieve this goal is to introduce a facade object that provides a single, simplified interface to the more general facilities of a subsystem

## Code Example

In the example given, we define 2 models: 
A character, that will use a potion 
A potion that we asume will be used by several parts of the code, in this example, by a character.

Potion has many different methods, but we define an interface iPotion that will define the contract we use to interact with the rest of the application. It doesnt have the full methods collection, just the necessary ones like "UsePotion" or "RechargePotion"

This way, we simplify the implementation that the rest of the code must use to interact with the potion interface.
To use a potion, a character must call the method "UsePotion" and doesnt need to deal with the rest of the functions.

## Applicability
 Use the Facade pattern when:
    * you need to have a limited but straightforward interface to a complex subsystem.
    * you want to structure a subsystem into layers.
    

## Structure

facade structure image


## Pros and Cons

### Pros
 isolates your code from the complexity of a subsystem.

### Cons
A facade can become a god object coupled to all classes of an app.