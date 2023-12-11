#!/usr/bin/env sh

day=$(printf '%02d' $1)
MAIN_TEMPLATE="./.templates/dayN.py"
TEST_TEMPLATE="./.templates/test_dayN.py"

cp $MAIN_TEMPLATE days/day${day}.py
cp $TEST_TEMPLATE tests/test_day${day}.py

sed -i '' "s/{day}/${day}/g" days/day${day}.py tests/test_day${day}.py

