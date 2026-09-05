# Project #: CS354 - P1

* Author: Broden
* Class: CS354 Section 
* Semester: Fall 2026

## Overview

Program simulates a basic overview of a bank with individual accounts which may accrue interest depending on the account type. Accounts 
store customer info and the account balance. Customers have a name and no other information stored in them. The main method simulates 
a Bank running an accrual on 3 accounts.
Extra Credit:
The extra credit makes 1000 small concurrent deposits into an account with an expected end balance of $1000.

## Reflection

For the main part of the assignment I feel it was fairly hard to get used to. I haven't used C a bunch
as i've only had one class in which we used C so getting used to the GO overall structure was hard to do.
Probably the first hard thing to implement was the account.go interface while keeping the abstract nature of 
the account 'class'. It took me a while to figure out that since an empty method counts as an implementation
the solution is just to not do anything with it and just leave it in the interface and the interface will force
its implementation later on. 
The hardest part of this assignment by far was the multithreading experiments. Really the only parallelized program
I've written one that involved piping data output over ssh to batch process some calculations on ras-pis. Multithreading
truly is a different way of thinking about programming because a big aspect of programming is understanding the flow of instructions
and for the longest time it's been one instruction and then the next but now you have to consider one instruction which branches into 
many different instructions at the same time and having to make sure the result of those doesn't interfere with each other or the next instruction
in the main control flow. I eventually figured out channels after looking at the examples provided and figured out mutexs by looking at the 
intro to Go webpage.


## Compiling and Using

Running Test Script:
Ensure that run-test.sh is set to be executable then use the command: ./run-test.sh

Running Program:
go run main.go

Creating Executable:
go build main.go

No user input is needed to run the program

## Results

The test script runs without issue, due to the extra credit it does point
out a difference in expected output. When run normally the output is as follows:
"
3: bob: 153.00
1: ann: 100.00
2: ann: 204.00
Total Interest: 7.00
4: bob: 1000.00
"

## Sources used
Very useful intro to go specifics:
https://go.dev/tour/concurrency/2
https://go.dev/tour/concurrency/9
Userful implementation examples of abstract functions:
https://www.reddit.com/r/golang/comments/163s90l/how_to_overcome_the_lack_of_abstract_functions_in/
https://nikhilakki.in/implement-an-abstract-class-in-go
Interesting read I found that directly relates to module one of the course:
https://noahs-blog.net/?p=377