#!/bin/bash

echo "post-start start" >> $HOME/status

# this runs each time the container starts
mkdir -p go/bin

echo "post-start complete" >> $HOME/status
