function shuffle(array) {
    var currentIndex = array.length, temporaryValue, randomIndex;
  
    // While there remain elements to shuffle...
    while (0 !== currentIndex) {
  
      // Pick a remaining element...
      randomIndex = Math.floor(Math.random() * currentIndex);
      currentIndex -= 1;
  
      // And swap it with the current element.
      temporaryValue = array[currentIndex];
      array[currentIndex] = array[randomIndex];
      array[randomIndex] = temporaryValue;
    }
  
    return array;
  }

function changeImage() {
    
    document.getElementById("img").src="imgs/"+images[indexCount];
    indexCount++
    if (indexCount >= images.length){
        indexCount=0
        images = shuffle(images)
    }
}

var image1="banksi.jpg"
var image2="dbz.png"
var image3="matrix.jpg"
var image4="monalisa.jpg"
var images = [image1,image2,image3,image4]
var images = shuffle(images)
var indexCount = 0