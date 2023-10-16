# Mnemosine
## Goal

The objective of this project is to design and develop a modern 
web application tailored for educational purposes. 
The aim is to establish a comprehensive system that empowers students, particularly those pursuing medicine and law, to efficiently review large volumes of information, mainly for exams.
The inspiration for this project is conceived from the renowned 'Anki' application, 
it utilizes famous study techniques to enhance memory retention via testing yourself. The whole app revolves around 
flashcards and testing your knowledge.

The primary focus is not solely on the application itself, but rather the implementation, design and challenges
that will be needed to achieve an MVP.

## Features:

- FlashCards: The main unit of data in this app are flashcards. They should give questions and answer and the frontend should be able to present the information with a quick and easy user experience
- Testing: testing should produce test results that could be used to feed or optimize the algorithm for memory retention.
- Decks: FlashCards can be stored in Decks, decks can be used internally by each user or they can be shared and added to other user's personal decks
- Semantic Search: you should be able to find data related to your topic, note that it's not a simple text search, meaning has to be present in the results.
- AI or NLP powered FlashCard generation: give the model data and it should return new flashcards for you to test yourself.

## Architecture:

Mainly Microservices Architecture.
I know it's overkill for the scope, but  this is an educational project.
This will also be a monorepo.

Diagram coming soon.

## Stack:

- Golang:  Services/ API
- Typescript:  FrontEnd with React
- Python:  LLM 
- C: quickly parsing and error checking incoming data.

## Databases:

- MongoDB - FlashCard service
- SQLite or Mysql for Users and data produced by the app.
- VectorDatabase: TBA


