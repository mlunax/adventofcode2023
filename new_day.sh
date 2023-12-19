#!/bin/sh
# Code borrowed from https://github.com/Daste745/aoc-2023/blob/master/new_day.sh
if [ -z "$1" ]
then
  cat << EOF
Usage: new_day.sh <day_number>

Example: new_day.sh 08
EOF
  exit 1
fi

day_number=$1
if [ -d "day$day_number" ]
then
  echo "Day$day_number already exists"
  exit 1
fi

mkdir -p ./day$day_number/
cp -v ./day_template/day.go ./day$day_number/day$day_number.go

