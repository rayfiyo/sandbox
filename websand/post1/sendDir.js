function sendDir() {
  var json_raw = {
    dir: "../",
  };

  var json_data = JSON.stringify(json_raw);

  //send
  xhr = new XMLHttpRequest();
  xhr.onload = function () {
    //response process
    var dirs = xhr.responseText;
    if (dirs.length > 0) alert(dirs);
  };
  xhr.onerror = function () {
    //error handling
    alert("error!");
  };
  xhr.open("post", "dummy.php", true);
  xhr.setRequestHeader("Content-Type", "application/json");
  xhr.send(json_data);
  console.log(dirs);
}
