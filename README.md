#codebag

simple CLI code snippet collect tool

##Usage

Add Snippets:

    cb add [-m=message] [-t=tags] [-r] [files/directory...]
    cb add -m "utf-8 conversion function" utf8.c
    cb add -m "simple hello world with vim"
    cb add -m "desc" -t "python, golang"
    
Remove Snippets:

    cb rm <id>...
    cb rm 1
    cb rm 1 2 3 4 5

List Snippets:

    cb ls


    e.g.
    $cb ls  
      	3	utf-8 conversion	2014-04-21 10:12:34	
      	2	hello world	2014-04-21 10:12:23	
      	1	Rpc server impl	2014-04-21 10:06:47
    
Get Snippet Conent:
    
    cb get <id>...
    
    e.g.
    $cb get 2
    #2 hello world 2014-04-21 10:12:23 
    
    hello world

