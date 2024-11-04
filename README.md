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

Install the packages from package.json

```bash
npm start
```

## 🍿 Step Five

Once you have the packages installed, set up your enviroment variables and are connected to your MongoDB you can run the application.

```bash
nodemon run start
```

This should then allow you to run the application in your web browser in the following location: http://localhost:3000.

From here you can set up a user and register and account and sign in.

Once you are signed in you can add reviews to books.

If you need to add books you will need to set up an Admin account in your MongoDB and set the user to isAdmin: true this will enable that user to add, edit and delte books.
