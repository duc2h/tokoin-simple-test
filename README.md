# tokoin-simple-test

## How to run?
1. To run process : run command line: go run main.go
2. To run unit test: go to floder testing, then run command line: go test + filename.go

## Explan what i do.
### Structure, design code:
* I apply structure this because it's quite easy to understand, clean and highly reuseable.
* We have folder datas to init data from file .json, and mock data to unit test. 
* Folder repositories are defind method can use to find data from memory. Folder repo-impl is implement method from repositories, it applies worker pool to make betterFolder performance (i'll explain it late). 
* Folder businesses are defind method to processing logic. Folder business-impl is implement method from businesses and interact with repositories to get data.
