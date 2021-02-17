Cnumber is a suite of programs designed to convert 0-9999 numbers into a Cistercian Numerals. Interactive local client inclue:

    Local - Interactive local client that is A CLI tool that can convert cnum arguements into a formatted Cistercian Numeral formatted for console output.
    HTML - Interactive server webpage for conversion. 
    API - Server endpoint that returns a body that contains the Cistercian Numeral.

_Useage:_


Local:

> cnum <local> <cnumber> 
    returns:
>           #######
>           cnumber
>           #######

HTML:

> cnum <html> <port>

API:

> cnum <api> <port>

'''
[serverip/dn]:[port]/?snum[number0-9999]
'''

Command 	Explanation

```
API 	    creates a API server for URI parameters returning cnumber
HTML        creates server that hosts a User Interface via HTML webpage
Local       runs a command line cnumber translation
```
