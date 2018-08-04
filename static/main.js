function init() {
    requestManagers(displayManagers);

    requestManager(12, displayManager);

    loadQuote();
}

function requestManagers(callback) {
    jsonRequest("/api/managers", callback);
}

function displayManagers(managersResponse) {
    document.getElementById("listManagers").innerHTML = JSON.stringify(managersResponse, null, 2);
}

function requestManager(id, callback) {
    jsonRequest("/api/managers/" + id, callback);
}

function displayManager(managerResponse) {
    document.getElementById("getManager").innerHTML = JSON.stringify(managerResponse, null, 2);
}

function loadQuote() {
    const button = document.getElementById("quoteButton");
    button.disabled = true;

    requestQuote(function (quote) {
        displayQuote(quote);
        button.disabled = false;
    });
}

function requestQuote(callback) {
    var response = jsonRequest("/api/random_quote", callback);
}

function displayQuote(quote) {
    document.getElementById("quote").style = "visibility: visible";
    let quoteElement = document.getElementById("quote");

    quoteElement.setAttribute("cite", quote.author);
    quoteElement.innerHTML = quote.text;
}

function jsonRequest(url, callback) {
    var xhttp = new XMLHttpRequest();
    xhttp.open("GET", url, true);
    xhttp.setRequestHeader("Content-type", "application/json");

    xhttp.onreadystatechange = function () {
        if (xhttp.readyState == 4) {
            if (xhttp.status == 200) {
                callback(JSON.parse(xhttp.responseText));
            } else {
                console.log("error. status: " + xhttp.status + ' - body: ' + xhttp.responseText)
            }
        }
    };

    xhttp.send();
}
