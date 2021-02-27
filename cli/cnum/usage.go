package main

const usage = `
<Cnum>
Cnum is a suite of programs designed to convert a number into Cistercian Numerals.
Usage:
        cnum <command> [arguments]
The commands are:
	<html>-creates server that hosts a User Interface via HTML webpage
		argument:[port#]
			sets a port number for the html index endpoint
			cnum <command> [portnumber] default port 9000																		
	<api>-creates a API server for URI parameters returning cnumber
		argument:[port#]
			sets a port number for api endpoint
			cnum <command> [portnumber] defaul port 9000
	<local>-creates a cnum based on the argument
		arguemnt:[0-9999]
			converts argument into cnum
			
Example:

$ ./cnum local 3725

output:
   ####
   # # 
   ##  
####   
   #   
   #   
#  #  #
 # #  #
  ##  #
   ####
`

// arguments
// 		random
// 			random Cistercian Numberal 0-9999
// 		countup
// 			Cistercian Numberals incrementally counting from 0
// 		countdown [number]
// 			Cistercian Numberals decrementing from a number to 0
