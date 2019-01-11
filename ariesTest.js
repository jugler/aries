window.onload = function start() {
    var images = window.imageList;
    var TypePage = window.TypePage;
    for(var i=0; i < images.length; i++){
        if (TypePage == "portrait/"){
            resizeImgsPortrait(i)
        }else{
            resizeImgs(i)
        }
    }

    reduceOpacity(window.imageList.length);
    canvas();
}
function reduceOpacity(imgNumber){
    for(var i=0;i<imgNumber-1;i++){
        document.getElementById(i).style.opacity=0;
    }
    document.getElementById(imgNumber-1).style.opacity=1;
}

function resizeImgs(id){
    var img = document.getElementById(id); 
    if (img == null){
        alert(id)
    }
    newResolution = calculateAspectRatioFit(img.width, img.height, window.screen.availWidth,screen.availHeight)
    img.width = newResolution.width
    img.height = newResolution.height

    //centerimage!!
    var left = (window.screen.availWidth - img.width)/2
    img.style.left = left;
    var top = (window.screen.availHeight - img.height)/2

    img.style.top = top;

}
function resizeImgsPortrait(id){
    var img = document.getElementById(id); 
    
    newResolution = calculateAspectRatioFit(img.width, img.height, window.screen.availWidth,screen.availHeight)
    img.width = newResolution.width
    img.height = newResolution.height

    //centerimage!!
    var top = (window.screen.availHeight - img.height)/2
    var left = (window.screen.availWidth - img.width)/2
    img.style.top = top;
    img.style.left = left;
}
function calculateAspectRatioFit(srcWidth, srcHeight, maxWidth, maxHeight) {

    var ratio = Math.min( (maxWidth+5) / srcWidth, (maxHeight+5) / srcHeight);
    return { width: srcWidth*ratio, height: srcHeight*ratio };
}

var currentImageId=window.imageList.length-1;

function nextImage(){
    nextImageId = currentImageId+1;
    if (nextImageId > window.imageList.length-1){
        nextImageId = 0
    }
    currImg = document.getElementById(currentImageId);
    nxtImg = document.getElementById(nextImageId);

    opacityLoop(currImg,currImg.style.opacity, false);  
    opacityLoop(nxtImg,0, true);  

    currentImageId=nextImageId;

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
}