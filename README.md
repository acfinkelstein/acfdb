# ACF Database

## Overview

This is a light-weight database developed in Go as part of a technical assessment.

## Setup and Execution

This project contains a Makefile that will facilitate in building and executing the database code.

`make build` will compile the database and generate an executable in the directory `cmd/acfdb`

`make run` will run the database program.

## Usage

ACFDB accepts a number of commands as input. The following commands and associated arguments are supported.

* SET [name] [value]
  * Sets the name in the database to the given value
  * Does not print any value
* GET [name]
  * Prints the value for the given name in the database
  * If the name does not exist, prints NULL
* DELETE [name]
  * Deletes the value for the given name in the database
  * Does not print any value, and does not notify if the name does not exist
* COUNT [value]
  * Returns the number of names that have the given value assigned to them.
  * If the value does not exist in the database, prints 0
* BEGIN
  * Begins a new transaction
  * Does not print any value
* ROLLBACK
  * Discards the changes from the most recent transaction
  * Prints TRANSACTION NOT FOUND if no transactions exist; does not print any value otherwise
* COMMIT
  * Commits all open transactions
  * Does not print any value
* END
  * Exits the database