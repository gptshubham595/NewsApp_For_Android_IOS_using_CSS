var clicks = 0;
var a = document.getElementById("bookmark");
var rC = document.getElementById("relevanceColor");
rC.style.color = "lightgreen";
function go() {
    clicks++;
    if (clicks % 2 === 1) {
        a.style.color = "lightgreen";
    }
    else {
        a.style.color = "#3498DB";
    }
};


