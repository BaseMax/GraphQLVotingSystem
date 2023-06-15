# GraphQL Voting System

This project is a GraphQL-based voting system developed in TypeScript. It provides functionalities for user authentication, registration, retrieving the list of votes, attending to a vote, submitting answers for questions within a vote, and administrative operations such as managing votes, questions, and viewing application details.

## Features

- User Authentication: Users can register and log in to the system.
- Vote Management:
  - Get List of Votes: Retrieve a list of all available votes.
  - Attend to a Vote: Attend a specific vote and save the start date and time.
  - Submit Answers: Users can submit their answers for the questions within a vote.
- Administrative Operations:
  - List of Applications: As an admin, view the list of applications for all votes.
  - Get List of Votes: Retrieve a list of all votes.
  - Create Vote: Admins can create a new vote.
  - Add Question: Add a new question to a vote.
  - Edit Question: Edit an existing question.
  - Delete Question: Delete a question from a vote.
  - View Answers: Admins can view all answers for a specific question.

## Technologies Used

- TypeScript: The project is developed using TypeScript, which provides static typing and enhances code maintainability.
- GraphQL: The API is built using GraphQL, a query language for APIs, which allows efficient data retrieval and flexible data querying.
- Node.js: The backend is developed using Node.js, a JavaScript runtime, which enables server-side execution of the GraphQL API.
- MongoDB: The project uses MongoDB as the database for storing vote details, questions, user information, and answers.

## Installation

Clone the repository:

```shell
git clone https://github.com/basemax/GraphQLVotingSystem.git
cd GraphQLVotingSystem
```

Install the dependencies:

```shell
npm install
```

Configure the environment variables:

- Create a .env file in the project's root directory.
- Define the required environment variables, such as database connection details and authentication tokens.

Build the TypeScript code:

```shell
npm run build
```

Start the server:

```shell
npm start
```

Access the GraphQL API at http://localhost:3000/graphql.

## Usage

Once the server is up and running, you can interact with the GraphQL API using a tool like GraphQL Playground or Insomnia.

Here are some example queries and mutations that you can execute:

**User Registration:**

```graphql
mutation {
  registerUser(username: "example_user", password: "password") {
    id
    username
  }
}
```

**User Login:**

```graphql
mutation {
  loginUser(username: "example_user", password: "password") {
    token
    user {
      id
      username
    }
  }
}
```

**Get List of Votes:**

```graphql
query {
  votes {
    id
    title
    startDate
    endDate
    questions {
      id
      text
    }
  }
}
```

**Attend to a Vote:**

```graphql
mutation {
  attendVote(voteId: "vote_id") {
    id
    title
    startDate
  }
}
```

**Submit Answer:**

```graphql
mutation {
  submitAnswer(
    voteId: "vote_id"
    questionId: "question_id"
    answer: "example_answer"
  ) {
    id
    text
    answer
  }
}
```

**Get List of Votes:**

```graphql
query {
  votes {
    id
    title
    startDate
    endDate
    questions {
      id
      text
    }
  }
}
```

**List of Applications (Admin Only):**

```graphql
query {
  applications {
    id
    vote {
      id
      title
    }
    user {
      id
      username
    }
    answers {
      id
      question {
        id
        text
      }
      answer
    }
  }
}
```

**Get List of All Votes (Admin Only):**

```graphql
query {
  allVotes {
    id
    title
    startDate
    endDate
    questions {
      id
      text
    }
  }
}
```

**View Answers for a Specific Question (Admin Only):**

```graphql
query {
  questionAnswers(questionId: "question_id") {
    id
    user {
      id
      username
    }
    answer
  }
}
```

### Mutations

**User Registration:**

```graphql
mutation {
  registerUser(username: "example_user", password: "password") {
    id
    username
  }
}
```

**User Login:**

```graphql
mutation {
  loginUser(username: "example_user", password: "password") {
    token
    user {
      id
      username
    }
  }
}
```

**Attend to a Vote:**

```graphql
mutation {
  attendVote(voteId: "vote_id") {
    id
    title
    startDate
  }
}
```

**Submit Answer:**

```graphql
mutation {
  submitAnswer(
    voteId: "vote_id"
    questionId: "question_id"
    answer: "example_answer"
  ) {
    id
    text
    answer
  }
}
```

**Create a New Vote (Admin Only):**

