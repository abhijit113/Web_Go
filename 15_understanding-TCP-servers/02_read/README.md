An open connection blocks.

The reader is reading from the open connection.

How does the reader know when it should stop reading?

Instructions: run this file, then in your browser go to:

http://localhost:8080/
http://localhost:8080/here I am/


Open cmd
Type netstat –ano
List of process with their ports will be opened
Search ‘process ID’ of the port you are unable to use (in my case port 11020)
Open task Manager and Stop that process
Now your port is ready to use :)
