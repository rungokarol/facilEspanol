# Application intended for learning a foreign language a'la Duolingo

## Motivation:

Main goal of this app is to learn new technologies by the creators. Decision has been made to start work in pair-programming mode, 100% remotely.
Creators agreed to begin with backed part of application. Golang has been selected as a main technology.
Frontend is postponed for now on. It's development will start when minimal working version of backed is ready.
The idea of the application is based on Duolingo.  Our goal is not to copy but to challenge ourselves in developing efficient features to keep going with learning Spanish language.

## Key features:
- main language is English
- foreign langue is Spanish
    - one language available to learn for now
- flashcards
- day strike rating
- interactive exercises
    - mainly writing
    - translating sentences both ways
    - at the beginning choosing the right answer would be sufficient
- explain grammar rules with basic examples
- user can store his/hers progress
- user can test latest gained knowledge

## Technical details
Programming language: Golang + Typescript
Communication with frontend: REST API
Data storage: Postgres
Unit testing framework: Testify/Jasmine
E2E testing: Protractor

##Features
- user registration & logging
- selecting words to flashcards by category
- lesson reviewing
- storing user's strike
- mail notifications when strike is in danger