```graphql
mutation {
  createVote(title: "New Vote", startDate: "2023-06-15T10:00:00Z", endDate: "2023-06-20T18:00:00Z") {
    id
    title
    startDate
    endDate
  }
}
```

**Add New Question to a Vote (Admin Only):**

```graphql
mutation {
  addQuestion(voteId: "vote_id", text: "New Question") {
    id
    text
  }
}
```

**Edit a Question (Admin Only):**

```graphql
mutation {
  editQuestion(questionId: "question_id", text: "Updated Question") {
    id
    text
  }
}
```

**Delete a Question (Admin Only):**

```graphql
mutation {
  deleteQuestion(questionId: "question_id") {
    id
    text
  }
}
```

**Get User Information:**

```graphql
query {
  user {
    id
    username
    votes {
      id
      title
    }
  }
}
```

**Get Vote Details:**

```graphql
query {
  vote(voteId: "vote_id") {
    id
    title
    startDate
    endDate
    questions {
      id
      text
    }
  }
}
```

**Create Admin User (Admin Only):**

```graphql
mutation {
  createAdminUser(username: "admin_user", password: "admin_password") {
    id
    username
    isAdmin
  }
}
```

**Delete Vote (Admin Only):**

```graphql
mutation {
  deleteVote(voteId: "vote_id") {
    id
    title
  }
}
```

**Change Answer (For Modifiable Votes):**

```graphql
mutation {
  changeAnswer(voteId: "vote_id", questionId: "question_id", answer: "new_answer") {
    id
    text
    answer
  }
}
```

**Change Answer (For Non-Modifiable Votes):**

```graphql
mutation {
  changeAnswer(voteId: "vote_id", questionId: "question_id", answer: "new_answer") {
    id
    text
    answer
    isModifiable
  }
}
```

These additional examples cover various scenarios such as retrieving user information, obtaining vote details, creating an admin user, deleting a vote, and changing answers for both modifiable and non-modifiable votes. Remember to replace the placeholder values such as vote_id and question_id with the actual IDs of the votes and questions you want to interact with.

Please refer to the API documentation or code comments for more detailed information about the available queries, mutations, and their respective arguments.

## Data Structure

Here's an example of how you can structure the models and schemas using Prisma with TypeScript for the GraphQL-based voting system:

Install Prisma:

```shell
npm install prisma --save-dev
npx prisma init
```

Define the Models:

In your Prisma schema file (schema.prisma), define the models for the voting system:

```prisma
model User {
  id        Int      @id @default(autoincrement())
  username  String   @unique
  password  String
  isAdmin   Boolean  @default(false)
  votes     Vote[]
}

model Vote {
  id          Int       @id @default(autoincrement())
  title       String
  startDate   DateTime
  endDate     DateTime
  questions   Question[]
  applications Application[]
}

model Question {
  id      Int    @id @default(autoincrement())
  text    String
  voteId  Int
  vote    Vote   @relation(fields: [voteId], references: [id])
}

model Application {
  id          Int         @id @default(autoincrement())
  voteId      Int
  userId      Int
  vote        Vote        @relation(fields: [voteId], references: [id])
  user        User        @relation(fields: [userId], references: [id])
  answers     Answer[]
}

model Answer {
  id          Int       @id @default(autoincrement())
  questionId  Int
  applicationId Int
  question    Question  @relation(fields: [questionId], references: [id])
  application Application @relation(fields: [applicationId], references: [id])
  answer      String
}
```

Generate Prisma Client:
Run the following command to generate the Prisma Client:

```shell
npx prisma generate
```

Use Prisma Client in TypeScript:

In your TypeScript files, you can import and use the generated Prisma Client to interact with the database. For example:

```typescript
import { PrismaClient } from '@prisma/client';

const prisma = new PrismaClient();

async function getUserVotes(userId: number) {
  const user = await prisma.user.findUnique({
    where: { id: userId },
    include: { votes: true },
  });
  return user.votes;
}

async function createVote(title: string, startDate: Date, endDate: Date) {
  const vote = await prisma.vote.create({
    data: {
      title,
      startDate,
      endDate,
    },
  });
  return vote;
}

// ... other functions using Prisma Client ...

export { getUserVotes, createVote };
```

Make sure to import the PrismaClient from @prisma/client and create an instance of it (prisma) to use the various Prisma methods for CRUD operations.

Copyright 2023, Max Base
