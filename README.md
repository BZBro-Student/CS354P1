# Project #: Project Name

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

Write a two paragraph reflection describing your experience with this
project.  Talk about what worked well and what was challenging.  
Did you run into an issue that took some time to figure out?  
Tell us about it. What did you enjoy, what was less desirable? Feel
free to add other items (within the two paragraph limit).

For the main part of the assignment I feel it was fairly hard to get used to. I haven't used C a bunch
as i've only had one class in which we used C so getting used to the GO overall structure was hard to do.
Probably the first hard thing to implement was the account.go interface while keeping the abstract nature of 
the account 'class'. It took me a while to figure out that since an empty method counts as an implementation
the solution is just to not do anything with it and just leave it in the interface and the interface will force
it's implementation later on. 

## Compiling and Using

This section should tell the user how to compile your code.  It is
also appropriate to instruct the user how to use your code. Does your
program require user input? If so, what does your user need to know
about it to use it as quickly as possible?

## Results

This section presents timing and other results of any experiments that
you were asked to perform as part of the project.

## Sources used

If you used any sources outside of the lecture notes, class lab files,
or text book you need to list them here. If you looked something up on
stackoverflow.com and fail to cite it in this section it will be
considered plagiarism and be dealt with accordingly. So be safe CITE!
https://go.dev/tour/concurrency/2
https://go.dev/tour/concurrency/9


----------

## Notes

* This README.md template is using Markdown. Here is some help on using Markdown:
  [markdown cheatsheet](https://github.com/adam-p/markdown-here/wiki/Markdown-Cheatsheet)


* Markdown can be edited and viewed natively in most IDEs such as Eclipse and VS Code. Just toggle
  between the Markdown source and preview tabs.

* To preview your README.md output online, you can copy your file contents to a Markdown editor/previewer
  such as [https://stackedit.io/editor](https://stackedit.io/editor).
