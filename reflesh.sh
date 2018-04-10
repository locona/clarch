#/bin/bash

set -e

for i in `ls  | egrep -v '(bindata|cmd|Gopkg*|vendor|clarch|Makefile|README.md|main.go|.git|reflesh.sh)'`;
do
  rm -rf $i
done
