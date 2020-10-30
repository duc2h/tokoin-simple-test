# tokoin-simple-test

## How to run?
1. To run process : run command line: go run main.go
2. To run unit test: go to floder testing, then run command line: go test + filename.go

## Explain what i do.
### Structure, design code:
* I apply structure this because it's quite easy to understand, clean and highly reuseable.
* We have folder datas to init data from file .json, and mock data to unit test. 
* Folder repositories are defined method can use to find data from memory. Folder repo-impl is implement method from repositories, it applies worker pool to make better performance (i'll explain it late). 
* Folder businesses are defined method to processing logic. Folder business-impl is implement method from businesses and interact with repositories to get data.

### Goroutine, Worker pool:
* Goroutine make golang become power. It can create hundreds of other thread processing but uses few resource. So i decide apply worker pool into my project to make processing find data faster.
![image](https://user-images.githubusercontent.com/36435846/97753759-e0421100-1b28-11eb-9b57-c491670496a4.png)
