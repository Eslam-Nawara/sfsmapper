#!/bin/sh

go run ./main.go mnt &
sleep 0.3

String=$(cat ./mnt/Text)
if [[ $String != 'This is a text content' ]]; then
    echo 'TEST FAILED: file "Text" does not match struct value'
    fusermount -zu ./mnt
    exit 1
fi

Int=$(cat ./mnt/Integer)
if [[ $Int != '2222' ]]; then
    echo 'TEST FAILED: file "Integer" does not match struct value'
    fusermount -zu ./mnt
    exit 1
fi

find ./mnt/Sub >> /dev/null
if [[ $? != 0 ]]; then
    echo 'TEST FAILED: dir "Sub" does not exist'
fi

value=$(cat ./mnt/Sub/SomeValue)
if [[ $value != 20 ]]; then
    echo 'TEST FAILED: file "SomeValue" does not match struct value'
    echo $value
    fusermount -zu ./mnt
    exit 1
fi

sleep 2
Text=$(cat ./mnt/Text)
if [[ $Text != 'This is a changed text content' ]]; then
    echo 'TEST FAILED: file "Text" does not match struct value'
    fusermount -zu ./mnt
    exit 1
fi

fusermount -zu ./mnt
echo 'TEST PASSED'
