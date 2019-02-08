var currentImagesLoaded=1;
var currentImageId=0;
var toogle=false;

window.onload = function start() {
    var images = window.imageList;
    var TypePage = window.TypePage;
    //load images
    for (var imagesIndex=0;imagesIndex<currentImagesLoaded;imagesIndex++){
        loadImage(imagesIndex);
    }
    document.getElementById(0).style.opacity=1;

    canvas();
}

function loadConfig() {
    var xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function() {
      if (this.readyState == 4 && this.status == 200) {
        var obj = JSON.parse(this.responseText);
        if (obj.TypeOfImage != window.TypeOfImage){
            window.imageList = obj.Images
            window.TypeOfImage = obj.TypeOfImage
            nextImage()
        }
        if (obj.NextImage != toogle){
            nextImage();
            toogle = obj.NextImage;
        }
      }
    };
    xhttp.open("GET", "configs/" + window.TypePage + ".config", true);
    xhttp.send();
  }

function loadImage(index){
    var img = new Image();
    img.src = window.imageList[index];
    img.id = index;
    img.onload= resize;
    var div = document.getElementById("fader");
    if (document.getElementById(index) == null){
        div.appendChild(img);
    }else{
        oldImg = document.getElementById(index);
        oldImg.src = img.src;
        oldImg.id = img.id;
        oldImg.onload = resize;
    }
   
    img.style.opacity=0;
}
function resize(e){
    var img = this; 
    newResolution = calculateAspectRatioFit(img.width, img.height, window.screen.availWidth,screen.availHeight)
    img.width = newResolution.width
    img.height = newResolution.height
    //centerimage!!
    img.style.top = (window.screen.availHeight - img.height)/2
    img.style.left = (window.screen.availWidth - img.width)/2
}

function calculateAspectRatioFit(srcWidth, srcHeight, maxWidth, maxHeight) {

    var ratio = Math.min( (maxWidth+5) / srcWidth, (maxHeight+5) / srcHeight);
    return { width: srcWidth*ratio, height: srcHeight*ratio };
}


function nextImage(){
    nextImageId = currentImageId+1;
    if (nextImageId >= window.imageList.length){
        nextImageId=0;
    }
    loadImage(nextImageId);
   
    if (nextImageId > currentImagesLoaded-1){
        currentImagesLoaded++;
        //loadImage(currentImagesLoaded-1);
        nextImageId=currentImagesLoaded-1;
    }
    console.log("loading image:" +nextImageId);
    currImg = document.getElementById(currentImageId);
    nxtImg = document.getElementById(nextImageId);

    opacityLoop(currImg,currImg.style.opacity, false);  
    opacityLoop(nxtImg,0,true);  
    setTimeout(function () {  
        clearOpacityOfRest(currentImageId);
    }, window.FadeDuration*1000)
    currentImageId=nextImageId;

}

function clearOpacityOfRest(indexCurrent){
    for (var imagesIndex=0;imagesIndex<currentImagesLoaded;imagesIndex++){
        if (imagesIndex != indexCurrent){
            document.getElementById(imagesIndex).style.opacity=0;
        }
    }
}


function opacityLoop (image,opacity,reverse) {  
    var rate = 0.05 / window.FadeDuration;        
    setTimeout(function () {  
       if(!reverse){
            image.style.opacity = opacity - rate;
            if (image.style.opacity > 0) opacityLoop(image,opacity - rate,false);
        }else{
            image.style.opacity = opacity + rate;
            if (image.style.opacity < 1) opacityLoop(image,opacity + rate,true);
        }
    }, 50)
 }


function canvas() {
    window.setInterval(function () {
        nextImage()
    }, imageRefresh);  

    window.setInterval(function(){
        loadConfig()
    },1000);
}