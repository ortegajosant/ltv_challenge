# Technical Challenge

### Info
This repo store the code which implement the required tasks to complete a technical challenge.
This project was developed in `Linux Ubuntu 20.04`

### Requirements

1. First of all, clone the repository using one of the following:
- `git clone` *https, ssh or Github CLI*
- `git remote add origin` *https, ssh or Github CLI*

2. Then, create a folder named `sql` in the root path of the project:
- `mkdir sql`

3. This project use a DB based on SQLite, you can download it from:
- https://s3.amazonaws.com/bv-challenge/jrdd.db.

4. This should download a file named `jrdd.db`, put this file into the `sql` folder created in step 3.

### To execute

1. Run the command: 
	-`go mod init`
	This will set the project as a Go module.
2. Run the :
	-`go mod tidy`
3. To compile the code and run it, execute:
	- `go build main.go`
	- `./ltv_challenge`
4. If you just want run:
	- `go run main.go` 
5. To run the tests, execute:
	- `go test ltv_challenge/tests`

The server will be running in the PORT 5000.

## API Reference

The developed API has the following endpoints:
1. ### Songs:
	- GET:
		- `/api/songs/artist/:artist` 
		This will give you a list of songs found by specific artist name.
		- `/api/songs/genre/:genre`
		This will give you a list of songs found by specific genre name.
		- `/api/songs/name/:name`
		This will give you a list of songs found by specific song name.
		- `/api/songs/length`
		This will give you a list of songs found by specific length range. You have to send also a body in JSON format as following:
		`{
			"max": int,
			"min": int
		}`
2. ### Genres:
	- GET:
		-  `/api/genres/search/:id`
		This will give you a genre found by specific ID.
		-  `/api/genres/allWithInfo`
		This will give you all the genres with the number of songs and the total length for each genre.

## Additional Questions
1. Are there any suggestions you could make about the structure of our data?
	- Some column names could be more descriptive, For example: changing the columns `Genre` to `GenreId`, `Song` to `SongName` or just `Name`.
	- Add one more table which could be named `Artists` and change the `Songs` table column named `artist` to `ArtistId`. Then, you can associate a lot of songs with the same artist.

2. What fields would you index in these tables?
	- The tables would be like:
		1. `Songs`:
			1. `Id`		
			2. `ArtistId`	
			3. `GenreId`	
			4. `Name`
			6. `Length`	
		2. `Genres`:
			1. `Id`
			2. `Name`
		3. `Artists`:
			1. `Id`
			2. `Name`

Developed by: Jose Antonio Ortega González
Email: ortega.josat@gmail.com
GitHubUser: ortegajosant
