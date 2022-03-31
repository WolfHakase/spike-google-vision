# spike-google-vision-api
This is a test project to see how to interact with the google vision api.

In this project I send some photos of Digimon cards to the Google Vision api and see  what can be done with the output data.

## What is a spike?
According to Wikipedia: 

`A spike is a product development method originating from extreme programming that uses the simplest possible program to explore potential solutions. It is used to determine how much work will be required to solve or work around a software issue. Typically, a "spike test" involves gathering additional information or testing for easily reproduced edge cases. The term is used in agile software development approaches like Scrum or Extreme Programming.`

In these spikes no use is made of TDD. In fact, the code is largely untested. 

The code isn't necessarily clean. 

It is merely a tool to understand the subject, so that the actual production code can more easily be tested and made clean.

An added benefit to spikes is that they don't contain anything worth stealing, therefore they're perfect to build a portfolio with.

## Google Vision
Google Vision is the AI image recognition api owned by Google.

Documentation on Google Vision can be found here: https://cloud.google.com/vision

## How to run
To run this project first change `credentials.template.json` to `credentials.json`.

Then fill in the Google Vision credentials for your project in that file.

How to setup your Vision API project you can read here: https://cloud.google.com/vision/docs/setup.

Once you've done that running this project is nothing special.

From Goland go to the `main.go` and use the hotkey `ctrl + shift + f10`.

Or if you don't use a fancy IDE just run it like you would run any other Go script.