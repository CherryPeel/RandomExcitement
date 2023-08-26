function refreshImage() {
    var imgElement = document.getElementById("randomImage");
    imgElement.src = "/v1/randomExcitement?" + new Date().getTime();
}
