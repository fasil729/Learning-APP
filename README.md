# Brilliant AI-Driven Learning Web Application for High School and Primary School

## Table of Contents
1. [Youtube Demo](#youtube-demo)
2. [Devpost Submission](#devpost-submission)
3. [Inspiration](#inspiration)
4. [What It Does](#what-it-does)
5. [How We Built the Project](#how-we-built-the-project)
6. [Challenges We Ran Into](#challenges-we-ran-into)
7. [Accomplishments We Are Proud Of](#accomplishments-we-are-proud-of)
8. [What We Learned](#what-we-learned)
9. [Built With](#built-with)

## Youtube Demo
[Youtube Demo Link](https://youtu.be/ht7JvNNaosQ?si=GnB5MjN-ZAGmlZPe)

## Devpost Submission
[Devpost Submission Link](https://devpost.com/software/brilliant-5cdvfl)

## Inspiration
Our inspiration for creating this project comes from the significant challenges faced by students in Ethiopia's educational system. Many schools across the country lack sufficient resources such as reference books and teaching aids, making it difficult for students to understand theoretical concepts through practical application. This contributes to generally low student performance, evidenced by the fact that only 3% of students pass the university entrance examination.

As a team, we've experienced these struggles firsthand. During our primary and high school years, we faced similar challenges due to the lack of laboratory facilities and other essential resources. This left us with gaps in understanding because we couldn't experiment or gain hands-on experience with scientific concepts.

Our project aims to address these issues by providing a web application that not only simplifies complex concepts but also encourages interactive learning. By focusing on high school and primary school students, we hope to improve the education system and ultimately create a better future for our community.

## What It Does
Our web application tackles educational challenges in Ethiopia by providing a suite of features that enhance the learning experience:

- **Clear Explanation of Science Concepts**: The website uses images and diagrams to explain complex scientific concepts, aiding students' understanding of experiments, processes, and results.
- **Simplified Science Experiments**: It offers step-by-step instructions for experiments, using locally available materials, making them accessible to students without laboratory resources.
- **AI-Powered Note-Taking**: The web app includes an AI-assisted note-taking system, allowing students to create visualized notes with embedded images and animations.
- **Interactive Quizzes**: To promote active learning, the platform includes quizzes and games that reinforce educational content and make learning fun.
- **Automated Summarization and Explanation**: The web app can automatically generate summaries, explanations, and quizzes, streamlining study sessions and exam preparation.
- **Tailored for High School and Lower-Level Students**: Content is specifically designed for high school curricula, offering resources and exam preparation materials that align with educational standards.

By combining these features, our web application not only addresses existing educational gaps but also aims to improve student performance and engagement. Our goal is to inspire creativity and a passion for learning among students.

## How We Built the Project
Our project journey began with a clear focus on solving the problem outlined in our inspiration. We used this foundation to establish clear objectives that guided us throughout the development process. The first step was designing the system architecture, carefully selecting the most suitable technologies to meet our requirements.

- **Technology Stack**: We chose TypeScript for the front-end, offering strong type safety and integration with modern JavaScript features. For the back-end, we utilized Go/Golang to manage server-side logic and API endpoints. This combination allowed us to create a robust and scalable system.
- **Project Planning**: To maintain an organized workflow, we outlined a detailed project timeline and assigned specific tasks to team members. Project management tools like Jira helped us track progress and manage deadlines effectively.
- **Development**: The development phase was a collaborative effort. Our team worked diligently to write clean, efficient code, following best practices and industry standards. We used Git for version control to manage code changes and facilitate smooth collaboration among team members.

## Challenges We Ran Into
Developing a new project from scratch presents unique challenges. The first hurdle was the limited time frame to turn our idea into a working solution. We had to brainstorm quickly, create a detailed plan, and assemble a collaborative team, all while managing our university coursework and other commitments.

A significant technical challenge arose when we couldn't access a free image generation API. Budget constraints prevented us from purchasing a premium service, and despite exploring various options, we couldn't find a suitable free image generation API. This obstacle prevented us from fully implementing the image generation feature in our web application. However, we've developed the underlying code, allowing for seamless integration once we find a suitable solution.

We also faced technical complexities in formatting responses from Gemini to meet the structure required for the frontend. This required additional effort to ensure smooth integration. Additionally, we encountered issues with rendering React components on markdown files, which added to the development complexity.

Despite these challenges, we remained flexible and found creative solutions to keep the project on track. While the lack of image generation API is a setback, it's an issue that can be addressed once we find the right resource. Overall, we're proud of the progress we've made and we overcame these obstacles and successfully delivered a project we are proud of.

## Accomplishments We Are Proud Of
Contributing to a project designed to improve education in our country is a significant accomplishment for our team. Our work aims to enhance the understanding and overall performance of students by providing a platform that helps them grasp complex concepts through interactive explanations, experiments, and visualizations. By doing so, we aim to positively impact the future generation, who will play a critical role in shaping Ethiopia's development.

Another aspect of our pride lies in the quality and precision of our web application. Using advanced technologies like Google's Gemini API, we have created a platform that can generate accurate lessons and chapter summaries, complete with images, visualizations, and step-by-step experiments. This level of detail and interactivity is rare in educational applications, making our project unique. The platform's ability to guide students through complex experiments with clear procedures and visuals is a testament to our team's dedication to excellence. We're proud to have built a tool that not only meets educational needs but also inspires curiosity and learning.

## What We Learned
During the project, we gained invaluable insights into both the technical and collaborative aspects of software development. We discovered how Artificial Intelligence (AI) can be applied to solve everyday problems, providing a powerful tool for addressing broader community issues. This project highlighted the significance of approaching complex challenges with a structured mindset and adapting to new technologies and frameworks as needed.

We also learned the critical role effective communication plays within a team, ensuring everyone is aligned and understands their responsibilities. Moreover, the project underscored the importance of flexibility, allowing us to pivot and adjust our approach when requirements evolved or unexpected obstacles arose.

Overall, this project has been a remarkable learning experience, reinforcing the value of teamwork, adaptability, and innovative problem-solving. We look forward to applying the lessons learned to future projects.

## Built With
- elephantsql
- gin
- go
- google-gemini (llm)
- next.js (front-end)
- postgresql
- render
- typescript
