Every Pattern´s Readme must follow the next structure:

# Facade Pattern

## Description
Command is a behavioral design pattern that turns a request into a stand-alone object that contains all information about the request. This transformation lets you pass requests as a method arguments, delay or queue a request’s execution, and support undoable operations. Good software design is often based on the principle of separation of concerns, which usually results in breaking an app into layers. For example: a layer for the graphical user interface another layer for the business logic and data handling. representing the front and back end architecture (Model View-Controller) in any web application. 

## Motivation
What problem solves the pattern
## Code Example
Imagine we want to sell the loot we got from our previous adventure in a dungeon. We would need to go to a merchant and be able to sell our items by selecting the item we want to sell and then clicking the button "sell", but we want to spend the less possible time on the merchant to go and get some more loot asap, so we need to create a new shortcut to sell the items we want without having to click twice, one for the item and one for the button by ctrl+clicking the item we wish to trade selling it instantly. Also we want to undo the previous transaction if we ever missclick a valuable or quest item by clicking the button "undo" or by pressing ctrl+Z. Then we must implement this 2 commands logic appart from the buttons.

## Applicability
When to apply this pattern
## Structure
the pattern´s structure 
## Pros and Cons


### Pros
why should i use this pattern
### Cons
why shouldn´t i