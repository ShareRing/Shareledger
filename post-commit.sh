#!/bin/bash


#branch=$(git rev-parse --symbolic --abbrev-ref $refname)


# IDENTITY
idFile="~/.ssh/sharering_server"
workdir="/data/shareldeger_dev/shareledger"
destdir="$workdir/bin"
datadir="$workdir/data"
sourcedir="./build"
exeFile="main_linux"
destFile="main_linux"
host="192.168.1.234"
port="22012"
username="trangtran"

branch=$(git branch | grep \* | cut -d ' ' -f2)

function moveExecutableFile(){
    if [ "develop" == "$branch" ]; then
        # Do something
        echo "Branch: Develop"

        echo "Build executable file for Linux"
        make build_linux

        if [[ $? != 0 ]]; then
            echo "Error in building executable for Ubuntu"
            exit
        fi

        echo "Copy to folder"
        scp -i "${idFile}" -P "${port}"  "${sourceDir}/${exeFile}" "${username}@${host}:${destDir}"

        if [[ $? != 0 ]]; then
            echo "Error in copy file"
        fi
    else
        echo "Not develop branch. Exit"
    fi

}


function executeCommand() {
    SSH_COMMAND="ssh -i ${idFile} ${username}@${host} -p ${port}"
    
    # Remove old data
    $SSH_COMMAND rm -rf "$dataDir"
    
    # Setup for single node
    $SSH_COMMAND "$dataDir/scritps/setup.sh" 1

    # Turn on all tags
    $SSH_COMMAND sed -i -e "/index_all_tags/ s/false/true/g" "$dataDir/main0/node0/config/config.toml"

    # Start
    $SSH_COMMAND "$dataDir/docker.sh 1 up"

}

#moveExecutableFile
executeCommand "$@"
