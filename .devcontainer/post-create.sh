#!/bin/bash

echo "post-create start" >> $HOME/status

# add your commands here
go install github.com/spf13/cobra/cobra@latest

echo "post-create complete" >> $HOME/status
