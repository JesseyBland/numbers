package html

//index html webpage
const index = `
<!DOCTYPE html>
<html>
<head>
  <meta charset="UTF-8" name="viewport" content="width=device-width, initial-scale=1">
  <style>
    .comment {
      resize: none;
      height: 150px;
      width: 65px;
    }
  </style>
   
</head>
<body>
<div>

    Enter a number 0-9999<br>
    <input  id="gist" cols="4" rows="1" type="number" inputmode="numeric" min="0" max="9999" >
    </input>
               <button>Submit</button>
    <div id="output" class="output">
        <div id="content"></div>
        <div id="status"></div>
        <pre id="log"><textarea id=text style="font-family:'Courier New' fontSize=25px" class="comment">
          
          
        </textarea>
      </pre>
      </div>
</div>
<div>
    <script>

      var is_mobile = !!navigator.userAgent.match(/iphone|android|blackberry/ig) || false;
        
        function snum(){
            fetch("/cnumber", {
            headers: {
             "Content-Type": "application/x-www-form-urlencoded",
                 },
            method: "post",
            body: document.getElementById("gist").value,
                }).then(response =>response.text()).then(data =>document.getElementById("text").value=data);
              }
            

    
            var submitBtn = document.querySelector('button');
            submitBtn.addEventListener('click', snum);
    

    </script>

</div>
</body>
</html>
`
