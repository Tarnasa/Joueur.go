#!bin/bash

cd ../
godoc
wget -m -k -q -erobots=off -X src/ --no-host-directories --no-use-server-timestamps http://localhost:6060
