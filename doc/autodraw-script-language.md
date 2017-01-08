# Autodraw Script Language
==========================

## Basic Idea

A **Graphic** is a combination of drawing operations.
Strictly speaking, a graphic is what you will see after applying the set of
operations on a white clean canvas.
But here, to make it easy to express, just call the set of operations Graphic.

A drawing operation is expressed explicitly by exactly a line of code.
The code is in form of [OPERATION] [ARGUMENTS].
For example

```
line 0 0 1 1
line 0 1 1 0
circle 0.5 0.5 0.707
```

The above codes simply draw a pair of crossing lines and a circle around them.
