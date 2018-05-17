#!/bin/sh

for test in $(find . -name "*_test.rb")
do
	ruby $test
done
