#!/bin/bash

echo "on-create start" >> $HOME/status

# add your commands here
go install github.com/spf13/cobra/cobra@latest

echo "on-create complete" >> $HOME/status
