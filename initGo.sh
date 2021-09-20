#!/bin/bash
check_empty(){
    if [ -z $3 ] || [ -z $4 ]
    then
        echo "[ERROR] $1 and $2 are required"
        exit 0
    fi
}

setup_gitconfig(){
    local setup=$1
    local token_name=$2
    local pat=$3

    if [ $setup = 1 ]
    then
        check_empty 'token name' 'pat' $token_name $pat
        echo '===== setup PAT ====='
        echo "token name: ${token_name}"
        echo "pat: ${pat}"
        git config --global url."https://$TOKEN_NAME:$PAT@$ORG_URL/".insteadOf "https://$ORG_URL/"
        echo '===== setup PAT SUCCESS ====='
    else
        git_config_path=~/.gitconfig
        grep "insteadOf = https://$ORG_URL/" $git_config_path > /dev/null
        if [ $? -eq 1 ]
        then
            echo '[ERROR] need to setup PAT'
            exit 0
        fi
    fi
}

setup_bashrc(){
    bashrc_path=~/.bashrc
    if [ -f $bashrc_path ]
    then
        grep '^export GO111MODULE=auto$' $bashrc_path > /dev/null
        if [ $? -eq 1 ]
        then
            echo 'export GO111MODULE=auto' >> $bashrc_path
        fi
        grep '^export PATH=$(go env GOPATH)/bin:$PATH$' $bashrc_path > /dev/null
        if [ $? -eq 1 ]
        then
            echo 'export PATH=$(go env GOPATH)/bin:$PATH' >> $bashrc_path
        fi
    else
        echo 'export GO111MODULE=auto' >> $bashrc_path
        echo 'export PATH=$(go env GOPATH)/bin:$PATH' >> $bashrc_path
    fi
}

read -p 'organization url:' ORG_URL
read -p 'project name:' PROJECT_NAME
read -p 'repo:' REPO
check_empty 'project name' 'repo' $PROJECT_NAME $REPO

read -p 'setup PAT (Y or N)' SETUP_PAT
if [ $SETUP_PAT = 'Y' ] || [ $SETUP_PAT = 'y' ] || [ $SETUP_PAT = 'N' ] || [ $SETUP_PAT = 'n' ]
then
    if [ $SETUP_PAT = 'Y' ] || [ $SETUP_PAT = 'y' ]
    then
        read -p 'input Token Name: ' TOKEN_NAME
        read -p 'input PAT: ' PAT
        setup_gitconfig 1 $TOKEN_NAME $PAT
    else
        setup_gitconfig 0 '' ''
    fi
else
    echo '[ERROR] not expected SETUP_PAT value'
    exit 0
fi

setup_bashrc

init_go_path=$(pwd)
source $bashrc_path

cd ~/go
path=$ORG_URL/$PROJECT_NAME/_git/$REPO.git
go get $path
go get -u github.com/swaggo/swag/cmd/swag

cd $init_go_path
cp -r ./* ~/go/src/$path
rm ~/go/src/$path/initGo.sh

cd ~/go/src/$path/
sed -i '' "s/testmodule/$REPO/g" ./go.mod ./internal/handlers/* ./main.go
go run main.go
