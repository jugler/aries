window.onload = function start() {
    fader = new Fader('fader',3);
    changeImage();
    canvas();
}

function canvas() {
    window.setInterval(function () {
        changeImage()
    }, 300000);  
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

function changeImage() {
    var imgCurrentId = "img" + (toogle ? 1 : 0);
    var imgNextId = "img" + (!toogle ? 1 : 0);

    document.getElementById(imgNextId).src = "imgs/" + images[indexCount];
    var img = document.getElementById(imgNextId); 
    var newHeight = calculateMisingHeight(1050)
    img.height = newHeight;
    img.width = 1050;
    
    if (++indexCount >= images.length){
        indexCount=0
        images = shuffle(images)
    }

   toogle=!toogle
   fader.setTarget(toogle ? 1 : 0)
}

function calculateMisingHeight(img_width){
    return (height / width) * img_width 
}


var height=1680
var width=1050

var image1="banksi.jpg"
var image2="dbz.png"
var image3="matrix.jpg"
var image4="monalisa.jpg"
var images = shuffle([image1,image2,image3,image4])
var indexCount = 1
var toogle=true;
var fader;