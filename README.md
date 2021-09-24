# lem-in
The task of the project is [here](https://github.com/01-edu/public/tree/master/subjects/lem-in)
# Chapters
- [Briefly about the project](Briefly-about-the-project)
- [Graphs in file](#Graphs-in-file)
- [How use program](#How-use-program)
- [How Program works](#How-Program-works)
- [Authors](#Authors)
# Briefly about the project
This is a console program that builds a graph based on the data transmitted by the user and distributes the ants along the paths so that everyone reaches point `B` in a minimum of several steps.
# Graphs in file
```txt
#comment
10
##start
start 1 5
##end
end 4 5
Room1 2 4
R2 3 4
Room3 2 6
Name 3 6

# Here is Paths
start-Room1
start-Room3
Room1-R2
Room1-Name
Room3-R2
Room3-Name
R2-end
Name-end
``` 
# How use program
> Program takes only 1 argument. 
It can be filename or flags.
#### Default:
- `"filename"` - your file path with graph data
#### Flags:
- `--file="filename"` - as default input. Just an explicit launch with a file
- `--http=":port"` - starts run server on your port. Using for visualization.

For run project:
```bash
go run . (filename | --http=:port | --file=filename)
```
### Example
# How Program works
[About algoritm](http://www.macfreek.nl/memory/Disjoint_Path_Finding)
# Authors
- [Dias1c](https://github.com/Dias1c)