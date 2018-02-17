var answer = prompt("Give me the url please?");
if (answer === null) {
  console.log("Invalid Response");
} else {
  $.get(window.location.href + name, function(response) {
  console.log(response);
  document.write(response);
});
}
