# **Upcoming Movies API** 🍿🎬

## ⭐️ Overview

This project uses a movies database api from [Rapid Api](https://rapidapi.com/hub). It tells you what movies are coming up based on the year or by genre, all depending what you search for in the command line.

It uses:

```text
🔹 Go 
🔹 Cobra Library
🔹 Movies API
🔹 Command Line Interface (CLI)
```

## ⚙️ Prerequisites

In order to run this application you will need to have the following:

```text
🔸 Basic knowledge of the CLI
🔸 Basic Knowledge of GO
🔸 VS code installed
```

## 🍿 Step One

Change your directory to where you wish to run this script and store the cloned repository:

```bash
cd <filename>
```

## 🐾 Step Two

Clone the repository from github and then move into the new directory.

```bash
git clone https://github.com/laurawarren88/go_movies_api.git
cd go_movies_api
```

## 🍿 Step Three

Take a look 👀 around the file 📂 structure and see what is happening with VS code.

```bash
code .
```

You may need to double check the API Key is still correct. Take a look in the config.json file for the location of the key. Compare this against the Key from [Rapid Api](https://rapidapi.com/SAdrian/api/moviesdatabase/playground/apiendpoint_b9f58be7-d8b9-405b-ad3a-48fc5117a2bf)

## 🍿 Step Four

To take a look at the up and coming movies by year or by genre, you need to call the right commands.

For the year:

```bash
go run main.go year
```

*You can only search movies from 2024 to 2029*

For the genre:

```bash
go run main.go genre
```

*Here's a list of searchable genres: Action, Adult, Adventure, Animation, Biography, Comedy, Crime, Documentary, Drama, Family, Fantasy, Film-Noir, History, Horror, Musical, Mystery, Romance, Sci-Fi, Short, Sport, Thriller, War, Western*

## 🍿 Step Five

Everything you search for in the CLI is saved to the data.log file.

You can view the file in your IDE or via the terminal.

To search the file in your terminal, use the view cmd and the number of lines you want to view.

```bash
go run main.go view 5
```

When the data.log file gets too crowded or you have finished with the information you can clear the data with the clear cmd.

```bash
go run main.go clear
```

This will clear the whole file so you can start again.
