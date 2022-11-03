#!/bin/sh

go run test/main.go test/mnt &
sleep 0.3

String=$(cat test/mnt/Text)
if [[ $String != 'This is a text content' ]]; then
    echo 'TEST FAILED: file "Text" does not match struct value'
    fusermount -zu ./mnt
    exit 1
fi

Int=$(cat test/mnt/Integer)
if [[ $Int != '2222' ]]; then
    echo 'TEST FAILED: file "Integer" does not match struct value'
    fusermount -zu ./mnt
    exit 1
fi

find test/mnt/Sub >> /dev/null
if [[ $? != 0 ]]; then
    echo 'TEST FAILED: dir "Sub" does not exist'
fi

value=$(cat test/mnt/Sub/SomeValue)
if [[ $value != 20 ]]; then
    echo 'TEST FAILED: file "SomeValue" does not match struct value'
    echo $value
    fusermount -zu ./mnt
    exit 1
fi

sleep 2
Text=$(cat test/mnt/Text)
if [[ $Text != 'This is a changed text content' ]]; then
    echo 'TEST FAILED: file "Text" does not match struct value'
    fusermount -zu ./mnt
    exit 1
fi

fusermount -zu test/mnt
echo 'TEST PASSED'
