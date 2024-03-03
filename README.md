# Blog Aggregator Project
## Overview
This Blog Aggregator is a robust backend system designed to periodically scrape RSS feeds from various blogs, aggregate posts, and store them in a Postgres database for easy access and management. Built with Go, it features user authentication, feed management, and an efficient way for users to follow their favorite blogs. 
This guide walks you through setting up the project, including database configuration, Go environment setup, and running the aggregator.

## Features
- RSS Feed Scraping: Automatically fetches and updates blog posts from specified RSS feed URLs at regular intervals.
- User Authentication: Supports user account creation and authentication, enabling personalized feed following.
- Feed Management: Allows users to add, follow, and unfollow blog feeds.
- Database Storage: Utilizes Postgres for storing user information, blog feeds, and posts.
- Concurrency Control: Implements go routines for concurrent RSS feed fetching, optimizing performance.
- Customizable Scraping Intervals: Offers the ability to set scraping frequency and concurrency levels.

## Prerequisites
Before starting, ensure you have the following:
- Basic understanding of SQL for querying relational databases.
- A text editor and command line access (VS Code with Zsh recommended).
- Go programming language installed (Go (Version 1.15 or later, check with command: ```go version```) if not, follow instructions at https://go.dev/doc/install).
- An HTTP client for testing (e.g., Thunder Client for VS Code, curl, Postman).
- Clone the Repository: Clone this project to your local machine.
- Dependencies: Install required Go dependencies using go get or manage them using Go Modules.
### PostgreSQL
- A preferred SQL client (e.g., PGAdmin for Postgres)
- Postgres Database Setup: Install Postgres and create a database for the project. Ensure the Postgres server is running.
- Environment Variables: Configure necessary environment variables, including database connection strings.

# Running the Project

## Run data migrations
Data migrations are crucial for setting up your database schema before running the Blog Aggregator. This section guides you through the process of creating necessary tables and performing initial data setup in your Postgres database

### Migration Tool
The project uses Goose for database migrations, a lightweight tool that helps manage your database schema. 
Goose documentation can be found here: https://github.com/pressly/goose.

Setting Up Migrations
Install Goose can be installed by running: ```go get -u github.com/pressly/goose/cmd/goose```

### Migration Files:
Migration files are located in the sql/schema directory of the project. Each file contains SQL commands for creating or altering database structures. 

Navigate to the project root directory and execute the following command to apply migrations ```goose postgres "your_database_connection_string" up``` This command runs all pending migrations, creating the necessary tables and structures in your database.

Rollback: If you need to revert the last migration, you can use the down command ```goose postgres "your_database_connection_string" down```. This is useful for undoing changes if a migration does not go as planned.

### Run server
After all migrations have been run (they are numbered from 001 to 006) your db structure is ready for API interactions.
- Start the Server: Navigate to the project directory and run ```go build``` followed by ./[project-name] to start the server.
- Feed Management: Use the provided API endpoints to add, follow, and unfollow RSS feeds.
- View Posts: Access aggregated posts through the designated endpoint, filtering by the feeds you follow.

## Extending the Project
- Frontend Integration: Develop a frontend interface or a CLI tool for better interaction with the aggregator.
- Enhanced Error Handling: Implement comprehensive error handling and logging for better debugging and stability.
- Feed Customization: Add features for users to customize feed update intervals and prioritize feeds.
