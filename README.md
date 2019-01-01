# aries
Digital Canvas 

Service to pull up images on a specific resolution and orientation, based on artists, period, art style, movie posters, etc.

Should be a webservice so its easily accessible from different canvas hardware with different options

The service will also do a morph from one image to the next one, the time to change images is configurable

The service should integrate to IFTTT to be able to integrate it to google home.

Possible solutions:
- Enable tweets from google home via IFTTT
- Have this service monitor a twitter account
- Have the service react to messages via twitter

e.g.
- Ok google, set the canvas to cubism
- IFTTT tweets {canvas:cubism}
- Service reads the "cubism" tag and applies filters on its images to show only those that match

# Current Functionality
1. Server loads the imgs directory, sends the html and js back with the filenames(random order) on the html
2. front end (aries.js) loads the images and circles through them with a fade utility.
3. front end scales the images to the appropiate resolution

