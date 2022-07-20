Every Pattern´s Readme must follow the next structure:

# Facade Pattern

## Description
a short description to help understand the pattern
what does the pattern do?
## Motivation
What problem solves the pattern
## Code Example
In the example we define a character whoose state is altered in many ways.
A character has a lot of attributes and components that varies over time like speed, momentum, position, etc. that conforms the character´s extrinsec state and should not be part of the character's information because it will slow down the whole application whenever we try to carry all that information around.
Instead we define the intrinsic state of the character that just has the escential attributes, those that will not change very much nor at all throughout the object's lifetime.

## Applicability
The Flyweight pattern is often used when you need to retrieve some object information but not an entire object. The Flyweight pattern is also called Cache.

## Structure
the pattern´s structure 
## Pros and Cons

### Pros
why should i use this pattern
### Cons
why shouldn´t i