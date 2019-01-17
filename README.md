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
4. front end uses AJAX to ask for config changes every 3 seconds
5. if front end receives 'next' image command it changes image
6. Lazy load of images
7. refreshes the page when all images have been cycled

#TODO
1. backend sepparate image tags when reading image
2. backend filter images based on config and tag of the images
3. read config file and update image list when its a different type
3. front end receive change of image type and act on it
2. front end either refresh the page with the new type, or ask for images of that type and redo the DOM

