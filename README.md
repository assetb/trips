# Trains and Towns
 Initial description of the problem see ...

 # How to execute
 1. Provide graph data in db/input1.dat file. First Test data provided.
 2. do
 cd cmd
 go build
 ./cmd

 # Install go(lang)
 https://golang.org/doc/install

# Code Description
Objects like Town, Edge, Route are introduced to describe problem and to be used on business level.
This version uses Town as a string with 1 length.
Edge describes a elementary direct route between two towns without in-between towns or stops.
Route is a collection of edges.

To solve problems it is used two approaches to handle a given collection of routes.
First one collects them in object collection.
Second one - in map collection.
Some tasks are convenient to solve by 1-st approach, others - by 2-nd.
Particularly search tasks are best addressed by second approach.

Algorithm decsriptions are done in comments of appropriate function
