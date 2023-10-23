#!/bin/sh

# note this is meant to run in the root of the repository

# remove previous running instances
make stop
make clean

# build the image and run it
make
make run
