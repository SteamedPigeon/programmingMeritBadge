
// JavaScript Document
// Note: lines that start with two backslashes are comments - not code!
 
var bePrepared = function () {
    // = = = = = = = declare all the variables = = = = = = = = 
    var tempF, tempC, myActionText, newText, tempWs, wc;

    //get the temp (F) from the HTML page
    tempF = document.getElementById('MyInputTemp').value;
tempWs = document.getElementById('MyInputWs').value;   // = = = = = = = convert the temp to Celsius (with only one decimal place)
    tempC = (5 / 9 * (tempF - 32)).toFixed(1);
    wc = 35.74+0.6215*tempF-35.75*Math.pow(tempWs,0.16)+0.4275*tempF*Math.pow(tempWs,0.16)

   // = = = = = = = evaluate the temp (three categories) = = = = = = =  
   if ((tempF < 60) || (wc<60)) {
         myActionText = " Take long-johns!";
     }
   if ((tempF >= 60) && (tempC < 75)) {
             myActionText = " Have Fun!";
         }
   if (tempF >= 75) {
             myActionText = " Take Sunscreen!";
         }

   // = = = = = = = build a complete sentence = = = = = = =  
   newText = "If the temperature is " + tempF + " &deg;F (" + tempC + " &deg;C): but it feels like " + wc.toFixed(1) + " &deg;F" + myActionText;
   //push the sentence back to the HTML page (using the ID of the markup element: 'myAnswer')
   document.getElementById('myAnswer').innerHTML = newText;
 };