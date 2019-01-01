var height
var width
var images
var indexCount = 0
var toogle=true;
var fader;
var imageRefresh;




window.onload = function start() {
    images = window.imageList
    height = window.screen.availHeight + 15 //compensate for scroll bars
    width = window.screen.availWidth - 25  //compensate for scroll bars
    //3815x2135 
    //vs la que segun trae 3840 x 2160
    //25 menos que la de specs
    //la real es
    //3840x2120
    resizeImgs()
    imageRefresh = window.imageRefresh
    fader = new Fader('fader', 2, height,width);
    canvas();
}
function resizeImgs(){
    for(i=0; i< images.length; i++){
        var img = document.getElementById(i); 
        newResolution = calculateAspectRatioFit(img.width,img.height,width,height)
        addCount = 10
        do{
            newResolution = calculateAspectRatioFit(img.width,img.height,width+addCount,height+addCount)
            addCount = addCount + 10
        }while(newResolution.height < height || newResolution.width < width)
        img.width = newResolution.width
        img.height = newResolution.height
    }
}

function changeImage() {   
    if (++indexCount >= images.length){
        indexCount = 0
    }

    fader.setTarget(indexCount)
}



function calculateAspectRatioFit(srcWidth, srcHeight, maxWidth, maxHeight) {
    var ratio = Math.min(maxWidth / srcWidth, maxHeight / srcHeight);
    return { width: srcWidth*ratio, height: srcHeight*ratio };
}

function canvas() {
    window.setInterval(function () {
        changeImage()
    }, imageRefresh);  
}


