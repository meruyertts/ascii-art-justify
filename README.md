# ascii-art-justify

Ascii-art is a program which consists in receiving a string as an argument and outputting the string in a graphic representation using ASCII. 

A graphic representation using ASCII, is to write the string received using ASCII characters, as you can see in the example below:  
```
@@@@@@BB@@@@``^^``^^``@@BB$$@@BB$$
@@%%$$$$^^^^WW&&8888&&^^""BBBB@@@@
@@@@@@""WW8888&&WW888888WW``@@@@$$
BB$$``&&&&WWWW8888&&&&8888&&``@@@@
$$``&&WW88&&88&&&&8888&&88WW88``$$
@@""&&&&&&&&88888888&&&&&&88&&``$$
``````^^``^^^^^^````""^^``^^``^^``
""WW^^@@@@^^``````^^BB@@^^``^^&&``
^^&&^^@@````^^``&&``@@````^^^^&&``
``WW&&^^""``^^WW&&&&""``^^^^&&88``
^^8888&&&&&&WW88&&88WW&&&&88&&WW``
@@``&&88888888WW&&WW88&&88WW88^^$$
@@""88&&&&&&&&888888&&``^^&&88``$$
@@@@^^&&&&&&""``^^^^^^8888&&^^@@@@
@@@@@@^^888888&&88&&&&MM88^^BB$$$$
@@@@@@BB````&&&&&&&&88""``BB@@BB$$
$$@@$$$$$$$$``````````@@$$@@$$$$$$
```
This project handles an input with numbers, letters, spaces, special characters and \n.  

Some banner files with a specific graphical template representation using ASCII will be given. The files are formatted in a way that is not necessary to change them.

* shadow
* standard
* thinkertoy

To change the alignment of the output you must use a flag --align=<type>, in which type can be :

* center
* left
* right
* justify

Program is adapted to the terminal size. If you reduce the terminal window the graphical representation is adapted to the terminal size.

Only text that fits the terminal size will be tested.

The flag must have exactly the same format as above, any other formats must return the following usage message:
```
Usage: go run . [OPTION] [STRING] [BANNER]

Example: go run . --align=right something standard
```
# Usage

```
$ go run . "something" thinkertoy --align=left
$ go run . "something" standard --align=center
```