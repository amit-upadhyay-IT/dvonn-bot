var test = require('./100000games.json');

var val1 = 0;
test.forEach(function(value){
  if (JSON.stringify(value).includes('["j5|w","i5|b","c1')) {
  	console.log(val1)
  }
});

console.log(val1)