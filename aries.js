var height
var width
var images
var indexCount = 1
var toogle=true;
var fader;
var imageRefresh;

window.onload = function start() {
    images = shuffle(window.imageList.split(","))
    height = window.screen.availHeight + 39 //compensate for scroll bars
    width = window.screen.availWidth - 15  //compensate for scroll bars
    imageRefresh = window.imageRefresh
    fader = new Fader('fader', 4);
    changeImage();
    canvas();
}

function changeImage() {
    var imgNextId = (!toogle ? 1 : 0);
    
    document.getElementById(imgNextId).src = images[indexCount];
    var img = document.getElementById(imgNextId); 

    img.height = calculateMisingHeight(width);
    img.width = width;
    
    if (++indexCount >= images.length){
        indexCount = 0
        images = shuffle(images)
    }

   toogle = !toogle
   fader.setTarget(toogle ? 1 : 0)
}

function calculateMisingHeight(img_width){
    return (height / width) * img_width 
}

function canvas() {
    window.setInterval(function () {
        changeImage()
    }, imageRefresh);  
}

function shuffle(array) {
    var currentIndex = array.length, temporaryValue, randomIndex;
  
    while (0 !== currentIndex) {
      randomIndex = Math.floor(Math.random() * currentIndex);
      currentIndex -= 1;
      temporaryValue = array[currentIndex];
      array[currentIndex] = array[randomIndex];
      array[randomIndex] = temporaryValue;
    }
    return array;
  }
