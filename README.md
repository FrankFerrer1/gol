# ngp-server
This is a server dedicated to our queen. Its actually a backend to the frontend application located here: https://github.com/Velaseriat/ost-player

# how to run
There are 2 makefiles which makes all this happen. One is at the root of the project and the other is in the src directory

The app is deployed via using docker so make sure that you install that for whatever OS you have. If on linux use your distros package manager to do it.
On windows you can probably just grab it from the actual website and install it as an exe: https://www.docker.com/products/docker-desktop/
It's called docker Desktop.

Once you ahve it installed you can use the makefile at the root in order to use it.

## Targets in makefile:

```
make
```
builds the image with the current code

```
make run
```
runs the built image (you have to make sure you run make so it has the board)
    
```
make stop
```
stop the currently running server
    
```
make clean
```
deletes the image built by make

You can also run it locally on your machine without docker using your computers go install by using the makefile in src. Basically
run 
```
make 
```
to build the binary and 

```
make clean
```
to delete it

# setting up
in the CLI run:
docker-compose up --build
will build everything and run the script(s)

# pg-admin

When setting up the pgadmin connection, set the following in Connection when registering a server

host name/address: 172.26.0.1 (should be the default IP that docker gives)
port: 5432
Maintenance database: assets
Username: postgres
password: postgres

### HAPPY COOMING