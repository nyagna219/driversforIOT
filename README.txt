The go folder contains the EDGE code
to run it 
$go run rest.go


The http folder contains the code for the cloud.
You will need a webserver ( eg apache with php enabled ) to run it.
Just point the server root to the http directory.

Xinu-code-BeagleBoneBlack_Listener contains the code to be RUN on the Sub commander BBB.
Sub commander BBB becomes the master in absence of any edge and cloud.


Xinu-code-BeagleBoneBlack_Listener_client contains the code to be run on the slave BBB.
Slave BBB accepts commands from both edge and the above Sub commander BBB.

ddl.json contains the DDL.
the DDL parser it located at XInu base directory named ddl_parser.py


