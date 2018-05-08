from __future__ import print_function
import json
from pprint import pprint


with open('../ddl.json') as data_file:    
        data = json.load(data_file)

print(data["ddl"]["devices"]["Led"]["off"])
#print(data)


with open('../include/ledputc.h', 'w') as f:
    print ("/*GENERATED FILE ; DO NOT EDIT*/", file=f)
    print ("volatile uint32 *l = devptr->dvcsr;", file=f)
    print ("*l = data;",file=f)
    print ("return OK;",file=f)


with open('../include/ledinit.h', 'w') as f:
    print ("volatile uint32 *l = devptr->dvcsr;", file=f)
    print ("*l = ",file=f , sep=' ', end='' )
    print (data["ddl"]["devices"]["Led"]["off"] ,file=f , sep=' ', end='' )
    print(";",file=f)
    print ("return OK;",file=f)

with open('../include/constants.h', 'w') as f:
    print ("/*GENERATED FILE ; DO NOT EDIT*/", file=f)
    print ("#define RLEDIP ",file=f , sep=' ', end='' )
    print (data["ddl"]["devices"]["Rled"]["Pin"] ,file=f , sep=' ', end='' )
    print(" ",file=f)

