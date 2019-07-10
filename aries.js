var currentImagesLoaded=2;
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
    document.getElementById(0+"Top").style.opacity=1;
    document.getElementById(0+"Bottom").style.opacity=1;
    document.getElementById(0+"BL").style.opacity=1;
    document.getElementById(0+"BR").style.opacity=1;
    document.getElementById(0+"Left").style.opacity=1;
    //alert("Opacity:"+ document.getElementById(0+"Left").style.opacity);
    document.getElementById(0+"Right").style.opacity=1;





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
            nextImage()
        }
        if (obj.NextImage != toogle){
            nextImage();
            toogle = obj.NextImage;
        }
        if (obj.Reload == true){
            alert("Reloading page!");
            location.reload();
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
    
    var imgFlippedTop = new Image();
    imgFlippedTop.id = index+"Top";
    imgFlippedTop.src = window.imageList[index];
    imgFlippedTop.onload= resizeFlipped;
    imgFlippedTop.className="flippedVertical"

    var imgFlippedBottom = new Image();
    imgFlippedBottom.id = index+"Bottom";
    imgFlippedBottom.src = window.imageList[index];
    imgFlippedBottom.onload= resizeFlipped;
    imgFlippedBottom.className="flippedVertical"
    
    var imgFlippedBottomLeft = new Image();
    imgFlippedBottomLeft.id = index+"BL";
    imgFlippedBottomLeft.src = window.imageList[index];
    imgFlippedBottomLeft.onload= resizeFlipped;
    imgFlippedBottomLeft.className="flippedBoth"
    var imgFlippedBottomRight = new Image();
    imgFlippedBottomRight.id = index+"BR";
    imgFlippedBottomRight.src = window.imageList[index];
    imgFlippedBottomRight.onload= resizeFlipped;
    imgFlippedBottomRight.className="flippedBoth"

    var imgFlippedLeft = new Image();
    imgFlippedLeft.id = index+"Left";
    imgFlippedLeft.src = window.imageList[index];
    imgFlippedLeft.onload= resizeFlipped;
    imgFlippedLeft.className="flippedHorizontal"

    var imgFlippedRight = new Image();
    imgFlippedRight.id = index+"Right";
    imgFlippedRight.src = window.imageList[index];
    imgFlippedRight.onload= resizeFlipped;
    imgFlippedRight.className="flippedHorizontal"

   
    var div = document.getElementById("fader");
    
    if (document.getElementById(index) != null){
        oldImg = document.getElementById(index);
        oldImg.parentNode.removeChild(oldImg);
    }
    div.appendChild(imgFlippedTop);
    div.appendChild(img);
    div.appendChild(imgFlippedBottom);
    div.appendChild(imgFlippedBottomLeft);
    div.appendChild(imgFlippedBottomRight);

    div.appendChild(imgFlippedLeft);
    div.appendChild(imgFlippedRight);



   
    img.style.opacity=0;
    imgFlippedTop.style.opacity=0;
    imgFlippedBottom.style.opacity=0;
    imgFlippedBottomLeft.style.opacity=0;
    imgFlippedBottomRight.style.opacity=0;
    imgFlippedLeft.style.opacity=0
    imgFlippedRight.style.opacity=0


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
function resizeFlipped(e){
    var img = this; 
    newResolution = calculateAspectRatioFit(img.width, img.height, window.screen.availWidth,screen.availHeight)
    img.width = newResolution.width
    img.height = newResolution.height
    //centerimage!!
    switch (true) {
        case img.id.includes("Top"):
            img.style.top = -img.height + (window.screen.availHeight - img.height)/2
            img.style.left = (window.screen.availWidth - img.width)/2
            break;
        case img.id.includes("Bottom"):
            img.style.top = (window.screen.availHeight + img.height)/2
            img.style.left = (window.screen.availWidth - img.width)/2
            break;
        case img.id.includes("Left"):
            img.style.top = (window.screen.availHeight - img.height)/2
            img.style.left = -img.width + (window.screen.availWidth - img.width)/2
            break;
        case img.id.includes("Right"):
            img.style.top = (window.screen.availHeight - img.height)/2
            img.style.left = (window.screen.availWidth + img.width)/2
            break;  
        case img.id.includes("TL"):
            img.style.top = -img.height + (window.screen.availHeight - img.height)/2
            img.style.left = -img.width + (window.screen.availWidth - img.width)/2
            break;
        case img.id.includes("TR"):
            img.style.top = -img.height + (window.screen.availHeight - img.height)/2
            img.style.left = (window.screen.availWidth + img.width)/2
            break;
        case img.id.includes("BL"):
            img.style.top = (window.screen.availHeight + img.height)/2
            img.style.left = -img.width + (window.screen.availWidth - img.width)/2
            break;
        case img.id.includes("BR"):
            img.style.top = (window.screen.availHeight + img.height)/2
            img.style.left = (window.screen.availWidth + img.width)/2
            break;
        default: 
            break;
    }
}

function calculateAspectRatioFit(srcWidth, srcHeight, maxWidth, maxHeight) {

    var ratio = Math.min( (maxWidth+5) / srcWidth, (maxHeight+5) / srcHeight);
    return { width: srcWidth*ratio, height: srcHeight*ratio };
}


function nextImage(){
    window.clearInterval(window.ImageRefreshInterval);
    nextImageId = currentImageId+1;
    if (nextImageId >= window.imageList.length){
        nextImageId=0;
    }
    loadImage(nextImageId+1);
    if (nextImageId > currentImagesLoaded-1){
        currentImagesLoaded++;
        //loadImage(currentImagesLoaded-1);
        nextImageId=currentImagesLoaded-1;
    }
    console.log("loading image:" +nextImageId);
    currImg = document.getElementById(currentImageId);
    currImgTop = document.getElementById(currentImageId+"Top");
    currImgBottom = document.getElementById(currentImageId+"Bottom");
    currImgBottomLeft = document.getElementById(currentImageId+"BL");
    currImgBottomRight = document.getElementById(currentImageId+"BR");
    currImgLeft = document.getElementById(currentImageId+"Left");
    currImgRight = document.getElementById(currentImageId+"Right");



    nxtImg = document.getElementById(nextImageId);
    nxtImgTop = document.getElementById(nextImageId+"Top");
    nxtImgBottom = document.getElementById(nextImageId+"Bottom");    
    nxtImgBottomLeft = document.getElementById(nextImageId+"BL");
    nxtImgBottomRight = document.getElementById(nextImageId+"BR");
    nxtImgLeft = document.getElementById(nextImageId+"Left");
    nxtImgRight = document.getElementById(nextImageId+"Right");





    opacityLoop(currImg,currImg.style.opacity, false);  
    opacityLoop(currImgTop,currImgTop.style.opacity, false);  
    opacityLoop(currImgBottom,currImgBottom.style.opacity, false);
    opacityLoop(currImgBottomLeft,currImgBottomLeft.style.opacity, false);
    opacityLoop(currImgBottomRight,currImgBottomRight.style.opacity, false);  
    opacityLoop(currImgLeft,currImgLeft.style.opacity, false);  
    opacityLoop(currImgRight,currImgRight.style.opacity, false);  



    opacityLoop(nxtImg,0,true);  
    opacityLoop(nxtImgTop,0,true);  
    opacityLoop(nxtImgBottom,0,true);  
    opacityLoop(nxtImgBottomLeft,0,true);  
    opacityLoop(nxtImgBottomRight,0,true);  
    opacityLoop(nxtImgLeft,0,true);  
    opacityLoop(nxtImgRight,0,true);  




    setTimeout(function () {  
        clearOpacityOfRest(currentImageId);
    }, window.FadeDuration*1000)
    currentImageId=nextImageId;

    window.ImageRefreshInterval = window.setInterval(function () {
        nextImage()
    }, imageRefresh);  

}

function clearOpacityOfRest(indexCurrent){
    //loop through all imgs on the web page, set opacity to 9
    var images = document.getElementsByTagName('img'); 
    for (var index=0; index < images.length; index++){
        images[index].style.opacity=0
        //console.log("removing opacity of:"+ images[index].id);
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
    window.ImageRefreshInterval = window.setInterval(function () {
        nextImage()
    }, imageRefresh);  

    window.setInterval(function(){
        loadConfig()
    },1000);
}