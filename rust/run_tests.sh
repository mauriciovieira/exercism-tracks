#!/bin/sh

cd rust

for exercise in $(find . -type d -mindepth 1 -maxdepth 1)
do
	cd $exercise
	cargo test
	cd ..
done
